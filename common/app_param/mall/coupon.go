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
		TimeNow base.TimeNormal        `json:"time_now" form:"time_now"`
		UserHid int64                  `json:"user_hid" form:"user_hid"`
		Amount  string                 `json:"amount" form:"amount"` //商品总金额
		ShopSpu []*ArgCanUseCouponShop `json:"shop_spu" form:"shop_spu"`
	}
	ArgCanUseCouponShop struct {
		ShopId    int64                 `json:"shop_id" form:"shop_id"` //店铺ID
		Amount    string                `json:"amount" form:"amount"`   //商品总金额
		SpuCoupon []*ArgCanUseCouponSpu `json:"spu_coupon" form:"spu_coupon"`
	}
	ArgCanUseCouponSpu struct {
		SpuId  string `json:"spu_id" form:"spu_id"` //商品ID
		Amount string `json:"amount" form:"amount"` //商品金额
	}

	ResultGetCanUseCoupon struct {
		PlatCoupon    *CanUsePlatCoupon           `json:"plat_coupon,omitempty"`     //平台券信息
		MapShopCoupon map[int64]*CanUsePlatCoupon `json:"map_shop_coupon,omitempty"` //店铺优惠券信息
		MapSpuCoupon  map[int64]*CanUsePlatCoupon `json:"map_spu_coupon,omitempty"`  //商品优惠券信息
		DecrAmount    string                      `json:"decr_amount,omitempty"`     //总扣减金额
	}

	CanUsePlatCoupon struct {
		CurrentUse *CouponInfo   `json:"current_use,omitempty"` //当前选中的最优秀优惠券
		CanUse     []*CouponInfo `json:"can_use,omitempty"`     //当前账号可使用的所有优惠券
		DecrAmount string        `json:"decr_amount,omitempty"` //总扣减金额
	}

	CouponInfo struct {
		ID             int64  `json:"id"`    //用户优惠券编号(用户ID 和优惠券ID组合的唯一号)
		Title          string `json:"title"` //优惠券名称
		TitleMark      string `json:"title_mark"`
		Status         uint8  `json:"status"`
		StatusName     string `json:"status_name"`
		CouponTypeName string `json:"coupon_type_name"`
		CouponID       string `json:"coupon_id"`  //用户优惠券编号(优惠券ID)
		StartTime      string `json:"start_time"` //有效期开始时间
		OverTime       string `json:"over_time"`  //有效期结束时间
		CreateTime     string `json:"create_time"`
		UserMark       string `json:"user_mark"` //使用说明
		UseAreaMark    string `json:"use_area_mark"`
		UseAreaLabel   string `json:"use_area_label"`
		Full           string `json:"full"`
		Decr           string `json:"decr"`
		Rebate         string `json:"rebate"`
		Cate           uint8  `json:"cate"`
		CateName       string `json:"cate_name"`
		EffectTimeDesc string `json:"effect_time_desc"`
		Disabled       bool   `json:"disabled"` //数据不合法(过期 或已删除等状态)
		CanUse         bool   `json:"can_use"`  //当前是否能够使用（优惠券使用期限未到false ）
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
