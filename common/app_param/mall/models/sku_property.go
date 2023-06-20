package models

import "github.com/juetun/base-wrapper/lib/base"

type (
	SkuProperty struct {
		ID        int64            `gorm:"column:id;primary_key" json:"id"`
		Label     string           `gorm:"column:label;type:varchar(40);uniqueIndex:idx_cate_label,priority:3;not null;default:'';comment:属性名称" json:"label"`
		CreatedAt base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
)
