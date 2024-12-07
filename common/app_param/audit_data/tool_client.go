package audit_data

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	//平台人工审核
	ToolClientAudit struct {
		CommonAudit
	}
)

func (r *ToolClientAudit) Do(item AuditParametersInterface) (result *ApplyResult, err error) {
	result = &ApplyResult{Status: DataChatStatusOk}
	if item.GetIsSynchronous() == IsSynchronousNo { //如果是异步审核
		result.Status = DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewToolClientAudit(Ctx *base.Context, Context context.Context) AuditClient {
	res := &ToolClientAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
