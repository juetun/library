package order_create

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall"
	"github.com/juetun/library/common/app_param/mall/freight"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

type (
	SkuRefund struct {
		*OrderSkuPk
		Pk       string          `json:"pk"`
		SkuPrice decimal.Decimal `json:"sku_price"` //商品单价
		SpuPlat  decimal.Decimal `json:"spu_plat"`  //商品维度每个Sku平台优惠
		SpuShop  decimal.Decimal `json:"spu_shop"`  //商品维度每个Sku店铺优惠
		ShopPlat decimal.Decimal `json:"shop_plat"` //店铺维度平台每个Sku优惠
		ShopShop decimal.Decimal `json:"shop_shop"` //店铺维度商家每个Sku优惠
		Plat     decimal.Decimal `json:"plat"`      //整个订单维度平台每个Sku优惠
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
		PayChargeDesc      string                             `json:"pay_charge_desc"`       //支付手续费说明
		PayNumber          string                             `json:"pay_number,omitempty"`  // 支付流水号
		Token              string                             `json:"token,omitempty"`       // 购物车全站唯一的key
		Count              int64                              `json:"count"`                 // 商品总数
		HaveError          bool                               `json:"have_error"`            //是否有错误
		ErrorMessage       string                             `json:"error_message"`         //错误信息
		OrderId            string                             `json:"order_id"`              //订单号
		EmsAddress         *freight.ResultGetByAddressIdsItem `json:"ems_address,omitempty"` //收货地址信息
		ShopDiscountAmount string                             `json:"shop_discount_amount"`  // 店铺优惠总金额 (店铺维度优惠+SPU维度优惠)
		PlatDiscountAmount string                             `json:"plat_discount_amount"`  // 平台优惠金额 (店铺维度优惠+SPU维度优惠)
		DeductionAmount    string                             `json:"deduction_amount"`      //代金券抵扣金额
		PlatCoupon         *mall.CanUseCouponItem             `json:"plat_coupon,omitempty"` //平台券信息 （店铺维度优惠）
		List               PreviewShopItems                   `json:"list"`                  // SKU
		ActType            uint8                              `json:"act_type"`              //1-首款（默认） 2-尾款（定金预售）
		DisabledPay        bool                               `json:"disabled_pay"`          //是否能够支付 true-不能支付 false-能支付
	}
	PreviewShopItems []*PreviewShopItem
	PreviewShopItem  struct {
		ShopId             int64              `json:"shop_id"`
		ShopIcon           string             `json:"shop_icon"`             // 店铺Icon
		ShopName           string             `json:"shop_name"`             // 店铺名称
		ShopHref           interface{}        `json:"shop_href"`             // 店铺链接地址
		ShopType           string             `json:"shop_type"`             // 店铺类型
		Count              int64              `json:"count"`                 // 商品总数
		PayType            uint8              `json:"pay_type"`              //支付方式（支付渠道）
		PayCharge          string             `json:"pay_charge"`            // 支付手续费
		DeductionAmount    string             `json:"deduction_amount"`      // 现金抵扣券
		ProductAmount      string             `json:"product_amount"`        // 商品金额
		TotalAmount        string             `json:"total_amount"`          // 该订单店铺总的金额
		Delivery           OrderSkuDelivery   `json:"delivery"`              // 邮费信息
		Mark               string             `json:"mark"`                  // 用户备注
		SubOrderId         string             `json:"sub_order_id"`          // 子单号
		ShopDiscountAmount string             `json:"shop_discount_amount"`  // 店铺优惠总金额 (店铺维度优惠+SPU维度优惠)
		PlatDiscountAmount string             `json:"plat_discount_amount"`  // 平台优惠金额 (店铺维度优惠+SPU维度优惠)
		SettlementAmount   string             `json:"settlement_amount"`     //结算金额
		ShopCoupon         *mall.CanUseCoupon `json:"shop_coupon,omitempty"` // 店铺券信息
		Products           PreviewSpuItems    `json:"products"`              // 商品列表
		ActType            uint8              `json:"act_type"`              //1-首款（默认） 2-尾款（定金预售）
		SortCreateTime     time.Time          `json:"-"`
		SortWeight         int64              `json:"-"` //排序权重
	}
	PreviewSpuItems []*PreviewSpuItem
	PreviewSkuItems []*PreviewSkuItem
	PreviewSpuItem  struct {
		SpuId          string             `json:"spu_id"`
		Skus           PreviewSkuItems    `json:"skus"`                 // SKU信息
		SpuCoupon      *mall.CanUseCoupon `json:"spu_coupon,omitempty"` // 店铺券信息
		Count          int64              `json:"count"`                // 商品总数
		SpuTotalAmount string             `json:"spu_total_amount"`     // 商品总额
		TotalAmount    string             `json:"total_amount"`         // 该订单店铺总的金额
		SaleType       uint8              `json:"sale_type"`
		SaleTypeName   string             `json:"sale_type_name"`
	}
	PreviewSkuItem struct {
		app_param.ArgOrderFromCartItem
		//Title          string          `json:"title"`
		SkuName         string            `json:"sku_name"`
		SkuPropertyName string            `json:"sku_property_name"`
		SkuId           string            `json:"sku_id"`      //购物车数据ID
		SkuPic          string            `json:"sku_pic"`     // 图片
		SkuStatus       int8              `json:"sku_status"`  // 商品状态
		StatusName      string            `json:"status_name"` // 商品状态名称 (已下架)
		TotalPrice      string            `json:"total_price"`
		SaleType        uint8             `json:"sale_type"`
		SaleTypeName    string            `json:"sale_type_name"`
		HaveVideo       bool              `json:"have_video"`   //是否有视频
		Mark            string            `json:"mark"`         //商品说明（如 比着加入有无车时降价多少）
		TimeMark        string            `json:"time_mark"`    //时间备注
		MarkSystem      string            `json:"mark_system"`  //数据不合法 系统说明(系统使用，记录更详细不合法原因)
		NotCanPay       bool              `json:"not_can_pay"`  //当前数据是否能够支付
		Invalidation    bool              `json:"invalidation"` //是否失效 true-已失效 false-未失效
		Checked         bool              `json:"checked"`      //是否选中
		SpecialTags     []*DataItemTag    `json:"special_tags"`
		TitleTags       []*models.PageTag `json:"title_tags"`
		ActType         uint8             `json:"act_type"`
		PriceLabel      string            `json:"price_label"`
		SortCreateTime  base.TimeNormal   `json:"-"`
		SortWeight      int64             `json:"-"`
		// 单一SKU的总价格
	}
	OrderSkuPk struct {
		SubOrderId string `json:"sub_order_id"`
		SpuId      string `json:"spu_id"`
		SkuId      string `json:"sku_id"`
	}
)

func newSkuRefund() (res *SkuRefund) {
	res = &SkuRefund{
		OrderSkuPk: &OrderSkuPk{},
		SkuPrice:   decimal.NewFromInt(0),
		SpuPlat:    decimal.NewFromInt(0),
		SpuShop:    decimal.NewFromInt(0),
		ShopPlat:   decimal.NewFromInt(0),
		ShopShop:   decimal.NewFromInt(0),
		Plat:       decimal.NewFromInt(0),
	}
	return
}

//计算每个SKU 需要的退款的金额
func (r *SkuRefund) calRefund() (res string) {
	res = r.SkuPrice.
		Sub(r.SpuPlat).
		Sub(r.SpuShop).
		Sub(r.ShopPlat).
		Sub(r.ShopShop).
		Sub(r.Plat).
		StringFixed(2)
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

// 计算全部商品扣减的优惠金额
// Param skuSetPrice 商品单价
// Param payAmount 订单支付金额
// Param couponAmount 优惠券金额
func getSkuPlatCoupon(skuSetPrice string, payAmount decimal.Decimal, couponAmount decimal.Decimal) (res decimal.Decimal, err error) {
	if skuSetPrice == "" {
		skuSetPrice = "0.00"
	}
	if payAmount.Equal(decimal.NewFromInt(0)) {
		err = fmt.Errorf("支付金额数据异常")
		return
	}
	var skuSetPriceDe decimal.Decimal
	if skuSetPriceDe, err = decimal.NewFromString(skuSetPrice); err != nil {
		return
	}
	//优惠金额/支付金额*商品单价
	res = couponAmount.Div(payAmount).Mul(skuSetPriceDe)
	return
}

func (r *OrderPreview) GetMapRefundSkuData() (mapRefundSkuData map[string]string, err error) {
	mapRefundSkuData = make(map[string]string, len(r.List)*20)
	var (
		skuRefund      *SkuRefund
		mapPlatSpuSku  map[string]decimal.Decimal
		mapShopSpuSku  map[string]decimal.Decimal
		mapPlatShopSku map[string]decimal.Decimal
		mapShopShopSku map[string]decimal.Decimal
		mapPlatCommon  map[string]decimal.Decimal
	)
	if mapPlatCommon, err = r.GetMapRefundCommonSkuData(); err != nil {
		return
	}
	//sku退款金额 等于 商品价格-spu平台优惠-spu商家优惠-店铺平台优惠-商家优惠-全平台优惠券分摊到商品的优惠
	for _, previewShopItems := range r.List {

		//计算SPU中每个SKU的优惠值
		if mapPlatShopSku, mapShopShopSku, err = previewShopItems.GetMapRefundSkuData(); err != nil {
			return
		}

		for _, spuInfo := range previewShopItems.Products {

			//计算SPU中每个SKU的优惠值
			if mapPlatSpuSku, mapShopSpuSku, err = spuInfo.GetMapRefundSkuData(previewShopItems.SubOrderId); err != nil {
				return
			}

			for _, sku := range spuInfo.Skus {
				skuRefund = newSkuRefund()
				skuRefund.SubOrderId = previewShopItems.SubOrderId
				skuRefund.SpuId = sku.SpuId
				skuRefund.SkuId = sku.SkuId
				skuRefund.InitPk()
				if sku.SkuSetPrice != "" {
					if skuRefund.SkuPrice, err = decimal.NewFromString(sku.SkuSetPrice); err != nil {
						return
					}
				}
				skuRefund.SetPlatCommonCoupon(mapPlatCommon)
				skuRefund.SetWithShopCoupon(mapPlatShopSku, mapShopShopSku)
				skuRefund.SetWithSpuCoupon(mapPlatSpuSku, mapShopSpuSku)
				mapRefundSkuData[skuRefund.Pk] = skuRefund.calRefund()

			}
		}
	}
	return
}

func (r *PreviewShopItem) getCouponAmount() (platCouponAmount, shopCouponAmount decimal.Decimal, err error) {
	platCouponAmount = decimal.NewFromInt(0)
	shopCouponAmount = decimal.NewFromInt(0)
	if r.ShopCoupon == nil {
		return
	}
	if r.ShopCoupon.Plat != nil && r.ShopCoupon.Plat.CurrentUse.CouponID != "" {
		if platCouponAmount, err = decimal.NewFromString(r.ShopCoupon.Plat.CurrentUse.Decr); err != nil {
			return
		}
	}
	if r.ShopCoupon.Shop != nil && r.ShopCoupon.Shop.CurrentUse.CouponID != "" {
		if shopCouponAmount, err = decimal.NewFromString(r.ShopCoupon.Shop.CurrentUse.Decr); err != nil {
			return
		}
	}
	return
}

func (r *PreviewSpuItem) getCouponAmount() (platCouponAmount, shopCouponAmount decimal.Decimal, err error) {
	platCouponAmount = decimal.NewFromInt(0)
	shopCouponAmount = decimal.NewFromInt(0)
	if r.SpuCoupon == nil {
		return
	}
	if r.SpuCoupon.Plat != nil && r.SpuCoupon.Plat.CurrentUse.CouponID != "" {
		if platCouponAmount, err = decimal.NewFromString(r.SpuCoupon.Plat.CurrentUse.Decr); err != nil {
			return
		}
	}
	if r.SpuCoupon.Shop != nil && r.SpuCoupon.Shop.CurrentUse.CouponID != "" {
		if shopCouponAmount, err = decimal.NewFromString(r.SpuCoupon.Shop.CurrentUse.Decr); err != nil {
			return
		}
	}
	return
}

//按照SKUID 分别
func (r *PreviewSpuItem) GetMapRefundSkuData(subOrderId string) (mapPlatSpuSku, mapShopSpuSku map[string]decimal.Decimal, err error) {
	var l = len(r.Skus)
	mapPlatSpuSku = make(map[string]decimal.Decimal, l)
	mapShopSpuSku = make(map[string]decimal.Decimal, l)
	if r.SpuCoupon == nil {
		return
	}
	var (
		pk                                                     string
		totalAmountDecimal, platCouponAmount, shopCouponAmount decimal.Decimal
	)

	if r.SpuTotalAmount == "" {
		r.SpuTotalAmount = "0.00"
	}

	if r.SpuTotalAmount == "0.00" {
		r.SpuTotalAmount = r.TotalAmount
	}
	if totalAmountDecimal, err = decimal.NewFromString(r.SpuTotalAmount); err != nil {
		return
	}
	if platCouponAmount, shopCouponAmount, err = r.getCouponAmount(); err != nil {
		return
	}
	for _, sku := range r.Skus {
		pk = (&OrderSkuPk{SubOrderId: subOrderId, SpuId: sku.SpuId, SkuId: sku.SkuId}).GetPk()
		if mapPlatSpuSku[pk], err = getSkuPlatCoupon(sku.SkuSetPrice, totalAmountDecimal, platCouponAmount); err != nil {
			return
		}

		if mapShopSpuSku[pk], err = getSkuPlatCoupon(sku.SkuSetPrice, totalAmountDecimal, shopCouponAmount); err != nil {
			return
		}
	}

	return
}

func (r *PreviewShopItem) GetMapRefundSkuData() (mapPlatShopSku map[string]decimal.Decimal, mapShopShopSku map[string]decimal.Decimal, err error) {
	var l = len(r.Products) * 10
	mapPlatShopSku = make(map[string]decimal.Decimal, l)
	mapShopShopSku = make(map[string]decimal.Decimal, l)
	var (
		pk                                                     string
		totalAmountDecimal, platCouponAmount, shopCouponAmount decimal.Decimal
	)
	if totalAmountDecimal, err = decimal.NewFromString(r.ProductAmount); err != nil {
		return
	}
	if platCouponAmount, shopCouponAmount, err = r.getCouponAmount(); err != nil {
		return
	}
	for _, spu := range r.Products {
		for _, sku := range spu.Skus {
			pk = (&OrderSkuPk{SubOrderId: r.SubOrderId, SpuId: sku.SpuId, SkuId: sku.SkuId}).GetPk()
			if mapPlatShopSku[pk], err = getSkuPlatCoupon(sku.SkuSetPrice, totalAmountDecimal, platCouponAmount); err != nil {
				return
			}

			if mapShopShopSku[pk], err = getSkuPlatCoupon(sku.SkuSetPrice, totalAmountDecimal, shopCouponAmount); err != nil {
				return
			}
		}
	}
	return
}

//计算平台券退款金额
func (r *OrderPreview) GetMapRefundCommonSkuData() (mapPlatCommon map[string]decimal.Decimal, err error) {
	mapPlatCommon = make(map[string]decimal.Decimal, len(r.List)*20)
	if r.PlatCoupon == nil || r.PlatCoupon.CurrentUse.CouponID == "" {
		return
	}

	if r.PlatCoupon.CurrentUse.Decr == "" {
		r.PlatCoupon.CurrentUse.Decr = "0.00"
	}

	var (
		couponAmount  decimal.Decimal
		productAmount decimal.Decimal
		pk            string
	)
	if couponAmount, err = decimal.NewFromString(r.PlatCoupon.CurrentUse.Decr); err != nil {
		return
	}

	if productAmount, err = decimal.NewFromString(r.ProductAmount); err != nil {
		return
	}
	for _, previewShopItems := range r.List {
		for _, spuInfo := range previewShopItems.Products {
			for _, sku := range spuInfo.Skus {
				pk = (&OrderSkuPk{SubOrderId: previewShopItems.SubOrderId, SpuId: sku.SpuId, SkuId: sku.SkuId}).GetPk()

				if mapPlatCommon[pk], err = getSkuPlatCoupon(sku.SkuSetPrice, productAmount, couponAmount); err != nil {
					return
				}
			}
		}
	}

	return
}

func (r *SkuRefund) SetPlatCommonCoupon(mapPlatCommon map[string]decimal.Decimal) {
	var (
		tmp decimal.Decimal
		ok  bool
	)
	if tmp, ok = mapPlatCommon[r.Pk]; ok {
		r.Plat = tmp
	}

	return
}

func (r *SkuRefund) SetWithShopCoupon(mapPlatShopSku map[string]decimal.Decimal, mapShopShopSku map[string]decimal.Decimal) {
	var (
		tmp decimal.Decimal
		ok  bool
	)
	if tmp, ok = mapPlatShopSku[r.Pk]; ok {
		r.ShopPlat = tmp
	}
	if tmp, ok = mapShopShopSku[r.Pk]; ok {
		r.ShopShop = tmp
	}
	return
}

func (r *SkuRefund) SetWithSpuCoupon(mapPlatSpuSku map[string]decimal.Decimal, mapShopSpuSku map[string]decimal.Decimal) {
	var (
		tmp decimal.Decimal
		ok  bool
	)
	if tmp, ok = mapPlatSpuSku[r.Pk]; ok {
		r.SpuPlat = tmp
	}
	if tmp, ok = mapShopSpuSku[r.Pk]; ok {
		r.SpuShop = tmp
	}
	return
}

func (r *SkuRefund) InitPk() {
	if r.OrderSkuPk == nil {
		r.OrderSkuPk = &OrderSkuPk{}
	}
	r.Pk = r.OrderSkuPk.GetPk()
}

func (r *OrderSkuPk) GetPk() (res string) {
	res = fmt.Sprintf("%v_%v_%v", r.SubOrderId, r.SpuId, r.SkuId)
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

//初始化支付方式
func (r *OrderPreview) InitPayTypeOption(info *common.HeaderInfo, payTypes ...string) (err error) {
	r.PayTypeOpt = SliceOrderPayTypeUse

	switch info.HTerminal {
	case app_param.TerminalMina: //如果是微信小程序
		switch info.HChannel {
		case "weixin": //如果是小程序微信使用
			r.getWeiXinMinaOpt()
		case "alipay":
			r.getAliPayMinaOpt() //支付宝小程序
		}
	case app_param.TerminalH5:
		r.getH5Opt()
	case app_param.TerminalAndroid, app_param.TerminalIos: //手机APP当前只支持
		r.getAppOpt()
	}

	if len(payTypes) > 0 && payTypes[0] != "" && payTypes[0] != "0" {
		r.PayType = payTypes[0]
		return
	}
	if len(r.PayTypeOpt) > 0 {
		r.PayType = fmt.Sprintf("%v", r.PayTypeOpt[0].Value)
	}
	return
}

//手机APP当前只支持支付宝
func (r *OrderPreview) getAppOpt() {
	mapPay, _ := SliceOrderPayType.GetMapAsKeyUint8()
	r.PayTypeOpt = base.ModelItemOptions{
		{
			Label: mapPay[OrderPayTypeAliPay],
			Value: OrderPayTypeAliPay,
		},
	}
	return
}

//H5当前只支持支付宝
func (r *OrderPreview) getH5Opt() {
	mapPay, _ := SliceOrderPayType.GetMapAsKeyUint8()
	r.PayTypeOpt = base.ModelItemOptions{
		{
			Label: mapPay[OrderPayTypeAliPay],
			Value: OrderPayTypeAliPay,
		},
	}
	return
}

//支付宝小程序小程序
func (r *OrderPreview) getAliPayMinaOpt() {
	mapPay, _ := SliceOrderPayType.GetMapAsKeyUint8()
	r.PayTypeOpt = base.ModelItemOptions{
		{
			Label: mapPay[OrderPayTypeWeiXin],
			Value: OrderPayTypeWeiXin,
		},
	}
	return
}

//微信小程序
func (r *OrderPreview) getWeiXinMinaOpt() {
	mapPay, _ := SliceOrderPayType.GetMapAsKeyUint8()
	r.PayTypeOpt = base.ModelItemOptions{
		{
			Label: mapPay[OrderPayTypeWeiXin],
			Value: OrderPayTypeWeiXin,
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
	if r.DeductionAmount == "" {
		r.DeductionAmount = "0.00"
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

//初始化支付手续费金额
func (r *OrderPreview) InitPayCharge() (err error) {
	var (
		payTypeNum uint64
		rabat      string
	)
	if payTypeNum, err = strconv.ParseUint(r.PayType, 10, 8); err != nil {
		return
	}
	if r.PayCharge, r.Amount, rabat, err = GetByPayTypeAndAmount(uint8(payTypeNum), r.Amount); err != nil {
		return
	}

	if rabat != "" {
		payTypeNumber, _ := strconv.ParseUint(r.PayType, 10, 8)
		r.PayChargeDesc = fmt.Sprintf("您当前使用的是%v支付,平台收取手续费为:%v", ParsePayType(uint8(payTypeNumber)), rabat)
	}
	return
}

//计算店铺结算金额
func (r *PreviewShopItem) CalSettlementAmount() (err error) {
	var (
		totalAmount, payCharge decimal.Decimal
	)
	if totalAmount, err = decimal.NewFromString(r.TotalAmount); err != nil {
		return
	}
	if payCharge, err = decimal.NewFromString(r.PayCharge); err != nil {
		return
	}
	r.SettlementAmount = totalAmount.Sub(payCharge).StringFixed(2)
	return
}

//计算优惠券
func (r *PreviewShopItem) CalCouponAndPayCharge() (err error) {
	var (
		platDecimal = decimal.NewFromInt(0)
		shopDecimal = decimal.NewFromInt(0) //店铺优惠券优惠的金额

	)

	if r.ShopCoupon != nil && r.ShopCoupon.Shop != nil && r.ShopCoupon.Shop.CurrentUse.ID > 0 && r.ShopCoupon.Shop.CurrentUse.Decr != "" {
		var spDecr decimal.Decimal
		spDecr, err = decimal.NewFromString(r.ShopCoupon.Shop.CurrentUse.Decr)
		shopDecimal = shopDecimal.Add(spDecr)
	}

	if r.ShopCoupon != nil && r.ShopCoupon.Plat != nil && r.ShopCoupon.Plat.CurrentUse.ID > 0 && r.ShopCoupon.Plat.CurrentUse.Decr != "" {
		var ptDecr decimal.Decimal
		ptDecr, err = decimal.NewFromString(r.ShopCoupon.Plat.CurrentUse.Decr)
		platDecimal = platDecimal.Add(ptDecr)
	}

	var spuShopDecr, spuPlatDecr decimal.Decimal
	if spuShopDecr, spuPlatDecr, err = r.getSpuDecrValue(); err != nil {
		return
	}

	shopDecimal = shopDecimal.Add(spuShopDecr)
	platDecimal = platDecimal.Add(spuPlatDecr)

	r.ShopDiscountAmount = shopDecimal.StringFixed(2)
	r.PlatDiscountAmount = platDecimal.StringFixed(2)

	//计算平台该收商家的手续费
	if err = r.calPayCharge(shopDecimal); err != nil {
		return
	}
	return
}

func (r *PreviewShopItem) calPayCharge(shopDecimal decimal.Decimal) (err error) {
	var (
		product, deductionAmount decimal.Decimal
		shopValue                string
	)
	if product, err = r.getProductAmount(); err != nil {
		return
	}
	if deductionAmount, err = r.getDeductionAmount(); err != nil {
		return
	}
	shopValue = product.Sub(deductionAmount).Sub(shopDecimal).StringFixed(2)
	if r.PayCharge, _, _, err = GetByPayTypeAndAmount(r.PayType, shopValue); err != nil {
		return
	}
	return
}

func (r *PreviewShopItem) getSpuDecrValue() (shopDecr, platDecr decimal.Decimal, err error) {
	shopDecr = decimal.NewFromInt(0)
	platDecr = decimal.NewFromInt(0)
	var (
		shopDecrVal, platDecrVal decimal.Decimal
	)
	for _, spuInfo := range r.Products {
		if spuInfo.SpuCoupon == nil {
			continue
		}

		if spuInfo.SpuCoupon.Plat != nil && spuInfo.SpuCoupon.Plat.CurrentUse.ID > 0 && spuInfo.SpuCoupon.Plat.CurrentUse.Decr != "" {
			if platDecrVal, err = decimal.NewFromString(spuInfo.SpuCoupon.Plat.CurrentUse.Decr); err != nil {
				return
			}
			platDecr = platDecr.Add(platDecrVal)
		}
		if spuInfo.SpuCoupon.Shop != nil && spuInfo.SpuCoupon.Shop.CurrentUse.ID > 0 && spuInfo.SpuCoupon.Shop.CurrentUse.Decr != "" {
			if shopDecrVal, err = decimal.NewFromString(spuInfo.SpuCoupon.Shop.CurrentUse.Decr); err != nil {
				return
			}
			shopDecr = shopDecr.Add(shopDecrVal)
		}
	}
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
			"SpuDecr":       r.getSpuDecr,       //spu优惠券扣减金额 优惠券+抵扣券（代金券）
			"ShopDecr":      r.getShopDecr,      //店铺优惠券抵扣金额 优惠券+抵扣券（代金券）
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
		if spuInfo.SpuCoupon.Plat != nil && spuInfo.SpuCoupon.Plat.CurrentUse.ID > 0 && spuInfo.SpuCoupon.Plat.CurrentUse.Decr != "" {
			if platDecr, err = decimal.NewFromString(spuInfo.SpuCoupon.Plat.CurrentUse.Decr); err != nil {
				return
			}
		}
		if spuInfo.SpuCoupon.Shop != nil && spuInfo.SpuCoupon.Shop.CurrentUse.ID > 0 && spuInfo.SpuCoupon.Shop.CurrentUse.Decr != "" {
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
		shopDecr        = decimal.NewFromInt(0)
		platDecr        = decimal.NewFromInt(0)
		deductionAmount = decimal.NewFromInt(0)
	)

	if r.ShopCoupon != nil && r.ShopCoupon.Shop != nil && r.ShopCoupon.Shop.CurrentUse.ID > 0 && r.ShopCoupon.Shop.CurrentUse.Decr != "" {
		if shopDecr, err = decimal.NewFromString(r.ShopCoupon.Shop.CurrentUse.Decr); err != nil {
			return
		}
	}

	if r.ShopCoupon != nil && r.ShopCoupon.Plat != nil && r.ShopCoupon.Plat.CurrentUse.ID > 0 && r.ShopCoupon.Plat.CurrentUse.Decr != "" {
		if platDecr, err = decimal.NewFromString(r.ShopCoupon.Plat.CurrentUse.Decr); err != nil {
			return
		}
	}

	if r.DeductionAmount != "" {
		if deductionAmount, err = decimal.NewFromString(r.DeductionAmount); err != nil {
			return
		}
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
	if r.ProductAmount, err = orderShopItem.GetProductAmount(); err != nil {
		return
	}
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
	r.SkuSetPrice = orderSkuItem.SkuSetPrice
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
	if r.DeductionAmount == "" {
		r.DeductionAmount = "0.00"
	}
	if r.SettlementAmount == "" {
		r.SettlementAmount = "0.00"
	}
	return
}
