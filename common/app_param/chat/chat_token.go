package chat

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgGetChatTokenOther struct {
		ChatType   uint8  `json:"chat_type" form:"chat_type"` // 聊天类型 私聊、群聊
		FromId     string `json:"from_id" form:"from_id"`
		ToId       string `json:"to_id" form:"to_id"`             // 聊天用户
		ToType     uint8  `json:"to_type" form:"to_type"`         // 聊天对象 1-用户 2-店铺 3-聊天室房间
		NeedAttend int    `json:"need_attend" form:"need_attend"` // 0:不关注 1:关注to_user_hid用户 2:彼此关注
	}
)

func (r *ArgGetChatTokenOther) Default(c *base.Context) (err error) {

	return
}
