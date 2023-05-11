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
		sKusFreight []*SkuFreightSingle                     `json:"-"`           //计算邮费的每个SKU需要的数据
		EmsAddress  *ResultGetByAddressIdsItem              `json:"ems_address"` //城市ID
		ToCityId    string                                  `json:"to_city_id"`
		dataGroup   map[int64]map[string][]SkuFreightSingle `json:"-"` //数据按照 店铺ID  SPU_ID分组
		Result      PriceFreightResult                      `json:"result"`
	}

	PriceFreightResult struct {
		Total    decimal.Decimal                 `json:"total"`     //总邮费
		TotalNum int64                           `json:"total_num"` //总邮费
		Shops    map[int64]*ShopCalResultFreight //邮费计算结果()
	}

	ShopCalResultFreight struct {
		ShopId       int64                 `json:"shop_id"`        //店铺ID
		ShopTotal    decimal.Decimal       `json:"shop_total"`     //店铺总邮费
		ShopTotalNum int                   `json:"shop_total_num"` //店铺商品总数
		SkuFreight   []SkuCalResultFreight `json:"sku_freight"`    //邮费价格
		summary      ShopSummary           `json:"-"`              //店铺数据汇总 总重量、总体积 总件数
	}
	ShopSummary struct { //店铺汇总数据
		Num    int64           `json:"num"`    //件数
		Weight decimal.Decimal `json:"weight"` //重量
		Volume decimal.Decimal `json:"volume"`
	}
	SkuCalResultFreight struct {
		SkuId        string          `json:"sku_id"`
		SpuId        string          `json:"spu_id"`
		ShopId       int64           `json:"shop_id"`
		TemplateId   int64           `json:"template_id"`   //邮费模板ID
		ToCityId     string          `json:"to_city_id"`    //邮寄城市ID
		FreightPrice decimal.Decimal `json:"freight_price"` //邮费价格
		Mark         string          `json:"mark"`          //备注
		//skuFreightSingle *SkuFreightSingle `json:"-"`             //计算需要的数据
	}

	SkuFreightSingle struct {
		Num               int64                      `json:"num"`                 //数量
		Sku               *models.Sku                `json:"sku"`                 //SKU信息
		SkuRelate         *models.SkuPropertyRelate  `json:"sku_relate"`          //SKU信息 relate
		Spu               *models.Product            `json:"spu"`                 //商品信息
		Shop              *models.Shop               `json:"shop"`                //店铺信息
		EmsAddressFreight *ResultGetByAddressIdsItem `json:"ems_address_freight"` //收货地址信息
		TemplateFreight   *TemplateFreight           `json:"template_freight"`    //运费模板
		FromCityId        string                     `json:"from_city_id"`        //发货城市,预留字段
	}

	TemplateFreight struct {
		Template models.FreightTemplate `json:"freight_model"`
		//Areas    []*models.FreightTemplateArea `json:"areas"`
	}

	ResFreightSku struct {
		Price decimal.Decimal `json:"price"`
		Mark  string          `json:"mark"`
	}
	OptionPriceFreight func(*PriceFreight)

	ResultGetByAddressIdsItem struct {
		Id           int64  `json:"id"`
		ProvinceId   string `json:"province_id"`
		CityId       string `json:"city_id"`
		AreaId       string `json:"area_id"`
		Province     string `json:"province"`
		City         string `json:"city"`
		Area         string `json:"area"`
		Title        string `json:"title"`
		Address      string `json:"address"`
		AreaAddress  string `json:"area_address"`
		ContactUser  string `json:"contact_user"`
		ContactPhone string `json:"contact_phone"`
		FullAddress  string `json:"full_address"`
	}
	CalCaseFreight struct {
		FreightSaleArea *models.FreightSaleAreaBase      //计算邮费条件条件基数
		ExtCase         *models.FreightFreeConditionBase //补充条件 （如 满多少包邮之类）
		PricingMode     uint8                            //计价方式
	}
)

func (r *ResultGetByAddressIdsItem) GetToCityId() (res string) {
	res = r.CityId
	return
}

//计算邮费动作
func (r *PriceFreight) Calculate() (res *PriceFreightResult, err error) {
	//先将数据按照店铺进行分组,数据初始化
	if err = r.groupData(); err != nil {
		return
	}

	//按店铺计算邮费
	if err = r.calculateShop(); err != nil {
		return
	}
	res = &r.Result
	return
}

func (r *PriceFreight) orgGroupParameters(skuItem *SkuFreightSingle, l int, dataItem *ShopCalResultFreight) (err error) {
	var (
		ok bool
	)
	if _, ok = r.dataGroup[skuItem.Sku.ShopId]; !ok {
		r.dataGroup[skuItem.Sku.ShopId] = map[string][]SkuFreightSingle{}
	}
	if _, ok := r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID]; !ok {
		r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID] = make([]SkuFreightSingle, 0, l)
	}

	//数量统计
	dataItem.summary.Num += skuItem.Num

	//重量汇总
	if skuItem.Sku.Weight != "" {
		var weight decimal.Decimal
		if weight, err = decimal.NewFromString(skuItem.Sku.Weight); err != nil {
			return
		}
		dataItem.summary.Weight = dataItem.summary.Weight.Add(weight)
	}

	r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID] = append(r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.ProductID],
		*skuItem)

	return
}

//先将数据按照店铺进行分组
func (r *PriceFreight) groupData() (err error) {
	var (
		ok bool
		l  = len(r.sKusFreight)
	)
	r.dataGroup = make(map[int64]map[string][]SkuFreightSingle, l)
	r.Result.Shops = make(map[int64]*ShopCalResultFreight, l)
	var (
		dataItem *ShopCalResultFreight
	)
	for _, skuItem := range r.sKusFreight {

		if dataItem, ok = r.Result.Shops[skuItem.Sku.ShopId]; !ok {
			dataItem = &ShopCalResultFreight{
				ShopTotal:  decimal.NewFromInt(0),
				ShopId:     skuItem.Shop.ShopID,
				SkuFreight: make([]SkuCalResultFreight, 0, l),
				summary: ShopSummary{
					Num:    0,
					Weight: decimal.NewFromInt(0),
					Volume: decimal.NewFromInt(0),
				},
			}
		}
		r.Result.Shops[skuItem.Sku.ShopId] = dataItem
		r.orgGroupParameters(skuItem, l, dataItem)
	}
	return
}

//添加参与邮费计算的SKU信息
func (r *PriceFreight) AppendNeedCalSKus(sKusFreight ...*SkuFreightSingle) (res *PriceFreight) {
	res = r
	var (
		l = len(sKusFreight)
	)
	r.sKusFreight = make([]*SkuFreightSingle, 0, l)
	r.dataGroup = map[int64]map[string][]SkuFreightSingle{}
	r.Result.Shops = map[int64]*ShopCalResultFreight{}
	r.sKusFreight = append(r.sKusFreight, sKusFreight...)
	return
}

func (r *PriceFreight) calculateShop() (err error) {

	for _, shopData := range r.Result.Shops {
		//每个店铺邮费逐一计算金额
		if err = r.calEveryShop(r.dataGroup[shopData.ShopId], shopData); err != nil {
			return
		}
	}
	return
}

func (r *PriceFreight) getFreightParameters(single *SkuFreightSingle) (notSupportSend, isFreeFreight bool, res *CalCaseFreight, err error) {
	if single.TemplateFreight == nil || single.TemplateFreight.Template.ID == 0 {
		err = fmt.Errorf("数据异常（spu_id:%v）未配置运费模板", single.Spu.ProductID)
		return
	}
	res = &CalCaseFreight{
		PricingMode: single.TemplateFreight.Template.PricingMode, //初始化计价方式 1-按件数 2-按重量
	}
	switch single.TemplateFreight.Template.FreeFreight {
	case models.FreightTemplateFree: //如果包邮
		isFreeFreight = true
		return
	}
	//初始化邮费计算条件
	if err = r.initFreightSaleArea(single, res); err != nil {
		return
	}

	//初始化包邮条件
	if err = r.initExtCase(single, res); err != nil {
		return
	}
	notSupportSend = true
	return
}

func (r *PriceFreight) initExtCase(single *SkuFreightSingle, res *CalCaseFreight) (err error) {

	var (
		cases []*models.FreightFreeCondition
	)
	if cases, err = single.TemplateFreight.Template.ParseFreeCondition(); err != nil {
		return
	}
	for _, item := range cases {
		var isBreak = false
		for _, it := range item.AreaCode {
			if it == r.EmsAddress.GetToCityId() {
				res.ExtCase = &item.FreightFreeConditionBase
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
	}
	return
}

func (r *PriceFreight) initFreightSaleArea(single *SkuFreightSingle, res *CalCaseFreight) (err error) {

	var (
		areas []*models.FreightSaleArea
	)
	r.ToCityId = r.EmsAddress.GetToCityId()
	if areas, err = single.TemplateFreight.Template.ParseSaleArea(); err != nil {
		return
	}
	for _, item := range areas {
		var isBreak = false
		for _, it := range item.AreaCode {
			if it == r.ToCityId {
				res.FreightSaleArea = &item.FreightSaleAreaBase
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
	}
	return
}

//计算每个店铺的邮费金额
func (r *PriceFreight) calEverySpu(spuCalResultFreight []SkuFreightSingle, shopCalResultFreight *ShopCalResultFreight, spuId string) (err error) {

	if len(spuCalResultFreight) == 0 {
		return
	}
	var (
		skuItemInfo    = spuCalResultFreight[0]
		notSupportSend bool
		freight        *CalCaseFreight
		isFreeFreight  bool //判断整个模板是否包邮
	)

	//根据当前SPU下任意一条SKU关联的快递信息获取运费模板（系统逻辑 同一个Spu的所选运费模板是相同的）
	//notSupportSend=false 表示此城市当前不支持发货
	if notSupportSend, isFreeFreight, freight, err = r.getFreightParameters(&skuItemInfo); err != nil {
		return
	}

	//分别计算每个SKU的邮费信息
	for _, skuCalResultFreight := range spuCalResultFreight {

		if dtm, e := r.orgSkuCalResultFreight(freight, notSupportSend, isFreeFreight, &skuCalResultFreight); e != nil {
			err = e
			return
		} else {
			shopCalResultFreight.SkuFreight = append(shopCalResultFreight.SkuFreight, dtm)
		}
	}
	return
}

func (r *PriceFreight) orgSkuCalResultFreight(freight *CalCaseFreight, notSupportSend, isFreeFreight bool, skuCalResultFreight *SkuFreightSingle) (dtm SkuCalResultFreight, err error) {
	var resFreightSku *ResFreightSku

	dtm = SkuCalResultFreight{
		SkuId:      skuCalResultFreight.Sku.GetHid(),
		SpuId:      skuCalResultFreight.SkuRelate.ProductId,
		ShopId:     skuCalResultFreight.Sku.ShopId,
		TemplateId: skuCalResultFreight.TemplateFreight.Template.ID,
		ToCityId:   r.EmsAddress.GetToCityId(),
	}

	if isFreeFreight {
		dtm.Mark = "包邮"
		return
	}
	//如果不支持邮寄城市
	if notSupportSend {
		dtm.Mark = fmt.Sprintf("该区域暂不支持邮寄")
		return
	}

	//计算每个SKU的邮费
	if resFreightSku, err = r.getFreightPrice(skuCalResultFreight, freight); err != nil {
		return
	}

	dtm.Mark = resFreightSku.Mark
	dtm.FreightPrice = resFreightSku.Price
	return
}

//计算每个店铺的邮费金额
//shopItemData 每个店铺的商品计算邮费的基本信息（按照SPU分组的映射）
func (r *PriceFreight) calEveryShop(shopItemData map[string][]SkuFreightSingle, shopCalResultFreight *ShopCalResultFreight) (err error) {

	for spuId, spuCalResultFreight := range shopItemData {
		if err = r.calEverySpu(spuCalResultFreight, shopCalResultFreight, spuId); err != nil {
			return
		}
	}

	return
}

//获取SKU的邮费
func (r *PriceFreight) getFreightPrice(freight *SkuFreightSingle, areas *CalCaseFreight) (res *ResFreightSku, err error) {

	//if freight.skuFreightSingle.FreightModel.FreeFreight
	//1-包邮 2-自定义运费 3-有条件包邮
	switch freight.TemplateFreight.Template.FreeFreight {
	case models.FreightTemplateFreeFreightFree: // "包邮"
		res = &ResFreightSku{}
		res.Price = decimal.NewFromInt(0)
		res.Mark = "包邮"
	case models.FreightTemplateFreeFreightPay, models.FreightTemplateFreeFreightFreeWithAsCondition: // "买家承担运费"// "有条件包邮"
		if res, err = r.getEverySkuFreightNeedPayPrice(freight, areas); err != nil {
			return
		}
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
func (r *PriceFreight) getEverySkuFreightNeedPayPrice(freight *SkuFreightSingle, calCaseFreight *CalCaseFreight) (res *ResFreightSku, err error) {
	res = &ResFreightSku{Price: decimal.NewFromInt(0)}

	switch calCaseFreight.PricingMode {
	case models.FreightTemplatePricingModeUnit: //按件数
		if res.Price, res.Mark, err = calCaseFreight.FreightSaleArea.GetPriceByUnit(freight.Num); err != nil {
			return
		}

	case models.FreightTemplatePricingModeWeight: //按重量
		var weight decimal.Decimal
		if weight, err = freight.Sku.GetWeightDecimal(); err != nil {
			return
		}
		if res.Price, res.Mark, err = calCaseFreight.FreightSaleArea.GetPriceByWeight(weight.Mul(decimal.NewFromInt(freight.Num))); err != nil {
			return
		}
	default:
		err = fmt.Errorf("暂不支持您选择的运费计价方式")
	}
	return
}

func OptionFreightContext(context *base.Context) OptionPriceFreight {
	return func(freight *PriceFreight) {
		freight.context = context
	}
}

func OptionFreightEmsAddress(EmsAddress *ResultGetByAddressIdsItem) OptionPriceFreight {
	return func(freight *PriceFreight) {
		freight.EmsAddress = EmsAddress
	}
}

func NewPriceFreight(options ...OptionPriceFreight) *PriceFreight {
	p := &PriceFreight{
		sKusFreight: make([]*SkuFreightSingle, 0, 16),
		dataGroup:   make(map[int64]map[string][]SkuFreightSingle, 16),
		Result: PriceFreightResult{
			Total:    decimal.NewFromInt(0),
			TotalNum: 0,
			Shops:    make(map[int64]*ShopCalResultFreight, 16),
		},
	}
	for _, option := range options {
		option(p)
	}
	return p
}
