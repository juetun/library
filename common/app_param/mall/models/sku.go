package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/recommend"
	"github.com/shopspring/decimal"
	"net/url"
	"time"

	"github.com/juetun/base-wrapper/lib/base"
)

const (
	SkuStatusManuscript = ProductStatusManuscript //草稿中(指定了spuID的数据)
	SkuStatusTmp        = ProductStatusTmp        //草稿中(ID初始化中)
	SkuStatusOnline     = ProductStatusOnline     //在线
	SkuStatusOffLine    = ProductStatusOffLine
	SkuStatusDeprecated = ProductStatusDeprecated //已删除
)

const (
	//是否有赠品
	SkuHaveGiftYes uint8 = iota + 1 //有赠品
	SkuHaveGiftNo                   //无赠品
)

const (
	SkuHaveBindSpuYes uint8 = iota + 1 //是否绑定SPU 已绑定
	SkuHaveBindSpuNo                   //未绑定
)

const ( //供货商供货渠道
	SkuProvideChannelAliBaba int64 = iota + 1
	SkuProvideChannelDown
)

var (
	//注意:此数据只能在后边添加,否则会影响数据结构
	SliceSkuProvideChannel = base.ModelItemOptions{
		{
			Value: SkuProvideChannelAliBaba,
			Label: "阿里巴巴",
		},
		{
			Value: SkuProvideChannelDown,
			Label: "线下渠道",
		},
	}

	SliceSkuHaveBindSpu = base.ModelItemOptions{
		{
			Value: SkuHaveBindSpuYes,
			Label: "已绑定",
		},
		{
			Value: SkuHaveBindSpuNo,
			Label: "未绑定",
		},
	}
	SliceSkuHaveGift = base.ModelItemOptions{
		{
			Value: SkuHaveGiftYes,
			Label: "有赠品",
		},
		{
			Value: SkuHaveGiftNo,
			Label: "无赠品",
		},
	}
	SliceSkuStatus = base.ModelItemOptions{

		{
			Value: SkuStatusManuscript,
			Label: "编辑中",
		},
		{
			Value: SkuStatusTmp,
			Label: "草稿中",
		},

		{
			Value: SkuStatusOnline,
			Label: "已上架",
		},
		{
			Value: SkuStatusOffLine,
			Label: "已下架",
		},
		{
			Value: SkuStatusDeprecated,
			Label: "已删除",
		},
	}

	//当前正在编辑或可查看的SKU
	SliceSkuStatusEditShow = []int8{SkuStatusManuscript, SkuStatusTmp, SkuStatusOffLine, SkuStatusOnline}

	//商品编辑SKU状态选项（界面展示）
	SliceSkuStatusEditPageShow = []SliceSkuStatusEditOption{
		{
			ModelItemOption: base.ModelItemOption{
				Value: SkuStatusOnline,
				Label: "上架",
			},
		},
		{
			ModelItemOption: base.ModelItemOption{
				Value: SkuStatusOffLine,
				Label: "下架",
			},
		},
	}
)

type (
	Sku struct {
		ID              string           `gorm:"column:id;primary_key;type:bigint(20);not null;default:0;comment:商品SkUID" json:"sku_id"`
		SkuName         string           `gorm:"column:sku_name;default:'';type:varchar(120);not null;comment:标题" json:"sku_name"`
		Thumbnail       string           `gorm:"column:thumbnail;type:varchar(255);not null;default:'';comment:封面图ID" json:"thumbnail"`
		ThumbnailURL    string           `json:"thumbnail_url" gorm:"-"`
		LockKey         string           `gorm:"column:lock_key;type:varchar(60);not null;default:'';comment:临时锁的KEy" json:"lock_key"`
		SkuAttRelateId  int64            `gorm:"column:sku_att_relate_id;default:0;type:bigint(20);not null;comment:商品属性关系ID" json:"sku_att_relate_id"`
		Image           string           `gorm:"column:image;type:varchar(800);not null;default:'';comment:图片json数组" json:"image"`
		Video           string           `gorm:"column:video;type:varchar(255);not null;default:'';comment:视频" json:"video"`
		UserHid         int64            `json:"user_hid" gorm:"column:user_hid;default:0;type:bigint(20);not null;comment:发布人用户ID"`
		ShopId          int64            `gorm:"column:shop_id;index:idx_pro_id,priority:1;default:0;type:bigint(20);not null;comment:店铺ID" json:"shop_id"`
		SkuStatus       int8             `gorm:"column:sku_status;default:4;type:tinyint(2);index:idx_pro_id,priority:3;not null;comment:状态 1-可用 2-下架 3-删除" json:"sku_status"`
		Weight          string           `gorm:"column:weight;default:0;type:decimal(10,2);not null;comment:重量 单位-千克" json:"weight"`
		Price           string           `gorm:"column:price;default:0;type:decimal(10,2);not null;comment:售价" json:"price"`
		MarketCost      string           `gorm:"column:market_cost;default:0;type:decimal(10,2);not null;comment:划线价" json:"market_cost"`
		PriceCost       string           `gorm:"column:price_cost;default:0;type:decimal(10,2);not null;comment:成本价" json:"price_cost"`
		ShopSaleCode    string           `gorm:"column:shop_sale_code;type:varchar(80);default:'';not null;comment:商家供货码" json:"shop_sale_code"`
		ProvideChannel  int64            `gorm:"column:provide_channel;type:bigint(20);not null;default:0;comment:供货渠道" json:"provide_channel"`
		ProvideSaleCode string           `gorm:"column:provide_sale_code;type:varchar(80);default:'';not null;comment:供货商供货码" json:"provide_sale_code"`
		SaleNum         int              `gorm:"column:sale_num;type:bigint(20);not null;default:0;comment:销量(数据可能不及时)" json:"sale_num"`
		SaleOnlineTime  base.TimeNormal  `gorm:"column:sale_online_time;not null;default:CURRENT_TIMESTAMP;comment:预售开始时间" json:"sale_online_time"`
		SaleOverTime    *base.TimeNormal `gorm:"column:sale_over_time;comment:预售结束时间" json:"sale_over_time"`
		Volume          string           `gorm:"column:volume;default:0;type:decimal(10,2);not null;comment:容积" json:"volume"`
		FlagTester      uint8            `gorm:"column:flag_tester;not null;type: tinyint(2);default:1;comment:是否为测试数据 1-不是 2-是"  json:"flag_tester"`
		HaveBindSpu     uint8            `gorm:"column:have_bind_spu;not null;type: tinyint(2);default:0;comment:是否绑定商品 1-是 2-不是"  json:"have_bind_spu"`

		CreatedAt base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
	ProductSKus              []*Sku
	SliceSkuStatusEditOption struct {
		base.ModelItemOption
		Checked bool `json:"checked"`
	}

	SkuPropertyRelate struct {
		ID                 int64  `gorm:"column:id;primary_key" json:"id"`
		ShopId             int64  `gorm:"column:shop_id;default:0;type:varchar(60);not null;comment:店铺ID" json:"shop_id"`
		ProductId          string `gorm:"column:product_id;uniqueIndex:uniquePK,priority:1;default:'';type:varchar(40);not null;comment:商品ID" json:"product_id"`
		CategoryId         int64  `gorm:"column:category_id;not null;type:bigint(20);default:0;comment:类目ID" json:"category_id"` // comment:用户类目类型;
		ParentId           int64  `gorm:"column:parent_id;not null;default:0;comment:skuID" json:"parent_id"`
		Pk                 string `gorm:"column:pk;uniqueIndex:uniquePK,priority:2;default:'';type:varchar(80);not null;comment:商品唯一Key" json:"pk"`
		SkuName            string `gorm:"column:sku_name;default:'';type:varchar(120);not null;comment:商品名称" json:"sku_name"`
		SkuId              string `gorm:"column:sku_id;default:'';type:varchar(40);not null;comment:skuID" json:"sku_id"`
		Price              string `gorm:"column:price;default:0;type:decimal(10,2);not null;comment:售价" json:"price"`
		IsNotAttrName      uint8  `gorm:"column:is_not_attr_name;type:tinyint(2);not null;default:2;comment:不是属性名 1-真-属性名 2-假-为属性"  json:"is_not_attr_name"`
		PropertyId         int64  `gorm:"column:property_id;default:0;not null;comment:属性ID" json:"property_id"`
		SpuStatus          int8   `gorm:"column:spu_status;default:0;type:tinyint(2);not null;comment:商品状态(具体与商品表对齐)" json:"spu_status"`
		SaleType           uint8  `gorm:"column:sale_type;not null;type: tinyint(2);default:1;comment:销售类型1-普通商品 2-全款预售 3-定金预售"  json:"sale_type"`
		IntentionalDeposit string `gorm:"column:intentional_deposit;default:0;type:decimal(10,2);not null;comment:意向金" json:"intentional_deposit"`
		IntentDepositVal   string `gorm:"column:intent_deposit_val;default:0;type:decimal(10,2);not null;comment:意向金抵扣" json:"intent_deposit_val"`
		DownPayment        string `gorm:"column:down_payment;default:0;type:decimal(10,2);not null;comment:定金" json:"down_payment"`
		DownPaymentVal     string `gorm:"column:down_payment_val;default:0;type:decimal(10,2);not null;comment:定金抵扣" json:"down_payment_val"`
		//FinalPayment       string           `gorm:"column:final_payment;default:0;type:decimal(10,2);not null;comment:尾款金额(商品为定金预售数据有效)" json:"final_payment"`
		FreightTemplate   int64            `gorm:"column:freight_template;type:varchar(80);default:0;not null;comment:运费模板ID" json:"freight_template"`
		SaleOnlineTime    base.TimeNormal  `gorm:"column:sale_online_time;not null;default:CURRENT_TIMESTAMP;comment:预售开始时间" json:"sale_online_time"`
		SaleOverTime      *base.TimeNormal `gorm:"column:sale_over_time;comment:预售结束时间" json:"sale_over_time"`
		FinalStartTime    base.TimeNormal  `gorm:"column:final_start_time;not null;default:CURRENT_TIMESTAMP;comment:尾款开始时间" json:"final_start_time"`
		FinalOverTime     base.TimeNormal  `gorm:"column:final_over_time;not null;default:CURRENT_TIMESTAMP;comment:尾款结束时间" json:"final_over_time"`
		SalesTaxRate      string           `gorm:"column:sales_tax_rate;not null;type:decimal(10,2);default:0;comment:销售税率(百分比)"  json:"sales_tax_rate"`
		SalesTaxRateValue string           `gorm:"column:sales_tax_rate_value;not null;type:decimal(10,2);default:0;comment:销售税（金额 单位元）"  json:"sales_tax_rate_value"`
		MaxLimitNum       int64            `gorm:"column:max_limit_num;default:0;type:bigint(20);not null;comment:限购数量，每人最多购买数量" json:"max_limit_num"`
		MinLimitNum       int64            `gorm:"column:min_limit_num;default:0;type:bigint(20);not null;comment:必购数量，如2件起购" json:"min_limit_num"`
		IsLeaf            uint8            `gorm:"column:is_leaf;default:2;type:tinyint(2);not null;comment:是否为叶子节点 1-是 2-否" json:"is_leaf"`
		SkuStatus         int8             `gorm:"column:sku_status;default:1;type:tinyint(2);not null;comment:状态 1-可用 2-下架 3-删除" json:"sku_status"`
		HaveGift          uint8            `gorm:"column:have_gift;default:2;type:tinyint(2);not null;comment:是有有赠品 1-有 2-无" json:"have_gift"`
		OpenSaleNumStatic uint8            `gorm:"column:open_sale_num_static;default:2;type:tinyint(2);not null;comment:是有有赠品 1-参与 2-不参与" json:"open_sale_num_static"`
		CreatedAt         base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt         base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt         *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
	}
	SkuPropertyRelatesCache []SkuPropertyRelate
)

func (r *SkuPropertyRelate) GetFinalPayment() (finalPayment string) {
	var price, intentDepositVal, downPay decimal.Decimal
	price, _ = decimal.NewFromString(r.Price)
	if r.IntentDepositVal != "" && r.IntentDepositVal != "0.00" {
		intentDepositVal, _ = decimal.NewFromString(r.IntentDepositVal)
	}
	if r.DownPaymentVal != "" && r.DownPaymentVal != "0.00" {
		downPay, _ = decimal.NewFromString(r.DownPaymentVal)
	}
	finalPayment = price.Sub(intentDepositVal).Sub(downPay).String()
	return
}

//前端编辑SKU时的选项
func GetSliceSkuStatusEditOption(skuStatus int8) (res []*SliceSkuStatusEditOption) {

	res = make([]*SliceSkuStatusEditOption, 0, len(SliceSkuStatusEditShow))

	var (
		data *SliceSkuStatusEditOption
	)

	for _, item := range SliceSkuStatusEditPageShow {
		data = &SliceSkuStatusEditOption{ModelItemOption: item.ModelItemOption,}
		if skuStatus == data.Value {
			data.Checked = true
		}
		res = append(res, data)
	}
	return
}

func (r *Sku) GetWeightDecimal() (weight decimal.Decimal, err error) {
	weight = decimal.NewFromFloat(0)
	if r.Weight != "" {
		weight, err = decimal.NewFromString(r.Weight)
	}
	return
}

func (r *Sku) Default() {

	if r.Price == "" {
		r.Price = "0"
	}
	if r.PriceCost == "" {
		r.PriceCost = "0"
	}

	if r.Weight == "" {
		r.Weight = "0"
	}
	if r.MarketCost == "" {
		r.MarketCost = "0.00"
	}

	if r.Volume == "" {
		r.Volume = "0"
	}
	if r.SkuStatus == 0 {
		r.SkuStatus = SkuStatusOffLine
	}
}

func (r *ProductSKus) UnmarshalBinary(data []byte) (err error) {
	if len(data) == 0 {
		*r = []*Sku{}
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *ProductSKus) MarshalBinary() (data []byte, err error) {
	if len(*r) == 0 {
		data = []byte{}
		return
	}
	data, err = json.Marshal(r)
	return
}
func (r *Sku) TableName() string {
	return fmt.Sprintf("%ssku", TablePrefix)
}

func (r *Sku) GetTableComment() (res string) {
	res = "商品sku表"
	return
}

func (r *Sku) GetHid() (res string) {
	res = r.ID
	return
}

func (r *Sku) ParseStatusName() (res string) {
	var ok bool
	MapSkuStatus, _ := SliceSkuStatus.GetMapAsKeyInt8()
	if res, ok = MapSkuStatus[r.SkuStatus]; ok {
		return
	}
	return
}

func ParseHaveBindSpu(HaveBindSpu uint8) (res string) {
	var ok bool
	MapSkuHaveBind, _ := SliceSkuHaveBindSpu.GetMapAsKeyUint8()
	if res, ok = MapSkuHaveBind[HaveBindSpu]; ok {
		return
	}
	res = fmt.Sprintf("未知类型(%v)", HaveBindSpu)
	return
}

func (r *Sku) ParseHaveBindSpu() (res string) {
	res = ParseHaveBindSpu(r.HaveBindSpu)
	return
}

func (r *Sku) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *Sku) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

// 判断商品(SKU)是否能够购买
// Return pageMessage 界面展示不能购买原因
// Return systemMark 系统记录不能购买原因 用于记录日志使用
func (r *Sku) JudgeCanBuyStatus(currentTimes ...time.Time) (canBuy bool, pageMessage, systemMark string) {
	var current time.Time
	if len(currentTimes) > 0 {
		current = currentTimes[0]
	} else {
		current = time.Now()
	}
	r.defaultSaleOverTime()

	switch r.SkuStatus {
	case ProductStatusOnline:
		if current.Before(r.SaleOnlineTime.Time) || current.After(r.SaleOverTime.Time) {
			systemMark = "商品(SKU)在售时间不正确"
			pageMessage = ""
			return
		}
		canBuy = true
	default:
		systemMark = "商品(SKU)状态不是在售中"
	}
	return
}

func (r *Sku) SetIdWithString(id string) (err error) {
	r.ID = id
	//r.ID, err = strconv.ParseInt(id, 10, 64)
	return
}

func (r *SkuPropertyRelate) GetProductHref(headerInfo *common.HeaderInfo) (res interface{}, err error) {
	var vals = &url.Values{}
	vals.Set("id", r.ProductId)
	vals.Set("sku_id", r.SkuId)
	res, err = recommend.GetPageLink(&recommend.LinkArgument{
		HeaderInfo: headerInfo,
		UrlValue:   vals,
		DataType:   recommend.AdDataDataTypeSpu,
		PageName:   recommend.PageNameSpu,
	})
	return
}

func (r *SkuPropertyRelate) ParseSaleType() (res string) {
	return ParseSaleType(r.SaleType)
}

// 默认销售截止时间
func (r *Sku) defaultSaleOverTime() {
	if r.SaleOverTime == nil {
		r.SaleOverTime = &base.TimeNormal{Time: r.SaleOnlineTime.Time.Add(100 * 24 * time.Hour)}
	}
}

//func (r *Sku) ParseSpuStatus() (res string) {
//	res = ProductParseStatus(r.SpuStatus)
//	return
//}
