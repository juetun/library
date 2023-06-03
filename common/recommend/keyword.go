package recommend

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
)

type (
	ArgKeywordSave struct {
		common.HeaderInfo
		UserHid int64           `json:"user_hid" form:"user_hid"`
		Keyword []string        `json:"keyword" form:"keyword"`
		TimeNow base.TimeNormal `json:"time_now" form:"time_now"`
	}
	ResultKeywordSave struct {
		Result bool `json:"result"`
	}
)

func (r *ArgKeywordSave) Default(ctx *base.Context) (err error) {
	var (
		l          = len(r.Keyword)
		mapKeyword = make(map[string]bool, l)
		keyWords   = make([]string, 0, l)
	)
	for _, item := range r.Keyword {
		if item == "" {
			continue
		}
		if _, ok := mapKeyword[item]; ok {
			return
		}
		mapKeyword[item] = true
		keyWords = append(keyWords, item)
	}
	r.Keyword = keyWords
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}