package upload_operate

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
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



func (r *UploadFile) GetEditorHtml(value string) (res string, err error) {
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

