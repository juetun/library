package recommend

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/common/response"
	"github.com/juetun/library/common/app_param"
	"strings"
)

type (
	ArgPageSearch struct {
		response.PageQuery
		app_param.RequestUser
		common.HeaderInfo
		base.GetDataTypeCommon
		TypeString       string          `json:"type" form:"type"`
		Types            []string        `json:"-" form:"-"`
		TimeNow          base.TimeNormal `json:"time_now" form:"time_now"`
		HeaderInfoString string          `json:"header_info" form:"header_info"`
		KeyWord          string          `json:"key_word" form:"key_word"`
	}
	ResultPageSearch struct {
		response.Pager
	}
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

//重置type类型
func (r *ArgPageSearch) ResetTypes() {
	r.TypeString = strings.Join(r.Types, ",")
}

func (r *ArgPageSearch) Default(c *base.Context) (err error) {
	if r.KeyWord == "" {
		err = fmt.Errorf("请选择您要搜索数据的关键词")
		return
	}
	if err = r.InitHeaderInfo(c.GinContext); err != nil {
		return
	}

	_ = r.InitRequestUser(c)

	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}

	types := strings.Split(r.TypeString, ",")
	r.Types = make([]string, 0, len(types))
	for _, item := range types {
		if item == "" {
			continue
		}
		r.Types = append(r.Types, item)
	}

	r.HeaderInfoString = c.GinContext.Request.Header.Get(app_obj.HttpHeaderInfo)
	return
}

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
