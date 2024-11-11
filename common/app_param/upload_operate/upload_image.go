package upload_operate

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	ProductImage struct {
		UploadFile
		IsThumbnail bool `json:"is_thumb,omitempty"` // 是否是缩略图
		Deleted     bool `json:"deleted,omitempty"`  //是否已删除
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
func (r ProductImages) GetThumbnail() (res ProductImage) {
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
			res = item
			return
		}
		i++
	}
	if res.ID == 0 && first.ID != 0 {
		res = first
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

func (r *ProductImages) UnmarshalBinary(data []byte) (err error) {
	if data == nil {
		return
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *ProductImages) MarshalBinary() (data []byte, err error) {
	if r == nil {
		return
	}
	data, err = json.Marshal(r)
	return
}
