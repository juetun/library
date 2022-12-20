package models

import (
	"fmt"
)

type ProductDesc struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`
	ProductId   string  `gorm:"column:product_id;uniqueIndex:idx_product_hid,priority:1;type:bigint(20);default:0;not null;comment:商品ID号" json:"product_id"`
	Description string `gorm:"column:description;type:text;not null;comment:商品描述" json:"description"`
}

func (r *ProductDesc) TableName() string {
	return fmt.Sprintf("%sproduct_desc", TablePrefix)
}

func (r *ProductDesc) GetTableComment() (res string) {
	res = "商品详情描述"
	return
}
