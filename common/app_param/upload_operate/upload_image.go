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
	UploadImage struct {
		ext_up.UploadCommon
		Src string `json:"src"`
	}
	ProductImage struct {
		UploadImage
		IsThumbnail bool `json:"is_thumb"` // 是否是缩略图
	}
	ImageHandler func(uploadImage *UploadImage)
)

func ImageContext(ctx *base.Context) ImageHandler {
	return func(uploadImage *UploadImage) {
		uploadImage.Context = ctx
	}
}

// NewUploadImage
func NewUploadImage(options ...ImageHandler) (res *UploadImage) {
	res = &UploadImage{}
	for _, option := range options {
		option(res)
	}
	return
}

func (r *UploadImage) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

func (r *UploadImage) GetEditorHtml(keys ...string) (res string, err error) {
	var (
		key = r.UploadCommon.GetKey(keys...)
	)
	res = fmt.Sprintf(`<img src="%s" alt="%s" data-href="%s" style=""/>`, r.Src, key, key)
	return
}

// GetShowUrl 获取图片地址的播放地址
func (r *UploadImage) GetShowUrl() (res *app_param.ResultExcelImportHeaderRelate, err error) {
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
func (r *UploadImage) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}
