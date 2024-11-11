package upload_operate

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"regexp"
	"strings"

	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

const (
	UploadFileIsImgYes uint8 = iota + 1 //是否为图片
	UploadFileIsImgNo                   //不是图片
)

var (
	SliceUploadFileIsImg = base.ModelItemOptions{
		{
			Value: UploadFileIsImgYes,
			Label: "图片",
		},
		{
			Value: UploadFileIsImgNo,
			Label: "非图片",
		},
	}
)

type (
	UploadFile struct {
		ext_up.UploadCommon
		IsImg uint8  `json:"is_img,omitempty"`
		Src   string `json:"src,omitempty"`
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

func (r *UploadFile) UnmarshalBinary(data []byte) (err error) {
	if data == nil {
		return
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *UploadFile) MarshalBinary() (data []byte, err error) {
	if r == nil {
		return
	}
	data, err = json.Marshal(r)
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
