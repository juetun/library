package models

import (
	"encoding/json"
	"fmt"

	"github.com/juetun/base-wrapper/lib/base"
)

const (
	FreightNumMax = 50
)

//修改店铺信息支持的字段
const (
	ShopCanUpdateColumnContactName  = "contact_user"
	ShopCanUpdateColumnContactPhone = "contact_phone"
	ShopCanUpdateColumnContactDesc = "desc"
)

type ShopExt struct {
	ShopID       int64            `gorm:"column:shop_id;primary_key" json:"shop_id"`
	Deposit      string           `gorm:"column:deposit;type:decimal(15,2);not null;default:0;comment:店铺保证金" json:"deposit"`
	FreightNum   int              `gorm:"column:freight_num;type:tinyint(2);not null;default:0;comment:运费模板数" json:"freight_num"`
	ContactUser  string           `gorm:"column:contact_user;type:varchar(30);not null;default:'';comment:店铺联系人"  json:"contact_user"`
	ContactPhone string           `gorm:"column:contact_phone;type:varchar(30);not null;default:'';comment:店铺联系电话"  json:"contact_phone"`
	Desc         string           `gorm:"column:desc;type:varchar(500);not null;default:'';comment:店铺描述" json:"desc"`
	CreatedAt    base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *base.TimeNormal `gorm:"column:deleted_at;" json:"deleted_at"`
}

func (r *ShopExt) TableName() string {
	return fmt.Sprintf("%sshop_ext", TablePrefix)
}

func (r *ShopExt) GetTableComment() (res string) {
	res = "店铺表"
	return
}

func (r *ShopExt) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *ShopExt) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

func (r *ShopExt) GetHid() (res string) {
	res = fmt.Sprintf("%d", r.ShopID)
	return
}
