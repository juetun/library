package upload_operate

import (
	"github.com/juetun/base-wrapper/lib/base"

	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UploadFile struct {
		ext_up.UploadCommon
		Src string `json:"src"`
	}
	FileHandler func(UploadFile *UploadFile)
)

func FileContext(ctx *base.Context) FileHandler {
	return func(uploadFile *UploadFile) {
		uploadFile.Context = ctx
	}
}

func NewUploadFile(options ...FileHandler) (res *UploadFile) {
	res = &UploadFile{}
	for _, option := range options {
		option(res)
	}
	return
}

func (r *UploadFile) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

func (r *UploadFile) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}

// GetShowUrl 获取音频的播放地址
func (r *UploadFile) GetShowUrl() (res ext_up.ShowData) {
	res = ext_up.ShowData{
		PlayAddress: map[string]string{},
	}
	return
}
