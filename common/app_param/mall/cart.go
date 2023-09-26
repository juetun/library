package mall

import "github.com/juetun/base-wrapper/lib/base"



const (
	CartPrSaleStatusInit      = iota + 1 //初始化
	CartPrSaleStatusPayFirst             //已付定金
	CartPrSaleStatusPayFinal             //已付尾款
	CartPrSaleStatusPayExpire            //支付尾款超时
)

var (
	//预售状态
	SlicePrSaleCartStatus = base.ModelItemOptions{
		{
			Label: "初始化",
			Value: CartPrSaleStatusInit,
		},
		{
			Label: "已付定金",
			Value: CartPrSaleStatusPayFirst,
		},
		{
			Label: "已付尾款",
			Value: CartPrSaleStatusPayFinal,
		},
		{
			Label: "尾款支付超时",
			Value: CartPrSaleStatusPayExpire,
		},
	}
)


type (
	ArgPreInfoInit struct {
		UserHid int64             `json:"user_hid" form:"user_hid"`
		TimeNow base.TimeNormal   `json:"time_now" form:"time_now"`
		SKus    []*PreInfoInitSku `json:"skus" form:"skus"`
	}
	PreInfoInitSku struct {
		SrcFirstOrderId    string          `json:"src_first_order_id" form:"src_first_order_id"`
		SrcFirstSubOrderId string          `json:"src_first_sub_order_id" form:"src_first_sub_order_id"`
		FinalOrderId       string          `json:"final_order_id" form:"final_order_id"`
		FinalSubOrderId    string          `json:"final_sub_order_id" form:"final_sub_order_id"`
		ShopId             string          `json:"shop_id" form:"shop_id"`
		SpuId              string          `json:"product_id" form:"product_id"`
		Gifts              string          `json:"gifts" form:"gifts"` //赠品信息
		SkuId              string          `json:"sku_id" form:"sku_id"`
		Num                int64           `json:"num" form:"num"`
		Price              string          `json:"price" form:"price"`
		SkuSetPrice        string          `json:"sku_set_price" form:"sku_set_price"`
		Status             uint8           `json:"status" form:"status"`
		SaleOnlineTime     base.TimeNormal `json:"sale_online_time" form:"sale_online_time"`
		SaleOverTime       base.TimeNormal `json:"sale_over_time" form:"sale_over_time"`
	}
	ResultPreInfoInit struct {
		Result bool `json:"result"`
	}
)

func (r *ArgPreInfoInit) Default(ctx *base.Context) (err error) {
	return
}
