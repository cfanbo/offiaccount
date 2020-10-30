package component

import (
	"time"
)

// 第三方平台 access_token
type App struct {
	// 第三方平台 appid
	AppId string `json:"component_appid"`

	// 第三方平台 appsecret
	AppSecret string `json:"component_appsecret"`
}

type Component struct {
	// 第三方平台 appid
	App

	// 微信后台推送的 ticket
	VerifyTicket string `json:"component_verify_ticket"`

	// 第三方平台 access_token
	AccessToken string `json:"component_access_token"`

	// 有效期，单位：秒
	ExpiresIn time.Time `json:"expires_in"`
}

func NewComponent() *Component {
	return &Component{}
}
