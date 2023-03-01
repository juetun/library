package mall

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/freight"
	"github.com/juetun/library/common/app_param/mall/models"
)

type (
	SkuData struct {
		Shop            *models.Shop              `json:"shop,omitempty"`
		ShopExt         *models.ShopExt           `json:"shop_ext,omitempty"`
		SkuRelate       *models.SkuPropertyRelate `json:"sku_relate"`
		ShopNotice      *models.ShopNotice        `json:"shop_notice,omitempty"`
		SKu             *models.Sku               `json:"sku,omitempty"`
		SkuStock        int64                     `json:"sku_stock,omitempty"`
		Product         *models.Product           `json:"product,omitempty"`          // 商品内容
		Brand           *models.Brand             `json:"brand,omitempty"`            // 品牌
		ProductDesc     *models.ProductDesc       `json:"product_desc,omitempty"`     // 商品详情
		Gift            []*SkuGiftItem            `json:"gift,omitempty"`             // 赠品
		FreightTemplate *freight.TemplateFreight  `json:"freight_template,omitempty"` // 运费模板
	}

	SkuGiftItem struct {
		SkuId              string           `json:"sku_id"`
		Thumbnail          string           `json:"thumbnail"`
		Image              string           `json:"image"`
		Video              string           `json:"video"`
		UserHid            int64            `json:"user_hid"`
		ShopId             int64            `json:"shop_id"`
		ProductId          string           `json:"product_id"`
		SkuStatus          int8             `json:"sku_status"`
		SkuName            string           `json:"sku_name"`
		Weight             string           `json:"weight"`
		MaxLimitNum        string           `json:"max_limit_num"`
		MinLimitNum        string           `json:"min_limit_num"`
		Price              string           `json:"price"`
		PriceCost          string           `json:"price_cost"`
		Stock              int64            `json:"stock"`
		ShopSaleCode       string           `json:"shop_sale_code"`
		ProvideSaleCode    string           `json:"provide_sale_code"`
		FreightTemplate    int64            `json:"freight_template"`
		SaleNum            int              `json:"sale_num"`
		SaleOnlineTime     base.TimeNormal  `json:"sale_online_time"`
		SaleOverTime       *base.TimeNormal `json:"sale_over_time"`
		FinalStartTime     base.TimeNormal  `json:"final_start_time"`
		FinalOverTime      base.TimeNormal  `json:"final_over_time"`
		SalesTaxRate       string           `json:"sales_tax_rate"`
		SalesTaxRateValue  string           `json:"sales_tax_rate_value"`
		FlagTester         uint8            `json:"flag_tester"`
		CreatedAt          base.TimeNormal  `json:"created_at"`
		UpdatedAt          base.TimeNormal  `json:"updated_at"`
		SkuGiftEffectSTime base.TimeNormal  `json:"sku_gift_effect_s_time"`
		SkuGiftEffectOTime base.TimeNormal  `json:"sku_gift_effect_o_time"`
		SkuGiftStatus      uint8            `json:"sku_gift_status"`
		SkuGiftCreatedAt   base.TimeNormal  `json:"sku_gift_created_at"`
		SkuGiftUpdatedAt   base.TimeNormal  `json:"sku_gift_updated_at"`
		SkuGiftId          int64            `json:"sku_gift_id"`
	}
)
