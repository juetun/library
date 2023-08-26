package mall_comment

import "github.com/juetun/base-wrapper/lib/base"

type (
	ArgSpuComment struct {
		TopNum            int64                  `json:"top_num" form:"top_num"`
		SpuIds            []string               `json:"spu_ids" form:"spu_ids"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"common"  form:"common"`
	}
	ResultSpuComment map[string]*OrderComment
)

func (r *ArgSpuComment) Default(ctx *base.Context) (err error) {
	if r.TopNum == 0 {
		r.TopNum = 2
	}
	return
}
