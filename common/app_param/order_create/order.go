package order_create

import (
	"github.com/juetun/base-wrapper/lib/base"
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
	ResultGetInfoByOrder map[string]ResultGetInfoByOrderItem

	ResultGetInfoByOrderItem struct {
		StatusName         string           `json:"status_name"` //订单状态中文描述
		ShopItems          []*OrderShopItem `json:"shop_items"`  //商品列表（按店铺分组）
		UserHid            int64            `json:"user_hid"`
		OrderId            string           `json:"order_id"` //订单ID号
		ProductNum         int64            `json:"product_num"`
		Status             uint8            `json:"status"` //订单状态
		SubStatus          uint8            `json:"sub_status"`
		AddressId          int64            `json:"address_id"`
		PayType            uint8            `json:"pay_type"`
		Amount             string           `json:"amount"` //支付总金额
		PayAmount          string           `json:"pay_amount"`
		TotalPostage       string           `json:"total_postage"`
		ProductAmount      string           `json:"product_amount"`
		ShopDiscountAmount string           `json:"shop_discount_amount"`
		PlatDiscountAmount string           `json:"plat_discount_amount"`
		Mark               string           `json:"mark"`
	}

	OrderShopItem struct {
		ShopId      int64               `json:"shop_id"`
		ShopIcon    string              `json:"shop_icon"`    // 店铺Icon
		ShopName    string              `json:"shop_name"`    // 店铺名称
		ShopType    string              `json:"shop_type"`    // 店铺类型
		Count       int64               `json:"count"`        // 商品总数
		ShopChecked bool                `json:"shop_checked"` // 店铺选择
		TotalAmount string              `json:"total_amount"` // 该订单店铺总的金额
		Products    []*OrderSkuItem     `json:"products"`     // 商品列表
		Delivery    OrderSkuDelivery    `json:"delivery"`     // 邮费信息
		Coupon      OrderShopItemCoupon `json:"coupon"`       // 优惠券信息
		SubOrderId  string              `json:"sub_order_id"` // 子单号
		Mark        string              `json:"mark"`         // 备注
		SortWeight  int64               `json:"-"`            // 排序权重
	}
	OrderShopItemCoupon struct {
		Mark     string `json:"mark"`
		CouponId string `json:"coupon_id"`
	}
	OrderSkuItem struct {
		//Title           string              `json:"title"`
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
