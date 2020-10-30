package offiaccount

import (
	"github.com/bitly/go-simplejson"
	"github.com/cfanbo/offiaccount/util"
)

// GetCallbackIp 获取微信服务器IP地址
func GetCallbackIp(accessToken string) ([]string, error) {
	url := "https://api.weixin.qq.com/cgi-bin/getcallbackip"
	params := map[string]string{
		"access_token": accessToken,
	}
	body, err := util.HttpGet(url, params)
	if err != nil {
		return nil, err
	}

	json, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	return json.Get("ip_list").StringArray()
}
