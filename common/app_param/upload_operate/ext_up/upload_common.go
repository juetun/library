package ext_up

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"net/url"
	"strconv"
	"strings"
)

const (
	UploadDivideString = "|"
	UploadDivideParams = "^#^"
)

type (
	UploadCommon struct {
		Context   *base.Context `json:"-" form:"-"`
		Type      string        `json:"type" form:"type"`
		Channel   string        `json:"channel" form:"channel"`
		ID        int64         `json:"id" form:"id"`
		UrlParams url.Values    `json:"url_params" form:"url_params"`
	}

	ShowData struct {
		DefaultKey  string      `json:"default_key"` //默认播放地址
		PlayAddress PlayAddress `json:"address"`     //播放地址列表 （高清、普清、源地址等）
	}
	PlayAddress map[string]string
)

func (r *UploadCommon) ToString() (res string) {
	res = r.GetFilePk()
	if r.UrlParams != nil && len(r.UrlParams) > 0 {
		res = base64.StdEncoding.EncodeToString([]byte(r.UrlParams.Encode())) + UploadDivideParams + res
	}
	return
}

func (r *UploadCommon) GetKey(keys ...string) (key string) {
	if len(keys) > 0 {
		key = keys[0]
	}
	if key == "" {
		key = r.ToString()
	}
	return
}

//获取文件的唯一KEY
func (r *UploadCommon) GetFilePk() (res string) {
	res = fmt.Sprintf("%s%s%s%s%d", r.Type, UploadDivideString, r.Channel, UploadDivideString, r.ID)
	return
}

func (r *UploadCommon) GetChannelPk() (res string) {
	res = fmt.Sprintf("%s_%s", r.Type, r.Channel)
	return
}

func (r *UploadCommon) parseCommon(saveUploadString string) (err error) {
	tmp := strings.Split(saveUploadString, UploadDivideString)
	var sliceString = make([]string, 0, 3)
	sliceString = append(sliceString, tmp[0:]...)
	switch len(sliceString) {
	case 0:
		sliceString = append(sliceString, []string{"", "", ""}...)
	case 1:
		sliceString = append(sliceString, []string{"", ""}...)
	case 2:
		sliceString = append(sliceString, "")
	}
	r.Type, r.Channel = sliceString[0], sliceString[1]
	if r.ID, err = strconv.ParseInt(sliceString[2], 10, 64); err != nil {
		err = fmt.Errorf("图片存储的数据格式不正确")
	}
	return
}

func (r *UploadCommon) ParseString(saveUploadString string) (err error) {
	if saveUploadString == "" {
		return
	}
	defer func() {
		if err == nil {
			return
		}
		if r.Context == nil {
			return
		}
		r.Context.Error(map[string]interface{}{
			"saveUploadString": saveUploadString,
			"err":              err.Error(),
		}, "UploadCommonParseString")
	}()
	paramString := strings.Split(saveUploadString, UploadDivideParams)
	switch len(paramString) {
	case 1:
		if err = r.parseCommon(paramString[0]); err != nil {
			return
		}
	case 2:
		if err = r.parseCommon(paramString[1]); err != nil {
			return
		}

		if paramString[0] != "" {
			if err = r.parseUrlValue(paramString[0]); err != nil {
				return
			}
		}

	}

	return
}

func (r *UploadCommon) parseUrlValue(urlArg string) (err error) {
	//encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	var decoded []byte
	if decoded, err = base64.StdEncoding.DecodeString(urlArg); err != nil {
		fmt.Println("decode error:", err)
		return
	}
	if r.UrlParams, err = url.ParseQuery(string(decoded)); err != nil {
		return
	}
	return
}

func (r *UploadCommon) ToJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}
