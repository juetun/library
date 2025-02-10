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
	//商品活动
	ArgPlatActivityByIds struct {
		base.ArgGetByStringIds
	}
	PlatIntranetActivity struct {
		Id           int64            `json:"id"`
		Title        string           `json:"title"`
		KeyVal       string           `json:"key_val"`
		BindDataType uint8            `json:"bind_data_type"`
		OtherAttr    string           `json:"other_attr"`
		CouponId     string           `json:"coupon_id"`
		WarmUpTime   *base.TimeNormal `json:"warm_up_time"`
		StartTime    *base.TimeNormal `json:"start_time"`
		OverTime     *base.TimeNormal `json:"over_time"`
	}
)

func (r *ArgPlatActivityByIds) Default(ctx *base.Context) (err error) {

	return
}

//获取平台活动信息
func GetPlatActivity(ctx *base.Context, args *ArgPlatActivityByIds) (res PlatIntranetActivity, err error) {
	if len(args.Ids) == 0 {
		return
	}
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameMallActivity,
		URI:         "/plat_activity/get_by_ids",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}

	if ctx.GinContext != nil {
		params.Header.Set(app_obj.HttpHeaderInfo, ctx.GinContext.GetHeader(app_obj.HttpHeaderInfo))
	}
	if params.BodyJson, err = json.Marshal(args); err != nil {
		return
	}

	req := rpc.NewHttpRpc(&params).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.Body; len(body) == 0 {
		return
	}

	var resResult struct {
		Code int                  `json:"code"`
		Data PlatIntranetActivity `json:"data"`
		Msg  string               `json:"message"`
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
