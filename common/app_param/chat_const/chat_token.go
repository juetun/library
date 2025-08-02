package chat_const

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/audit_data"
)

const (
	ChatConsoleUserHIdDefault = -999 //官方客服默认用户ID
)
const (
	ConstChatTokenToUser uint8 = iota + 1 //用户聊天
	ConstChatTokenToShop                  //用户-店铺
	ConstChatTokenToRoom                  //用户-房间
)
const (
	ConstChatTokenNeedAttendNo        uint8 = iota //不关注
	ConstChatTokenNeedAttendUser                   //关注to_user_hid用户
	ConstChatTokenNeedAttendEachOther              //彼此关注
)

const (
	WebsocketDataTypeBase   uint8 = iota + 1 //基础数据类型，数据连接
	WebsocketDataTypeChat                    //聊天类型
	WebsocketDataTypeCommon                  //通用类型
)

const (
	//消息烈性
	ChatMsgTypeText  = (iota + 11) * 10000 // 聊天数据类型 110000-文本
	ChatMsgTypeImg                         // 聊天数据类型 110001-图片
	ChatMsgTypeCard                        // 聊天数据卡片 110003-数据卡片 (如商品数据)
	ChatMsgTypeVideo                       // 聊天数据类型 视频
	ChatMsgTypeMusic                       // 聊天数据类型 音频
)
const (
	ChatMsgChatTypeSingle uint8 = iota + 1 // 单聊
	ChatMsgChatTypeRoom                    // 群聊
)
const (
	CurrentUserRoleConsole      = "console"      //官方客服
	CurrentUserRoleCustomerHelp = "console_help" //官方客服托管
	CurrentUserRoleCustomer     = "customer"     //客服
	CurrentUserRoleUser         = "user"         //用户
)

const (
	ChatMsgSendTypeEachOther uint8 = iota + 1 // 全双工 彼此发送
	ChatMsgSendTypeFrom                       // 发送方可见
	ChatMsgSendTypeTo                         // 接收方可见
)

const (
	ShowLocLeft  uint8 = iota + 1 //左侧显示
	ShowLocRight                  //右侧显示
)

const (
	WebsocketSuccessMessage   = iota + 50001
	WebsocketSuccess          //建立连接成功
	WebsocketSuccessParameter //建立连接成功,发送参数参数不合法
	WebsocketEnd              //断开连接
	WebsocketOnlineReply      //在线回复
	WebsocketOfflineReply     //离线回复
	WebsocketLimit            //连接数超过限制
	WebsocketChatList         //聊天信息列表
	WebsocketChatDelete       //删除聊天信息
)

const (
	ClientReplayTypeLoadMore uint8 = iota + 1 //通知客户端刷新聊天内容
	ClientUpdateMsg                           //更新消息
	ClientRemoveMsg                           //删除消息
)
const (
	CustomerManagerDataTypeShop     uint8 = iota + 1 //店铺实体
	CustomerManagerDataTypePlatform                  //系统客服
)

var (
	SliceShopEntryType = base.ModelItemOptions{
		{
			Label: "店铺实体",
			Value: CustomerManagerDataTypeShop,
		},
		{
			Label: "系统客服",
			Value: CustomerManagerDataTypePlatform,
		},
	}
	SliceClientReplayType = base.ModelItemOptions{
		{
			Value: ClientReplayTypeLoadMore,
			Label: "加载列表中更多的数据",
		},
		{
			Value: ClientUpdateMsg,
			Label: "更新消息",
		},
		{
			Value: ClientRemoveMsg,
			Label: "移除消息",
		},
	}
	SliceWebsocketStatus = base.ModelItemOptions{
		{
			Value: WebsocketSuccessMessage,
			Label: "response ok",
		},
		{
			Value: WebsocketSuccess,
			Label: "connected ok",
		},
		{
			Value: WebsocketSuccessParameter,
			Label: "parameter is Invalid",
		},
		{
			Value: WebsocketEnd,
			Label: "connect close",
		},
		{
			Value: WebsocketOnlineReply,
			Label: "connect close",
		},
		{
			Value: WebsocketOfflineReply,
			Label: "connect close",
		},
		{
			Value: WebsocketLimit,
			Label: "connect expand limit", //连接数超过限制
		},
	}

	SliceShowLoc = base.ModelItemOptions{
		{
			Label: "",
			Value: ShowLocLeft,
		},
		{
			Label: "",
			Value: ShowLocRight,
		},
	}
	SliceSendMsgDataType = base.ModelItemOptions{
		{
			Value: WebsocketDataTypeChat,
			Label: "聊天",
		},
		{
			Value: WebsocketDataTypeBase,
			Label: "连接信息",
		},
		{
			Value: WebsocketDataTypeCommon,
			Label: "通用信息",
		},
	}

	SliceCurrentUserRole = base.ModelItemOptions{
		{
			Label: "客服",
			Value: CurrentUserRoleCustomer,
		},
		{
			Label: "官方客服",
			Value: CurrentUserRoleConsole,
		},
		{
			Label: "托管到官方客服",
			Value: CurrentUserRoleCustomerHelp,
		},
		{
			Label: "用户",
			Value: CurrentUserRoleUser,
		},
	}

	SliceUserChatType = base.ModelItemOptions{
		{
			Value: ChatMsgChatTypeSingle,
			Label: "单聊",
		},
		{
			Value: ChatMsgChatTypeRoom,
			Label: "群聊",
		},
	}
	SliceChatMsgType = base.ModelItemOptions{
		{
			Label: "文本",
			Value: ChatMsgTypeText,
		},
		{
			Label: "图片",
			Value: ChatMsgTypeImg,
		},
		{
			Label: "数据卡片",
			Value: ChatMsgTypeCard,
		},
		{
			Label: "视频",
			Value: ChatMsgTypeVideo,
		},
		{
			Label: "音频",
			Value: ChatMsgTypeMusic,
		},
	}

	SliceChatTokenTo = base.ModelItemOptions{
		{Label: "用户", Value: ConstChatTokenToUser},
		{Label: "店铺", Value: ConstChatTokenToShop},
		{Label: "房间", Value: ConstChatTokenToRoom},
	}

	SliceConstChatTokenNeedAttend = base.ModelItemOptions{
		{Label: "不关注", Value: ConstChatTokenNeedAttendNo},
		{Label: "关注用户", Value: ConstChatTokenNeedAttendUser},
		{Label: "彼此关注", Value: ConstChatTokenNeedAttendEachOther},
	}
	SliceChatMsgSendType = base.ModelItemOptions{
		{
			Label: "全双工",
			Value: ChatMsgSendTypeEachOther,
		},
		{
			Label: "发送方可见",
			Value: ChatMsgSendTypeFrom,
		},
		{
			Label: "接收方可见",
			Value: ChatMsgSendTypeTo,
		},
	}
)

type (
	ArgGetChatTokenOther struct {
		App        string          `json:"app" form:"app"`
		ChatType   uint8           `json:"chat_type" form:"chat_type"` // 聊天类型 私聊、群聊
		FromId     string          `json:"from_id" form:"from_id"`
		FromType   uint8           `json:"from_type" form:"from_type"`
		ToId       string          `json:"to_id" form:"to_id"`             // 聊天用户
		ToType     uint8           `json:"to_type" form:"to_type"`         // 聊天对象 1-用户 2-店铺 3-聊天室房间
		NeedAttend uint8           `json:"need_attend" form:"need_attend"` // 0:不关注 1:关注to_user_hid用户 2:彼此关注
		TimeNow    base.TimeNormal `json:"time_now" form:"time_now"`
	}

	//审核状态
	ApplyResult struct {
		Status        uint8  `json:"status"`         //审核状态
		Message       string `json:"msg"`            //审核结果
		ErrorType     string `json:"e_type"`         //审核失败类型
		ApplyId       string `json:"apply_id"`       //审核请求ID
		ApplyType     uint8  `json:"apply_type"`     //审核类型
		ApplyResponse string `json:"apply_response"` //审核请求的响应
	}
)

func (r *ApplyResult) ParseApplyType() (typeName string) {
	typeName = ParseApplyType(r.ApplyType)
	return
}

func ParseApplyType(applyType uint8) (typeName string) {
	mapApplyType, _ := audit_data.SliceDataChatApplyToolType.GetMapAsKeyUint8()
	var ok bool
	if typeName, ok = mapApplyType[applyType]; !ok {
		typeName = fmt.Sprintf("未知审核类型(%v)", applyType)
		return
	}
	return
}

func (r *ArgGetChatTokenOther) GetIsCustomer() (isCustomer bool) {
	switch r.FromType {
	case ConstChatTokenToShop: //店铺
		isCustomer = true
		return
	}
	switch r.ToType {
	case ConstChatTokenToShop: //店铺
		isCustomer = true
		return
	}

	return
}

func (r *ArgGetChatTokenOther) Default(c *base.Context) (err error) {

	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}

func (r *ArgGetChatTokenOther) ToString() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}
