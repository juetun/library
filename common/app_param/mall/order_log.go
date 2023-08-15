package mall

//操作订单日志的用户类型
const (
	OrderOperateLogUser = iota + 1 //用户
	OrderOperateLogShop            //店铺管理员
	OrderOperateLogPlat            //平台客服
)
