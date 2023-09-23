package app_param

import "github.com/juetun/base-wrapper/lib/base"

//订单详情表标记订单商品是否参加了活动的活动信息
const (
	PlatActivityTypeCoupon uint8 = iota + 1 //平台优惠券满减
)

var SlicePlatActivityType = base.ModelItemOptions{
	{
		Label: "平台优惠券满减或打折",
		Value: PlatActivityTypeCoupon,
	},
}
