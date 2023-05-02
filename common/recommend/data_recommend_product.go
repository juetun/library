package recommend

import "github.com/juetun/base-wrapper/lib/base"

type DataRecommendProduct struct {
	ProductID  string          `json:"id" gorm:"column:id;"`
	ValidStart base.TimeNormal `json:"valid_start" gorm:"-"` //有效期开始使劲啊
	ValidOver  base.TimeNormal `json:"valid_over" gorm:"-"`  //有效期结束时间
	Status     uint8           `json:"status" gorm:"-"`      //状态  recommend.AdDataStatusCanUse //广告可用  recommend.AdDataStatusOffLine //不可用
	UpdatedAt  base.TimeNormal `gorm:"column:updated_at;" json:"updated_at"`
}
