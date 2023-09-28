package mall

import "github.com/juetun/base-wrapper/lib/base"

const (
	CartPrSaleStatusInit      = iota + 1 //初始化
	CartPrSaleStatusPayFirst             //已付定金
	CartPrSaleStatusPayFinal             //已付尾款
	CartPrSaleStatusPayExpire            //支付尾款超时
)

const (
	OrderActTypeFirst uint8 = iota + 1 //首款
	OrderActTypeFinal                  //尾款
)

var (
	SliceOrderActType = base.ModelItemOptions{
		{
			Label: "首付款",
			Value: OrderActTypeFirst,
		},
		{

			Label: "尾款",
			Value: OrderActTypeFinal,
		},
	}
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
	//移除购物车数据参数
	ArgRemoveCart struct {
		UserHid   int64                    `json:"user_hid" form:"user_hid"`
		UserToken string                   `json:"user_token" form:"user_token"`
		SkuItems  []*ArgRemoveCartDataItem `json:"sku_items" form:"sku_items"`
		TimeNow   base.TimeNormal          `json:"time_now" form:"time_now"`
	}
	ArgRemoveCartDataItem struct {
		ActType         uint8  `json:"act_type" form:"act_type"`
		FinalOrderId    string `json:"final_order_id" form:"final_order_id"`
		FinalSubOrderId string `json:"final_sub_order_id" form:"final_sub_order_id"`
		FirstOrderId    string `json:"first_order_id" form:"first_order_id"`
		FirstSubOrderId string `json:"first_sub_order_id" form:"first_sub_order_id"`
		SkuId           string `json:"sku_id" form:"sku_id"`
		SpuId           string `json:"spu_id" form:"spu_id"`
	}
	ResultRemoveCartDataItem struct {
		Result bool `json:"result"`
	}
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

func (r *ArgRemoveCart) Default(ctx *base.Context) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}
