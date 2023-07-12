package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
)

const (
	SkuDataTypeSku   = "sku_data"  //SKU信息
	SkuDataTypeStock = "sku_stock" //库存
	SkuUploadSkuImg  = "sku_img"   // SKU图片
)

type (
	ArgGetSkuDataStringIds struct {
		common.HeaderInfo
		base.GetDataTypeCommon
		IdsCompose []*SkuAndSpuIdCompose `json:"ids_compose"` //ID参数组合
		DataTypes  []string              `json:"data_types"`  //获取数据类型 取值范围和获取SPU的数据类型一致
	}
	SkuAndSpuIdCompose struct {
		ShopId int64  `json:"shop_id"` //店铺ID非必须参数，特定场景使用
		SpuId  string `json:"spu_id"`  //SPU_id 必有参数,
		SkuId  string `json:"sku_id"`  //sku_id 必有参数,
	}
)

func (r *ArgGetSkuDataStringIds) Default(ctx *base.Context) (err error) {
	_ = r.InitHeaderInfo(ctx.GinContext)
	return
}

func (r *ArgGetSkuDataStringIds) GetAllIds() (shopIds []int64, spuIds, skuIds []string) {

	var (
		l         = len(r.IdsCompose)
		mapSkuId  = make(map[string]bool, l)
		mapShopId = make(map[int64]bool, l)
		mapSpuId  = make(map[string]bool, l)
	)
	shopIds = make([]int64, 0, l)
	spuIds = make([]string, 0, l)
	skuIds = make([]string, 0, l)
	for _, item := range r.IdsCompose {
		if _, ok := mapSkuId[item.SkuId]; !ok {
			mapSkuId[item.SkuId] = true
			skuIds = append(skuIds, item.SkuId)
		}

		if _, ok := mapShopId[item.ShopId]; !ok {
			mapShopId[item.ShopId] = true
			shopIds = append(shopIds, item.ShopId)
		}

		if _, ok := mapSpuId[item.SpuId]; !ok {
			mapSpuId[item.SpuId] = true
			spuIds = append(spuIds, item.SpuId)
		}
	}
	return
}

func (r *ArgGetSkuDataStringIds) GetAllSkuIds() (res []string) {

	var (
		l        = len(r.IdsCompose)
		mapSkuId = make(map[string]bool, l)
	)

	res = make([]string, 0, l)
	for _, item := range r.IdsCompose {
		if _, ok := mapSkuId[item.SkuId]; !ok {
			mapSkuId[item.SkuId] = true
			res = append(res, item.SkuId)
		}
	}
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
