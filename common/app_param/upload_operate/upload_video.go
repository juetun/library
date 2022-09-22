package upload_operate

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UploadVideo struct {
		ext_up.UploadCommon
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
	res = fmt.Sprintf("%s%s%d", r.Channel, ext_up.UploadDivideString, r.ID)
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
