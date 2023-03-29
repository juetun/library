package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	DtAreaRegionNameTypeDefault  = ""
	DtAreaRegionNameTypeComplete = "complete"
)

var (
	SliceDtAreaRegionNameType = base.ModelItemOptions{
		{
			Label: "普通",
			Value: DtAreaRegionNameTypeDefault,
		},
		{
			Label: "完整",
			Value: DtAreaRegionNameTypeComplete,
		},
	}
)

type (
	ArgGetByCodes struct {
		base.ArgGetByStringIds
		NameType string `json:"name_type" form:"name_type"`
	}

	ArgAreaGet struct {
	}
	ResultAreaGet     []*ResultAreaGetItem
	ResultAreaGetItem struct {
		Value         string               `json:"value"`
		Label         string               `json:"label"`
		ProvinceId    string               `json:"province_id"` // 省的编码
		Level         int                  `json:"level"`
		SelectCount   int                  `json:"select_count"`
		Checked       bool                 `json:"checked"`
		OtherChecked  bool                 `json:"other_checked"` //其他地方是否已经选中过了
		Indeterminate bool                 `json:"indeterminate"`
		Disabled      bool                 `json:"disabled"`
		Children      []*ResultAreaGetItem `json:"children"`
	}
)

func (r *ArgAreaGet) Default(context *base.Context) (err error) {

	return
}

func (r *ArgGetByCodes) Default(ctx *base.Context) (err error) {

	mapV, _ := SliceDtAreaRegionNameType.GetMapAsKeyString()
	if _, ok := mapV[r.NameType]; !ok {
		err = fmt.Errorf("对不起，当前不支持你选择的类型(name_type:%s)", r.NameType)
		return
	}
	return
}
