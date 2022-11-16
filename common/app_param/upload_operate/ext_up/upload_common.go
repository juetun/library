package ext_up

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
	"strings"
)

const (
	UploadDivideString = "|"
)

type (
	UploadCommon struct {
		Context *base.Context `json:"-" form:"-"`
		Type    string        `json:"type" form:"type"`
		Channel string        `json:"channel" form:"channel"`
		ID      int64         `json:"id" form:"id"`
	}

	ShowData struct {
		DefaultKey  string      `json:"default_key"`
		PlayAddress PlayAddress `json:"address"`
	}
	PlayAddress map[string]string
)

func (r *UploadCommon) ToString() (res string) {
	res = fmt.Sprintf("%s%s%s%s%d", r.Type, UploadDivideString, r.Channel, UploadDivideString, r.ID)
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

func (r *UploadCommon) ParseString(saveUploadString string) (err error) {
	if saveUploadString == "" {
		return
	}
	defer func() {
		if err == nil {
			return
		}
		r.Context.Error(map[string]interface{}{
			"saveUploadString": saveUploadString,
			"err":              err.Error(),
		}, "UploadCommonParseString")
	}()

	tmp := strings.Split(saveUploadString, UploadDivideString)
	var sliceString =tmp[0:]
	switch len(tmp) {
	case 0:
		sliceString[0], sliceString[1], sliceString[2] = "", "", ""
	case 1:
		sliceString[1], sliceString[2] = "", ""
	case 2:
		sliceString[2] = ""
	}
	r.Type, r.Channel = sliceString[0], sliceString[1]
	if r.ID, err = strconv.ParseInt(sliceString[2], 10, 64); err != nil {
		err = fmt.Errorf("图片存储的数据格式不正确")
	}
	return
}

func (r *UploadCommon) ToJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}
