package app_param

import "github.com/shopspring/decimal"

type (
	ArgOrderFromCartItem struct {
		ShopId        int64  `json:"shop_id" form:"shop_id"`               // 店铺ID
		SpuId         string `json:"spu_id" form:"spu_id"`                 // 商品ID
		SkuId         string `json:"sku_id" form:"sku_id"`                 // sku地址
		Num           int64  `json:"num" form:"num"`                       // 商品数量
		SaleType      uint8  `json:"sale_type" form:"sale_type"`           // 销售类型
		SkuPrice      string `json:"sku_price" form:"sku_price"`           // SPU项目本次要支付的单价(定金预售定金金额或尾款金额 sku_price)
		SkuSetPrice   string `json:"sku_set_price" form:"sku_set_price"`   // SPU项目本的单价
		FreightTplId  int64  `json:"freight_tpl_id" form:"freight_tpl_id"` // 运费模板
		SubOrderId    string `json:"sub_order_id" form:"sub_order_id"`
		FreightAmount string `json:"freight_amount" form:"freight_amount"` // 邮费
	}
)

func (r *ArgOrderFromCartItem) GetPrice() (res decimal.Decimal, err error) {
	if r.SkuPrice == "" {
		r.SkuPrice = "0.00"
	}
	if res, err = decimal.NewFromString(r.SkuPrice); err != nil {
		return
	}
	return
}

func (r *ArgOrderFromCartItem) GetTotalSkuPrice() (res decimal.Decimal, err error) {
	var price decimal.Decimal
	if price, err = r.GetPrice(); err != nil {
		return
	}
	res = price.Mul(decimal.NewFromInt(r.Num))
	return
}
