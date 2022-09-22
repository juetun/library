package ext_up

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
	"strings"
)

const (
	UploadDivideString = "|"
)

type (
	UploadCommon struct {
		Context *base.Context
		Channel string `json:"channel" form:"channel"`
		ID      int64  `json:"id" form:"id"`
	}

	ShowData struct {
		DefaultKey  string      `json:"default_key"`
		PlayAddress PlayAddress `json:"address"`
	}
	PlayAddress map[string]string
)

func (r *UploadCommon) ParseString(saveUploadString string) (err error) {
	tmp := strings.Split(saveUploadString, UploadDivideString)
	switch len(tmp) {
	case 1:
		tmp[1] = "0"
	}
	r.Channel = tmp[0]
	r.ID, err = strconv.ParseInt(tmp[1], 10, 64)
	return
}

func (r *UploadCommon) ToJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}
