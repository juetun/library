package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
)

type (
	//平台自审程序
	PrivateAudit struct {
		CommonAudit
	}
)

func (r *PrivateAudit) Do(item app_param.AuditParametersInterface) (result *app_param.ApplyResult, err error) {
	result = &app_param.ApplyResult{}
	return
}

func NewPrivateAudit(Ctx *base.Context, Context context.Context) app_param.AuditClient {
	res := &PrivateAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
