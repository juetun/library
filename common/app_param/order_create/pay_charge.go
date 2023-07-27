package order_create

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/shopspring/decimal"
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

//根据支付类型和金额计算支付手续费
func GetByPayTypeAndAmount(payType uint8, amount string) (payCharge, totalAmount string, err error) {
	var rabat string
	if rabat, err = GetPayChargeRabat(payType); err != nil {
		return
	}
	payCharge, totalAmount, err = CalPayCharge(rabat, amount)
	return
}

//计算支付手续费
func GetPayChargeRabat(payType uint8) (rabat string, err error) {
	switch payType {
	case OrderPayTypeAliPay: //支付宝支付
		rabat = ConfigPay.Pay.AliPay.FlatRabat
	case OrderPayTypeWeiXin: // 微信支付
		rabat = ConfigPay.Pay.WeiXinPay.FlatRabat
	default:
		err = fmt.Errorf("当前只支持微信及支付宝支付")
		return
	}
	return
}

//计算支付平台手续费
// Param rebat 费率小数()
func CalPayCharge(rabat, amount string) (payCharge, totalAmount string, err error) {
	totalAmount = amount
	payCharge = "0.00"
	if rabat == "" || rabat == "0.00" {
		return
	}
	var rabatDecimal, amountDecimal decimal.Decimal
	if rabatDecimal, err = decimal.NewFromString(rabat); err != nil {
		return
	}
	var oneDecimal = decimal.NewFromInt(1)
	if rabatDecimal.LessThan(decimal.NewFromInt(0)) || rabatDecimal.GreaterThanOrEqual(oneDecimal) {
		err = fmt.Errorf("支付手续费费率必须为0-1之间")
		return
	}
	if amountDecimal, err = decimal.NewFromString(amount); err != nil {
		return
	}
	var hundred = decimal.NewFromInt(100)
	var totalDecimal = amountDecimal.Div(oneDecimal.Sub(rabatDecimal)).Mul(hundred).Ceil().Div(hundred)
	totalAmount = totalDecimal.StringFixed(2)
	payCharge = totalDecimal.Sub(amountDecimal).StringFixed(2)
	return
}
