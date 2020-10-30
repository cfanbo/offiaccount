package component

import (
	"encoding/json"

	"github.com/cfanbo/offiaccount"

	"github.com/levigross/grequests"
)

type AuthorizerOptionItem struct {
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
}

type AuthorizerOption struct {
	AuthorizerAppid string `json:"authorizer_appid"`
	AuthorizerOptionItem
}

type AuthorizerOptionResponse struct {
	AuthorizerOption
}

// @link https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_get_authorizer_option.html
func (c *Component) GetAuthorizerOption(authorizerAppid string, optionName string) (AuthorizerOptionResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option"

	params := map[string]string{
		"component_access_token": c.AccessToken,
	}

	data := map[string]string{
		"component_appid":  c.AppId,
		"authorizer_appid": authorizerAppid,
		"option_name":      optionName,
	}

	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return AuthorizerOptionResponse{}, err
	}
	defer resp.Close()

	result := AuthorizerOptionResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return AuthorizerOptionResponse{}, err
	}

	return result, nil
}

func (c *Component) SetAuthorizerOption(authorizerAppid string, option AuthorizerOptionItem) (offiaccount.SuccessResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option"

	params := map[string]string{
		"component_access_token": c.AccessToken,
	}

	data := map[string]string{
		"component_appid":  c.AppId,
		"authorizer_appid": authorizerAppid,
		"option_name":      option.OptionName,
		"option_value":     option.OptionValue,
	}

	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return offiaccount.SuccessResponse{}, err
	}
	defer resp.Close()

	result := offiaccount.SuccessResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return offiaccount.SuccessResponse{}, err
	}

	return result, nil
}
