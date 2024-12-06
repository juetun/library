package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	//数美审核
	ShuMeiAudit struct {
		CommonAudit
	}
)

func (r *ShuMeiAudit) Do(item AuditParametersInterface) (result *ApplyResult, err error) {
	result = &ApplyResult{Status: DataChatStatusOk}
	if item.GetIsSynchronous() == IsSynchronousNo { //如果是异步审核
		result.Status = DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewShuMeiAudit(Ctx *base.Context, Context context.Context) AuditClient {
	res := &ShuMeiAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
