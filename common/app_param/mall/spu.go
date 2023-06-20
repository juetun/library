package mall

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/juetun/library/common/app_param/upload_operate"
)

type (
	SpuData struct {
		Shop            *models.Shop                `json:"shop"`
		ShopExt         *models.ShopExt             `json:"shop_ext"`
		ShopNotice      *models.ShopNotice          `json:"shop_notice"`
		Brand           *models.Brand               `json:"brand"`
		SKu             []*SpuSkuItem               `json:"sku"`
		SKuProperty     SkuPropertyList             `json:"s_ku_property"`
		Product         *models.Product             `json:"product"` // 商品内容
		Video           *upload_operate.UploadVideo `json:"video"`
		ProductDesc     *models.ProductDesc         `json:"product_desc"`     //商品详情
		FreightTemplate *models.FreightTemplate     `json:"freight_template"` //运费模板
		SpuId           string                      `json:"spu_id"`
	}
	SkuPropertyList []*SkuProperty
	SkuProperty     struct {
		models.SkuPropertyRelate
		models.SkuProperty
	}
	Property struct {
		SkuId string `json:"sku_id"`
		*models.SkuProperty
		models.SkuPropertyRelate
		*models.Sku
	}
	SpuSkuItem struct {
		Sku         models.Sku                `json:"sku"`
		SkuProperty []*Property               `json:"sku_property"` //sku属性
		SkuRelate   *models.SkuPropertyRelate `json:"sku_relate"`
		//SkuStock    *models.SkuStock       `json:"sku_stock"`    //库存数据
		Stock uint64            `json:"stock"`
		Gifts []*models.SkuGift `json:"gifts"` //赠品数据
	}
)

//根据SpuId获取Spu相关信息（sku spu spu详情 赠品）
//Param skuDataType 获取SKU相关信息的类型取值范围
//				SpuDataShop
//				SpuDataBrand
// 				SpuDataTypeProductDesc
// 				SpuDataTypeSKus
// 				SpuDataTypeFreightTemplate
// Return  map[int64]*wrappers.SpuData
func GetSpusWithDataTypeBySpuId(spuIds *base.ArgGetByStringIds, dataTypes ...string) (res map[string]*SpuData, err error) {

	return
}
