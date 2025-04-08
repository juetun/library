package chat_const

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/audit_data"
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
	ChatMsgTypeMusic                       //聊天数据类型 音频
)
const (
	ChatMsgChatTypeSingle uint8 = iota + 1 // 单聊
	ChatMsgChatTypeRoom                    // 群聊
)
const (
	CurrentUserRoleConsole  = "console"  //官方客服
	CurrentUserRoleCustomer = "customer" //客服
	CurrentUserRoleUser     = "user"     //用户
)

var (
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
