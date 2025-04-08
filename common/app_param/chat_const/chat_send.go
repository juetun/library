package chat_const

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	SendMsgChat struct {
		MsgId           string `json:"msg_id,omitempty" form:"msg_id"`                     // 消息ID
		ChatType        uint8  `json:"chat_type,omitempty" form:"chat_type"`               // 聊天类型 单聊 、群聊
		FromId          string `json:"from_id,omitempty" form:"from_id"`                   // 发送消息方
		FromType        uint8  `json:"from_type,omitempty" form:"from_type"`               //
		ToId            string `json:"to_id,omitempty" form:"to_id"`                       // 发送给的目标用户或群
		ToType          uint8  `json:"to_type,omitempty" form:"to_type"`                   //
		ManagerUserHid  string `json:"manager_user_hid,omitempty" form:"manager_user_hid"` // 管理员ID
		Content         string `json:"content,omitempty" form:"content"`                   // 消息内容
		IsCustomer      bool   `json:"is_customer,omitempty" form:"is_customer"`           // 是否是客服聊天
		CustomerType    uint8  `json:"customer_type,omitempty" form:"customer_type"`       // 客服主体(CustomerManagerDataTypeShop-店铺实体  CustomerManagerDataTypePlatform-系统客服)
		ToCustomerUser  string `json:"to_customer_user,omitempty" form:"to_customer_user"` // 聊天信息产生时的管理员用户ID
		CurrentUserRole string `json:"current_user_role" form:"current_user_role"`         //当前用户角色
		SendType        uint8  `json:"send_type,omitempty" form:"send_type"`               // 发送方式 1-全双工 2-发送方可见 3-接收方可见
		ReplayType      uint8  `json:"replay_type,omitempty" form:"replay_type"`           //客户端响应数据类型 ，提醒客户端刷新消息内容
	}
	ArgGetChatToken struct {
		ArgGetChatTokenOther
		CurrentUserRole string `json:"current_user_role,omitempty" form:"current_user_role"` //当前用户聊天中的角色 customer-客服;user-普通用户
		IsCustomer      bool   `json:"is_customer,omitempty" form:"is_customer"`             //当前数据是否是与客服聊天
	}
	ArgChatSend struct {
		app_param.RequestUser
		Data []*struct {
			ArgGetChatToken
			SendMsgChat    SendMsgChat `json:"send_msg_chat" form:"send_msg_chat"`
			MsgType        int         `json:"msg_type" form:"msg_type"`
			CurrentUserStr string      `json:"cus" form:"cus"`
		}
	}
)

func (r *ArgGetChatToken) needAttendArea() (res bool) {
	area := []uint8{ConstChatTokenNeedAttendNo, ConstChatTokenNeedAttendUser, ConstChatTokenNeedAttendEachOther}
	for _, i2 := range area {
		if i2 == r.NeedAttend {
			res = true
			return
		}
	}
	return
}

func (r *ArgChatSend) Default(ctx *base.Context) (err error) {
	for _, item := range r.Data {
		if err = item.ArgGetChatToken.Default(ctx); err != nil {
			return
		}
		if item.MsgType == 0 {
			item.MsgType = ChatMsgTypeText
		}
		if item.SendMsgChat.ChatType == 0 {
			item.SendMsgChat.ChatType = ChatMsgChatTypeSingle
		}
		if err = item.SendMsgChat.Default(); err != nil {
			return
		}
		item.SendMsgChat.InitMsgId(item.MsgType)
		var (
			fromEntityInfo *FromEntityInfo
		)
		if fromEntityInfo, err = GetFromEntityInfo(item.CurrentUserRole, &r.RequestUser); err != nil {
			return
		}
		item.FromType = fromEntityInfo.FromType
		item.FromId = fromEntityInfo.FromId
		item.CurrentUserStr = fromEntityInfo.CurrentUserStr
	}

	return
}
