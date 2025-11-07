package queue

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/utils"
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
		ID         int64            `gorm:"column:id;primary_key" json:"id"`
		TopicId    int64            `json:"topic_id" gorm:"column:topic_id;Index:topicInfo,priority:1;not null;type:bigint(20);default:0;comment:topic id"`
		MessageId  string           `json:"message_id" gorm:"column:message_id;UniqueIndex:msg_id,priority:1;not null;type:varchar(100);default:'';comment:消息ID"`
		Status     uint8            `gorm:"column:status;not null;type: tinyint(2);Index:topicInfo,priority:2;default:1;comment:使用状态 1-使用中 2-已停用"  json:"status,omitempty"`
		Data       string           `gorm:"column:data;type:text;not null;comment:商品描述" json:"data"`
		SyncStatus uint8            `gorm:"column:sync_status;not null;type: tinyint(2);Index:topicInfo,priority:2;default:1;comment:使用状态 1-使用中 2-已停用"  json:"sync_status,omitempty"`
		SyncKey    string           `json:"sync_key" gorm:"column:sync_key;not null;type:varchar(100);default:'';comment:锁数据的消费Key"`
		SyncAt     base.TimeNormal  `gorm:"column:sync_at;not null;default:'2000-01-01 00:00:00';comment:同步时间" json:"sync_at" `
		CreatedAt  base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
		OnlineAt   int64            `gorm:"column:online_at;not null;type:bigint(20);default:0" json:"online_at,omitempty"`
		DeletedAt  *base.TimeNormal `gorm:"column:deleted_at" json:"-"`
	}
)

func (r *MessageQueueData) ParseStatus() (res string) {
	mapStatus, _ := SliceQueueDataStatus.GetMapAsKeyUint8()
	if tmp, ok := mapStatus[r.Status]; ok {
		res = tmp
		return
	}
	res = "未知状态"
	return
}

func (r *MessageQueueData) ParseSyncStatus() (res string) {
	mapStatus, _ := SliceMessageQueueDataSyncStatus.GetMapAsKeyUint8()
	if tmp, ok := mapStatus[r.SyncStatus]; ok {
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
	if r.SyncStatus == 0 {
		r.SyncStatus = MessageQueueDataSyncStatusNo
	}
	if r.SyncAt.IsZero() {
		t, _ := utils.DateParse(utils.DateNullStringDefault, utils.DateTimeGeneral)
		r.SyncAt = base.GetNowTimeNormal(t)
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
