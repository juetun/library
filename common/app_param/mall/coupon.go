package mall

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

type (
	ArgGetCouponBindData struct {
		CouponId string   `json:"coupon_id" form:"coupon_id"`
		SpuIds   []string `json:"spu_ids" form:"spu_ids"`
		DataType string   `json:"data_type" form:"data_type"`
	}
	ResGetCouponBindData map[string]bool

	ArgGetCanUseCoupon struct {
		TimeNow base.TimeNormal      `json:"time_now" form:"time_now"`
		UserHid int64                `json:"user_hid" form:"user_hid"`
		ShopSpu map[int64]ArgShopSpu `json:"shop_spu" form:"shop_spu"`
	}
	ArgShopSpu struct {
		ShopId int64    `json:"shop_id" form:"shop_id"`
		SpuId  []string `json:"spu_id" form:"spu_id"`
	}

	ResultGetCanUseCoupon struct {
		PlatCoupon    ResultGetCanUsePlatCoupon `json:"plat_coupon"`     //平台券信息
		MapShopCoupon map[int64]ShopCouponList  `json:"map_shop_coupon"` //店铺优惠券信息
	}
	ShopCouponList            []*ResultCanUseCoupon
	ResultGetCanUsePlatCoupon struct {
		UserCouponId int64  `json:"user_coupon_id"`
		Label        string `json:"label"`
	}
	ResultCanUseCoupon struct {
		UserCouponId int64  `json:"user_coupon_id"`
		Label        string `json:"label"`
	}
)

func (r *ArgGetCouponBindData) Default(ctx *base.Context) (err error) {
	return
}

func (r *ArgGetCanUseCoupon) Default(ctx *base.Context) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}

func (r *ShopCouponList) GetLabels(res []string) {
	res = make([]string, 0, len(*r))
	for _, item := range *r {
		res = append(res, item.Label)
	}
	return
}

//根据参数获取可使用的优惠券
func GetCouponInfo(ctx *base.Context, argStruct *ArgGetCanUseCoupon) (res *ResultGetCanUseCoupon, err error) {
	res = &ResultGetCanUseCoupon{}

	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameMallActivity,
		URI:         "/coupon/get_can_use_coupon",
		Value:       url.Values{},
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}
	if params.BodyJson, err = json.Marshal(argStruct); err != nil {
		return
	}

	var (
		body      []byte
		resResult struct {
			Code int                    `json:"code"`
			Data *ResultGetCanUseCoupon `json:"data"`
			Msg  string                 `json:"message"`
		}
		req = rpc.NewHttpRpc(&params).
			Send().
			GetBody()
	)

	if err = req.Error; err != nil {
		return
	}

	if body = req.Body; len(body) == 0 {
		return
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
