package app_param

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
)

type (
	ArgParamsUserHaveConsoleImport struct {
		User      *RequestUser  `json:"user" form:"user"`             //用户
		ImportKey string        `json:"import_key" form:"import_key"` //接口KEY
		Ctx       *base.Context `json:"-" form:"-"`
	}
	ResultConsoleHaveImportPermit struct {
		StatusCode    int    `json:"status_code"`     //异常的状态码
		NotHavePermit bool   `json:"not_have_permit"` //true-没有权限 false-有权限
		IsSuper       bool   `json:"is_super"`        //是否为超级管理员
		ErrorMsg      string `json:"error_msg"`       //错误提示内容
	}
)

func (r *ArgParamsUserHaveConsoleImport) Default(c *base.Context) (err error) {
	return
}

//判断用户是否有接口权限
func GetUserHaveConsoleImportPermit(arg *ArgParamsUserHaveConsoleImport) (res *ResultConsoleHaveImportPermit, err error) {
	res = &ResultConsoleHaveImportPermit{NotHavePermit: true, ErrorMsg: "接口异常"}
	var value = url.Values{}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     AppNameAdmin,
		URI:         "/get_have_permit",
		Header:      http.Header{},
		Value:       value,
		Context:     arg.Ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = json.Marshal(arg); err != nil {
		return
	}
	var data = struct {
		Code int                            `json:"code"`
		Data *ResultConsoleHaveImportPermit `json:"data"`
		Msg  string                         `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	if data.Data != nil {
		res = data.Data
	}

	return
}
