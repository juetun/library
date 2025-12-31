package sns

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
)

const (
	EditOptionTypeRing = "ring"
	EditOptionTypeUser = "usr"
)

type (
	ArgRingsId struct {
		common.HeaderInfo
		app_param.RequestUser
		base.ArgGetByNumberIds
	}
	ResultRingMap  map[string]ResultRingItem
	EditOptionItem struct {
		ID    int64  `gorm:"column:id" json:"id"`
		Title string `gorm:"column:title" json:"title,omitempty"`
	}
	LoadEditorItem struct {
		Value  string `json:"value"`
		Label  string `json:"label"`
		Prefix string `json:"prefix"`
	}
	ResultRingItem struct {
		ID             int64  `json:"id"`
		Title          string `json:"title,omitempty"`
		Thumbnail      string `json:"thumbnail,omitempty"`
		ThumbnailURL   string `json:"thumbnail_url"`
		ShowLevel      uint8  `json:"show_level"`
		IsRecommend    uint8  `json:"is_recommend"` //是否推荐圈子
		Key            string `json:"key,omitempty"`
		PingYin        string `json:"ping_yin,omitempty"`
		PingYinFirst   string `json:"ping_yin_first,omitempty"`
		Description    string `json:"description"`
		SubDescription string `json:"sub_description"`
		Status         int8   `json:"status,omitempty"`
		UserHid        int64  `json:"user_hid"`
	}
)

func (r *LoadEditorItem) ParseFromEditOptionItem(option *EditOptionItem, optionType string) {
	r.Value = option.GetPk(optionType)
	r.Label = option.Title
	switch optionType {
	case EditOptionTypeRing:
		r.Prefix = "#"
	case EditOptionTypeUser:
		r.Prefix = "@"
	}
	return
}

func (r *EditOptionItem) GetPk(optionType string) (res string) {
	res = fmt.Sprintf("%v|%v", optionType, r.ID)
	return
}

func (r *ArgRingsId) Default(ctx *base.Context) (err error) {
	return
}
