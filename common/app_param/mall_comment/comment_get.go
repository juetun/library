package mall_comment

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgSpuComment struct {
		SpuIds            []string               `json:"spu_ids" form:"spu_ids"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"common"  form:"common"`
	}
	ResultSpuComment map[string]*OrderComment
)

func (r *ArgSpuComment) Default(ctx *base.Context) (err error) {

	return
}
