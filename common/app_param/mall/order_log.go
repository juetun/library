package mall

import "github.com/juetun/base-wrapper/lib/base"

//操作订单日志的用户类型
const (
	OrderOperateLogUser   = iota + 1 //用户
	OrderOperateLogShop              //店铺管理员
	OrderOperateLogPlat              //平台客服
	OrderOperateLogSystem            //系统
)

var (
	SliceOrderOperateLogUserType = base.ModelItemOptions{
		{
			Label: "用户",
			Value: OrderOperateLogUser,
		},
		{
			Label: "店铺管理员",
			Value: OrderOperateLogShop,
		},
		{
			Label: "平台客服",
			Value: OrderOperateLogPlat,
		},
		{
			Label: "系统",
			Value: OrderOperateLogSystem,
		},
	}
)
