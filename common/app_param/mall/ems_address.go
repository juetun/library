package mall

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/freight"
	"strconv"
	"strings"
)

type (
	ArgGetByAddressIds struct {
		NotNeedEncrypt bool   `json:"not_need_encrypt" form:"not_need_encrypt"` //手机号是否需要加密 默认false-需要
		IdString       string `json:"ids" form:"ids"`
		base.GetDataTypeCommon
		Ids []int64 `json:"-" form:"-"`
	}

	ResultGetByAddressIds map[int64]*freight.ResultGetByAddressIdsItem
)

func (r *ArgGetByAddressIds) Default(ctx *base.Context) (err error) {
	idString := strings.Split(r.IdString, ",")
	r.Ids = make([]int64, 0, len(idString))
	var id int64
	for _, idStr := range idString {
		if idStr == "" {
			continue
		}
		if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
			err = fmt.Errorf("地址ID格式不正确")
			return
		}
		r.Ids = append(r.Ids, id)
	}

	return
}
