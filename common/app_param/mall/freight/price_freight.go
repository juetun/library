package freight

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/shopspring/decimal"
	"strings"
)

type (
	//邮费计算结构体
	PriceFreight struct {
		context     *base.Context                          `json:"-"`
		sKusFreight []*SkuFreightSingle                    `json:"-"`           //计算邮费的每个SKU需要的数据
		EmsAddress  *ResultGetByAddressIdsItem             `json:"ems_address"` //城市ID
		ToCityId    string                                 `json:"to_city_id"`
		dataGroup   map[int64]map[int64][]SkuFreightSingle `json:"-"` //数据按照 店铺ID  SPU_ID分组
		Result      PriceFreightResult                     `json:"result"`
	}

	PriceFreightResult struct {
		Total       decimal.Decimal                 `json:"-"`         //总邮费
		TotalString string                          `json:"total"`     //总费用
		TotalNum    int64                           `json:"total_num"` //总邮费
		Shops       map[int64]*ShopCalResultFreight `json:"shops"`     //邮费计算结果()
	}

	ShopCalResultFreight struct {
		ShopId             int64                  `json:"shop_id"` //店铺ID
		FreightTotal       decimal.Decimal        `json:"-"`       //店铺总邮费
		FreightTotalString string                 `json:"freight_total"`
		ShopTotalNum       int64                  `json:"shop_total_num"` //店铺商品总数
		SkuFreight         []*SkuCalResultFreight `json:"sku_freight"`    //邮费价格
		Summary            AttrSummary            `json:"summary"`        //店铺数据汇总 总重量、总体积 总件数
		CalParameter       *CalCaseFreight        `json:"cal_parameter"`  //计算邮费的基本参数
	}
	CalResultFreight struct {
		FreightId          int64                  `json:"shop_id"` //店铺ID
		FreightTotal       decimal.Decimal        `json:"-"`       //店铺总邮费
		FreightTotalString string                 `json:"shop_total"`
		FreightTotalNum    int64                  `json:"shop_total_num"` //店铺商品总数
		SkuFreight         []*SkuCalResultFreight `json:"sku_freight"`    //邮费价格
		Summary            AttrSummary            `json:"summary"`        //店铺数据汇总 总重量、总体积 总件数
	}
	AttrSummary struct { //店铺汇总数据
		SkuTotalPrice       decimal.Decimal `json:"-"` //价格
		SkuTotalPriceString string          `json:"sku_total_price"`
		Num                 int64           `json:"num"`    //件数
		Weight              decimal.Decimal `json:"-"`      //重量
		WeightString        string          `json:"weight"` //重量
		Volume              decimal.Decimal `json:"-"`
		VolumeString        string          `json:"volume"`
	}
	SkuCalResultFreight struct {
		SkuId              string          `json:"sku_id"`
		SpuId              string          `json:"spu_id"`
		ShopId             int64           `json:"shop_id"`
		TemplateId         int64           `json:"template_id"` //邮费模板ID
		ToCityId           string          `json:"to_city_id"`  //邮寄城市ID
		FreightPrice       decimal.Decimal `json:"-"`           //邮费价格
		FreightPriceString string          `json:"freight_price"`
		Mark               string          `json:"mark"` //备注
		AttrSummary
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
		SkuTotalPrice decimal.Decimal `json:"sku_total_price"`
		FreightPrice  decimal.Decimal `json:"freight_price"`
		Mark          string          `json:"mark"`
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
		FreightSaleArea *models.FreightSaleAreaBase      `json:"base,omitempty"` //计算邮费条件基本规则
		ExtCase         *models.FreightFreeConditionBase `json:"ext,omitempty"`  //补充条件 （如 满多少包邮之类）
		PricingMode     uint8                            `json:"pri_mode"`       //计价方式
	}
)

func (r *SkuCalResultFreight) Default() {
	r.FreightPriceString = r.FreightPrice.StringFixed(2)
	r.AttrSummary.Default()
	return
}

//重新初始化数据设置包邮
func (r *CalResultFreight) ReInitSkuFreightFree(desc string) {
	return
	r.FreightTotal = decimal.NewFromInt(0)
	for key, item := range r.SkuFreight {
		if item.Mark != "" {
			item.Mark = fmt.Sprintf("%v,%v", item.Mark, desc)
		} else {
			item.Mark = desc
		}
		item.FreightPrice = decimal.NewFromInt(0)
		item.AttrSummary.Default()
		r.SkuFreight[key] = item
	}
	return
}

func (r *ShopCalResultFreight) Default() {
	r.Summary.Default()
	r.FreightTotalString = r.FreightTotal.StringFixed(2)
	return
}

func (r *SkuFreightSingle) GetWeight() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)
	r.Sku.Default()
	var weight decimal.Decimal
	if weight, err = decimal.NewFromString(r.Sku.Weight); err != nil {
		return
	}
	res = weight.Mul(decimal.NewFromInt(r.Num))
	return
}

func (r *SkuFreightSingle) GetVolume() (res decimal.Decimal, err error) {
	res = decimal.NewFromInt(0)
	r.Sku.Default()
	var volume decimal.Decimal
	if volume, err = decimal.NewFromString(r.Sku.Volume); err != nil {
		return
	}
	res = volume.Mul(decimal.NewFromInt(r.Num))
	return
}

func (r *AttrSummary) Default() {
	r.SkuTotalPriceString = r.SkuTotalPrice.StringFixed(2)
	r.WeightString = r.Weight.StringFixed(2)
	r.VolumeString = r.Volume.StringFixed(2)
}

//func (r *PriceFreightResult) MarshalJSON() (res []byte, err error) {
//	r.TotalString = r.Total.StringFixed(2)
//	res, err = json.Marshal(r)
//	return
//}
//
//func (r *ShopCalResultFreight) MarshalJSON() (res []byte, err error) {
//	r.ShopTotalString = r.ShopTotal.StringFixed(2)
//	res, err = json.Marshal(r)
//	return
//}

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

	r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.FreightTemplate] = append(r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.FreightTemplate],
		*skuItem)

	return
}

//先将数据按照店铺进行分组
func (r *PriceFreight) groupData() (err error) {
	var (
		ok bool
		l  = len(r.sKusFreight)
	)
	r.dataGroup = make(map[int64]map[int64][]SkuFreightSingle, l)
	r.Result.Shops = make(map[int64]*ShopCalResultFreight, l)
	var (
		dataItem *ShopCalResultFreight
	)
	for _, skuItem := range r.sKusFreight {

		if dataItem, ok = r.Result.Shops[skuItem.Sku.ShopId]; !ok {
			dataItem = &ShopCalResultFreight{
				FreightTotal:       decimal.NewFromInt(0),
				FreightTotalString: "0.00",
				ShopId:             skuItem.Shop.ShopID,
				SkuFreight:         make([]*SkuCalResultFreight, 0, l),
				Summary: AttrSummary{
					Num:           0,
					SkuTotalPrice: decimal.NewFromInt(0),
					Weight:        decimal.NewFromInt(0),
					Volume:        decimal.NewFromInt(0),
				},
			}
			dataItem.Summary.Default()
		}
		r.Result.Shops[skuItem.Sku.ShopId] = dataItem
		if _, ok = r.dataGroup[skuItem.Sku.ShopId]; !ok {
			r.dataGroup[skuItem.Sku.ShopId] = map[int64][]SkuFreightSingle{}
		}
		if _, ok = r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.FreightTemplate]; !ok {
			r.dataGroup[skuItem.Sku.ShopId][skuItem.Spu.FreightTemplate] = make([]SkuFreightSingle, 0, l)
		}
		if err = r.orgGroupParameters(skuItem, l, dataItem); err != nil {
			return
		}
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
	r.dataGroup = map[int64]map[int64][]SkuFreightSingle{}
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
		r.Result.Total = r.Result.Total.Add(shopData.FreightTotal)
		r.Result.TotalNum += shopData.ShopTotalNum
	}
	r.Result.TotalString = r.Result.Total.StringFixed(2)
	return
}

func (r *CalCaseFreight) ToJson() (res string) {
	if r == nil {
		return
	}
	var bt []byte
	bt, _ = json.Marshal(r)
	res = string(bt)
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
	if notSupportSend, err = r.initFreightSaleArea(single, res); err != nil {
		return
	}
	//初始化包邮条件
	if err = r.initExtCase(single, res); err != nil {
		return
	}

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

func (r *PriceFreight) initFreightSaleArea(single *SkuFreightSingle, res *CalCaseFreight) (notSupportSend bool, err error) {

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
	if res.FreightSaleArea == nil {
		notSupportSend = true
		return
	}

	return
}

//邮费基本信息计算
func (r *PriceFreight) freightCalBase(freightCalResultFreight []SkuFreightSingle, shopCalResultFreight *ShopCalResultFreight, ) (freight *CalCaseFreight, err error) {
	var (
		notSupportSend   bool
		isFreeFreight    bool //判断整个模板是否包邮
		calResultFreight = NewCalResultFreight(len(freightCalResultFreight))
	)
	var (
		skuItemInfo = freightCalResultFreight[0]
	)

	//根据当前SPU下任意一条SKU关联的快递信息获取运费模板（系统逻辑 同一个Spu的所选运费模板是相同的）
	//notSupportSend=false 表示此城市当前不支持发货
	if notSupportSend, isFreeFreight, freight, err = r.getFreightParameters(&skuItemInfo); err != nil {
		return
	}
	//分别计算每个SKU的邮费信息
	for _, skuCalResultFreight := range freightCalResultFreight {
		shopCalResultFreight.ShopTotalNum += skuCalResultFreight.Num
		if dtm, e := r.orgSkuCalResultFreight(freight, notSupportSend, isFreeFreight, &skuCalResultFreight); e != nil {
			err = e
			return
		} else {
			calResultFreight.Summary.SkuTotalPrice = shopCalResultFreight.Summary.SkuTotalPrice.Add(dtm.AttrSummary.SkuTotalPrice)
			calResultFreight.Summary.Num += skuCalResultFreight.Num
			calResultFreight.Summary.Volume = shopCalResultFreight.Summary.Volume.Add(dtm.Volume)
			calResultFreight.Summary.Weight = shopCalResultFreight.Summary.Weight.Add(dtm.Weight)
			calResultFreight.FreightTotal = shopCalResultFreight.FreightTotal.Add(dtm.FreightPrice)
			dtm.Default()
			calResultFreight.SkuFreight = append(calResultFreight.SkuFreight, dtm)
		}
	}
	var (
		isFree bool
		desc   string
	)
	//实现有条件包邮逻辑
	if isFree, desc, err = r.freightCalExt(freight.ExtCase, calResultFreight); err != nil {
		return
	}

	if isFree { //如果满足有条件包邮
		calResultFreight.ReInitSkuFreightFree(desc)
	}

	r.mergeShopAndFreightData(shopCalResultFreight, calResultFreight)
	return
}

func (r *PriceFreight) mergeShopAndFreightData(shopCalResultFreight *ShopCalResultFreight, calResultFreight *CalResultFreight) {
	shopCalResultFreight.Summary.SkuTotalPrice = shopCalResultFreight.Summary.SkuTotalPrice.Add(calResultFreight.Summary.SkuTotalPrice)
	shopCalResultFreight.Summary.Num += calResultFreight.Summary.Num
	shopCalResultFreight.Summary.Volume = shopCalResultFreight.Summary.Volume.Add(calResultFreight.Summary.Volume)
	shopCalResultFreight.Summary.Weight = shopCalResultFreight.Summary.Weight.Add(calResultFreight.Summary.Weight)
	shopCalResultFreight.FreightTotal = shopCalResultFreight.FreightTotal.Add(calResultFreight.FreightTotal)
	shopCalResultFreight.SkuFreight = append(shopCalResultFreight.SkuFreight, calResultFreight.SkuFreight...)
	return
}

func NewCalResultFreight(num int) (res *CalResultFreight) {
	res = &CalResultFreight{
		SkuFreight: make([]*SkuCalResultFreight, 0, num),
		Summary: AttrSummary{
			Num:           0,
			SkuTotalPrice: decimal.NewFromInt(0),
			Weight:        decimal.NewFromInt(0),
			Volume:        decimal.NewFromInt(0),
		},
	}
	return
}

//实现有条件包邮逻辑
func (r *PriceFreight) freightCalExt(freight *models.FreightFreeConditionBase, calResultFreight *CalResultFreight) (isFree bool, desc string, err error) {
	if freight == nil { //如果没有配置一条满足条件包邮
		return
	}
	switch freight.FreightType {
	case models.FreeConditionOptAnd: //且 逻辑 两个条件都满足
		numberCondition := decimal.NewFromInt(calResultFreight.Summary.Num).GreaterThan(decimal.NewFromInt(int64(freight.FullNumber)))
		if freight.FullPrice != "" {
			desc = "不包邮"
			return
		}
		var FullPrice decimal.Decimal
		if FullPrice, err = decimal.NewFromString(freight.FullPrice); err != nil {
			return
		}
		priceCondition := calResultFreight.Summary.SkuTotalPrice.GreaterThan(FullPrice)
		if numberCondition && priceCondition {
			isFree = true
			desc = fmt.Sprintf("满(%v)且满(%v)件包邮", freight.FullNumber, freight.FullPrice)
		}

	case models.FreeConditionOptOr: //或 逻辑 两个条件满足一个即可
		var (
			numberCondition bool
			priceCondition  bool
			descList        = make([]string, 0, 3)
		)
		if freight.FullNumber != 0 {
			numberCondition = decimal.NewFromInt(calResultFreight.Summary.Num).GreaterThan(decimal.NewFromInt(int64(freight.FullNumber)))
		}
		if freight.FullPrice != "" {
			var FullPrice decimal.Decimal
			if FullPrice, err = decimal.NewFromString(freight.FullPrice); err != nil {
				return
			}
			priceCondition = calResultFreight.Summary.SkuTotalPrice.GreaterThan(FullPrice)
		}
		if numberCondition {
			descList = append(descList, fmt.Sprintf("满(%v件)", freight.FullNumber))
		}
		if priceCondition {
			descList = append(descList, fmt.Sprintf("满(￥%v)", freight.FullPrice))
		}
		if numberCondition || priceCondition {
			desc = strings.Join(descList, "或")
			desc = fmt.Sprintf("%v包邮", desc)
			isFree = true
		}
	default:
		err = fmt.Errorf("数据异常,对不起当前只支持您配置的数据逻辑")
	}

	return
}

//计算每个店铺的邮费金额,按照相同运费模板的合并一起算
func (r *PriceFreight) calEveryFreight(freightCalResultFreight []SkuFreightSingle, shopCalResultFreight *ShopCalResultFreight, freightId int64) (err error) {

	if len(freightCalResultFreight) == 0 {
		return
	}
	var (
		freight *CalCaseFreight
	)
	//邮费基本条件计算
	if freight, err = r.freightCalBase(freightCalResultFreight, shopCalResultFreight); err != nil {
		return
	}
	shopCalResultFreight.CalParameter = freight
	shopCalResultFreight.Default()

	return
}

func (r *PriceFreight) initSkuCalResultFreight(skuCalResultFreight *SkuFreightSingle) (dtm *SkuCalResultFreight, err error) {
	dtm = &SkuCalResultFreight{
		SkuId:      skuCalResultFreight.Sku.GetHid(),
		SpuId:      skuCalResultFreight.SkuRelate.ProductId,
		ShopId:     skuCalResultFreight.Sku.ShopId,
		TemplateId: skuCalResultFreight.TemplateFreight.Template.ID,
		ToCityId:   r.EmsAddress.GetToCityId(),
		AttrSummary: AttrSummary{
			Weight: decimal.NewFromInt(0),
			Volume: decimal.NewFromInt(0),
			Num:    skuCalResultFreight.Num,
		},
	}
	if dtm.Weight, err = skuCalResultFreight.GetWeight(); err != nil {
		return
	}
	if dtm.Volume, err = skuCalResultFreight.GetVolume(); err != nil {
		return
	}
	dtm.AttrSummary.Default()
	return
}

func (r *PriceFreight) orgSkuCalResultFreight(freight *CalCaseFreight, notSupportSend, isFreeFreight bool, skuCalResultFreight *SkuFreightSingle) (dtm *SkuCalResultFreight, err error) {
	var resFreightSku *ResFreightSku
	if dtm, err = r.initSkuCalResultFreight(skuCalResultFreight); err != nil {
		return
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
	dtm.AttrSummary.SkuTotalPrice = resFreightSku.SkuTotalPrice
	dtm.Mark = resFreightSku.Mark
	dtm.FreightPrice = resFreightSku.FreightPrice
	dtm.AttrSummary.Default()
	return
}

//计算每个店铺的邮费金额
//shopItemData 每个店铺的商品计算邮费的基本信息（按照SPU分组的映射）
func (r *PriceFreight) calEveryShop(shopItemData map[int64][]SkuFreightSingle, shopCalResultFreight *ShopCalResultFreight) (err error) {
	for freightId, spuCalResultFreight := range shopItemData {

		//按照【店铺ID】【运费模板维度】分组计算邮费
		if err = r.calEveryFreight(spuCalResultFreight, shopCalResultFreight, freightId); err != nil {
			return
		}

	}

	return
}

//获取SKU的邮费
func (r *PriceFreight) getFreightPrice(freight *SkuFreightSingle, areas *CalCaseFreight) (res *ResFreightSku, err error) {
	var (
		skuPrice decimal.Decimal
	)
	if skuPrice, err = decimal.NewFromString(freight.SkuRelate.Price); err != nil {
		return
	}
	res = &ResFreightSku{
		SkuTotalPrice: skuPrice.Mul(decimal.NewFromInt(freight.Num)),
		FreightPrice:  decimal.NewFromInt(0),
	}
	//if freight.skuFreightSingle.FreightModel.FreeFreight
	//1-包邮 2-自定义运费 3-有条件包邮
	switch freight.TemplateFreight.Template.FreeFreight {
	case models.FreightTemplateFreeFreightFree: // "包邮"

		res.Mark = "包邮"
	case models.FreightTemplateFreeFreightPay, models.FreightTemplateFreeFreightFreeWithAsCondition: // "买家承担运费"// "有条件包邮"
		if err = r.getEverySkuFreightNeedPayPrice(res, freight, areas); err != nil {
			return
		}
	default:

		err = fmt.Errorf("数据异常,当前不支持你选择的邮费类型")
		r.context.Error(map[string]interface{}{
			"err":     err.Error(),
			"freight": freight,
		}, "PriceFreightGetFreightPrice")
	}
	return
}

//获取每个SKU邮费
func (r *PriceFreight) getEverySkuFreightNeedPayPrice(resFreightSku *ResFreightSku, freight *SkuFreightSingle, calCaseFreight *CalCaseFreight) (err error) {

	switch calCaseFreight.PricingMode {
	case models.FreightTemplatePricingModeUnit: //按件数
		if resFreightSku.FreightPrice, resFreightSku.Mark, err = calCaseFreight.FreightSaleArea.GetPriceByUnit(freight.Num); err != nil {
			return
		}

	case models.FreightTemplatePricingModeWeight: //按重量
		var weight, totalWeight decimal.Decimal
		if weight, err = freight.Sku.GetWeightDecimal(); err != nil {
			return
		}
		totalWeight = weight.Mul(decimal.NewFromInt(freight.Num))
		if resFreightSku.FreightPrice, resFreightSku.Mark, err = calCaseFreight.FreightSaleArea.GetPriceByWeight(totalWeight); err != nil {
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
		dataGroup:   make(map[int64]map[int64][]SkuFreightSingle, 16),
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
