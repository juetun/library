package order_create

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	ArgCreateOrderFromCart struct {
		RequestUser       app_param.RequestUser             `json:"-" form:"-"`
		SkuString         string                            `json:"sku_item" form:"sku_item"`
		Amount            string                            `json:"amount" form:"amount"` // 总金额
		App               string                            `json:"app" form:"app"`
		BuyChannel        string                            `json:"buy_channel" form:"buy_channel"` // 终端渠道号
		BuyClient         string                            `json:"buy_client" form:"buy_client"`   // 终端类型
		AppVersion        string                            `json:"app_version" form:"app_version"` // app版本
		Status            uint8                             `json:"status" form:"status"`           // 订单状态
		AddressId         int64                             `json:"address_id" form:"address_id"`   // 收货地址
		Express           string                            `json:"express" form:"express"`         // 默认快递信息
		PayType           uint8                             `json:"pay_type" form:"pay_type"`       // 支付类型
		FlagTester        uint8                             `json:"flag_tester" form:"flag_tester"`
		Type              string                            `json:"type" form:"Type"` //数据操作路径
		ReceiptUserInfo   *ReceiptUserInfo                  `json:"receipt_user_info" form:"receipt_user_info"`
		SkuItems          []*app_param.ArgOrderFromCartItem `json:"-" form:"-"`
		TimeNow           base.TimeNormal                   `json:"-" form:"-"`
		GetDataTypeCommon base.GetDataTypeCommon            `json:"-" form:"-"`
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

func (r *ArgCreateOrderFromCart) validateSku() (err error) {
	if r.SkuString == "" {
		err = fmt.Errorf("请选择要付款的商品")
		return
	}
	if err = json.Unmarshal([]byte(r.SkuString), &r.SkuItems); err != nil {
		err = fmt.Errorf("参数异常")
		return
	}
	if len(r.SkuItems) == 0 {
		err = fmt.Errorf("未选择要付款的商品")
		return
	}

	for _, item := range r.SkuItems {
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
	}
	return
}
