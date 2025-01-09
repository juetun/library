package mall

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/juetun/library/common/app_param/upload_operate"
	"github.com/juetun/library/common/recommend"
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
	ArgSpuDetail struct {
		common.HeaderInfo
		app_param.RequestUser
		SkuId string `json:"sku_id" form:"sku_id"`
		SpuId string `json:"spu_id" form:"spu_id"`
		base.GetDataTypeCommon
		TimeNow       base.TimeNormal `json:"-" form:"-"`
		PriceCategory string          `json:"-" form:"-"`
	}

	ResultSpuDetail struct {
		SpuId            string                   `json:"spu_id"`
		Title            string                   `json:"title"`
		Thumbnail        string                   `json:"thumbnail"`
		ThumbnailUrl     string                   `json:"thumbnail_url"`
		UserHid          int64                    `json:"user_hid"`           //当前访问的用户ID
		Media            []*ResultSpuDetailMedia  `json:"media"`              //多媒体 图片 和视频
		PreTags          []*recommend.DataItemTag `json:"pre_tags,omitempty"` //前缀标签
		BrandId          int64                    `json:"brand_id"`           //品牌ID
		Notice           *DetailNotice            `json:"notice,omitempty"`
		BrandName        string                   `json:"brand_name"` //品牌名称
		ShopId           int64                    `json:"shop_id"`
		Status           int8                     `json:"status"`
		StatusName       string                   `json:"status_name"`
		SubTitle         string                   `json:"sub_title"`
		Price            string                   `json:"price"`      //价格
		PriceCost        string                   `json:"price_cost"` //划线价
		TagIdsArray      []int64                  `json:"tag_ids_array"`
		ServiceIds       string                   `json:"service_ids"`
		Keywords         string                   `json:"keywords"`
		FreightType      uint8                    `json:"freight_type"`
		FreightTemplate  int64                    `json:"freight_template"`
		FreightTitle     string                   `json:"freight_title"` //运费模板名称
		TotalStock       int64                    `json:"total_stock"`   // 总库存
		CategoryId       int64                    `json:"category_id"`
		PreheatTimeStart string                   `json:"preheat_time_start"`      //预热开始时间
		PreheatTimeOver  string                   `json:"preheat_time_over"`       //预热结束时间
		SaleType         uint8                    `json:"sale_type"`               //销售类型（定金预售、预售、普通商品）
		PullOnTime       string                   `json:"pull_on_time"`            //上架时间
		PullOffTime      string                   `json:"pull_off_time"`           //下架时间
		DownPayInfo      *DownPayInfo             `json:"down_pay_info,omitempty"` //定金预售信息
		PreSale          *PreSaleInfo             `json:"pre_sale,omitempty"`      //预售数据
		CanBuy           bool                     `json:"can_buy"`                 //是否能够购买
		ShopManager      bool                     `json:"shop_manager"`            //是否店铺管理员
		ShowError        bool                     `json:"show_error"`              //是否展示错误提示，不显示商品其他内容
		Mark             string                   `json:"mark"`                    //备注
		SaleCountShow    uint8                    `json:"sale_count_show"`
		Relate           *DetailRelate            `json:"relate,omitempty"`
		Shop             *DetailShop              `json:"shop,omitempty"`
		//Comment          *mall_comment.OrderComment         `json:"comment"`  //评论信息
		SkuNum      int                       `json:"sku_num"`  //SKU数量 当只有一个SKU时 可用不用选择
		CartNum     int64                     `json:"cart_num"` //购物车商品数量
		SettleType  uint8                     `json:"settle_type"`
		FlagTester  uint8                     `json:"flag_tester"`
		Description string                    `json:"description"`
		Goods       *ResultDetailSkuInfoGoods `json:"goods"`
		PriceMark   string                    `json:"price_mark"` //价格说明
	}
	DetailNotice struct {
		ShopNoticeUrl string `json:"shop_notice_url"` //店铺公告
		PlatNoticeUrl string `json:"plat_notice_url"` //平台公告
	}
	ResultSpuDetailMedia struct {
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}
	DetailShop struct {
		ShopId      int64       `json:"shop_id"`
		ShopName    string      `json:"shop_name"`
		ShopLogoUrl string      `json:"shop_logo_url"`
		ShopHref    interface{} `json:"shop_href"`
	}
	DetailRelate struct {
		RelateType      uint8  `json:"relate_type"`
		RelateItemId    string `json:"relate_item_id"`
		RelateBuyCount  int64  `json:"relate_buy_count"`
		RelateBuyAMount string `json:"relate_buy_amount"`
		SaleNum         int    `json:"sale_num"` //销量
	}
	ResultDetailSkuInfoGoods struct {
		SkuID     string `json:"skuId"`
		Price     string `json:"price"`
		Num       int64  `json:"num"`
		Min       int64  `json:"min"`
		Max       int64  `json:"max"`
		SkuName   string `json:"sku_name"` //属性名称
		Weight    string `json:"weight"`   //重量
		ImagePath string `json:"imagePath"`
		Disable   bool   `json:"disable"`
		Pk        string `json:"pk"`
		CanBuy    bool   `json:"can_buy"` //当前是否能够购买
		Mark      string `json:"mark"`
		Stock     int64  `json:"stock"`
	}
	PreSaleInfo struct {
		DeliveryTime     string `json:"delivery_time"`
		PreheatTimeStart string `json:"preheat_time_start"` //预热开始时间
		PreheatTimeOver  string `json:"preheat_time_over"`  //预热结束时间
	}

	DownPayInfo struct {
		DownPayment    string `json:"down_payment"`     //定金
		SaleOnlineTime string `json:"sale_online_time"` //预售开始时间
		SaleOverTime   string `json:"sale_over_time"`   //预售结束时间
		FinalStartTime string `json:"final_start_time"`
		FinalOverTime  string `json:"final_over_time"`
		DeliveryTime   string `json:"delivery_time"`

		CountdownStart  string `json:"countdown_start"`  //倒计时开始时间
		NowTime         string `json:"now_time"`         //倒计时开始时间
		CountdownEnd    string `json:"countdown_end"`    //倒计时结束时间
		CountdownFormat string `json:"countdown_format"` //倒计时时间格式
		StatusDesc      string `json:"status_desc"`
		ShowMsg         string `json:"show_msg"` //显示提示信息

		ShowType uint8  `json:"show_type"`
		BtnShow  string `json:"btn_show"` //按钮显示
	}
)

func (r *ArgGetSpuDataWithSpuId) Default(c *base.Context) (err error) {

	return
}

func (r *ArgGetSpuDataWithSpuId) GetJsonByte() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}



func GetSpusDetailDataTypeBySpuId(ctx *base.Context, spuIds *ArgGetSpuDataWithSpuId) (res map[string]*SpuData, err error) {
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



//商品详情页数据
func GetSpusPageDetailDataTypeBySpuId(ctx *base.Context, spuIds *ArgSpuDetail) (res ResultSpuDetail, err error) {
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameMall,
		URI:         "/product/detail_for_page",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}

	if ctx.GinContext != nil {
		params.Header.Set(app_obj.HttpHeaderInfo, ctx.GinContext.GetHeader(app_obj.HttpHeaderInfo))
	}
	if 	params.BodyJson,err=json.Marshal(spuIds);err!=nil{
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
		Data ResultSpuDetail `json:"data"`
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
