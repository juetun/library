package mall

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	ArgShopRefundAddr struct {
		ShopId int64 `json:"shop_id" form:"shop_id"`
	}

	ResultShopRefundAddr struct {
		ID           int64  `json:"id"`
		ShopID       int64  `json:"shop_id"`
		ProvinceId   int64  `json:"province_id"`
		CityId       int64  `json:"city_id"`
		AreaId       int64  `json:"area_id"`
		ProvinceName string `json:"province_name"`
		CityName     string `json:"city_name"`
		AreaName     string `json:"area_name"`
		Address      string `json:"address"`
		AcceptUser   string `json:"accept_user"`
		AcceptPhone  string `json:"accept_phone"`
		FullAddr     string `json:"full_addr"` //地址全名
	}
)

func (r *ArgShopRefundAddr) Default(ctx *base.Context) (err error) {

	return
}

func (r *ResultShopRefundAddr) InitFullAddress() (res string) {
	r.FullAddr = fmt.Sprintf("%s %s %s %s", r.ProvinceName, r.CityName, r.AreaName, r.Address)
	res = r.FullAddr
	return
}
