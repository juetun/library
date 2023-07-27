package order_create

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall"
	"github.com/juetun/library/common/app_param/mall/freight"
	"github.com/juetun/library/common/app_param/pay_parameter"
	"github.com/shopspring/decimal"
	"time"
)

const (
	OrderFromTypeCart   = "cart"   //购物车
	OrderFromTypeDirect = "detail" //直接购买
)

var (
	SliceOrderFromType = base.ModelItemOptions{
		{
			Label: "购物车",
			Value: OrderFromTypeCart,
		},
		{
			Label: "直接购买",
			Value: OrderFromTypeDirect,
		},
	}
)

type (
	//使用优惠券信息
	ArgUseCouponData struct {
		UserHid int64             `json:"user_hid"` //使用优惠券的用户
		Data    UseCouponDataList `json:"data"`
	}
	UseCouponDataList []*UseCouponDataItem
	UseCouponDataItem struct {
		UserCouponId int64  `json:"user_coupon_id"` //用户优惠券存储的数据ID
		CouponId     string `json:"coupon_id"`      //优惠券ID
		Num          int    `json:"num"`            //使用优惠券数据量
	}
	ArgGetUserCouponByShopId struct {
		UserHid int64   `json:"user_hid"` //使用优惠券的用户
		ShopIds []int64 `json:"shop_ids"`
	}

	ResultGetUserCouponByShopId struct {
		Data []*ResultGetUserCouponByShopIdItem `json:"data"`
	}
	OrderPreview struct {
		Amount             string                             `json:"amount"`                // 总金额
		AmountDecimal      decimal.Decimal                    `json:"-"`                     // 总金额 数字格式 值与Amount一致
		Preferential       string                             `json:"preferential"`          // 优惠金额
		PayCharge          string                             `json:"pay_charge"`            //支付需要手续费
		TotalPostage       string                             `json:"total_postage"`         //邮费
		ProductAmount      string                             `json:"product_amount"`        //商品总额
		Status             uint8                              `json:"status,omitempty"`      // 订单状态
		AddressId          string                             `json:"address_id,omitempty"`  // 收货地址
		Express            string                             `json:"express,omitempty"`     // 快递信息
		PayTypeOpt         base.ModelItemOptions              `json:"pay_type_opt"`          // 支付信息
		PayType            string                             `json:"pay_type"`              // 支付类型
		PayNumber          string                             `json:"pay_number,omitempty"`  // 支付流水号
		Token              string                             `json:"token,omitempty"`       // 购物车全站唯一的key
		Count              int64                              `json:"count"`                 // 商品总数
		HaveError          bool                               `json:"have_error"`            //是否有错误
		ErrorMessage       string                             `json:"error_message"`         //错误信息
		OrderId            string                             `json:"order_id"`              //订单号
		EmsAddress         *freight.ResultGetByAddressIdsItem `json:"ems_address,omitempty"` //收货地址信息
		ShopDiscountAmount string                             `json:"shop_discount_amount"`  // 店铺优惠总金额 (店铺维度优惠+SPU维度优惠)
		PlatDiscountAmount string                             `json:"plat_discount_amount"`  // 平台优惠金额 (店铺维度优惠+SPU维度优惠)
		PlatCoupon         *mall.CanUseCouponItem             `json:"plat_coupon,omitempty"` //平台券信息 （店铺维度优惠）
		List               PreviewShopItems                   `json:"list"`                  // SKU
	}
	PreviewShopItems []*PreviewShopItem
	PreviewShopItem  struct {
		ShopId             int64              `json:"shop_id"`
		ShopIcon           string             `json:"shop_icon"`             // 店铺Icon
		ShopName           string             `json:"shop_name"`             // 店铺名称
		ShopType           string             `json:"shop_type"`             // 店铺类型
		Count              int64              `json:"count"`                 // 商品总数
		PayCharge          string             `json:"pay_charge"`            // 支付手续费
		DeductionAmount    string             `json:"deduction_amount"`      // 现金抵扣券
		ProductAmount      string             `json:"product_amount"`        // 商品金额
		TotalAmount        string             `json:"total_amount"`          // 该订单店铺总的金额
		Delivery           OrderSkuDelivery   `json:"delivery"`              // 邮费信息
		Mark               string             `json:"mark"`                  // 用户备注
		SubOrderId         string             `json:"sub_order_id"`          // 子单号
		ShopDiscountAmount string             `json:"shop_discount_amount"`  // 店铺优惠总金额 (店铺维度优惠+SPU维度优惠)
		PlatDiscountAmount string             `json:"plat_discount_amount"`  // 平台优惠金额 (店铺维度优惠+SPU维度优惠)
		ShopCoupon         *mall.CanUseCoupon `json:"shop_coupon,omitempty"` // 店铺券信息
		Products           PreviewSpuItems    `json:"products"`              // 商品列表

		SortCreateTime time.Time `json:"-"`
		SortWeight     int64     `json:"-"` //排序权重
	}
	PreviewSpuItems []*PreviewSpuItem
	PreviewSkuItems []*PreviewSkuItem
	PreviewSkuItem  struct {
		app_param.ArgOrderFromCartItem
		//Title          string          `json:"title"`
		SkuName         string `json:"sku_name"`
		SkuPropertyName string `json:"sku_property_name"`
		SkuId           string `json:"sku_id"`      //购物车数据ID
		SkuPic          string `json:"sku_pic"`     // 图片
		SkuStatus       int8   `json:"sku_status"`  // 商品状态
		StatusName      string `json:"status_name"` // 商品状态名称 (已下架)
		TotalPrice      string `json:"total_price"`
		SaleType        uint8  `json:"sale_type"`
		SaleTypeName    string `json:"sale_type_name"`
		HaveVideo       bool   `json:"have_video"`   //是否有视频
		Mark            string `json:"mark"`         //商品说明（如 比着加入有无车时降价多少）
		MarkSystem      string `json:"mark_system"`  //数据不合法 系统说明(系统使用，记录更详细不合法原因)
		NotCanPay       bool   `json:"not_can_pay"`  //当前数据是否能够支付
		Invalidation    bool   `json:"invalidation"` //是否失效 true-已失效 false-未失效
		Checked         bool   `json:"checked"`      //是否选中

		SpecialTags []*DataItemTag `json:"special_tags"`

		SortCreateTime base.TimeNormal `json:"-"`
		SortWeight     int64           `json:"-"`
		// 单一SKU的总价格
	}
	DataItemTag struct {
		Type      string `json:"type"`                //标签类型，可选值为primary success danger warning	默认	default
		Label     string `json:"label"`               //类型名称
		Color     string `json:"color,omitempty"`     //标签颜色
		TextColor string `json:"textColor,omitempty"` //文本颜色，优先级高于color属性	String	white
		Plain     bool   `json:"plain"`               //是否为空心样式	Boolean	false
		Round     bool   `json:"round"`               //是否为圆角样式	Boolean	false
		Mark      bool   `json:"mark"`                //是否为标记样式
	}
	PreviewSpuItem struct {
		SpuId        string             `json:"spu_id"`
		Skus         PreviewSkuItems    `json:"skus"`                 // SKU信息
		SpuCoupon    *mall.CanUseCoupon `json:"spu_coupon,omitempty"` // 店铺券信息
		Count        int64              `json:"count"`                // 商品总数
		TotalAmount  string             `json:"total_amount"`         // 该订单店铺总的金额
		SaleType     uint8              `json:"sale_type"`
		SaleTypeName string             `json:"sale_type_name"`
	}
	ResultGetUserCouponByShopIdItem struct {
		UseCouponDataItem
		ShopId int64 `json:"shop_id"`

		Desc string `json:"desc"` //优惠券信息描述
	}

	ResultGetInfoByOrder map[string]ResultGetInfoByOrderItem

	ResultGetInfoByOrderItem struct {
		StatusName         string           `json:"status_name"`          //订单状态中文描述
		ShopItems          []*OrderShopItem `json:"shop_items"`           //商品列表（按店铺分组）
		UserHid            int64            `json:"user_hid"`             //用户
		OrderId            string           `json:"order_id"`             //订单ID号
		ProductNum         int64            `json:"product_num"`          //商品数
		Status             uint8            `json:"status"`               //订单状态
		SubStatus          uint8            `json:"sub_status"`           //子单状态
		AddressId          int64            `json:"address_id"`           //收货地址信息
		PayType            uint8            `json:"pay_type"`             //支付方式
		Amount             string           `json:"amount"`               //支付总金额
		PayAmount          string           `json:"pay_amount"`           //支付金额
		TotalPostage       string           `json:"total_postage"`        //邮费
		ProductAmount      string           `json:"product_amount"`       //商品金额
		Preferential       string           `json:"preferential"`         //优惠金额
		ShopDiscountAmount string           `json:"shop_discount_amount"` //店铺优惠金额
		PlatDiscountAmount string           `json:"plat_discount_amount"` //平台优惠金额
		PayCharge          string           `json:"pay_charge"`           //支付需要手续费
		Mark               string           `json:"mark"`                 //备注
	}

	OrderShopItem struct {
		ShopId             int64            `json:"shop_id"`              // 店铺ID
		ShopIcon           string           `json:"shop_icon"`            // 店铺Icon
		ShopName           string           `json:"shop_name"`            // 店铺名称
		ShopType           string           `json:"shop_type"`            // 店铺类型
		Count              int64            `json:"count"`                // 商品总数
		ShopChecked        bool             `json:"shop_checked"`         // 店铺选择
		TotalAmount        string           `json:"total_amount"`         // 该订单店铺总的金额
		PayCharge          string           `json:"pay_charge"`           // 支付手续费
		Products           []*OrderSkuItem  `json:"products"`             // 商品列表
		Delivery           OrderSkuDelivery `json:"delivery"`             // 邮费信息
		ShopDiscountAmount string           `json:"shop_discount_amount"` // 店铺优惠总金额 (店铺维度优惠+SPU维度优惠)
		PlatDiscountAmount string           `json:"plat_discount_amount"` // 平台优惠总金额 (店铺维度优惠+SPU维度优惠)
		SubOrderId         string           `json:"sub_order_id"`         // 子单号
		Mark               string           `json:"mark"`                 // 备注
		SortWeight         int64            `json:"-"`                    // 排序权重
	}

	OrderSkuItem struct {
		SkuName         string              `json:"sku_name"`
		SpuId           string              `json:"spu_id"`
		SkuId           string              `json:"sku_id"`      //购物车数据ID
		SkuPic          string              `json:"sku_pic"`     // 图片
		SkuStatus       int8                `json:"sku_status"`  // 商品状态
		StatusName      string              `json:"status_name"` // 商品状态名称 (已下架)
		Price           string              `json:"price"`       //单价
		Num             int64               `json:"num"`         //商品数量
		PriceCate       uint8               `json:"price_cate"`  //定金类型当前商品类型为定金预售时1-首付款 2-尾款
		PriceCateStr    string              `json:"price_cate_str"`
		PriceCateName   string              `json:"price_cate_name"`
		TotalPrice      string              `json:"total_price"`
		SaleTypeName    string              `json:"sale_type_name"`
		SaleType        uint8               `json:"sale_type"`
		SkuPropertyName string              `json:"sku_property_name"` //SKU属性名
		DownPayment     string              `json:"down_payment"`      //定金
		HaveVideo       bool                `json:"have_video"`        //是否有视频
		Mark            string              `json:"mark"`              //商品说明（如 比着加入有无车时降价多少）
		MarkSystem      string              `json:"mark_system"`       //数据不合法 系统说明(系统使用，记录更详细不合法原因)
		Checked         bool                `json:"checked"`           //是否选中
		SortCreateTime  base.TimeNormal     `json:"-"`
		SpecialTags     []*OrderDataItemTag `json:"special_tags"`
		SortWeight      int64               `json:"-"`
	}
	OrderSkuDelivery struct {
		PostageMark string `json:"postage_mark"` //备注
		Cost        string `json:"cost"`         //费用
		IsFree      bool   `json:"is_free"`      //是否包邮
	}
	OrderDataItemTag struct {
		Type      string `json:"type"`                //标签类型，可选值为primary success danger warning	默认	default
		Label     string `json:"label"`               //类型名称
		Color     string `json:"color,omitempty"`     //标签颜色
		TextColor string `json:"textColor,omitempty"` //文本颜色，优先级高于color属性	String	white
		Plain     bool   `json:"plain"`               //是否为空心样式	Boolean	false
		Round     bool   `json:"round"`               //是否为圆角样式	Boolean	false
		Mark      bool   `json:"mark"`                //是否为标记样式
	}
)

func (r *ArgGetUserCouponByShopId) Default(ctx *base.Context) (err error) {
	if r.UserHid == 0 {
		err = fmt.Errorf("请选择查看优惠券的用户")
		return
	}

	return
}

func (r *OrderShopItem) Default() (err error) {

	if r.TotalAmount == "" {
		r.TotalAmount = "0.00"
	}
	if r.PayCharge == "" {
		r.PayCharge = "0.00"
	}
	if r.Delivery.Cost == "" {
		r.Delivery.Cost = "0.00"
	}
	if r.ShopDiscountAmount == "" {
		r.ShopDiscountAmount = "0.00"
	}
	if r.PlatDiscountAmount == "" {
		r.PlatDiscountAmount = "0.00"
	}
	return
}

func (r *ArgUseCouponData) Default(ctx *base.Context) (err error) {
	if r.UserHid == 0 {
		err = fmt.Errorf("请选择使用优惠券的用户")
		return
	}

	//默认优惠券数量
	for _, item := range r.Data {
		if item.Num == 0 {
			item.Num = 1
		}
	}
	return
}

//初始化支付方式
func (r *OrderPreview) InitPayTypeOption(info *common.HeaderInfo) (err error) {
	r.PayTypeOpt = pay_parameter.SliceOrderPayType

	switch info.HTerminal {
	case app_param.TerminalMina: //如果是微信小程序
		switch info.HChannel {
		case "weixin": //如果是小程序微信使用
			r.getWeiXinMinaOpt()
		}

	}
	r.PayType = fmt.Sprintf("%v", r.PayTypeOpt[0].Value)
	return
}

func (r *OrderPreview) getWeiXinMinaOpt() {
	mapPay, _ := pay_parameter.SliceOrderPayType.GetMapAsKeyUint8()
	r.PayTypeOpt = base.ModelItemOptions{
		{
			Label: mapPay[pay_parameter.OrderPayTypeWeiXin],
			Value: pay_parameter.OrderPayTypeWeiXin,
		},
	}
	return
}

func (r *OrderPreview) Default() (err error) {
	if r.Preferential == "" {
		r.Preferential = "0.00"
	}
	if r.ShopDiscountAmount == "" {
		r.ShopDiscountAmount = "0.00"
	}
	if r.PlatDiscountAmount == "" {
		r.PlatDiscountAmount = "0.00"
	}
	if r.PayCharge == "" {
		r.PayCharge = "0.00"
	}
	if r.PayCharge == "" {
		r.PayCharge = "0.00"
	}
	if r.TotalPostage == "" {
		r.TotalPostage = "0.00"
	}
	if r.Amount == "" {
		r.Amount = "0.00"
	}
	if r.ProductAmount == "" {
		r.ProductAmount = "0.00"
	}
	return
}

func (r *OrderPreview) AmountDecr(decr string) (err error) {
	var (
		decimalDec decimal.Decimal
	)
	if decimalDec, err = decimal.NewFromString(decr); err != nil {
		return
	}
	r.AmountDecimal = r.AmountDecimal.Sub(decimalDec)
	r.Amount = r.AmountDecimal.StringFixed(2)
	return
}

//计算店铺实际优惠金额
func (r *PreviewShopItem) CalTotalWithCoupon() (err error) {
	type (
		CostValueHandler func() (value decimal.Decimal, err error)
	)
	var (
		mapValueHandler = map[string]CostValueHandler{
			"ProductAmount": r.getProductAmount, //商品金额
			"SpuDecr":       r.getSpuDecr,       //spu扣减金额 优惠券+抵扣券（代金券）
			"ShopDecr":      r.getShopDecr,      //店铺抵扣金额 优惠券+抵扣券（代金券）
			"PayCharge":     r.getPayCharge,     //支付手续费
			"DeliveryCost":  r.getDeliveryCost,  //邮费金额
		}
		mapValue = make(map[string]decimal.Decimal, len(mapValueHandler))
	)
	for key, handler := range mapValueHandler {
		if mapValue[key], err = handler(); err != nil {
			return
		}
	}
	//商品价格+邮费+支付渠道手续费-平台券抵扣-店铺券抵扣-代金券抵扣
	r.TotalAmount = mapValue["ProductAmount"].Add(mapValue["DeliveryCost"]).Add(mapValue["PayCharge"]).Sub(mapValue["SpuDecr"]).Sub(mapValue["ShopDecr"]).StringFixed(2)
	return
}

func (r *PreviewShopItem) getSpuDecr() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)

	var (
		shopDecr, platDecr, deductionAmount decimal.Decimal
	)
	for _, spuInfo := range r.Products {
		if spuInfo.SpuCoupon == nil {
			continue
		}
		shopDecr = decimal.NewFromInt(0)
		platDecr = decimal.NewFromInt(0)
		deductionAmount = decimal.NewFromInt(0)

		if spuInfo.SpuCoupon.DeductionAmount != "" {
			if deductionAmount, err = decimal.NewFromString(spuInfo.SpuCoupon.DeductionAmount); err != nil {
				return
			}
		}
		if spuInfo.SpuCoupon.Plat != nil && spuInfo.SpuCoupon.Plat.CurrentUse.ID > 0 {
			if platDecr, err = decimal.NewFromString(spuInfo.SpuCoupon.Plat.CurrentUse.Decr); err != nil {
				return
			}
		}
		if spuInfo.SpuCoupon.Shop != nil && spuInfo.SpuCoupon.Shop.CurrentUse.ID > 0 {
			if shopDecr, err = decimal.NewFromString(spuInfo.SpuCoupon.Shop.CurrentUse.Decr); err != nil {
				return
			}
		}
		res = res.Add(shopDecr).Add(platDecr).Add(deductionAmount)
	}
	return
}

func (r *PreviewShopItem) getShopDecr() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)
	var (
		shopDecr, platDecr, deductionAmount decimal.Decimal
	)

	shopDecr = decimal.NewFromInt(0)
	if r.ShopCoupon != nil && r.ShopCoupon.Shop != nil && r.ShopCoupon.Shop.CurrentUse.ID > 0 && r.ShopCoupon.Shop.CurrentUse.Decr != "" {
		shopDecr, err = decimal.NewFromString(r.ShopCoupon.Shop.CurrentUse.Decr)
	}

	platDecr = decimal.NewFromInt(0)
	if r.ShopCoupon != nil && r.ShopCoupon.Plat != nil && r.ShopCoupon.Plat.CurrentUse.ID > 0 && r.ShopCoupon.Plat.CurrentUse.Decr != "" {
		platDecr, err = decimal.NewFromString(r.ShopCoupon.Plat.CurrentUse.Decr)
	}

	deductionAmount = decimal.NewFromInt(0)
	if r.ShopCoupon != nil && r.ShopCoupon.Plat != nil && r.ShopCoupon.Plat.CurrentUse.ID > 0 && r.ShopCoupon.Plat.CurrentUse.Decr != "" {
		deductionAmount, err = decimal.NewFromString(r.ShopCoupon.Plat.CurrentUse.Decr)
	}
	res = res.Add(shopDecr).Add(platDecr).Add(deductionAmount)
	return
}

func (r *PreviewShopItem) getDeductionAmount() (productAmount decimal.Decimal, err error) {
	productAmount = decimal.NewFromInt(0)
	if r.DeductionAmount != "" {
		productAmount, err = decimal.NewFromString(r.DeductionAmount)
	}
	return
}

func (r *PreviewShopItem) getPayCharge() (productAmount decimal.Decimal, err error) {
	productAmount = decimal.NewFromInt(0)
	if r.PayCharge != "" {
		productAmount, err = decimal.NewFromString(r.PayCharge)
	}
	return
}

func (r *PreviewShopItem) getDeliveryCost() (productAmount decimal.Decimal, err error) {
	productAmount = decimal.NewFromInt(0)
	if r.Delivery.Cost != "" {
		productAmount, err = decimal.NewFromString(r.Delivery.Cost)
	}
	return
}

func (r *PreviewShopItem) getProductAmount() (productAmount decimal.Decimal, err error) {
	productAmount = decimal.NewFromInt(0)
	if r.ProductAmount != "" {
		productAmount, err = decimal.NewFromString(r.ProductAmount)
	}
	return
}

func (r *PreviewShopItem) SetShopItem(orderShopItem *OrderShopItem) (err error) {
	if err = orderShopItem.Default(); err != nil {
		return
	}
	r.ShopId = orderShopItem.ShopId
	r.ShopIcon = orderShopItem.ShopIcon
	r.ShopName = orderShopItem.ShopName
	r.ShopType = orderShopItem.ShopType
	r.Count = orderShopItem.Count
	r.TotalAmount = orderShopItem.TotalAmount
	r.ProductAmount = orderShopItem.TotalAmount
	r.SortWeight = orderShopItem.SortWeight
	if orderShopItem.Delivery.Cost == "" {
		orderShopItem.Delivery.Cost = "0.00"
	}
	r.PlatDiscountAmount = orderShopItem.PlatDiscountAmount
	r.ShopDiscountAmount = orderShopItem.ShopDiscountAmount
	r.Delivery = orderShopItem.Delivery
	//r.Coupon = orderShopItem.Coupon
	r.Mark = orderShopItem.Mark
	r.PayCharge = orderShopItem.PayCharge
	r.SubOrderId = orderShopItem.SubOrderId

	return
}

func (r *PreviewSkuItem) ParseOrderSkuItem(orderSkuItem *OrderSkuItem) {
	r.SkuId = orderSkuItem.SkuId
	r.SpuId = orderSkuItem.SpuId
	r.Category = orderSkuItem.PriceCateStr
	r.SaleType = orderSkuItem.SaleType
	r.SaleTypeName = orderSkuItem.SaleTypeName
	r.SkuPrice = orderSkuItem.Price
	r.SkuPropertyName = orderSkuItem.SkuPropertyName
	r.Num = orderSkuItem.Num
	return
}

func (r *PreviewShopItem) Default() (err error) {
	if r.PayCharge == "" {
		r.PayCharge = "0.00"
	}
	if r.DeductionAmount == "" {
		r.DeductionAmount = "0.00"
	}
	if r.ProductAmount == "" {
		r.ProductAmount = "0.00"
	}
	if r.TotalAmount == "" {
		r.TotalAmount = "0.00"
	}
	if r.Delivery.Cost == "" {
		r.Delivery.Cost = "0.00"
	}
	if r.ShopDiscountAmount == "" {
		r.ShopDiscountAmount = "0.00"
	}
	if r.PlatDiscountAmount == "" {
		r.PlatDiscountAmount = "0.00"
	}
	return
}

func NewOrderPreview() (res *OrderPreview) {
	res = &OrderPreview{
		Amount:     "0.00",
		List:       []*PreviewShopItem{},
		PlatCoupon: &mall.CanUseCouponItem{},
	}
	_ = res.Default()
	return
}

func NewPreviewShopItem() (res *PreviewShopItem) {
	res = &PreviewShopItem{
		ShopCoupon: &mall.CanUseCoupon{},
		Products:   PreviewSpuItems{},
	}
	_ = res.Default()
	_ = res.ShopCoupon.Default()
	return
}
