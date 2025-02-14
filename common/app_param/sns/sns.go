package sns

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
)

type (
	ArgRingsId struct {
		common.HeaderInfo
		app_param.RequestUser
		base.ArgGetByNumberIds
	}
	ResultRingMap  map[string]ResultRingItem
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

func (r *ArgRingsId) Default(ctx *base.Context) (err error) {

	return
}
