package pay_parameter

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/order_create"
)

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
func ParsePayType(payType uint8) (res string) {
	MapOrderPayType, _ := order_create.SliceOrderPayType.GetMapAsKeyUint8()
	if _, ok := MapOrderPayType[payType]; ok {
		res = MapOrderPayType[payType]
		return
	}
	return fmt.Sprintf("未知支付类型(%d)", payType)
}
