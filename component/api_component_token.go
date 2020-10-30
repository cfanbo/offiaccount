package component

import (
	"encoding/json"

	"github.com/cfanbo/offiaccount/util"
)

type AccessTokenResponse struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

// GetComponentAccessToken
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_access_token.html
func (c *Component) GetComponentAccessToken(componentVerifyTicket string) (AccessTokenResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"
	params := map[string]string{}
	data := map[string]string{
		"component_appid":         c.AppId,
		"component_appsecret":     c.AppSecret,
		"component_verify_ticket": componentVerifyTicket,
	}

	body, err := util.HttpPost(url, params, data)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	result := AccessTokenResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return result, nil
}
