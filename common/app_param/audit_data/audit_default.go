package audit_data

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	//默认为审核
	CommonAudit struct {
		Ctx     *base.Context   `json:"-"`
		Context context.Context `json:"-"`
	}
	DefaultAudit struct {
		CommonAudit
	}
)

func (r *DefaultAudit) Do(item AuditParametersInterface) (result *ApplyResult, err error) {
	result = &ApplyResult{Status: DataChatStatusOk}
	if item.GetIsSynchronous() == IsSynchronousNo { //如果是异步审核
		result.Status = DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewDefaultAudit(Ctx *base.Context, Context context.Context) AuditClient {
	res := &DefaultAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
