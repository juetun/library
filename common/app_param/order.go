package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/const_apply"
	"github.com/shopspring/decimal"
)

const (
	OrderPageCategoryFirst  = "first"  //第一次付款或定金付款
	OrderPageCategorySecond = "second" //定金预售付尾款
)

const (
	OrderShopDetailPriceCateFirst  uint8 = iota + 1 //第一次付款或定金付款
	OrderShopDetailPriceCateSecond                  //定金预售付尾款
)

var (
	MapOrderCategoryActType = map[string]uint8{
		OrderPageCategoryFirst:  OrderShopDetailPriceCateFirst,
		OrderPageCategorySecond: OrderShopDetailPriceCateSecond,
	}
	SliceOrderShopDetailPriceCate = base.ModelItemOptions{
		{
			Label: "普通商品付款或定金付款",
			Value: OrderShopDetailPriceCateFirst,
		},
		{
			Label: "定金预售尾款",
			Value: OrderShopDetailPriceCateSecond,
		},
	}
	SliceOrderPageCategory = base.ModelItemOptions{
		{
			Label: "普通商品付款或定金付款",
			Value: OrderPageCategoryFirst,
		},
		{
			Label: "定金预售尾款",
			Value: OrderPageCategorySecond,
		},
	}
)

type (
	ArgOrderFromCartItem struct {
		ShopId          int64       `json:"shop_id,omitempty" form:"shop_id"`                 // 店铺ID
		CategoryId      int64       `json:"-" form:"-"`                                       // 商品类目
		SpuId           string      `json:"spu_id,omitempty" form:"spu_id"`                   // 商品ID
		SkuId           string      `json:"sku_id,omitempty" form:"sku_id"`                   // sku地址
		CartId          string      `json:"cart_id" form:"cart_id"`                           //购物车（或定金预售）数据ID
		SkuImg          string      `json:"sku_img,omitempty" form:"sku_img"`                 //商品图片
		Num             int64       `json:"num,omitempty" form:"num"`                         // 商品数量
		SaleType        uint8       `json:"sale_type,omitempty" form:"sale_type"`             // 销售类型
		SkuPrice        string      `json:"sku_price,omitempty" form:"sku_price"`             // SPU项目本次要支付的单价(定金预售定金金额或尾款金额 sku_price)
		SkuSetPrice     string      `json:"sku_set_price,omitempty" form:"sku_set_price"`     // SPU项目本的单价
		FreightTplId    int64       `json:"freight_tpl_id,omitempty" form:"freight_tpl_id"`   // 运费模板
		SubOrderId      string      `json:"sub_order_id,omitempty" form:"sub_order_id"`       //子订单号
		Category        string      `json:"category,omitempty" form:"category"`               //数据类型 first-首款 -或 second-尾款
		Checked         bool        `json:"checked" form:"checked"`                           //是否选中
		Pk              string      `json:"pk" form:"pk"`                                     //数据唯一性标记参数
		SpuFlagTester   uint8       `json:"spu_flag_tester,omitempty" form:"spu_flag_tester"` //spu是否为测试数据
		SkuFlagTester   uint8       `json:"sku_flag_tester,omitempty" form:"sku_flag_tester"` //sku是否为测试数据
		FreightAmount   string      `json:"freight_amount,omitempty" form:"freight_amount"`   // 邮费
		ShopSaleCode    string      `json:"shop_sale_code,omitempty" form:"shop_sale_code"`
		SkuPropertyName string      `json:"sku_property_name,omitempty" form:"sku_property_name"`
		ProvideChannel  int64       `json:"provide_channel,omitempty" form:"provide_channel"`
		ProvideSaleCode string      `json:"provide_sale_code,omitempty" form:"provide_sale_code"`
		OrderSrcChannel string      `json:"order_src_channel" form:"order_src_channel"` //订单来源渠道
		OrderSrcLoc     string      `json:"order_src_loc" form:"order_src_loc"`         //订单来源展示坑位
		Link            interface{} `json:"link"`                                       //商品链接
	}
)

func (r *ArgOrderFromCartItem) GetPrice() (res decimal.Decimal, err error) {
	if r.SkuPrice == "" {
		r.SkuPrice = "0.00"
	}
	if res, err = decimal.NewFromString(r.SkuPrice); err != nil {
		return
	}
	return
}

func SetOrderCategoryWith(actType uint8) (category string, err error) {
	var mapCategory = make(map[uint8]string, len(MapOrderCategoryActType))
	for key, value := range MapOrderCategoryActType {
		mapCategory[value] = key
	}
	if tmp, ok := mapCategory[actType]; ok {
		category = tmp
		return
	}
	err = fmt.Errorf("操作类型系统暂不支持")
	return
}

func (r *ArgOrderFromCartItem) SetOrderCategoryWith(actType uint8) (err error) {
	r.Category, err = SetOrderCategoryWith(actType)
	return
}

func ParseOrderActTypeWithCategory(category string) (actType uint8, err error) {
	if tmp, ok := MapOrderCategoryActType[category]; ok {
		actType = tmp
		return
	}
	err = fmt.Errorf("操作类型系统暂不支持")
	return
}

func (r *ArgOrderFromCartItem) ParseOrderActTypeWithCategory() (actType uint8, err error) {
	return ParseOrderActTypeWithCategory(r.Category)
}

func (r *ArgOrderFromCartItem) ValidateCategory() (err error) {
	if r.Category == "" {
		err = fmt.Errorf("请选择付款时机")
		return
	}
	var mapCategory map[string]string
	if mapCategory, err = SliceOrderPageCategory.GetMapAsKeyString(); err != nil {
		return
	}
	if _, ok := mapCategory[r.Category]; !ok {
		err = fmt.Errorf("请选择正确的付款时机")
		return
	}
	return
}
func (r *ArgOrderFromCartItem) GetTotalSkuPrice() (res decimal.Decimal, err error) {
	var price decimal.Decimal
	if price, err = r.GetPrice(); err != nil {
		return
	}
	res = price.Mul(decimal.NewFromInt(r.Num))
	return
}

func (r *ArgOrderFromCartItem) Default() {
	if r.SkuFlagTester == 0 {
		r.SkuFlagTester = const_apply.FlagTesterNo
	}
	if r.SpuFlagTester == 0 {
		r.SkuFlagTester = const_apply.FlagTesterNo
	}
}
