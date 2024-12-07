package action

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/audit_data"
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

func (r *DefaultAudit) Do(item audit_data.AuditParametersInterface) (result *audit_data.ApplyResult, err error) {
	result = &audit_data.ApplyResult{Status: audit_data.DataChatStatusOk}
	if item.GetIsSynchronous() == audit_data.IsSynchronousNo { //如果是异步审核
		result.Status = audit_data.DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewDefaultAudit(Ctx *base.Context, Context context.Context) audit_data.AuditClient {
	res := &DefaultAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
