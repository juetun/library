package mall

import (
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	//商品活动
	ArgPlatActivityByIds struct {
		base.ArgGetByNumberIds
	}
	PlatIntranetActivity struct {
		Id           int64            `json:"id"`
		Title        string           `json:"title"`
		KeyVal       string           `json:"key_val"`
		BindDataType uint8            `json:"bind_data_type"`
		OtherAttr    string           `json:"other_attr"`
		CouponId     string           `json:"coupon_id"`
		WarmUpTime   *base.TimeNormal `json:"warm_up_time"`
		StartTime    *base.TimeNormal `json:"start_time"`
		OverTime     *base.TimeNormal `json:"over_time"`
	}
)

func (r *ArgPlatActivityByIds) Default(ctx *base.Context) (err error) {

	return
}
