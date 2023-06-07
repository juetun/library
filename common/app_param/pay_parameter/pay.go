package pay_parameter

import "github.com/juetun/base-wrapper/lib/base"

const (
	OrderPayTypeAliPay uint8 = iota + 1 // 支付宝支付
	OrderPayTypeWeiXin                  // 微信支付
)

var (
	SliceOrderPayType = base.ModelItemOptions{
		{
			Label: "支付宝",
			Value: OrderPayTypeAliPay,
		},
		{
			Label: "微信支付",
			Value: OrderPayTypeWeiXin,
		},
	}

)
