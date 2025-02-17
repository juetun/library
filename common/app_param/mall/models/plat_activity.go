package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"time"
)

var (
	ActivityBeforeTime = time.Hour * 30
)

const (
	PlatActivityDataTypeSpu uint8 = iota + 1 //平台活动数据来源
)
const (
	PlatActivityNone                uint8 = iota //未参加活动
	PlatActivityBindDataTypeCoupon               //活动绑定优惠券
	PlatActivityBindDataTypeAtPrice              //指定价个购买
)
const (
	PlatActivityStatusUse        uint8 = iota + 1 //有效
	PlatActivityStatusLapse                       //失效
	PlatActivityStatusManuscript                  //草稿
	PlatActivityStatusInit                        //初始化中
	PlatActivityStatusWaiting                     //待上架
	PlatActivityStatusDelete                      //已删除
	PlatActivityStatusPre                         //活动预热中
	PlatActivityWaitingStart                      //活动待开始
)

var (
	PlatActivityStatusCanUpdateStatus = map[uint8][]uint8{
		PlatActivityStatusUse:        {PlatActivityStatusWaiting}, //可使用
		PlatActivityStatusWaiting:    {PlatActivityStatusUse},     //待上架
		PlatActivityStatusLapse:      {},                          //已失效
		PlatActivityStatusManuscript: {PlatActivityStatusUse},     //草稿中
		PlatActivityStatusInit:       {},                          //初始化
		PlatActivityStatusDelete:     {},                          //已删除
	}
	SlicePlatActivityStatus = base.ModelItemOptions{
		{
			Value: PlatActivityStatusUse,
			Label: "可使用",
		},
		{
			Value: PlatActivityStatusWaiting,
			Label: "待上架",
		},
		{
			Value: PlatActivityStatusLapse,
			Label: "已失效",
		},
		{
			Value: PlatActivityStatusManuscript,
			Label: "草稿中",
		},
		{
			Value: PlatActivityStatusInit,
			Label: "初始化",
		},
		{
			Value: PlatActivityStatusDelete,
			Label: "已删除",
		},
		{
			Value: PlatActivityStatusPre,
			Label: "预热中",
		},
		{
			Value: PlatActivityWaitingStart,
			Label: "未开始",
		},
	}
	SlicePlatActivityDataType = base.ModelItemOptions{
		{
			Label: "电商SPU",
			Value: PlatActivityDataTypeSpu,
		},
	}
	//活动绑定数据（如优惠券 、代金券、1元购 0元购等）
	PlatActivityBindDataType = base.ModelItemOptions{
		{
			Label: "",
			Value: PlatActivityNone,
		},
		{
			Label: "绑定优惠券",
			Value: PlatActivityBindDataTypeCoupon,
		},
		{
			Label: "指定价包邮",
			Value: PlatActivityBindDataTypeAtPrice,
		},
	}
)

type (
	//商品活动
	ArgPlatActivityByIds struct {
		base.ArgGetByStringIds
	}
	PlatActivity struct {
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

	OtherAttrValue struct {
		Price string `json:"price"` //价格
	}
)

func (r *PlatActivity) GetOtherAttr() (otherAttrValue *OtherAttrValue, err error) {
	otherAttrValue = &OtherAttrValue{}
	defer func() {
		if err != nil {
			return
		}
		err = otherAttrValue.Default()
	}()
	if r.OtherAttr == "" {
		return
	}
	err = json.Unmarshal([]byte(r.OtherAttr), otherAttrValue)
	return
}

func (r *PlatActivity) InitOtherAttr(price string, otherAttrValues ...*OtherAttrValue) (err error) {
	var otherAttrValue *OtherAttrValue
	if len(otherAttrValues) > 0 {
		otherAttrValue = otherAttrValues[0]
	} else {
		otherAttrValue = &OtherAttrValue{}
	}
	switch r.BindDataType {
	case PlatActivityBindDataTypeAtPrice: //指定价格
		otherAttrValue.Price = price

	default:

	}
	var bt []byte
	if bt, err = json.Marshal(otherAttrValue); err != nil {
		return
	}
	r.OtherAttr = string(bt)
	return
}

func (r *OtherAttrValue) Default() (err error) {
	if r.Price == "" {
		r.Price = "0.00"
	}
	return
}

func (r *ArgPlatActivityByIds) Default(ctx *base.Context) (err error) {

	return
}

//获取平台活动信息
func GetPlatActivity(ctx *base.Context, args *ArgPlatActivityByIds) (res map[string]*PlatActivity, err error) {
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
		Code int                      `json:"code"`
		Data map[string]*PlatActivity `json:"data"`
		Msg  string                   `json:"message"`
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
