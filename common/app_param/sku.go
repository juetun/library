package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	SkuDataTypeSku   = "sku_data"  //SKU信息
	SkuDataTypeStock = "sku_stock" //库存
	SkuUploadSkuImg  = "sku_img"   // SKU图片
)

type (
	ArgGetSkuDataStringIds struct {
		base.GetDataTypeCommon
		IdsCompose []*SkuAndSpuIdCompose `json:"ids_compose"` //ID参数组合
		DataTypes  []string              `json:"data_types"`  //获取数据类型 取值范围和获取SPU的数据类型一致
	}
	SkuAndSpuIdCompose struct {
		SpuId string `json:"spu_id"`
		SkuId string `json:"sku_id"`
	}
)
func (r *ArgGetSkuDataStringIds) Default(ctx *base.Context) (err error) {

	return
}

func (r *ArgGetSkuDataStringIds) GetAllSpuIds() (res []string) {

	var (
		l        = len(r.IdsCompose)
		mapSpuId = make(map[string]bool, l)
	)

	res = make([]string, 0, l)
	for _, item := range r.IdsCompose {
		if _, ok := mapSpuId[item.SpuId]; !ok {
			mapSpuId[item.SpuId] = true
			res = append(res, item.SpuId)
		}
	}
	return
}

func (r *ArgGetSkuDataStringIds) GetAllPks() (res []string) {
	res = make([]string, 0, len(r.IdsCompose))
	for _, item := range r.IdsCompose {
		res = append(res, item.GetPk())
	}
	return
}

func (r *SkuAndSpuIdCompose) GetPk() (res string) {
	res = fmt.Sprintf("%s_%s", r.SpuId, r.SkuId)
	return
}
