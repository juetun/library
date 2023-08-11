package mall

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgShopRefundAddr struct {
		ShopId int64 `json:"shop_id" form:"shop_id"`
	}

	ResultShopRefundAddr struct {
		ID          int64  `json:"id"`
		ShopID      int64  `json:"shop_id"`
		ProvinceId  int64  `json:"province_id"`
		CityId      int64  `json:"city_id"`
		AreaId      int64  `json:"area_id"`
		Address     string `json:"address"`
		AcceptUser  string `json:"accept_user"`
		AcceptPhone string `json:"accept_phone"`
		FullAddr    string `json:"full_addr"` //地址全名
	}
)

func (r *ArgShopRefundAddr) Default(ctx *base.Context) (err error) {

	return
}
