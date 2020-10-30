package component

import (
	"encoding/json"
	"log"

	"github.com/levigross/grequests"
)

type ApiQueryAuthResponse struct {
	AuthorizationInfo AuthorizerInfo `json:"authorization_info"`
}

type AuthorizerInfo struct {
	AuthorizerAppid        string         `json:"authorizer_appid"`
	AuthorizerAccessToken  string         `json:"authorizer_access_token"`
	ExpiresIn              int            `json:"expires_in"`
	AuthorizerRefreshToken string         `json:"authorizer_refresh_token"`
	FuncInfo               []FuncInfoItem `json:"func_info"`
}

type FuncInfoItem struct {
	FuncscopeCategory map[string]string `json:"funcscope_category"`
}

// GetApiQueryAuth
func (c *Component) GetApiQueryAuth(preAuthCode string) (ApiQueryAuthResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"
	params := map[string]string{
		"component_access_token": c.AccessToken,
	}
	data := map[string]string{
		"component_appid":    c.AppId,
		"authorization_code": preAuthCode,
	}

	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		log.Fatalf("api_query_auth err:%s", err)
		return ApiQueryAuthResponse{}, nil
	}
	defer resp.Close()

	result := ApiQueryAuthResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return ApiQueryAuthResponse{}, err
	}

	return result, nil
}
