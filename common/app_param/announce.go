package app_param

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

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

//发送消息通知
func PublicUserMessage(arg *ArgAddUserMessage, ctx *base.Context) (resData *ResultAddUserMessage, err error) {
	resData = &ResultAddUserMessage{}
	var value = url.Values{}
	var bodyByte []byte

	//判断参数是否为空
	if arg == nil {
		return
	}

	if bodyByte, err = json.Marshal(arg); err != nil {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameNotice,
		URI:         "/announce/add_user_message",
		Header:      http.Header{},
		Value:       value,
		BodyJson:    bodyByte,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                   `json:"code"`
		Data *ResultAddUserMessage `json:"data"`
		Msg  string                `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	if data.Data != nil {
		resData = data.Data
	}

	return
}

func (r *ArgAddUserMessage) Default(c *base.Context) (err error) {
	_ = c
	return
}
