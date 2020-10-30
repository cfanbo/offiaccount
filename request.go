package offiaccount

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"log"
	"strings"
	"sync"
)

type RootRequest struct {
	AppId   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}

// NewRootRequest 将请求xml格式解析为请求结构体
func NewRootRequest(rawData []byte) (RootRequest, error) {
	payload := RootRequest{}
	err := xml.Unmarshal(rawData, &payload)
	if err != nil {
		return RootRequest{}, err
	}

	return payload, nil
}

// GetAppId
func (r RootRequest) GetAppId() string {
	return r.AppId
}

// GetData
func (r RootRequest) GetEncryptData() string {
	return r.Encrypt
}

var (
	ERR_APPID error = errors.New("无效appId")
)

var once sync.Once

type Request struct {
	AppId  string
	AESKey []byte
}

var request *Request

func NewRequest(appId, key string) *Request {
	// aeskey
	data, err := base64.StdEncoding.DecodeString(key + "=")
	if err != nil {
		log.Fatal(err)
	}

	once.Do(func() {
		request = &Request{
			AppId:  appId,
			AESKey: []byte(data),
		}
	})

	return request
}

// parse 将接收的数据流解析出真实有效的数据
func (r *Request) parse(payload []byte) ([]byte, error) {
	// 一、请求内容解析为struct
	rootReq, err := NewRootRequest(payload)
	if err != nil {
		log.Fatal(err)
	}

	// 二、内容解密
	body, err := Decrypt(rootReq.GetEncryptData(), r.AESKey)
	if err != nil {
		log.Fatal(err)
	}

	// 三、解析字符串里真实有效数据
	return r.GetPayload(body, r.AppId)
}

// GetTicket 解析Ticket请求
func (r *Request) NewTicket(payload []byte) (Ticket, error) {
	body, err := r.parse(payload)
	if err != nil {
		return Ticket{}, err
	}

	// 将 xml 解析为 struct
	return NewTicket(body)
}

// GetTicketPayload 获取解密后的有效payload
func (r *Request) GetPayload(body []byte, appId string) ([]byte, error) {
	if !strings.HasSuffix(string(body), appId) {
		return nil, ERR_APPID
	}

	// 4 个字节(网络字节序)
	b := body[4:]

	// 尾串appid
	b = b[:len(b)-len(appId)]

	return b, nil
}
