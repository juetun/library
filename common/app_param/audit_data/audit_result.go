package audit_data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

type (
	AuditResult struct {
		ActionType int             `json:"action_type"` //当前审核的数据类型（order_comment:订单评论;comment:普通数据评论 聊天信息评论）
		Ctx        *base.Context   `json:"-"`
		Context    context.Context `json:"-"`
	}
	ArgSyncAuditRecord struct {
		MsgId       string                 `json:"msg_id,omitempty" form:"msg_id"`
		OtherParams map[string]interface{} `json:"other_params,omitempty" form:"other_params"`
		ApplyResult *ApplyResult           `json:"apply_result,omitempty" form:"apply_result"`
	}
	AuditResultOption func(property *AuditResult)
)

//审核结果同步
func (r *AuditResult) ResultSyncInfo(parameters *ArgSyncAuditRecord, appNames ...string) (err error) {
	var appName string
	if len(appNames) > 0 {
		appName = appNames[0]
	}
	switch r.ActionType {
	case AuditTypeOrderComment: //订单评论
		appName = app_param.AppNameMallOrderComment
	case AuditTypeComment: //社交评论
		appName = app_param.AppNameComment
	case AuditTypeChat: //聊天信息
		appName = app_param.AppNameChat
	}
	if appName == "" {
		err = fmt.Errorf("没有设置请求的服务")
		return
	}
	if err = r.syncInfo(appName, parameters); err != nil {
		return
	}
	return
}

func (r *AuditResult) syncInfo(appName string, param *ArgSyncAuditRecord) (err error) {
	if param == nil {
		return
	}

	var bodyContent []byte
	if bodyContent, err = json.Marshal(param); err != nil {
		return
	}

	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     r.Ctx,
		Method:      http.MethodPost,
		AppName:     appName,
		URI:         "/audit_data/sync_result",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		BodyJson:    bodyContent,
		Header:      http.Header{},
	}

	req := rpc.NewHttpRpc(&params).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.Body; len(body) == 0 {
		return
	}

	var resResult struct {
		Code int `json:"code"`
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
		Msg string `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}
	return
}

func NewAuditResult(options ...AuditResultOption) (res *AuditResult) {
	res = &AuditResult{}
	for _, item := range options {
		item(res)
	}
	return
}

func AuditResultActionType(actionType int) AuditResultOption {
	return func(property *AuditResult) {
		property.ActionType = actionType
	}
}

func AuditResultContext(context context.Context) AuditResultOption {
	return func(property *AuditResult) {
		property.Context = context
	}
}

func AuditResultCtx(ctx *base.Context) AuditResultOption {
	return func(property *AuditResult) {
		property.Ctx = ctx
	}
}
