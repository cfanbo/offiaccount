package oauth

// @link https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Official_Accounts/official_account_website_authorization.html

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/cfanbo/offiaccount/util"

	"github.com/levigross/grequests"
)

type AuthorizerRedirect struct {
	componentAppid                          string
	appId, redirectUri, responseType, scope string
	state                                   string
}

func NewAuthorizerRedirect(componentAppid string) *AuthorizerRedirect {
	return &AuthorizerRedirect{
		componentAppid: componentAppid,
		responseType:   "code",
	}
}

func (a *AuthorizerRedirect) SetAppId(appId string) *AuthorizerRedirect {
	a.redirectUri = appId

	return a
}
func (a *AuthorizerRedirect) SetRedirectUri(url string) *AuthorizerRedirect {
	a.appId = url

	return a
}

func (a *AuthorizerRedirect) SetScope(scope string) *AuthorizerRedirect {
	a.scope = scope

	return a
}

func (a *AuthorizerRedirect) SetResponseType(t string) *AuthorizerRedirect {
	a.responseType = t

	return a
}

func (a *AuthorizerRedirect) SetState(state string) *AuthorizerRedirect {
	a.state = state

	return a
}

func (a *AuthorizerRedirect) String() string {
	buf := strings.Builder{}

	buf.WriteString("https://open.weixin.qq.com/connect/oauth2/authorize")
	buf.WriteByte('?')
	buf.WriteString("&appid=" + a.appId)
	buf.WriteString("&redirect_uri=" + url.QueryEscape(a.redirectUri))
	buf.WriteString("&response_type=" + a.responseType)
	buf.WriteString("&scope=" + a.scope)
	buf.WriteString("&state=" + a.state)
	buf.WriteString("&component_appid=" + a.componentAppid)
	buf.WriteByte('#')
	buf.WriteString("wechat_redirect")

	return buf.String()
}

type AuthorizerResult struct {
	Code  string `json:"code"`
	State string `json:"state"`
	Appid string `json:"appid"`
}

func GetCodeByCallbackUrl(rawQuery string) (url.Values, error) {
	return url.ParseQuery(rawQuery)
}

type AccessTokenByCodeRequest struct {
	AppId, Code, GrantType               string
	ComponentAppid, ComponentAccessToken string
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func GetAccessTokenByCode(r AccessTokenByCodeRequest) (AccessTokenResponse, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/component/access_token"

	params := map[string]string{
		"component_access_token": r.ComponentAccessToken,
	}
	data := map[string]string{
		"component_appid": r.ComponentAppid,
		"appid":           r.AppId,
		"code":            r.Code,
		"grant_type":      r.GrantType,
	}
	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer resp.Close()

	result := AccessTokenResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return result, nil
}

type RefreshAccessTokenRequest struct {
	AppId, RefreshToken, GrantType       string
	ComponentAppid, ComponentAccessToken string
}

func RefreshAccessToken(r RefreshAccessTokenRequest) (AccessTokenResponse, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/component/refresh_token"

	data := map[string]string{
		"appid":                  r.AppId,
		"grant_type":             r.GrantType,
		"refresh_token":          r.RefreshToken,
		"component_appid":        r.ComponentAppid,
		"component_access_token": r.ComponentAccessToken,
	}
	requestOptions := &grequests.RequestOptions{
		Data: data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer resp.Close()

	result := AccessTokenResponse{}
	err = json.Unmarshal(resp.Bytes(), &result)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return result, nil
}

type SnsUserInfoResponse struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege  string `json:"privilege"`
	Unionid    string `json:"unionid"`
}

func GetSnsUserInfo(accessToken, openid, lang string) (SnsUserInfoResponse, error) {
	url := "https://api.weixin.qq.com/sns/userinfo"
	params := map[string]string{
		"access_token": accessToken,
		"openid":       openid,
		"lang":         lang,
	}
	body, err := util.HttpGet(url, params)
	if err != nil {
		return SnsUserInfoResponse{}, err
	}

	result := SnsUserInfoResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SnsUserInfoResponse{}, err
	}

	return result, nil
}
