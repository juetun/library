package mall

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strings"
)

const (
	StockActTypeAdd  = "add"  //添加库存
	StockActTypeDecr = "decr" //减少库存
)

type (
	StockOperateItem struct {
		SpuId        string `json:"spu_id"`
		SkuId        string `json:"sku_id_op"`
		SrcPk        string `json:"src_pk"`       //源SKU的PK值
		Num          int64  `json:"num"`          //需要加上或减去的库存数
		ActualityNum int64  `json:"-"`            //实际扣减库存数
		NotForce     bool   `json:"not_force"`    //非强制扣减或加上库存 （false-强制减库存, true-非强制减，扣到0即止 定金预售功能使用）
		ActType      string `json:"act_type"`     //添加库存，还是减少库存 add:添加库存 decr:减少库存
		HasUpDb      bool   `json:"has_up_db"`    //true-更新了数据的数据
		HasUpCache   bool   `json:"has_up_cache"` //true-更新了缓存的数据
	}
	ArgAddOrDecrStock struct {
		SkuInfoString string `json:"sku_info_string"`
		SkuStockItems []*StockOperateItem
	}

	ResultAddOrDecrStock struct {
		Result     bool                    `json:"result"`
		ResultList ResultStockOperateItems `json:"result_list"`
	}

	ResultStockOperateItems []StockOperateResultItem

	StockOperateResultItem struct {
		SpuId        string `json:"spu_id"`
		SkuId        string `json:"sku_id"`
		SrcPk        string `json:"src_pk"`        //源SKU的PK值
		Num          int64  `json:"num"`           //需要加上或减去的库存数
		ActualityNum int64  `json:"actuality_num"` //实际扣减库存数
		NotForce     bool   `json:"not_force"`     //非强制扣减或加上库存 （false-强制减库存, true-非强制减，扣到0即止 定金预售功能使用）
		ActType      string `json:"act_type"`      //添加库存，还是减少库存 add:添加库存 decr:减少库存
		HaveError    bool   `json:"have_error"`    //是否有错误
		ErrorMessage string `json:"error_message"` //错误提示
	}
)

func (r *ArgAddOrDecrStock) Default(ctxt *base.Context) (err error) {

	if r.SkuInfoString == "" {
		err = fmt.Errorf("请选择你要操作的SKU信息")
		return
	}

	if err = json.Unmarshal([]byte(r.SkuInfoString), &r.SkuStockItems); err != nil {
		return
	}
	return
}

func (r *StockOperateResultItem) SetStockOperateItem(data *StockOperateItem) {
	r.SkuId = data.SkuId
	r.SpuId = data.SpuId
	r.Num = data.Num
	r.SrcPk = data.SrcPk
	r.ActualityNum = data.ActualityNum
	r.NotForce = data.NotForce
	r.ActType = data.ActType
}

func (r ResultStockOperateItems) Error() (err error) {
	var msg = make([]string, 0, len(r))
	for _, item := range r {
		if !item.HaveError {
			continue
		}
		msg = append(msg, item.ErrorMessage)
	}
	if len(msg) == 0 {
		return
	}
	err = fmt.Errorf(strings.Join(msg, ","))
	return
}
