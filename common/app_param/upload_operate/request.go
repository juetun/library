package upload_operate

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

func GetDataByUserIdsFromUploadServer(ctx *base.Context, arg *ArgUploadGetInfo) (resData *ResultMapUploadInfo, err error) {
	resData = &ResultMapUploadInfo{}
	var value = url.Values{}
	var bodyByte []byte

	//判断参数是否为空
	if arg == nil || arg.IsNull() {
		return
	}

	if bodyByte, err = json.Marshal(arg); err != nil {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUpload,
		URI:         "/upload/get_upload_address",
		Header:      http.Header{},
		Value:       value,
		BodyJson:    bodyByte,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                  `json:"code"`
		Data *ResultMapUploadInfo `json:"data"`
		Msg  string               `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	resData = data.Data
	return
}