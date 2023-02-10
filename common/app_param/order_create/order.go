package order_create

import "github.com/juetun/base-wrapper/lib/base"

const (
	OrderFromTypeCart   = "cart"   //购物车
	OrderFromTypeDirect = "detail" //直接购买
)

var (
	SliceOrderFromType = base.ModelItemOptions{
		{
			Label: "购物车",
			Value: OrderFromTypeCart,
		},
		{
			Label: "直接购买",
			Value: OrderFromTypeDirect,
		},
	}
)
