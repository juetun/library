package upload_operate

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
	"path"
	"strings"
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
		Cover           string `json:"cover"`             //封面图
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

//默认DefaultType
func (r *UploadVideo) InitDefaultType() {
	if r.DefaultType != "" {
		return
	}

	if r.HD != "" {
		r.DefaultType = "hd"
		return
	}
	r.DefaultType = "src"
}

func (r *UploadVideo) getSrc() (src string, err error) {
	src = r.Src
	switch r.DefaultType {
	case "src", "":
		src = r.Src
	case "hd":
		src = r.HD
	case "sd":
		src = r.SD
	case "ld":
		src = r.LD
	default:
		err = fmt.Errorf("当前不支持你选择的商品转码类型(%s)", r.DefaultType)
	}
	return
}

func (r *UploadVideo) GetEditorHtml(keys ...string) (res string, err error) {
	var (
		key                  = r.UploadCommon.GetKey(keys...)
		src, source, extName string
	)
	if src, err = r.getSrc(); err != nil {
		return
	}
	if src != "" {
		extName = strings.ToLower(path.Ext(src))
	}
	if extName == "" {
		res = fmt.Sprintf(`<video poster="%s" controls="true" width="auto" height="auto">
暂不支持Video标签。</video>`, key)
	}
	switch extName {
	case "mp4":
		source = fmt.Sprintf(`<source src="%s" type="video/mp4"/>`, src)
		res = fmt.Sprintf(`<video poster="%s" controls="true" width="auto" height="auto">%s
您的浏览器不支持Video标签。</video>`, key, source)
	case "ogv":
		source = fmt.Sprintf(`<source src="%s" type="video/ogg" />`, src)
		res = fmt.Sprintf(`<video poster="%s" controls="true" width="auto" height="auto">%s
您的浏览器不支持Video标签。</video>`, key, source)
	case "webm":
		source = fmt.Sprintf(`<source src="%s" type="video/webM" />`, src)
		res = fmt.Sprintf(`<video poster="%s" controls="true" width="auto" height="auto">%s
您的浏览器不支持Video标签。</video>`, key, source)
	default:
		res = fmt.Sprintf(`<video poster="%s" controls="true" width="auto" height="auto">暂不支持该视频播放</video>`, key)
	}

	return
}

func (r *UploadVideo) ToString() (res string) {
	res = r.UploadCommon.ToString()
	return
}

// GetShowUrl 获取视频的播放地址
func (r *UploadVideo) GetShowUrl() (res ext_up.ShowData) {
	var (
		tmp string
		ok  bool
	)
	res = ext_up.ShowData{
		PlayAddress: make(map[string]string, 7),
	}
	res.PlayAddress["src"] = r.Src

	if r.HD != "" {
		res.PlayAddress["hd"] = r.HD
	}
	if r.SD != "" {
		res.PlayAddress["sd"] = r.SD
	}
	if r.LD != "" {
		res.PlayAddress["ld"] = r.LD
	}
	if r.Cover != "" {
		res.PlayAddress["cover"] = r.Cover
	}
	if r.DefaultType != "" {
		res.PlayAddress["default_type"] = r.DefaultType
		if tmp, ok = res.PlayAddress[r.DefaultType]; ok {
			res.PlayAddress["play_src"] = tmp
		}
		if tmp == "" {
			res.PlayAddress["play_src"] = res.PlayAddress["src"]
		}
	}
	return
}

func (r *UploadVideo) ParseString(saveUploadString string) (err error) {
	err = r.UploadCommon.ParseString(saveUploadString)
	return
}

func (r *UploadVideo) UnmarshalBinary(data []byte) (err error) {
	if data == nil {
		return
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *UploadVideo) MarshalBinary() (data []byte, err error) {
	if r == nil {
		return
	}
	data, err = json.Marshal(r)
	return
}
