package app_param

import (
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
)

type (
	ArgParamsUserHaveConsoleImport struct {
		User      *RequestUser  `json:"user" form:"user"`             //用户
		PageName  string        `json:"page_name" form:"page_name"`   //页面名称
		System    string        `json:"system" form:"system"`         //所属系统
		ImportKey string        `json:"import_key" form:"import_key"` //接口KEY
		Ctx       *base.Context `json:"-" form:"-"`
	}
	ResultConsoleHaveImportPermit struct {
		Result   bool   `json:"result"`
		ErrorMsg string `json:"error_msg"`
	}
)

//判断用户是否有接口权限
func GetUserHaveConsoleImportPermit(arg *ArgParamsUserHaveConsoleImport) (res *ResultConsoleHaveImportPermit, err error) {
	res = &ResultConsoleHaveImportPermit{Result: false, ErrorMsg: "接口异常"}
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
	res = data.Data
	return
}
