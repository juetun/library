package upload_operate

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UploadVideo struct {
		ext_up.UploadCommon
		ParseCodeStatus uint8  `json:"parse_code_status"` //转码状态
		Src             string `json:"src"`               //源站地址
		HD              string `json:"hd,omitempty"`      //高清
		SD              string `json:"sd,omitempty"`      //标清
		LD              string `json:"ld,omitempty"`      //普清
		DefaultType     string `json:"default_type"`      //hd,sd,ld,src
	}
	VideoHandler func(uploadVideo *UploadVideo)
)

func VideoContext(ctx *base.Context) VideoHandler {
	return func(uploadVideo *UploadVideo) {
		uploadVideo.Context = ctx
	}
}

func NewUploadVideo(options ...VideoHandler) (res *UploadVideo) {
	res = &UploadVideo{}
	for _, option := range options {
		option(res)
	}
	return
}

func (r *UploadVideo) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

// GetShowUrl 获取视频的播放地址
func (r *UploadVideo) GetShowUrl() (res ext_up.ShowData) {
	res = ext_up.ShowData{
		PlayAddress: map[string]string{},
	}
	return
}

func (r *UploadVideo) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}
