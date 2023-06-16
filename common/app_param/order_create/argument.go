package order_create

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall/freight"
)

type (
	ArgCreateOrderFromCart struct {
		RequestUser app_param.RequestUser `json:"-" form:"-"`
		common.HeaderInfo
		OrderId string `json:"order_id" form:"order_id"` //订单ID号
		//SkuString          string                            `json:"sku_item,omitempty" form:"sku_item"`
		Amount             string                            `json:"amount,omitempty" form:"amount"`         // 总金额
		Status             uint8                             `json:"status,omitempty" form:"status"`         // 订单状态
		AddressId          int64                             `json:"address_id,omitempty" form:"address_id"` // 收货地址
		Express            string                            `json:"express,omitempty" form:"express"`       // 默认快递信息
		PayType            uint8                             `json:"pay_type,omitempty" form:"pay_type"`     // 支付类型
		Type               string                            `json:"type,omitempty" form:"Type"`             //数据操作路径
		ReceiptUserInfo    *ReceiptUserInfo                  `json:"receipt_user_info,omitempty" form:"receipt_user_info"`
		PriceFreightResult *freight.PriceFreightResult       `json:"freight_result" form:"freight_result"`       //订单的邮费计算结果 api-mall服务侧计算
		SkuItems           []*app_param.ArgOrderFromCartItem `json:"sku_item,omitempty" form:"sku_item"`         //商品详情
		OrderSrcChannel    string                            `json:"order_src_channel" form:"order_src_channel"` //订单来源渠道
		OrderSrcLoc        string                            `json:"order_src_loc" form:"order_src_loc"`
		TimeNow            base.TimeNormal                   `json:"-" form:"-"`
		GetDataTypeCommon  base.GetDataTypeCommon            `json:"dt_common" form:"dt_common"`
	}

	ArgGetInfoByOrderId struct {
		RequestUser app_param.RequestUser `json:"-" form:"-"`
		OrderId     string                `json:"order_id" form:"order_id"`
	}
	ReceiptUserInfo struct {
		ProvinceId   string `json:"province_id"`
		CityId       string `json:"city_id"`
		AreaId       string `json:"area_id"`
		Address      string `json:"address"`
		ZipCode      string `json:"zip_code"`
		ContactUser  string `json:"contact_user"`
		ContactPhone string `json:"contact_phone"`

		FullAddress string `json:"full_address"`
	}
)

func (r *ReceiptUserInfo) GetJson() (res string, err error) {
	if r == nil {
		return
	}
	var bt []byte
	if bt, err = json.Marshal(r); err != nil {
		return
	}
	res = string(bt)
	return
}
func (r *ArgCreateOrderFromCart) GetShopIds() (shopIds []int64) {
	shopIds = make([]int64, 0, len(r.SkuItems))
	mapShopIds := make(map[int64]bool, len(r.SkuItems))
	for _, item := range r.SkuItems {

		if _, ok := mapShopIds[item.ShopId]; !ok {
			shopIds = append(shopIds, item.ShopId)
			mapShopIds[item.ShopId] = true
		}
	}
	return
}

func (r *ArgCreateOrderFromCart) Default(c *base.Context) (err error) {
	if err = r.InitHeaderInfo(c.GinContext); err != nil {
		return
	}
	if err = r.RequestUser.InitRequestUser(c); err != nil {
		return
	}

	r.TimeNow = base.GetNowTimeNormal()

	if err = r.validateSku(); err != nil {
		return
	}

	if err = r.validateType(); err != nil {
		return
	}

	return
}

func (r *ArgCreateOrderFromCart) GetSkuIdMap() (res map[string]string) {
	res = make(map[string]string, len(r.SkuItems))
	for _, sku := range r.SkuItems {
		res[sku.SkuId] = sku.SkuId
	}
	return
}

func (r *ArgCreateOrderFromCart) validateType() (err error) {
	var MapOrderFromType map[string]string
	if MapOrderFromType, err = SliceOrderFromType.GetMapAsKeyString(); err != nil {
		return
	}
	if _, ok := MapOrderFromType[r.Type]; !ok {
		err = fmt.Errorf("创建订单来源参数错误")
		return
	}
	return
}

//重新组织SkuString的值
func (r *ArgCreateOrderFromCart) RestSkuStringWithSkuItems() (err error) {

	if r.SkuItems == nil {
		return
	}
	//var bt []byte
	//if bt, err = json.Marshal(r.SkuItems); err != nil {
	//	return
	//}
	//r.SkuString = string(bt)
	return
}

func (r *ArgCreateOrderFromCart) validateSku() (err error) {
	//if r.SkuString == "" {
	//	err = fmt.Errorf("请选择要付款的商品")
	//	return
	//}
	//if err = json.Unmarshal([]byte(r.SkuString), &r.SkuItems); err != nil {
	//	err = fmt.Errorf("参数异常")
	//	return
	//}
	if len(r.SkuItems) == 0 {
		err = fmt.Errorf("未选择要付款的商品")
		return
	}
	var skuItems = make([]*app_param.ArgOrderFromCartItem, 0, len(r.SkuItems))
	for _, item := range r.SkuItems {

		item.Default()
		if item.SkuId == "" {
			err = fmt.Errorf("您选择的商品数据异常(sku_id)")
			return
		}
		if item.SpuId == "" {
			err = fmt.Errorf("您选择的商品数据异常(spu_id)")
			return
		}
		if err = item.ValidateCategory(); err != nil {
			return
		}
		if !item.Checked {
			continue
		}
		skuItems = append(skuItems, item)
	}
	r.SkuItems = skuItems
	return
}
