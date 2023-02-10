package order_create

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	ArgCreateOrderFromCart struct {
		app_param.RequestUser
		SkuString  string                            `json:"sku_items" form:"sku_items"`
		SkuItems   []*app_param.ArgOrderFromCartItem `json:"-"`
		Amount     string                            `json:"amount" form:"amount"` // 总金额
		BuyChannel string                            `json:"buy_channel"`          // 终端渠道号
		BuyClient  string                            `json:"buy_client"`           // 终端类型
		AppVersion string                            `json:"app_version"`          // app版本
		Status     uint8                             `json:"status"`               // 订单状态
		AddressId  string                            `json:"address_id"`           // 收货地址
		Express    string                            `json:"express"`              // 默认快递信息
		PayType    uint8                             `json:"pay_type"`             // 支付类型
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
	if err = json.Unmarshal([]byte(r.SkuString), &r.SkuItems); err != nil {
		err = fmt.Errorf("参数异常")
		return
	}
	if len(r.SkuItems) == 0 {
		err = fmt.Errorf("未选择要操作的商品")
		return
	}

	for _, item := range r.SkuItems {
		if item.SkuId == "" {
			err = fmt.Errorf("您选择的商品数据异常")
			return
		}
	}

	return
}