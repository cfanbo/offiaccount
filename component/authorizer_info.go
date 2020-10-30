package component

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type AuthorizerInfoResponse struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
	AuthorizerInfo    AuthorizerInfos   `json:"authorizer_info"`
}

type AuthorizationInfo struct {
	NickName        string         `json:"nick_name"`
	HeadImg         string         `json:"head_img"`
	ServiceTypeInfo map[string]int `json:"service_type_info"`
	VerityTypeInfo  map[string]int `json:"verify_type_info"`
	UserName        string         `json:"user_name"`
	PrincipalName   string         `json:"principal_name"`
	Alias           string         `json:"alias"`
	BusinessInfo    map[string]int `json:"business_info"`
	QrcodeUrl       string         `json:"qrcode_url"`
}

type AuthorizerInfos struct {
	NickName        string         `json:"nick_name"`
	HeadImg         string         `json:"head_img"`
	ServiceTypeInfo map[string]int `json:"service_type_info"`
	VerityTypeInfo  map[string]int `json:"verify_type_info"`
	UserName        string         `json:"user_name"`
	PrincipalName   string         `json:"principal_name"`
	Alias           string         `json:"alias"`
	BusinessInfo    map[string]int `json:"business_info"`
	QrcodeUrl       string         `json:"qrcode_url"`
}

// GetAuthorizerInfo
// @link https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_get_authorizer_info.html
func (c *Component) GetAuthorizerInfo(authorizerAppid string) (AuthorizerInfoResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"

	params := map[string]string{
		"component_access_token": c.AccessToken,
	}

	data := map[string]string{
		"component_appid":  c.AppId,
		"authorizer_appid": authorizerAppid,
	}

	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return AuthorizerInfoResponse{}, err
	}
	defer resp.Close()

	result := AuthorizerInfoResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return AuthorizerInfoResponse{}, err
	}

	return result, nil
}
