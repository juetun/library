package action

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/audit_data"
)

type (
	//平台人工审核
	ToolClientAudit struct {
		CommonAudit
	}
)

func (r *ToolClientAudit) Do(item audit_data.AuditParametersInterface) (result *audit_data.ApplyResult, err error) {
	result = &audit_data.ApplyResult{Status: audit_data.DataChatStatusOk}
	if audit_data.GetIsSynchronous() == audit_data.IsSynchronousNo { //如果是异步审核
		result.Status = audit_data.DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewToolClientAudit(Ctx *base.Context, Context context.Context) audit_data.AuditClient {
	res := &ToolClientAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
