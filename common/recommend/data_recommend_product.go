package recommend

import "github.com/juetun/base-wrapper/lib/base"

type DataRecommendProduct struct {
	ProductID string          `json:"id" gorm:"column:id;"`
	UpdatedAt base.TimeNormal `gorm:"column:updated_at;" json:"updated_at"`
}
