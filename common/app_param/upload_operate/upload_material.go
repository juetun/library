package upload_operate

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
	"net/http"
	"net/url"
)

type (
	UploadMaterial struct {
		ext_up.UploadCommon
		Src string `json:"src"`
	}

	MaterialHandler func(uploadImage *UploadMaterial)
)

func MaterialContext(ctx *base.Context) ImageHandler {
	return func(uploadImage *UploadImage) {
		uploadImage.Context = ctx
	}
}

// NewUploadMaterial
func NewUploadMaterial(options ...MaterialHandler) (res *UploadMaterial) {
	res = &UploadMaterial{}
	for _, option := range options {
		option(res)
	}
	return
}

func (r *UploadMaterial) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

// GetShowUrl 获取图片地址的播放地址
func (r *UploadMaterial) GetShowUrl() (res *app_param.ResultExcelImportHeaderRelate, err error) {
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     r.Context,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUpload,
		URI:         "/upload/get_img_address",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}

	if params.BodyJson, err = r.UploadCommon.ToJson(); err != nil {
		return
	}
	req := rpc.NewHttpRpc(&params).
		Send()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.GetBody().Body; len(body) == 0 {
		return
	}

	var resResult struct {
		Code int                                     `json:"code"`
		Data app_param.ResultExcelImportHeaderRelate `json:"data"`
		Msg  string                                  `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}

	res = &resResult.Data
	return
}
func (r *UploadMaterial) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}
