package app_param

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgAddQueueMessage struct {
		TimeNow base.TimeNormal     `json:"time_now" form:"time_now"`
		Data    []*MessageQueueData `json:"data" form:"data"`
	}
	MessageQueueData struct {
		Topic     string          `json:"topic" form:"topic"`
		MessageId string          `json:"message_id" form:"message_id"`
		Data      string          `json:"data" form:"data"`
		OnlineAt  base.TimeNormal `json:"online_at" form:"online_at"`
	}
	ResultAddQueueMessage struct {
		Result bool `json:"result"`
	}
)

func (r *ArgAddQueueMessage) Default(ctx *base.Context) (err error) {
	return
}
