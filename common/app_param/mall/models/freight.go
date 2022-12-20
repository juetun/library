package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	FreightTemplateFree uint8 = iota + 1
	FreightTemplateUserPay
)
const (
	FreightTemplatePricingModeUnit   uint8 = iota + 1 //按件数
	FreightTemplatePricingModeWeight                  //按重量
)

var (
	//注意:此数据只能在后边添加,否则会影响数据结构
	SliceFreightTemplateFree = base.ModelItemOptions{
		{
			Value: FreightTemplateFree,
			Label: "包邮",
		}, {
			Value: FreightTemplateUserPay,
			Label: "买家承担运费",
		},
	}
	SliceFreightTemplatePricingMode = base.ModelItemOptions{
		{
			Value: FreightTemplatePricingModeUnit,
			Label: "按件数",
		}, {
			Value: FreightTemplatePricingModeWeight,
			Label: "按重量",
		},
	}
)

type (
	FreightTemplate struct {
		ID          int64            `gorm:"column:id;primary_key" json:"id"`
		ShopId      int64            `gorm:"column:shop_id;type:bigint(20);not null;default:0;comment:店铺id" json:"shop_id"`
		Title       string           `gorm:"column:title;type:varchar(200);not null;default:'';comment:模板名称" json:"title"`
		ProvinceId  int64            `gorm:"column:province_id;type:bigint(10);not null;default:0;comment:发货省" json:"province_id"`
		CityId      int64            `gorm:"column:city_id;type:bigint(10);not null;default:0;comment:发货市" json:"city_id"`
		AreaId      int64            `gorm:"column:area_id;type:int(10);not null;default:0;comment:发货区或县" json:"area_id"`
		FreeFreight uint8            `gorm:"column:free_freight;type:tinyint(2);not null;default:1;comment:是否包邮1-包邮 2-买家承担运费 3-有条件包邮" json:"free_freight"`
		PricingMode uint8            `gorm:"column:pricing_mode;type:tinyint(2);not null;default:1;comment:计价方式1-按件数 2-按重量" json:"pricing_mode"`
		CreatedAt   base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt   base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt   *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
	FreightTemplatesCache []*FreightTemplate
)

func (r *FreightTemplatesCache) UnmarshalBinary(data []byte) (err error) {
	if len(data) == 0 {
		*r = []*FreightTemplate{}
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *FreightTemplatesCache) MarshalBinary() (data []byte, err error) {
	if len(*r) == 0 {
		data = []byte{}
		return
	}
	data, err = json.Marshal(r)
	return
}

func (r *FreightTemplate) TableName() string {
	return fmt.Sprintf("%sfreight_template", TablePrefix)
}
func (r *FreightTemplate) GetTableComment() (res string) {
	res = "运费模板"
	return
}

func (r *FreightTemplate) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *FreightTemplate) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

func (r *FreightTemplate) ParseFreeFreight() (res string) {
	var mapSlice map[uint8]string
	mapSlice, _ = SliceFreightTemplateFree.GetMapAsKeyUint8()
	if tmp, ok := mapSlice[r.FreeFreight]; ok {
		res = tmp
		return
	}
	res = fmt.Sprintf("未知类型(%d)", r.FreeFreight)
	return
}

func (r *FreightTemplate) ParsePricingMode() (res string) {
	var mapSlice map[uint8]string
	mapSlice, _ = SliceFreightTemplatePricingMode.GetMapAsKeyUint8()
	if tmp, ok := mapSlice[r.PricingMode]; ok {
		res = tmp
		return
	}
	res = fmt.Sprintf("未知类型(%d)", r.PricingMode)
	return
}
