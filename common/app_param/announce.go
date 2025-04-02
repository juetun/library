package app_param

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgAddUserMessage struct {
		Data   []*ArgAddUserMessageItem `json:"data" form:"data"`
		Common base.GetDataTypeCommon   `json:"-" form:"-"`
	}
	ArgAddUserMessageItem struct {
		NoticeTemplateKey string                 `json:"notice_template_key"` //通知模版KEY
		ToUserHid         int64                  `json:"to_user_hid" form:"to_user_hid"`
		Content           map[string]interface{} `json:"content" form:"content"` //内容映射(配合模版使用）
	}
	ResultAddUserMessage struct {
		Result bool `json:"result"`
	}
)

func (r *ArgAddUserMessage) Default(c *base.Context) (err error) {
	_ = c
	return
}