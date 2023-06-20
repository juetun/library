package mall

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/juetun/library/common/app_param/upload_operate"
	"net/http"
	"net/url"
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
	ArgGetSpuDataWithSpuId struct {
		ArgGetByStringIds *base.ArgGetByStringIds `json:"spu_id_info" form:"spu_id_info"`
		DataTypes         []string                `json:"data_types" form:"data_types"`
	}
)

func (r *ArgGetSpuDataWithSpuId) Default(c *base.Context) (err error) {

	return
}

func (r *ArgGetSpuDataWithSpuId) GetJsonByte() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

//根据SpuId获取Spu相关信息（sku spu spu详情 赠品）
//Param skuDataType 获取SKU相关信息的类型取值范围
//				SpuDataShop
//				SpuDataBrand
// 				SpuDataTypeProductDesc
// 				SpuDataTypeSKus
// 				SpuDataTypeFreightTemplate
// Return  map[int64]*wrappers.SpuData
func GetSpusWithDataTypeBySpuId(ctx *base.Context, spuIds *ArgGetSpuDataWithSpuId) (res map[string]*SpuData, err error) {
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameMall,
		URI:         "/product/get_spu_map_by_ids",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}
	if ctx.GinContext != nil {
		params.Header.Set(app_obj.HttpHeaderInfo, ctx.GinContext.GetHeader(app_obj.HttpHeaderInfo))
	}
	if params.BodyJson, err = spuIds.GetJsonByte(); err != nil {
		return
	}
	req := rpc.NewHttpRpc(&params).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.Body; len(body) == 0 {
		return
	}

	var resResult struct {
		Code int                 `json:"code"`
		Data map[string]*SpuData `json:"data"`
		Msg  string              `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}
	res = resResult.Data
	return
}
