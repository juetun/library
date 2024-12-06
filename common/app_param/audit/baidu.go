package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	//百度审核
	BaiDuAudit struct {
		CommonAudit
	}
)

func (r *BaiDuAudit) Do(item app_param.AuditParametersInterface) (result *app_param.ApplyResult, err error) {
	result = &app_param.ApplyResult{}
	return
}

func NewBaiDuAudit(Ctx *base.Context, Context context.Context) app_param.AuditClient {
	res := &BaiDuAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
