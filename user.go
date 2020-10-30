package offiaccount

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cfanbo/offiaccount/util"

	"github.com/bitly/go-simplejson"
	"github.com/levigross/grequests"
)

type User struct {
	Subscribe      uint8  `json:"subscribe"`
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            uint8  `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	Headimgurl     string `json:"headimgurl"`
	SubscribeTime  uint64 `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	Groupid        uint32 `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        uint32 `json:"qr_scene"`
	QrScene_str    string `json:"qr_scene_str"`
}

type UserList struct {
	users []User
}

// @link https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func GetUserInfo(accessToken, openid string) (User, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info"
	params := map[string]string{
		"access_token": accessToken,
		"openid":       openid,
		"lang":         "zh_CN",
	}

	result, err := util.HttpGet(url, params)
	if err != nil {
		return User{}, err
	}

	// 检查是否出错
	js, err := simplejson.NewJson(result)
	if err != nil {
		return User{}, err
	}
	if ret, ok := js.CheckGet("errcode"); ok {
		errcode, _ := ret.Int()
		return User{}, newError(errcode)
	}

	var user User
	json.Unmarshal(result, &user)

	return user, err
}

type UserListItem struct {
	Openid string `json:"openid"`
	Lang   string `json:"lang"`
}

type BatchGetUserRequest struct {
	UserList []UserListItem `json:"user_list"`
}

func (r *BatchGetUserRequest) addUserItem(item UserListItem) {
	r.UserList = append(r.UserList, item)
}

type BatchGetUserResponse struct {
	UserInfoList []User `json:"user_info_list"`
}

// @link
func BatchGetUser(accessToken string, UserListItem ...UserListItem) (BatchGetUserResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	params := map[string]string{
		"access_token": accessToken,
	}

	req := &BatchGetUserRequest{}
	for _, user := range UserListItem {
		req.addUserItem(user)
	}

	jsonData, _ := json.Marshal(req)
	requestOption := &grequests.RequestOptions{
		Params: params,
		JSON:   jsonData,
	}
	resp, err := grequests.Post(url, requestOption)
	if err != nil {
		return BatchGetUserResponse{}, err
	}
	defer resp.Close()

	result, _ := ioutil.ReadAll(resp.RawResponse.Body)
	fmt.Println(string(result))

	// 检查是否出错
	js, err := simplejson.NewJson(result)
	if err != nil {
		return BatchGetUserResponse{}, err
	}
	if ret, ok := js.CheckGet("errcode"); ok {
		errcode, _ := ret.Int()
		return BatchGetUserResponse{}, newError(errcode)
	}

	var users BatchGetUserResponse
	json.Unmarshal(result, &users)

	return users, nil
}

type GetUserListResponse struct {
	Total      uint32   `json:"totle"`
	Count      uint32   `json:"count"`
	Data       []string `json:"data"`
	NextOpenid string   `json:"next_openid"`
}

func GetUserList(accessToken, NextOpenid string) (GetUserListResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info"
	params := map[string]string{
		"access_token": accessToken,
		"next_openid":  NextOpenid,
	}

	result, err := util.HttpGet(url, params)
	if err != nil {
		return GetUserListResponse{}, err
	}

	// 检查是否出错
	js, err := simplejson.NewJson(result)
	if err != nil {
		return GetUserListResponse{}, err
	}
	if ret, ok := js.CheckGet("errcode"); ok {
		errcode, _ := ret.Int()
		return GetUserListResponse{}, newError(errcode)
	}

	var ret GetUserListResponse
	json.Unmarshal(result, &ret)

	return ret, err
}
