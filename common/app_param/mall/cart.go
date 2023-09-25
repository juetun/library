package mall

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgPreInfoInit struct {
		UserHid int64            `json:"user_hid" form:"user_hid"`
		TimeNow base.TimeNormal  `json:"time_now" form:"time_now"`
		SKus    []PreInfoInitSku `json:"skus" form:"skus"`
	}
	PreInfoInitSku struct {
		FirstOrderId   string          `json:"first_order_id" form:"first_order_id"`
		FinalOrderId   string          `json:"final_order_id" form:"final_order_id"`
		ShopId         string          `json:"shop_id" form:"shop_id"`
		SpuId          string          `json:"product_id" form:"product_id"`
		SkuId          string          `json:"sku_id" form:"sku_id"`
		Num            int64           `json:"num" form:"num"`
		Price          string          `json:"price" form:"price"`
		FinalPrice     string          `json:"final_price" form:"final_price"`
		Gifts          string          `json:"gifts" form:"gifts"`
		Status         uint8           `json:"status" form:"status"`
		SaleOnlineTime base.TimeNormal `json:"sale_online_time" form:"sale_online_time"`
		SaleOverTime   base.TimeNormal `json:"sale_over_time" form:"sale_over_time"`
	}
)

func (r *ArgPreInfoInit) Default(ctx *base.Context) (err error) {
	return
}
