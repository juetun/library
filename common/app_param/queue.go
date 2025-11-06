package app_param

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
)

type (
	ArgMessageConsume    []*MessageQueueDataParam
	ResultMessageConsume struct {
		Result bool `json:"result"`
	}
	ArgAddQueueMessage struct {
		TimeNow base.TimeNormal          `json:"time_now" form:"time_now"` //发送时间
		Data    []*MessageQueueDataParam `json:"data" form:"data"`
	}
	MessageQueueDataParam struct {
		Topic     string           `json:"topic" form:"topic"`                   //消息的主题
		MessageId string           `json:"message_id" form:"message_id"`         //消息ID
		Data      string           `json:"data" form:"data"`                     //消息数据
		IsDelay   bool             `json:"is_delay"form:"is_delay"`              //是否延迟发送
		OnlineAt  *base.TimeNormal `json:"online_at,omitempty" form:"online_at"` //发送消息的时刻节点（此时刻后发送）
	}
	ResultAddQueueMessage struct {
		Result bool `json:"result"`
	}

	//消费者信息
	MessageConsumer interface {
		Action(data ArgMessageConsume) (res *ResultMessageConsume)
	}

	//标记队列消费成功使用
	ArgAddQueueFlagOk struct {
		TopicId   int64    `json:"topic_id" form:"topic_id"`
		ConsumeId int64    `json:"consume_id" form:"consume_id"`
		MessageId []string `json:"message_id" form:"message_id"`
	}
	ResultAddQueueFlagOk struct {
		Result bool `json:"result"`
	}
	QueueDataInfo struct {
		Data       string          `json:"data"`
		MessageId  string          `json:"message_id"`
		TopicId    int64           `json:"topic_id"`
		ConsumeId  int64           `json:"consume_id"`
		ConsumeKey string          `json:"consume_key"`
		OnlineAt   int64           `json:"online_at"`
		CreatedAt  base.TimeNormal `json:"created_at"`
	}
)

func (r *QueueDataInfo) ToString() (res string) {
	var (
		bt []byte
		e  error
	)
	if bt, e = json.Marshal(r); e != nil {
		return
	}
	res = string(bt)
	return
}

func (r *ArgAddQueueFlagOk) Default(ctx *base.Context) (err error) {
	return
}

func (r *ArgAddQueueMessage) Default(ctx *base.Context) (err error) {
	return
}

//添加消息到队列中
func QueueAddMessage(arg *ArgAddQueueMessage, ctx *base.Context) (res *ResultAddQueueMessage, err error) {
	res = &ResultAddQueueMessage{}
	if arg == nil || len(arg.Data) == 0 {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     AppNameNotice,
		URI:         "/queue/add_message",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	ro.BodyJson, _ = json.Marshal(arg)
	var data = struct {
		Code int                    `json:"code"`
		Data *ResultAddQueueMessage `json:"data"`
		Msg  string                 `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	if data.Data != nil {
		res = data.Data
	}

	return
}

//标记队列消费成功
func QueueFlagConsumeOk(arg *ArgAddQueueFlagOk, ctx *base.Context) (res *ResultAddQueueFlagOk, err error) {
	res = &ResultAddQueueFlagOk{}
	if arg == nil || arg.ConsumeId == 0 || len(arg.MessageId) == 0 {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     AppNameNotice,
		URI:         "/queue/add_message",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	ro.BodyJson, _ = json.Marshal(arg)
	var data = struct {
		Code int                   `json:"code"`
		Data *ResultAddQueueFlagOk `json:"data"`
		Msg  string                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	if data.Data != nil {
		res = data.Data
	}

	return
}
