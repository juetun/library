package order_create

import (
	"fmt"
	"github.com/shopspring/decimal"
)

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

