package operate_admin

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

type (
	ArgAddAdminLog struct {
		Data []*AddAdminLog `json:"data"`
	}
	AddAdminLog struct {
		UserHid     int64           `json:"user_hid"`
		DataType    string          `json:"data_type"` //数据类型
		DataId      string          `json:"data_id"`   //数据ID
		Module      string          `json:"module"`
		Description string          `json:"description"`
		CreatedAt   base.TimeNormal `json:"created_at"`
	}
	ResultAddAdminLog struct {
		Result bool `json:"result"`
	}
)

//客服后台写入日志
func AddAdminLogAct(ctx *base.Context, arg *ArgAddAdminLog) (err error) {
	var bodyByte []byte
	if bodyByte, err = json.Marshal(arg); err != nil {
		return
	}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameAdmin,
		URI:         "/add_log",
		Value:       url.Values{},
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
		BodyJson:    bodyByte,
	}

	req := rpc.NewHttpRpc(&params).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.Body; len(body) == 0 {
		return
	}
	var resResult struct {
		Code int               `json:"code"`
		Data ResultAddAdminLog `json:"data"`
		Msg  string            `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}
	return
}

func (r *ArgAddAdminLog) Default(c *base.Context) (err error) {

	return
}
