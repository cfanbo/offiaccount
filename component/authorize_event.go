package component

// @link https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/authorize_event.html

import (
	"encoding/json"

	"github.com/bitly/go-simplejson"
)

type AuthorizeEvent struct {
	AppId           string `xml:"AppId"`
	CreateTime      int    `xml:"CreateTime"`
	InfoType        string `xml:"InfoType"`
	AuthorizerAppid string `xml:"AuthorizerAppid"`
}

// 授权成功 更新授权
type AuthorizeEventSuccess struct {
	AuthorizeEvent
	AuthorizationCode            string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode"`
}

// 取消授权
type AuthorizeEventCancel struct {
	AuthorizeEvent
}

// NewAuthorizeEvent
func NewAuthorizeEvent(xmlData []byte) (interface{}, error) {
	js, err := simplejson.NewJson(xmlData)
	if err != nil {
		return nil, err
	}

	// 成功授权 和 更新授权
	_, ok := js.CheckGet("PreAuthCode")
	if ok {
		successResponse := AuthorizeEventSuccess{}
		json.Unmarshal(xmlData, &successResponse)

		// 统一返回成功
		return successResponse, nil
	}

	cancelResponse := AuthorizeEventCancel{}
	json.Unmarshal(xmlData, &cancelResponse)

	return cancelResponse, nil
}
