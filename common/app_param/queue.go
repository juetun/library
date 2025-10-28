package app_param

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgAddQueueMessage struct {
		Data []*MessageQueueData
	}
	MessageQueueData struct {
		Topic     string          `json:"topic" form:"topic"`
		MessageId string          `json:"message_id" form:"message_id"`
		Data      string          `json:"data" form:"data"`
		CreatedAt base.TimeNormal `json:"created_at" form:"created_at"`
		OnlineAt  base.TimeNormal `json:"online_at" form:"online_at"`
	}
	ResultAddQueueMessage struct {
		Result bool `json:"result"`
	}
)

func (r *ArgAddQueueMessage) Default(ctx *base.Context) (err error) {
	return
}
