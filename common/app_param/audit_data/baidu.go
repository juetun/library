package audit_data

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	//百度审核
	BaiDuAudit struct {
		CommonAudit
	}
)

func (r *BaiDuAudit) Do(item AuditParametersInterface) (result *ApplyResult, err error) {
	result = &ApplyResult{Status: DataChatStatusOk}
	if item.GetIsSynchronous() == IsSynchronousNo { //如果是异步审核
		result.Status = DataChatStatusWaiting
		result.Message = "审核中..."
	}
	return
}

func NewBaiDuAudit(Ctx *base.Context, Context context.Context) AuditClient {
	res := &BaiDuAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
