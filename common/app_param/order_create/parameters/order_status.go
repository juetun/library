package parameters

import (
	"github.com/juetun/base-wrapper/lib/base"
	"time"
)

// 订单状态 放防止实现代码系统 此处每个状态的值强制定义，请勿随意修改
const (
	OrderStatusInit         uint8 = 10 // 订单初始化
	OrderStatusPaying       uint8 = 20 // 付款中
	OrderStatusCancel       uint8 = 24 // 已取消
	OrderStatusPayFailure   uint8 = 25 // 付款失败
	OrderStatusPayExpire    uint8 = 26 // 下单超时
	OrderStatusPayingFinish uint8 = 30 // 付款完成(待发货)

	OrderStatusGoodWaiting          uint8 = 35 // 待发货
	OrderStatusGoodSending          uint8 = 40 // 已发货
	OrderStatusGoodSendFinished     uint8 = 50 // 已收货
	OrderStatusGoodSendAutoFinished uint8 = 51 // 自动收货

	OrderStatusHasComment     uint8 = 55 // 已评价
	OrderStatusHasCommentAuto uint8 = 56 // 自动评价

	OrderStatusReturnMoney             uint8 = 60 // 申请退款
	OrderStatusReturnMoneyApproving    uint8 = 65 // 退款待确认
	OrderStatusReturnMoneyClientSure   uint8 = 70 // 申请退款平台客服确认
	OrderStatusReturnMoneyShopSure     uint8 = 80 // 申请退款商家确认
	OrderStatusReturnMoneyGoodSending  uint8 = 90 // 申请退款货物退还发货中
	OrderStatusReturnMoneyGoodFinished uint8 = 95 // 退款货物签收
	OrderStatusReturnMoneyPaying       uint8 = 11 // 申请退款中
	OrderStatusReturnMoneyFinished     uint8 = 12 // 退款完成
	OrderStatusReturnMoneyCancel       uint8 = 13 // 退款取消

	OrderStatusFinished uint8 = 100 // 订单结束
	OrderStatusError    uint8 = 110 // 订单支付数据异常
)

const (
	OrderExpireComment = 30 * 24 * time.Hour //当前发货后30天用户未确认发货或未评论 自动标记订单完成
)

var (
	SliceStatusLabel = base.ModelItemOptions{
		{
			Value: OrderStatusInit,
			Label: "待支付",
		},
		{
			Value: OrderStatusPaying,
			Label: "付款中",
		},
		{
			Value: OrderStatusPayingFinish,
			Label: "付款完成",
		},
		{
			Value: OrderStatusCancel,
			Label: "已取消",
		},
		{
			Value: OrderStatusPayExpire,
			Label: "下单超时",
		},
		{
			Value: OrderStatusPayFailure,
			Label: "付款失败",
		},
		{
			Value: OrderStatusGoodWaiting,
			Label: "待发货",
		},
		{
			Value: OrderStatusGoodSending,
			Label: "已发货",
		},
		{
			Value: OrderStatusGoodSendAutoFinished,
			Label: "自动收货",
		},
		{
			Value: OrderStatusGoodSendFinished,
			Label: "待评价",
		},
		{
			Value: OrderStatusHasComment, //手动评价
			Label: "已评价",
		},
		{
			Value: OrderStatusHasCommentAuto, //自动评价
			Label: "已评价",
		},
		{
			Value: OrderStatusReturnMoney,
			Label: "申请退款",
		},
		{
			Value: OrderStatusReturnMoneyApproving,
			Label: "退款待确认",
		},
		{
			Value: OrderStatusReturnMoneyClientSure,
			Label: "平台已确认退款",
		},
		{
			Value: OrderStatusReturnMoneyShopSure,
			Label: "已确认退款单",
		},
		{
			Value: OrderStatusReturnMoneyGoodSending,
			Label: "退款发货中",
		},
		{
			Value: OrderStatusReturnMoneyPaying,
			Label: "申请退款中",
		},
		{
			Value: OrderStatusReturnMoneyCancel,
			Label: "取消退款",
		},
		{
			Value: OrderStatusReturnMoneyGoodFinished,
			Label: "退货已签收",
		},
		{
			Value: OrderStatusReturnMoneyFinished,
			Label: "退款完成",
		},
		{
			Value: OrderStatusFinished,
			Label: "订单完成",
		},
		{
			Value: OrderStatusError,
			Label: "支付异常",
		},
	}
)

// StatusRelation 订单各状态间关系描述，用于数据校验使用
var StatusRelation = map[uint8]StatusNext{
	OrderStatusInit: { // 订单初始化状态的下一个状态
		OrderStatusPaying,
		OrderStatusCancel,
	},
	OrderStatusPaying: { // 付款中
		OrderStatusPayingFinish,
		OrderStatusPayFailure,
		OrderStatusCancel,
	},
	OrderStatusPayFailure: { //付款失败
		OrderStatusPayingFinish, //手工改成付款成功
		OrderStatusPaying,       //手工改成付款中
	},
	OrderStatusPayingFinish: { // 付款完成
		OrderStatusGoodSending, // 发货
		OrderStatusReturnMoney, // 退款
	},
	OrderStatusReturnMoney: { // 申请退款
		OrderStatusReturnMoneyClientSure, // 申请退款平台客服确认
		OrderStatusReturnMoneyShopSure,   // 申请退款商家客服确认
	},
	OrderStatusReturnMoneyClientSure: { // 申请退款平台客服确认
		OrderStatusReturnMoneyPaying,
	},
	OrderStatusReturnMoneyShopSure: { // 申请退款商家确认
		OrderStatusReturnMoneyPaying,
	},
	OrderStatusReturnMoneyGoodSending: { // 申请退款发货中
		OrderStatusReturnMoneyGoodFinished,
	},
	OrderStatusReturnMoneyGoodFinished: { // 退款货物签收
		OrderStatusFinished,
	},
	OrderStatusReturnMoneyPaying: { // 申请退款中
		OrderStatusReturnMoneyFinished,
	},
	OrderStatusGoodWaiting: { //待发货
		OrderStatusGoodSending,
	},
	OrderStatusGoodSendFinished: { //已收货
		OrderStatusHasComment,     //手工评论
		OrderStatusHasCommentAuto, //自动评论
		OrderStatusFinished,
		OrderStatusReturnMoney,
	},
	OrderStatusGoodSending: { // 已发货
		OrderStatusGoodSendFinished,
		OrderStatusReturnMoney,
	},
}

// CanSetStatus判断当前的订单状态是否支持修改成下一个状态
func CanSetStatus(newStatus, nowStatus uint8) (canSetStatus bool) {
	if dt, ok := StatusRelation[nowStatus]; ok {
		for _, value := range dt {
			if value == newStatus {
				canSetStatus = true
				return
			}
		}
	}
	return
}

type StatusNext []uint8
