package offiaccount

import (
	"errors"
	"strconv"
)

func newError(code int) error {
	var msg string
	switch code {
	case -1:
		msg = "系统繁忙，此时请开发者稍候再试"
	case 0:
		msg = "请求成功"
	case 40001:
		msg = "AppSecret错误或者AppSecret不属于这个公众号，请开发者确认AppSecret的正确性"
	case 40002:
		msg = "请确保grant_type字段值为client_credential"
	case 40013:
		msg = "invalid appid"
	case 40164:
		msg = "调用接口的IP地址不在白名单中，请在接口IP白名单中进行设置。（小程序及小游戏调用不要求IP地址在白名单内。）"
	case 89503:
		msg = "此IP调用需要管理员确认,请联系管理员"
	case 89501:
		msg = "此IP正在等待管理员确认,请联系管理员"
	case 89506:
		msg = "24小时内该IP被管理员拒绝调用两次，24小时内不可再使用该IP调用"
	case 89507:
		msg = "1小时内该IP被管理员拒绝调用一次，1小时内不可再使用该IP调用"
	default:
		msg = "未知错误，请查看官方错误编号：" + strconv.Itoa(code)
	}
	return errors.New(strconv.Itoa(code) + msg)
}
