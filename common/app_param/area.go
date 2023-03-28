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
)

func (r *ArgGetByCodes) Default(ctx *base.Context) (err error) {

	mapV, _ := SliceDtAreaRegionNameType.GetMapAsKeyString()
	if _, ok := mapV[r.NameType]; !ok {
		err = fmt.Errorf("对不起，当前不支持你选择的类型(name_type:%s)", r.NameType)
		return
	}
	return
}
