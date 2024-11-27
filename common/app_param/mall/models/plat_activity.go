package models

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
	"time"
)

var (
	ActivityBeforeTime = time.Hour * 30
)

const (
	PlatActivityDataTypeSpu uint8 = iota + 1 //平台活动数据来源
)
const (
	PlatActivityBindDataTypeCoupon  uint8 = iota + 1 //活动绑定优惠券
	PlatActivityBindDataTypeAtPrice                  //指定价个购买
)
const (
	PlatActivityStatusUse        uint8 = iota + 1 //有效
	PlatActivityStatusLapse                       //失效
	PlatActivityStatusManuscript                  //草稿
	PlatActivityStatusInit                        //初始化中
	PlatActivityStatusWaiting                     //待上架
	PlatActivityStatusDelete                      //已删除
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
