package models

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	ShopNoticeStatusUse         = iota + 1 //使用中
	ShopNoticeStatusInit                   //初始化数据
	ShopNoticeStatusIneffective            //失效
)

var (
	SliceShopNoticeStatus = base.ModelItemOptions{
		{
			Value: ShopNoticeStatusUse,
			Label: "使用中",
		},
		{
			Value: ShopNoticeStatusInit,
			Label: "初始化",
		},
		{
			Value: ShopNoticeStatusIneffective,
			Label: "失效",
		},
	}
)

type (
	ShopNotice struct {
		ID        int64            `gorm:"column:id;primary_key" json:"-"`
		ShopID    int64            `gorm:"column:shop_id;default:0;" json:"shop_id"`
		Image     string           `gorm:"column:image;type:varchar(255);not null;default:'';comment:公告图片" json:"image"`
		StartTime base.TimeNormal  `gorm:"column:start_time;not null;default:CURRENT_TIMESTAMP" json:"start_time"`
		OverTime  base.TimeNormal  `gorm:"column:over_time;not null;default:CURRENT_TIMESTAMP" json:"over_time"`
		Status    uint8            `gorm:"column:status;type:tinyint(2);not null;default:2;comment:当前使用状态" json:"status"`
		CreatedAt base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt *base.TimeNormal `gorm:"column:deleted_at;" json:"deleted_at"`
	}
)

func (r *ShopNotice) TableName() (res string) {
	return fmt.Sprintf("%sshop_notice", TablePrefix)
}

func (r *ShopNotice) GetTableComment() (res string) {
	res = "店铺公告表"
	return
}

func (r *ShopNotice) ParseStatus() (res string, err error) {
	var mapKey map[uint8]string
	if mapKey, err = SliceShopNoticeStatus.GetMapAsKeyUint8(); err != nil {
		return
	}
	if dt, ok := mapKey[r.Status]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知状态(%s)", r.Status)
	return
}
