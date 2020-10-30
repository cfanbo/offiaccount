package component

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type PreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetPreAuthCode 获取预授权码
// @link https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/pre_auth_code.html
func (c *Component) GetPreAuthCode() (*PreAuthCode, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"
	params := map[string]string{
		"component_access_token": c.AccessToken,
	}
	data := map[string]string{
		"component_appid": c.AppId,
	}
	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return &PreAuthCode{}, nil
	}
	defer resp.Close()

	authCode := PreAuthCode{}
	json.Unmarshal(resp.Bytes(), &authCode)

	return &authCode, nil
}
