package order_create

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/shopspring/decimal"
	"strconv"
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

	DataItemTag struct {
		Type      string `json:"type"`                //标签类型，可选值为primary success danger warning	默认	default
		Label     string `json:"label"`               //类型名称
		Color     string `json:"color,omitempty"`     //标签颜色
		TextColor string `json:"textColor,omitempty"` //文本颜色，优先级高于color属性	String	white
		Plain     bool   `json:"plain"`               //是否为空心样式	Boolean	false
		Round     bool   `json:"round"`               //是否为圆角样式	Boolean	false
		Mark      bool   `json:"mark"`                //是否为标记样式
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
		ActType            uint8            `json:"act_type"`             //1-首付款 或2-尾款
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
		ActType            uint8            `json:"act_type"`             //1-首付款 或2-尾款
		SortWeight         int64            `json:"-"`                    // 排序权重
	}

	OrderSkuItem struct {
		SkuName         string              `json:"sku_name"`
		SpuId           string              `json:"spu_id"`
		SkuId           string              `json:"sku_id"`        //购物车数据ID
		SkuPic          string              `json:"sku_pic"`       // 图片
		SkuStatus       int8                `json:"sku_status"`    // 商品状态
		StatusName      string              `json:"status_name"`   // 商品状态名称 (已下架)
		SkuSetPrice     string              `json:"sku_set_price"` //sku商品单价价(如果是定金预售，则为定金+尾款金额  )
		Price           string              `json:"price"`         //单价
		Num             int64               `json:"num"`           //商品数量
		PriceCate       uint8               `json:"price_cate"`    //定金类型当前商品类型为定金预售时1-首付款 2-尾款
		PriceCateStr    string              `json:"price_cate_str"`
		PriceCateName   string              `json:"price_cate_name"`
		TotalPrice      string              `json:"total_price"`
		SaleTypeName    string              `json:"sale_type_name"`
		SaleType        uint8               `json:"sale_type"`
		SkuPropertyName string              `json:"sku_property_name"` //SKU属性名
		HaveVideo       bool                `json:"have_video"`        //是否有视频
		Mark            string              `json:"mark"`              //商品说明（如 比着加入有无车时降价多少）
		MarkSystem      string              `json:"mark_system"`       //数据不合法 系统说明(系统使用，记录更详细不合法原因)
		Checked         bool                `json:"checked"`           //是否选中
		ActType         uint8               `json:"act_type"`          //
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

func (r *OrderShopItem) GetProductAmount() (res string, err error) {
	var productAmount = decimal.NewFromInt(0)
	var pAmount decimal.Decimal
	for _, skuItem := range r.Products {
		if skuItem.TotalPrice != "" {
			if pAmount, err = decimal.NewFromString(skuItem.TotalPrice); err != nil {
				return
			}
			productAmount = productAmount.Add(pAmount)
		}
	}
	res = productAmount.StringFixed(2)
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

func (r *ResultGetInfoByOrderItem) DataOrderUpdatePayType(payTypeString string) (err error) {
	if payTypeString != "" {
		var payType uint64
		if payType, err = strconv.ParseUint(payTypeString, 10, 8); err != nil {
			return
		}
		r.PayType = uint8(payType)
	}
	return
}
