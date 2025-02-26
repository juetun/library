package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	SkuGiftStatusEffect     uint8 = iota + 1 // 有效
	SkuGiftStatusEffectInit                  //初始化
	SkuGiftStatusEffectNo                    //失效
)

var (
	SliceSkuGiftStatus = base.ModelItemOptions{
		{
			Label: "有效",
			Value: SkuGiftStatusEffect,
		},
		{
			Label: "待生效",
			Value: SkuGiftStatusEffectInit,
		},
		{
			Label: "失效",
			Value: SkuGiftStatusEffectNo,
		},
	}
)

type (
	SkuGift struct {
		ID          int64            `gorm:"column:id;primary_key" json:"-"`
		MasterSkuId string           `gorm:"column:master_sku_id;not null;index:idx_shop_id,priority:2;type:bigint(20);default:0;comment:主SKU" json:"master_sku_id"`
		GiftSkuId   string           `gorm:"column:gift_sku_id;not null;type:bigint(20);default:0;comment:主SKU" json:"gift_sku_id"`
		EffectSTime base.TimeNormal  `gorm:"column:effect_s_time;not null;default:CURRENT_TIMESTAMP;comment:有效期开始时间" json:"effect_s_time"`
		EffectOTime base.TimeNormal  `gorm:"column:effect_o_time;not null;default:CURRENT_TIMESTAMP;comment:有效期结束时间" json:"effect_o_time"`
		ShopId      int64            `gorm:"column:shop_id;index:idx_shop_id,priority:1;default:0;type:bigint(20);not null;comment:店铺ID" json:"shop_id"`
		Status      uint8            `gorm:"column:status;not null;type:tinyint(1);index:idx_shop_id,priority:3;default:3;comment:赠品状态 1-有效 2-失效 3-初始化" json:"status" `
		Mark        string           `gorm:"column:mark;not null;type:varchar(255);not null;default:'';comment:备注" json:"mark"`
		CreatedAt   base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt   base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt   *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
	SkuGiftCache []*SkuGift
	SkuGiftModel struct {
		SkuGift
		GiftSkuImg   string `json:"gift_sku_img"`   //赠品的图片链接
		GiftSkuTitle string `json:"gift_sku_title"` //赠品的标题
	}
	SkuGiftsInfo struct {
		Src         string `json:"src"`           //赠品的封面图链接
		Tip         string `json:"tip"`           //赠品名称
		SkuId       string `json:"sku_id"`        //赠品的商品skuId
		SkuSrcPrice string `json:"sku_src_price"` //赠品的商品原价
	}
)

func (r *SkuGiftCache) UnmarshalBinary(data []byte) (err error) {
	if len(data) == 0 {
		*r = []*SkuGift{}
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *SkuGiftCache) MarshalBinary() (data []byte, err error) {
	if len(*r) == 0 {
		data = []byte{}
		return
	}
	data, err = json.Marshal(r)
	return
}

func (r *SkuGift) TableName() string {
	return fmt.Sprintf("%ssku_gift", TablePrefix)
}

func (r *SkuGift) GetTableComment() (res string) {
	res = "SKU的赠品信息"
	return
}

func (r *SkuGift) ParseFromSku(sku *Sku) (res *SkuGift) {
	res = &SkuGift{}

	return
}

func (r *SkuGift) ParseStatus() (res string) {
	var mapStatus, _ = SliceSkuGiftStatus.GetMapAsKeyUint8()
	if tp, ok := mapStatus[r.Status]; ok {
		res = tp
		return
	}
	res = fmt.Sprintf("未知赠品状态类型(%d)", r.Status)
	return
}
