package component

import (
	"encoding/json"
	"strconv"

	"github.com/cfanbo/offiaccount/util"
)

type AuthorizerListItem struct {
	AuthorizerAppid string `xml:"authorizer_appid"`
	RefreshToken    string `xml:"refresh_token"`
	AuthTime        int    `xml:"auth_time"`
}

type AuthorizerListResponse struct {
	TotalCount string `xml:"total_count"`
	List       []AuthorizerListItem
}

// NewAuthorizeEvent
func (c *Component) GetAuthorizerList(offset, count uint) (AuthorizerListResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list"

	if count > 500 {
		count = 500
	}

	params := map[string]string{
		"component_access_token": c.AccessToken,
	}

	data := map[string]string{
		"component_appid": c.AppId,
		"offset":          strconv.Itoa(int(offset)),
		"count":           strconv.Itoa(int(count)),
	}

	body, err := util.HttpPost(url, params, data)
	if err != nil {
		return AuthorizerListResponse{}, nil
	}

	result := AuthorizerListResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return AuthorizerListResponse{}, err
	}

	return result, nil
}
