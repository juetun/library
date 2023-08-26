package mall_comment

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgSpuComment struct {
		SpuIds            []string               `json:"spu_ids" form:"spu_ids"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"common"  form:"common"`
	}
	ResultSpuComment struct {
		Code int                      `json:"code"`
		Data map[string]*OrderComment `json:"data"`
		Msg  string                   `json:"message"`
	}
)

func (r *ArgSpuComment) Default(ctx *base.Context) (err error) {

	return
}
