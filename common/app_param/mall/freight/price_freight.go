package freight

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/shopspring/decimal"
)

type (
	//邮费计算结构体
	PriceFreight struct {
		context     *base.Context                           `json:"-"`
		sKusFreight []*SkuFreightSingle                     `json:"-"`          //计算邮费的每个SKU需要的数据
		ToCityId    int64                                   `json:"to_city_id"` //城市ID
		dataGroup   map[int64]map[string][]SkuFreightSingle `json:"-"`          //数据按照 店铺ID  SPU_ID分组
		Result      PriceFreightResult                      `json:"result"`
	}

	PriceFreightResult struct {
		Total    decimal.Decimal                `json:"total"`     //总邮费
		TotalNum int64                          `json:"total_num"` //总邮费
		Shops    map[int64]ShopCalResultFreight //邮费计算结果()
	}

	ShopCalResultFreight struct {
		ShopId       int64                 `json:"shop_id"`    //店铺ID
		ShopTotal    decimal.Decimal       `json:"shop_total"` //店铺总邮费
		ShopTotalNum int                   `json:"shop_total_num"`
		SkuFreight   []SkuCalResultFreight `json:"sku_freight"` //邮费价格
	}
	SkuCalResultFreight struct {
		SkuId        string          `json:"sku_id"`
		SpuId        string          `json:"spu_id"`
		ShopId       int64           `json:"shop_id"`
		TemplateId   int64           `json:"template_id"`   //邮费模板ID
		ToCityId     int64           `json:"to_city_id"`    //邮寄城市ID
		FreightPrice decimal.Decimal `json:"freight_price"` //邮费价格
		Mark         string          `json:"mark"`          //备注
		//skuFreightSingle *SkuFreightSingle `json:"-"`             //计算需要的数据
	}

	SkuFreightSingle struct {
		Num             int64                     `json:"num"` //数量
		Sku             *models.Sku               `json:"sku"` //SKU信息
		SkuRelate       *models.SkuPropertyRelate `json:"sku_relate"`
		Spu             *models.Product           `json:"spu"`          //商品信息
		Shop            *models.Shop              `json:"shop"`         //店铺信息
		FromCityId      int64                     `json:"from_city_id"` //邮寄城市ID
		TemplateFreight *TemplateFreight          //运费模板
	}

	TemplateFreight struct {
		Template models.FreightTemplate        `json:"freight_model"`
		Areas    []*models.FreightTemplateArea `json:"areas"`
	}

	ResFreightSku struct {
		Price decimal.Decimal `json:"price"`
		Mark  string          `json:"mark"`
	}
	OptionPriceFreight func(*PriceFreight)
)

//计算邮费动作
func (r *PriceFreight) Calculate() (err error) {
	//先将数据按照店铺进行分组
	r.groupData()

	//按店铺计算邮费
	err = r.calculateShop()
	return
}

func (r *PriceFreight) orgGroupParameters(skuItem *SkuFreightSingle, l int) {
	var (
		ok bool
	)
	if _, ok = r.dataGroup[skuItem.Sku.ShopId]; !ok {
		r.dataGroup[skuItem.Sku.ShopId] = map[string][]SkuFreightSingle{}
		return
	}

	if _, ok := r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID]; !ok {
		r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID] = make([]SkuFreightSingle, 0, l)
	}

	r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID] = append(r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID],
		*skuItem)

	return
}

//先将数据按照店铺进行分组
func (r *PriceFreight) groupData() {
	var (
		ok bool
		l  = len(r.sKusFreight)
	)

	for _, skuItem := range r.sKusFreight {

		r.orgGroupParameters(skuItem, l)

		if _, ok = r.Result.Shops[skuItem.Sku.ShopId]; !ok {
			r.Result.Shops[skuItem.Sku.ShopId] = ShopCalResultFreight{ShopId: skuItem.Shop.ShopID, SkuFreight: make([]SkuCalResultFreight, 0, l)}
		}
	}
	return
}

//添加参与邮费计算的SKU信息
func (r *PriceFreight) AppendNeedCalSKus(sKusFreight ...*SkuFreightSingle) (res *PriceFreight) {
	res = r
	r.sKusFreight = append(r.sKusFreight, sKusFreight...)
	return
}

func (r *PriceFreight) calculateShop() (err error) {

	for _, shopData := range r.Result.Shops {
		//每个店铺邮费逐一计算金额
		if err = r.calEveryShop(r.dataGroup[shopData.ShopId], shopData.ShopId); err != nil {
			return
		}
	}
	return
}

func (r *PriceFreight) calSku(skuCalResultFreight *SkuFreightSingle, freight *models.FreightTemplateArea) (resFreightSku *ResFreightSku, err error) {

	if resFreightSku, err = r.getFreightPrice(skuCalResultFreight, freight); err != nil {
		return
	}
	r.Result.TotalNum += skuCalResultFreight.Num
	r.Result.Total = r.Result.Total.Add(resFreightSku.Price)
	return
}

func (r *PriceFreight) getFreightParameters(single *SkuFreightSingle) (notSupportSend bool, res *models.FreightTemplateArea) {

	for _, item := range single.TemplateFreight.Areas {
		if item.CityId == r.ToCityId {
			res = item
			return
		}
	}
	notSupportSend = true
	return
}

//计算每个店铺的邮费金额
func (r *PriceFreight) calEverySpu(spuCalResultFreight []SkuFreightSingle, shopId int64) (err error) {

	if len(spuCalResultFreight) == 0 {
		return
	}

	notSupportSend, freight := r.getFreightParameters(&(spuCalResultFreight[0]))

	var dtm SkuCalResultFreight

	for _, skuCalResultFreight := range spuCalResultFreight {
		if dtm, err = r.orgSkuCalResultFreight(shopId, freight, notSupportSend, skuCalResultFreight); err != nil {
			return
		}
		if shopResult, ok := r.Result.Shops[shopId]; ok {
			shopResult.SkuFreight = append(shopResult.SkuFreight, dtm)
			r.Result.Shops[shopId] = shopResult
		}
	}
	return
}

func (r *PriceFreight) orgSkuCalResultFreight(shopId int64, freight *models.FreightTemplateArea, notSupportSend bool, skuCalResultFreight SkuFreightSingle) (dtm SkuCalResultFreight, err error) {
	var resFreightSku *ResFreightSku

	dtm = SkuCalResultFreight{
		//skuFreightSingle: skuItem,
		SkuId:      skuCalResultFreight.Sku.GetHid(),
		SpuId:      skuCalResultFreight.SkuRelate.ProductId,
		ShopId:     skuCalResultFreight.Sku.ShopId,
		TemplateId: skuCalResultFreight.TemplateFreight.Template.ID,
		ToCityId:   r.ToCityId,
	}

	//如果不支持邮寄城市
	if notSupportSend {
		dtm.Mark = fmt.Sprintf("该区域暂不支持邮寄")
		return
	}

	//计算每个SKU的邮费
	if resFreightSku, err = r.calSku(&skuCalResultFreight, freight); err != nil {
		return
	}

	dtm.Mark = resFreightSku.Mark
	dtm.FreightPrice = resFreightSku.Price
	return
}

//计算每个店铺的邮费金额
func (r *PriceFreight) calEveryShop(freight map[string][]SkuFreightSingle, shopId int64) (err error) {
	for spuId, spuCalResultFreight := range freight {
		_ = spuId
		if err = r.calEverySpu(spuCalResultFreight, shopId); err != nil {
			return
		}
	}

	return
}

//获取SKU的邮费
func (r *PriceFreight) getFreightPrice(freight *SkuFreightSingle, areas *models.FreightTemplateArea) (res *ResFreightSku, err error) {

	//if freight.skuFreightSingle.FreightModel.FreeFreight
	//1-包邮 2-自定义运费 3-有条件包邮
	switch freight.TemplateFreight.Template.FreeFreight {
	case models.FreightTemplateFreeFreightFree: // "包邮"
		res = &ResFreightSku{}
		res.Price = decimal.NewFromInt(0)
		res.Mark = "包邮"
	case models.FreightTemplateFreeFreightPay, models.FreightTemplateFreeFreightFreeWithAsCondition: // "买家承担运费"// "有条件包邮"
		res, err = r.getEverySkuFreightNeedPayPrice(freight, areas)
	default:
		res = &ResFreightSku{}
		err = fmt.Errorf("数据异常,当前不支持你选择的邮费类型")
		r.context.Error(map[string]interface{}{
			"err":     err.Error(),
			"freight": freight,
		}, "PriceFreightGetFreightPrice")
	}
	return
}

//获取每个SKU邮费
func (r *PriceFreight) getEverySkuFreightNeedPayPrice(freight *SkuFreightSingle, areas *models.FreightTemplateArea) (res *ResFreightSku, err error) {
	res = &ResFreightSku{}

	return
}

func OptionFreightContext(context *base.Context) OptionPriceFreight {
	return func(freight *PriceFreight) {
		freight.context = context
	}
}

func OptionFreightToCityId(toCityId int64) OptionPriceFreight {
	return func(freight *PriceFreight) {
		freight.ToCityId = toCityId
	}
}

func NewPriceFreight(options ...OptionPriceFreight) *PriceFreight {
	p := &PriceFreight{sKusFreight: make([]*SkuFreightSingle, 0, 16)}
	for _, option := range options {
		option(p)
	}
	return p
}
