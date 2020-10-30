package offiaccount

import (
	"encoding/xml"
	"time"
)

/*

<xml>
<AppId>some_appid</AppId>
<CreateTime>1413192605</CreateTime>
<InfoType>component_verify_ticket</InfoType>
<ComponentVerifyTicket>some_verify_ticket</ComponentVerifyTicket>
</xml>

*/
type Ticket struct {
	AppId                 string `xml:"AppId"`
	CreateTime            int    `xml:"CreateTime"`
	InfoType              string `xml:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"`
}

// NewTicket
func NewTicket(xmlData []byte) (Ticket, error) {
	var ticket = Ticket{}
	err := xml.Unmarshal(xmlData, &ticket)
	if err != nil {
		return Ticket{}, err
	}

	return ticket, nil
}

// GetAppId appId
func (t Ticket) GetAppId() string {
	return t.AppId
}

// GetCreateTime UNIX日期
func (t Ticket) GetCreateTime() int {
	return t.CreateTime
}

// GetCreateTimeString 字符串格式
func (t Ticket) GetCreateTimeString() string {
	return time.Unix(int64(t.CreateTime), 0).Format("2006-01-02 15:04:05")
}

// GetTicket  ticket
func (t Ticket) GetTicket() string {
	return t.ComponentVerifyTicket
}
