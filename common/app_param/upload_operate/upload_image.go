package upload_operate

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type (

	ProductImage struct {
		UploadFile
		IsThumbnail bool `json:"is_thumb"` // 是否是缩略图
		Deleted     bool `json:"deleted"`  //是否已删除
	}
	ImageHandler  func(uploadImage *UploadFile)
	ProductImages []ProductImage
)

func ImageContext(ctx *base.Context) ImageHandler {
	return func(uploadImage *UploadFile) {
		uploadImage.Context = ctx
	}
}

//获取封面图数据
func (r ProductImages) GetThumbnail() (res *ProductImage) {
	var (
		first ProductImage
		i     int
	)
	for _, item := range r {
		if item.Deleted {
			continue
		}
		if i == 0 {
			first = item
		}
		if item.IsThumbnail {
			res = &item
		}
		i++
	}
	if (res == nil || res.ID == 0) && first.ID != 0 {
		res = &first
	}
	return
}

//获取没有删除的图片
func (r *ProductImages) GetNotDeleteData() (res []ProductImage) {
	res = make([]ProductImage, 0, len(*r))
	for _, item := range *r {
		if item.Deleted {
			continue
		}
		res = append(res, item)
	}

	return
}

func (r *UploadImage) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

func (r *UploadImage) GetEditorHtml(value string) (res string, err error) {
	var (
		reg *regexp.Regexp
	)
	res = value
	res = strings.ReplaceAll(res, "%", "%%")
	if reg, err = regexp.Compile(`src="[^"]+"`); err != nil {
		return
	}

	repl := fmt.Sprintf(`src="%s"`, r.Src)
	res = reg.ReplaceAllStringFunc(value, func(s string) (res string) {
		res = repl
		return
	})
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
