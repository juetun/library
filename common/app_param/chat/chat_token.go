package chat

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
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

var (
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
)

func (r *ArgGetChatTokenOther) GetIsCustomer() (isCustomer bool) {
	switch r.FromType {
	case ConstChatTokenToShop: //店铺
		isCustomer = true
		return
	}
	switch r.ToType {
	case ConstChatTokenToShop: //房间
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
