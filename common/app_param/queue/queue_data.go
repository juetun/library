package queue

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	MessageQueueDataInit uint8 = iota + 1
	MessageQueueDataUsing
	MessageQueueDataUsed
)
const (
	MessageQueueDataSyncStatusNo uint8 = iota + 1
	MessageQueueDataSyncStatusSyncing
	MessageQueueDataSyncStatusOk
)

var (
	SliceMessageQueueDataSyncStatus = base.ModelItemOptions{
		{
			Label: "未同步",
			Value: MessageQueueDataSyncStatusNo,
		},
		{
			Label: "同步中",
			Value: MessageQueueDataSyncStatusSyncing,
		},
		{
			Label: "同步结束",
			Value: MessageQueueDataSyncStatusOk,
		},
	}
	SliceQueueDataStatus = base.ModelItemOptions{
		{
			Label: "待执行",
			Value: MessageQueueDataInit,
		},
		{
			Label: "使用中",
			Value: MessageQueueDataUsing,
		},
		{
			Label: "使用结束",
			Value: MessageQueueDataUsed,
		},
	}
)

type (
	MessageQueueData struct {
		ID        int64            `gorm:"column:id;primary_key" json:"id"`
		TopicId   int64            `json:"topic_id" gorm:"column:topic_id;Index:topicInfo,priority:1;not null;type:bigint(20);default:0;comment:topic id"`
		MessageId string           `json:"message_id" gorm:"column:message_id;UniqueIndex:msg_id,priority:1;not null;type:varchar(100);default:'';comment:消息ID"`
		Status    uint8            `gorm:"column:status;not null;type: tinyint(2);Index:topicInfo,priority:2;default:1;comment:使用状态 1-使用中 2-已停用"  json:"status,omitempty"`
		Data      string           `gorm:"column:data;type:text;not null;comment:商品描述" json:"data"`
		CreatedAt base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
		OnlineAt  int64            `gorm:"column:online_at;not null;type:bigint(20);default:0" json:"online_at,omitempty"`
		DeletedAt *base.TimeNormal `gorm:"column:deleted_at" json:"-"`
	}
	ConsumeQueueDataIndex struct {
		Topic     string `json:"tc"`
		TopicId   int64  `json:"tcid"`
		ConsumeId int64  `json:"cid"`
	}
)

func (r *ConsumeQueueDataIndex) GetIndex() (res string) {
	res = fmt.Sprintf("%v_%v", r.TopicId, r.ConsumeId)
	return
}

func (r *ConsumeQueueDataIndex) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

func (r *ConsumeQueueDataIndex) UnmarshalBinary(data []byte) (err error) {
	if r == nil {
		r = &ConsumeQueueDataIndex{}
	}
	err = json.Unmarshal(data, r)
	return
}

func (r *ConsumeQueueDataIndex) ToString() (res string) {
	if r == nil {
		return
	}
	var bt []byte
	bt, _ = json.Marshal(r)
	res = string(bt)
	return
}
func (r *MessageQueueData) ParseStatus() (res string) {
	mapStatus, _ := SliceQueueDataStatus.GetMapAsKeyUint8()
	if tmp, ok := mapStatus[r.Status]; ok {
		res = tmp
		return
	}
	res = "未知状态"
	return
}

func (r *MessageQueueData) GetTableComment() (res string) {
	res = "队列队列数据表"
	return
}

func (r *MessageQueueData) Default() (err error) {
	if r.Status == 0 {
		r.Status = MessageQueueDataInit
	}

	return
}

func (r *MessageQueueData) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

func (r *MessageQueueData) MarshalBinary() (data []byte, err error) {
	if r == nil {
		data = []byte("{}")
		return
	}
	data, err = json.Marshal(r)
	return
}
