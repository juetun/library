package mall

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgGetCouponBindData struct {
		CouponId string   `json:"coupon_id" form:"coupon_id"`
		SpuIds   []string `json:"spu_ids" form:"spu_ids"`
	}
	ResGetCouponBindData map[string]bool
)

func (r *ArgGetCouponBindData) Default(ctx *base.Context) (err error) {
	return
}
