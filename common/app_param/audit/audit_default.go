package audit

import (
	"context"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param"
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

func (r *DefaultAudit) Do(item app_param.AuditParametersInterface) (result *app_param.ApplyResult, err error) {
	result = &app_param.ApplyResult{}
	return
}

func NewDefaultAudit(Ctx *base.Context, Context context.Context) ( app_param.AuditClient) {
	res := &DefaultAudit{}
	res.CommonAudit.Ctx = Ctx
	res.CommonAudit.Context = Context
	return res
}
