package offiaccount

import (
	simplejson "github.com/bitly/go-simplejson"
	"github.com/cfanbo/offiaccount/util"
)

type AccessTokenResult struct {
	AccessToken string
	ExpiresIn   uint64
}

func AccessToken(appId, secret string) (AccessTokenResult, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token"

	params := map[string]string{
		"grant_type": "client_credential",
		"appid":      appId,
		"secret":     secret,
	}

	body, err := util.HttpGet(url, params)
	if err != nil {
		return AccessTokenResult{}, err
	}

	json, err := simplejson.NewJson(body)
	if err != nil {
		return AccessTokenResult{}, err
	}

	if js, ok := json.CheckGet("errcode"); ok {
		errcode, _ := js.Int()
		return AccessTokenResult{}, newError(errcode)
	}

	ac, _ := json.Get("access_token").String()
	in, _ := json.Get("expires_in").Uint64()
	a := AccessTokenResult{
		AccessToken: ac,
		ExpiresIn:   in,
	}

	return a, nil
}
