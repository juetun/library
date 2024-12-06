package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	//数美审核
	ShuMeiAudit struct {
		CommonAudit
	}
)

func (r *ShuMeiAudit) Do(item app_param.AuditParametersInterface) (result *app_param.ApplyResult, err error) {
	result = &app_param.ApplyResult{}
	return
}

func NewShuMeiAudit(Ctx *base.Context, Context context.Context) app_param.AuditClient {
	res := &ShuMeiAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
