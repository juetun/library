package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	MobileCountryCodeChina = "86"
)
const (
	ShortMessageTypeUpdateLogin    int = iota + 1 //登录
	ShortMessageTypeUpdateRegister                //注册
	ShortMessageTypeUpdateMobile                  //修改手机号
)

var (
	SliceShortMessageType = base.ModelItemOptions{
		{
			Label: "登录短信验证",
			Value: ShortMessageTypeUpdateLogin,
		},
		{
			Label: "注册短信验证",
			Value: ShortMessageTypeUpdateRegister,
		},
		{
			Label: "手机修改短信验证",
			Value: ShortMessageTypeUpdateMobile,
		},
	}
)

type (
	//短信验证码基本逻辑
	ValidateShortMessage struct {
		Type        int    `json:"type" form:"type"` //验证码类型
		UserHid     int64  `json:"user_hid" form:"user_hid"`
		Mobile      string `json:"mobile" form:"mobile"`
		CountryCode string `json:"country_code" form:"country_code"`
	}
)

func (r *ValidateShortMessage) Default() (err error) {
	if r.CountryCode == "" {
		r.CountryCode = MobileCountryCodeChina
	}
	if r.Mobile == "" {
		err = fmt.Errorf("请输入手机号")
		return
	}
	var (
		mapType map[int]string
	)
	if mapType, err = SliceShortMessageType.GetMapAsKeyInt(); err != nil {
		return
	}
	if _, ok := mapType[r.Type]; !ok {
		err = fmt.Errorf("验证码类型暂不支持(%v)", r.Type)
		return
	}

	return
}
