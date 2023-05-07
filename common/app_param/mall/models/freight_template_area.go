package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strings"
)

const (
	FreightTemplateTitleMaxLength = 580 //运费模板最大长度支持
)
const (
	FreightTemplateFreeFreightFree                = iota + 1 // 包邮
	FreightTemplateFreeFreightPay                            // 买家承担运费
	FreightTemplateFreeFreightFreeWithAsCondition            // 有条件包邮
)

const (
	FreightTemplateHasUseYes        = iota + 1 // 已使用
	FreightTemplateHasUseInit                  // 未使用 3
	FreightTemplateHasUseDeprecated            // 已弃用
)

const (
	FreightTemplateFreightCalWeight = iota + 1 // 按重量
	FreightTemplateFreightCalVolume            // 按容积
	FreightTemplateFreightCalNum               // 按数量
)

const (
	// FreightTemplateSendTypeExpressDelivery 运送方式 1-快递 2-EMS 3-平邮
	FreightTemplateSendTypeExpressDelivery = iota + 1 // 快递 快递中等
	FreightTemplateSendTypeEms                        // EMS 飞机空运
	FreightTemplateSendTypeGeneral                    // 平邮 平邮最慢
)

var MapFreightTemplateFreeFreight = map[uint8]string{
	FreightTemplateFreeFreightFree:                "包邮",
	FreightTemplateFreeFreightPay:                 "不包邮",
	FreightTemplateFreeFreightFreeWithAsCondition: "有条件包邮",
}

var MapFreightTemplateHasUse = map[uint8]string{
	FreightTemplateHasUseInit:       "未使用",
	FreightTemplateHasUseYes:        "已使用",
	FreightTemplateHasUseDeprecated: "已弃用",
}

var MapFreightTemplateFreightCal = map[uint8]string{
	FreightTemplateFreightCalWeight: "按重量",
	FreightTemplateFreightCalVolume: "按容积",
	FreightTemplateFreightCalNum:    "按数量",
}

var MapFreightTemplateSendType = map[uint8]string{
	FreightTemplateSendTypeExpressDelivery: "快递",
	FreightTemplateSendTypeEms:             "EMS",
	FreightTemplateSendTypeGeneral:         "平邮",
}

// FreightTemplateArea 运费模板设置
type FreightTemplateArea struct {
	ID                   int64            `gorm:"column:id;primary_key" json:"id"`
	TemplateId           int64            `gorm:"column:template_id;type:bigint(20);not null;default:0;comment:模板id"  json:"template_id"`
	ShopId               int64            `gorm:"column:shop_id;type:bigint(20);not null;default:0;comment:店铺id" json:"shop_id"`
	ProvinceId           string           `gorm:"column:province_id;type:varchar(30);not null;default:'';comment:收货省" json:"province_id"`
	CityId               string           `gorm:"column:city_id;type:varchar(30);not null;default:'';comment:收货市" json:"city_id"`
	AreaId               string           `gorm:"column:area_id;type:varchar(30);not null;default:'';comment:收货区或县" json:"area_id"`
	FreeFreight          uint8            `gorm:"column:free_freight;type:tinyint(1);not null;default:1;comment:是否包邮 1-包邮 2-自定义运费 3-有条件包邮" json:"free_freight"`
	FreightCal           uint8            `gorm:"column:freight_cal;type:tinyint(1);not null;default:1;comment:计价方式 1-按重量计算 2-按件数计算" json:"freight_cal"`
	SendType             uint8            `gorm:"column:send_type;type:tinyint(1);not null;default:1;comment:运送方式 1-快递 2-EMS 3-平邮" json:"send_type"`
	HasUse               uint8            `gorm:"column:has_use;type:tinyint(1);not null;default:1;comment:模板是否已使用 1-已使用 2-未使用 3-已弃用(弃用模板主要用于历史数据查询关联使用)" json:"has_use"`
	PermitSaleArea       string           `gorm:"column:permit_sale_area;type:varchar(3000);not null;default:'';comment:允许发货区域json" json:"permit_sale_area"`
	FreeFreightCondition string           `gorm:"column:free_freight_condition;type:varchar(3000);not null;default:'';comment:指定包邮条件"  json:"free_freight_condition"`
	OtherArea            string           `gorm:"column:other_area;type:varchar(800);not null;default:'';comment:发货区县下级信息(额外信息)" json:"other_area"`
	FreightDesc          string           `gorm:"column:freight_desc;type:varchar(800);not null;default:'';comment:邮费描述" json:"freight_desc"`
	CreatedAt            base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt            *base.TimeNormal `gorm:"column:deleted_at;" json:"deleted_at"`
}

func (r *FreightTemplateArea) TableName() string {
	return fmt.Sprintf("%sfreight_template", TablePrefix)
}

func (r *FreightTemplateArea) ParseFreeFreight() (res string) {
	res, _ = MapFreightTemplateFreeFreight[r.FreeFreight]
	return
}

func (r *FreightTemplateArea) ParseFreeSendType() (res string) {
	res, _ = MapFreightTemplateSendType[r.SendType]
	return
}

func (r *FreightTemplateArea) ParseFreeFreightCal() (res string) {
	res, _ = MapFreightTemplateFreightCal[r.FreightCal]
	return
}

func (r *FreightTemplateArea) ParseFreeHasUse() (res string) {
	res, _ = MapFreightTemplateHasUse[r.HasUse]
	return
}

func (r *FreightTemplateArea) GetTableComment() (res string) {
	res = "运费模板"
	return
}

func (r *FreightTemplateArea) SetFreeFreightCondition(data *FreeFreightCondition) (err error) {
	var bt []byte
	bt, err = json.Marshal(data)
	r.FreeFreightCondition = string(bt)
	return
}

func (r *FreightTemplateArea) SetPermitSaleAreaStruct(data *PermitSaleAreaStruct, err error) {
	var bt []byte
	bt, err = json.Marshal(data)
	r.FreeFreightCondition = string(bt)
}

func (r *FreightTemplateArea) ParseFreeFreightCondition() (res *FreeFreightCondition, err error) {
	if r.FreeFreightCondition != "" {
		err = json.Unmarshal([]byte(r.FreeFreightCondition), res)
	}
	return
}

func (r *FreightTemplateArea) ParsePermitSaleAreaStruct() (res *PermitSaleAreaStruct, err error) {
	if r.PermitSaleArea != "" {
		err = json.Unmarshal([]byte(r.PermitSaleArea), res)
	}
	return
}

type (
	// FreeFreightCondition 包邮条件
	FreeFreightCondition struct {
		OrCondition  []FreeFreightCondition `json:"or_condition"`
		AndCondition []FreeFreightCondition `json:"and_condition"`
		Type         string                 `json:"type"`
		Value        string                 `json:"value"`
	}

	// PermitSaleAreaStruct 可配送区域
	PermitSaleAreaStruct struct {
		AreaCode  int64 `json:"area_code"` //城市区域
		Condition *FreeFreightCondition
	}
)

const (
	FreightConditionEqual = "eq"   //等于
	FreightConditionList  = "list" //大于
)

var (
	FreightConditionMap = map[string]string{
		FreightConditionEqual: "等于",
		FreightConditionList:  "在列表中",
	}
)

func (r *FreeFreightCondition) GetDesc() (res string) {

	var eqString = ""
	switch r.Type {
	case FreightConditionEqual: //等于
		eqString = fmt.Sprintf("等于%s", r.Value)
	case FreightConditionList: //在列表中
	}
	var dataString = []string{
		eqString,
		r.getOrAndDesc(),
	}
	res = strings.Join(dataString, " 且 ")
	return

}

// GetDesc 获取邮费的中文描述
func (r *FreeFreightCondition) getOrAndDesc() (res string) {
	orConditionString := ""
	andConditionString := ""
	if len(r.OrCondition) > 0 {
		conditions := make([]string, 0, len(r.OrCondition))
		for _, value := range r.OrCondition {
			conditions = append(conditions, fmt.Sprintf("(%s)", value.GetDesc()))
		}
		orConditionString = strings.Join(conditions, " 或 ")
		return
	}
	if len(r.AndCondition) > 0 {
		conditions := make([]string, 0, len(r.AndCondition))
		for _, value := range r.AndCondition {
			conditions = append(conditions, fmt.Sprintf("(%s)", value.GetDesc()))
		}
		andConditionString = strings.Join(conditions, " 且 ")
		return
	}

	if orConditionString != "" && andConditionString != "" {
		res = fmt.Sprintf("%s 且 %s", orConditionString, andConditionString)
		return
	}
	res = fmt.Sprintf("%s%s", orConditionString, andConditionString)
	return
}
