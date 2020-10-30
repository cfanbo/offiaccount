package component

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type AuthorizerAccessTokenResponse struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// GetAuthorizerToken
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_authorizer_token.html
func (c *Component) GetAuthorizerAccessToken(authorizerAppid, authorizerRefreshToken string) (AuthorizerAccessTokenResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"
	params := map[string]string{
		"component_access_token": c.AccessToken,
	}
	data := map[string]string{
		"component_appid":          c.AppId,
		"authorizer_appid":         authorizerAppid,
		"authorizer_refresh_token": authorizerRefreshToken,
	}
	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return AuthorizerAccessTokenResponse{}, err
	}
	defer resp.Close()

	result := AuthorizerAccessTokenResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return AuthorizerAccessTokenResponse{}, err
	}

	return result, nil
}
