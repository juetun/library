package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/shopspring/decimal"
	"strings"
)

const (
	FreightTemplateFree uint8 = iota + 1
	FreightTemplateUserPay
)
const (
	FreightTemplatePricingModeUnit   uint8 = iota + 1 //按件数
	FreightTemplatePricingModeWeight                  //按重量
)
const (
	FreeConditionOptAnd uint8 = iota + 1 //且 逻辑 两个条件都满足
	FreeConditionOptOr                   //或 逻辑 两个条件满足一个即可
)

var (
	SliceFreeConditionOpt = base.ModelItemOptions{
		{
			Label: "所有条件",
			Value: FreeConditionOptAnd,
		},
		{
			Label: "任一条件",
			Value: FreeConditionOptOr,
		},
	}
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

var (
	MapFreightTemplateFreeFreight = map[uint8]string{
		FreightTemplateFreeFreightFree:                "包邮",
		FreightTemplateFreeFreightPay:                 "不包邮",
		FreightTemplateFreeFreightFreeWithAsCondition: "有条件包邮",
	}
	MapFreightTemplateHasUse = map[uint8]string{
		FreightTemplateHasUseInit:       "未使用",
		FreightTemplateHasUseYes:        "已使用",
		FreightTemplateHasUseDeprecated: "已弃用",
	}
	MapFreightTemplateFreightCal = map[uint8]string{
		FreightTemplateFreightCalWeight: "按重量",
		FreightTemplateFreightCalVolume: "按容积",
		FreightTemplateFreightCalNum:    "按数量",
	}
	MapFreightTemplateSendType = map[uint8]string{
		FreightTemplateSendTypeExpressDelivery: "快递",
		FreightTemplateSendTypeEms:             "EMS",
		FreightTemplateSendTypeGeneral:         "平邮",
	}
)

type (
	FreightTemplate struct {
		ID               int64            `gorm:"column:id;primary_key" json:"id"`
		ShopId           int64            `gorm:"column:shop_id;type:bigint(20);not null;default:0;comment:店铺id" json:"shop_id"`
		Title            string           `gorm:"column:title;type:varchar(200);not null;default:'';comment:模板名称" json:"title"`
		ProvinceId       int64            `gorm:"column:province_id;type:bigint(10);not null;default:0;comment:发货省" json:"province_id"`
		CityId           int64            `gorm:"column:city_id;type:bigint(10);not null;default:0;comment:发货市" json:"city_id"`
		AreaId           int64            `gorm:"column:area_id;type:int(10);not null;default:0;comment:发货区或县" json:"area_id"`
		FreeFreight      uint8            `gorm:"column:free_freight;type:tinyint(2);not null;default:1;comment:是否包邮1-包邮 2-买家承担运费 3-有条件包邮" json:"free_freight"`
		PricingMode      uint8            `gorm:"column:pricing_mode;type:tinyint(2);not null;default:1;comment:计价方式1-按件数 2-按重量" json:"pricing_mode"`
		HaveUse          uint8            `gorm:"column:have_use;type:tinyint(2);not null;default:2;comment:是否已使用1-已使用 2-未使用" json:"have_use"`
		SaleArea         string           `gorm:"column:sale_area;type:text;not null;comment:允许发货区域json" json:"sale_area"`
		FreeCondition    string           `gorm:"column:free_condition;type:text;not null;comment:指定包邮条件"  json:"free_condition"`
		PostageCondition uint8            `gorm:"column:postage_condition;type:tinyint(2);not null;default:1;comment:是否包邮1-包邮 2-买家承担运费 3-有条件包邮" json:"postage_condition"`
		CreatedAt        base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt        base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt        *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
	FreightTemplatesCache []*FreightTemplate

	FreightFreeCondition struct {
		AreaCode []string `json:"a"` //区域
		FreightFreeConditionBase
	}
	FreightFreeConditionBase struct {
		FreightType uint8  `json:"ft"`
		FullPrice   string `json:"fp"`
		FullNumber  uint32 `json:"fn"`
	}
	FreightSaleArea struct {
		AreaCode []string `json:"a"` //区域
		FreightSaleAreaBase
	}
	FreightSaleAreaBase struct {
		FirstGoods string `json:"fg"` //首件数
		FirstPay   string `json:"fp"` //首费
		ExtGoods   string `json:"eg"` //续件数
		ExtPrice   string `json:"ep"` //续费
	}
)

func (r *FreightSaleAreaBase) GetFirstPay() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)
	if r.FirstPay != "" {
		if res, err = decimal.NewFromString(r.FirstPay); err != nil {
			return
		}
	}
	return
}

func (r *FreightSaleAreaBase) GetExtPrice() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)
	if r.ExtPrice != "" {
		if res, err = decimal.NewFromString(r.ExtPrice); err != nil {
			return
		}
	}
	return
}

func (r *FreightSaleAreaBase) GetFirstGoods() (firstGoods decimal.Decimal, err error) {
	firstGoods = decimal.NewFromInt(0)
	if r.FirstGoods == "" {
		return
	}
	if firstGoods, err = decimal.NewFromString(r.FirstGoods); err != nil {
		return
	}
	return
}

func (r *FreightSaleAreaBase) GetExtGoods() (extGoods decimal.Decimal, err error) {
	extGoods = decimal.NewFromInt(0)
	if r.ExtGoods == "" {
		return
	}
	if extGoods, err = decimal.NewFromString(r.ExtGoods); err != nil {
		return
	}

	return
}

func (r *FreightSaleAreaBase) GetPriceByUnit(num int64) (res decimal.Decimal, desc string, err error) {
	var descList = make([]string, 0, 10)
	defer func() {
		desc = strings.Join(descList, ",")
	}()
	numDecimal := decimal.NewFromInt(num)
	res = decimal.NewFromInt(0)
	var zeroDecimal = decimal.NewFromInt(0)
	var firstGoods decimal.Decimal
	if firstGoods, err = r.GetFirstGoods(); err != nil {
		descList = append(descList, "参数异常(首件数)")
		return
	}
	var firstPay decimal.Decimal
	if firstPay, err = r.GetFirstPay(); err != nil {
		descList = append(descList, "参数异常(首费数)")
		return
	}
	res = res.Add(firstPay)
	if numDecimal.LessThanOrEqual(firstGoods) { //未超过首件付基础运费
		descList = append(descList, "首件基础运费")
		res = firstPay
		return
	}
	descList = append(descList, fmt.Sprintf("基础费￥%v", firstPay.StringFixed(2)))
	var extGoods decimal.Decimal
	if extGoods, err = r.GetExtGoods(); err != nil {
		descList = append(descList, "参数异常(续件数)")
		return
	}
	if extGoods.Equal(zeroDecimal) { //如果续件为0 则不参与计算
		descList = append(descList, "续件未配置")
		return
	}
	var ExtPrice decimal.Decimal
	if ExtPrice, err = r.GetExtPrice(); err != nil {
		descList = append(descList, "参数异常(续费数)")
		return
	}
	//如果续费数为0页可用推出不参与计算
	if ExtPrice.Equal(decimal.NewFromInt(0)) {
		descList = append(descList, "续件(续费数)未配置")
		return
	}
	step := numDecimal.Sub(firstGoods).Div(extGoods).Ceil()
	descList = append(descList, fmt.Sprintf("续费:%v*%v", step, ExtPrice.StringFixed(2)))
	res = res.Add(step.Mul(ExtPrice))
	return
}

func (r *FreightSaleAreaBase) GetPriceByWeight(weightSummary decimal.Decimal) (res decimal.Decimal, desc string, err error) {
	res = decimal.NewFromInt(0)
	var descList = make([]string, 0, 10)
	defer func() {
		desc = strings.Join(descList, ",")
	}()
	var zeroDecimal = decimal.NewFromInt(0)
	var firstGoods decimal.Decimal
	if firstGoods, err = r.GetFirstGoods(); err != nil {
		descList = append(descList, "参数异常(首重数)")
		return
	}
	var firstPay decimal.Decimal
	if firstPay, err = r.GetFirstPay(); err != nil {
		descList = append(descList, "参数异常(首费数)")
		return
	}
	res = res.Add(firstPay)

	//如果总重量小于等于首重，则只付基础运费
	if weightSummary.LessThanOrEqual(firstGoods) {
		res = firstPay
		descList = append(descList, "首重基础运费")
		return
	}
	descList = append(descList, fmt.Sprintf("基础费￥%v", firstPay.StringFixed(2)))
	var extGoods decimal.Decimal
	if extGoods, err = r.GetExtGoods(); err != nil {
		descList = append(descList, "参数异常(续重数)")
		return
	}
	if extGoods.Equal(zeroDecimal) { //如果续重为0 则不参与计算
		descList = append(descList, "续重未配置")
		return
	}
	var ExtPrice decimal.Decimal
	if ExtPrice, err = r.GetExtPrice(); err != nil {
		descList = append(descList, "参数异常(续费数)")
		return
	}
	//如果续费数为0页可用推出不参与计算
	if ExtPrice.Equal(decimal.NewFromInt(0)) {
		descList = append(descList, "续重(续费数)未配置")
		return
	}
	step := weightSummary.Sub(firstGoods).Div(extGoods).Ceil()
	descList = append(descList, fmt.Sprintf("续费:%v*%v", step, ExtPrice.StringFixed(2)))
	res = res.Add(step.Mul(ExtPrice))
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

func (r *FreightTemplate) SetFreeCondition(data []*FreightFreeCondition) (err error) {
	if len(data) == 0 {
		return
	}
	var res []byte
	if res, err = json.Marshal(data); err != nil {
		return
	}
	r.FreeCondition = string(res)
	return
}

func (r *FreightTemplate) ParseFreeCondition() (res []*FreightFreeCondition, err error) {
	if r.FreeCondition == "" {
		return
	}
	err = json.Unmarshal([]byte(r.FreeCondition), &res)
	return
}

func (r *FreightTemplate) SetSaleArea(data []*FreightSaleArea) (err error) {
	if len(data) == 0 {
		return
	}
	var res []byte
	if res, err = json.Marshal(data); err != nil {
		return
	}
	r.SaleArea = string(res)
	return
}

func (r *FreightTemplate) ParseSaleArea() (res []*FreightSaleArea, err error) {
	if r.SaleArea == "" {
		return
	}
	err = json.Unmarshal([]byte(r.SaleArea), &res)
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

func (r *FreightTemplate) Default() {

	return
}
