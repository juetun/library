package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	//平台人工审核
	ToolClientAudit struct {
		CommonAudit
	}
)

func (r *ToolClientAudit) Do(item app_param.AuditParametersInterface) (result *app_param.ApplyResult, err error) {
	result = &app_param.ApplyResult{}
	return
}

func NewToolClientAudit(Ctx *base.Context, Context context.Context) app_param.AuditClient {
	res := &ToolClientAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
