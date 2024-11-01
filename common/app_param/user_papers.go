package app_param

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/const_apply"
)

//用户资质是否需要填写类型

const (
	PapersDateHave      uint8 = iota + 1 // Papers.DateExpiry有时间
	PapersDateHasNot                     // Papers.DateExpiry没有时间
	PapersDateMustInput                  // 必填
)

var SliceMustDate = base.ModelItemOptions{
	{
		Value: PapersDateHave, //有时间
		Label: "有",
	},
	{
		Value: PapersDateHasNot, //没有时间
		Label: "无",
	},
	{
		Value: PapersDateMustInput, //有且必填
		Label: "必填",
	},
}

//	map[uint8]string{
//	PaperMustDateNotNeed: "不填",
//	PaperMustDateYes:     "必填",
//	PaperMustDateNo:      "可不填",
//}

const (
	DataPapersGroupShopPropertyRadio    = "radio"    //单选
	DataPapersGroupShopPropertyCheckbox = "checkbox" //多选
	DataPapersGroupShopPropertySelect   = "select"   //下拉菜单
)

const (
	AuditingStatusOk      = const_apply.ApplyStatusOk       // 品牌状态初始化
	AuditingStatusInit    = const_apply.ApplyStatusInit     // 品牌状态审核通过
	AuditingStatusFailure = const_apply.ApplyStatusFailure  // 品牌状态审核失败
	AuditingStatusWaiting = const_apply.ApplyStatusAuditing //待审核
)

const (
	PapersIsNeverExpiresYes uint8 = iota + 1 //永久有效
	PapersIsNeverExpiresNo                   //有过期时间
)
const (
	ShopIsFactoryYes uint8 = iota + 1 //如果是厂家在
	ShopIsFactoryNo                   //不是厂家
)

var (
	SliceShopIsFactory = base.ModelItemOptions{
		{Value: ShopIsFactoryYes, Label: "生产厂家",},
		{Value: ShopIsFactoryNo, Label: "经销商",},
	}
	SlicePapersIsNeverExpires = base.ModelItemOptions{
		{
			Value: PapersIsNeverExpiresNo,
			Label: "有过期时间",
		},
		{
			Value: PapersIsNeverExpiresYes,
			Label: "永久有效",
		},
	}
	Slice = base.ModelItemOptions{
		{
			Label: "不填",
			Value: PapersDateHasNot,
		},
		{
			Label: "选填",
			Value: PapersDateHave,
		},
		{
			Label: "必填",
			Value: PapersDateMustInput,
		},
	}

	SliceAuditingStatus = base.ModelItemOptions{
		{
			Value: AuditingStatusOk,
			Label: "初始化",
		},
		{
			Value: AuditingStatusInit,
			Label: "审核通过",
		},
		{
			Value: AuditingStatusFailure,
			Label: "审核失败",
		},
		{
			Value: AuditingStatusWaiting,
			Label: "待审核",
		},
	}
	SliceDataPapersGroupShopProperty = base.ModelItemOptions{
		{
			Value: DataPapersGroupShopPropertyRadio,
			Label: "单选",
		},
		{
			Value: DataPapersGroupShopPropertyCheckbox,
			Label: "多选",
		},
		{
			Value: DataPapersGroupShopPropertySelect,
			Label: "下拉菜单",
		},
	}
)

type (
	DataPapersGroupShopProperty struct {
		ShowType     string `json:"show_type"` //DataPapersGroupShopPropertyRadio DataPapersGroupShopPropertyCheckbox  DataPapersGroupShopPropertySelect
		ShowTypeName string `json:"show_type_name,omitempty"`
	}

	ResultGetUpdateBatch struct {
		UserHid        int64  `json:"user_hid"`
		ShopId         int64  `json:"shop_id"`
		BatchId        string `json:"batch_id"`
		Status         int8   `json:"status"`
		Type           string `json:"type"`
		IsInit         uint8  `json:"is_init"`
		CanSubmitApply bool   `json:"can_submit_apply"`
	}
)

func (r *DataPapersGroupShopProperty) ParseShowType() (res string) {
	if r.ShowType == "" { //默认类型
		r.ShowType = DataPapersGroupShopPropertyRadio
	}
	MapDataPapersGroupShopProperty, _ := SliceDataPapersGroupShopProperty.GetMapAsKeyString()
	var ok bool
	if res, ok = MapDataPapersGroupShopProperty[r.ShowType]; ok {
		r.ShowTypeName = res
		return
	}

	res = fmt.Sprintf("未知类型(%s)", r.ShowType)
	r.ShowTypeName = res
	return
}

func (r *DataPapersGroupShopProperty) ToString() (res string, err error) {
	var bt []byte
	bt, err = json.Marshal(r)
	res = string(bt)
	return
}
