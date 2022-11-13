package upload_operate

import (
	"github.com/juetun/base-wrapper/lib/base"

	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UploadMusic struct {
		ext_up.UploadCommon
		Src string `json:"src"`
	}
	MusicHandler func(uploadMusic *UploadMusic)
)

func MusicContext(ctx *base.Context) MusicHandler {
	return func(uploadMusic *UploadMusic) {
		uploadMusic.Context = ctx
	}
}

func NewUploadMusic(options ...MusicHandler) (res *UploadMusic) {
	res = &UploadMusic{}
	for _, option := range options {
		option(res)
	}
	return
}

func (r *UploadMusic) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

func (r *UploadMusic) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}

// GetShowUrl 获取音频的播放地址
func (r *UploadMusic) GetShowUrl() (res ext_up.ShowData) {
	res = ext_up.ShowData{
		PlayAddress: map[string]string{},
	}
	return
}
