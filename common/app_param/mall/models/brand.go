package models

import (
	"encoding/json"
	"fmt"

	"github.com/juetun/base-wrapper/lib/base"
)

const (
	BrandStatusInit    uint8 = iota + 1 // 品牌状态初始化
	BrandStatusOk                       // 品牌状态审核通过
	BrandStatusFailure                  // 品牌状态审核失败
)

var (
	SliceBrandStatus = base.ModelItemOptions{
		{
			Value: BrandStatusInit,
			Label: "初始化",
		},
		{
			Value: BrandStatusOk,
			Label: "审核通过",
		},
		{
			Value: BrandStatusFailure,
			Label: "审核失败",
		},
	}
)

type Brand struct {
	ID            int64            `gorm:"column:id;primary_key" json:"-"`
	Name          string           `gorm:"column:name;not null;type: varchar(80);default:'';comment:品牌名称" json:"name"`
	Logo          string           `gorm:"column:logo;type:varchar(255);not null;default:'';comment:品牌logo" json:"logo"`
	LogoUrl       string           `json:"logo_url" gorm:"-"`
	CreateUserHid int64            `gorm:"column:create_user_hid;default:0;index:idx_userHid,priority:1;type:bigint(20);not null;comment:发布人用户ID"`
	ShopId        int64            `gorm:"column:name;not null;type: bigint(20);default:0;comment:店铺ID" json:"shop_id" `
	Status        uint8            `gorm:"column:status;type:tinyint(2);not null;default:1;comment:店铺审核状态2-审核通过 1-待审核 3-审核失败" json:"status"`
	CreatedAt     base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *base.TimeNormal `gorm:"column:deleted_at;" json:"deleted_at"`
}

func (r *Brand) TableName() string {
	return fmt.Sprintf("%sbrand", TablePrefix)
}

func (r *Brand) GetTableComment() (res string) {
	res = "商品sku表"
	return
}

func (r *Brand) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *Brand) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

// ParseStatus 店铺状态
func (r *Brand) ParseStatus() (res string) {
	mapBrandStatus, _ := SliceBrandStatus.GetMapAsKeyUint8()
	if dt, ok := mapBrandStatus[r.Status]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知状态(%d)", r.Status)
	return
}

func (r *Brand) DefaultLogo() (res string) {
	if r.Logo == "" {
		r.Logo = DefaultImageShow
	}
	res = r.Logo
	return
}
