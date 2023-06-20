package mall

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgGetCouponBindData struct {
		CouponId string   `json:"coupon_id" form:"coupon_id"`
		SpuIds   []string `json:"spu_ids" form:"spu_ids"`
		DataType string   `json:"data_type" form:"data_type"`
	}
	ResGetCouponBindData map[string]bool
)

func (r *ArgGetCouponBindData) Default(ctx *base.Context) (err error) {
	return
}
