package app_param

import (
	"context"
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
)

//审核工具类型
const (
	ApplyToolTypeDefault uint8 = iota //默认为审核
	ApplyToolTypePrivate              //平台自审程序
	ApplyToolTypeBaiDu                //百度审核
	ApplyToolTypeShuMei               //数美审核
	ApplyToolTypeClient               //平台人工审核
)
const (
	DataChatStatusOk      uint8 = iota + 1 //审核通过
	DataChatStatusWaiting                  //待审核
	DataChatStatusFailure                  //审核失败
)

var (
	SliceDataChatStatus = base.ModelItemOptions{
		{
			Value: DataChatStatusOk,
			Label: "审核通过",
		},
		{
			Value: DataChatStatusWaiting,
			Label: "待审核",
		},
		{
			Value: DataChatStatusFailure,
			Label: "审核失败",
		},
	}
	SliceDataChatApplyToolType = base.ModelItemOptions{
		{
			Value: ApplyToolTypeDefault,
			Label: "~",
		},
		{
			Value: ApplyToolTypePrivate,
			Label: "平台自审",
		},
		{
			Value: ApplyToolTypeBaiDu,
			Label: "百度审核",
		},
		{
			Value: ApplyToolTypeShuMei,
			Label: "数美审核",
		},
		{
			Value: ApplyToolTypeClient,
			Label: "客服人工审核",
		},
	}
)

//审核信息
type (
	Client interface {
		Do(argByte []byte) (result *ApplyResult, err error)
	}
	AuditParametersInterface interface {
		GetArg() (res []byte)                                        //获取参数
		Audit(auditData *AuditData) (result *ApplyResult, err error) //审核逻辑
	}
	AuditData struct {
		client     Client                     `json:"-"`
		Ctx        base.Context               `json:"-"`
		Context    context.Context            `json:"-"`
		Parameters []AuditParametersInterface `json:"arg"`
	}
	AuditParametersText struct {
		MsgId  string       `json:"msg_id"`
		Text   []string     `json:"text"` //审核的图片列表
		Result *ApplyResult `json:"result"`
	}
	AuditParametersImg struct {
		MsgId   string   `json:"msg_id"`
		ImgUrls []string `json:"img_urls"` //审核的图片列表
	}
	AuditParametersVideoUrls struct {
		MsgId     string   `json:"msg_id"`
		VideoUrls []string `json:"video_urls"` //审核的视频链接
	}
	AuditParametersMusicUrls struct {
		MsgId     string   `json:"msg_id"`
		MusicUrls []string `json:"music_urls"` //审核的音频链接
	}
	ApplyResult struct {
		Status        uint8  `json:"status"`         //审核状态
		Message       string `json:"msg"`            //审核结果
		ErrorType     string `json:"e_type"`         //审核失败类型
		ApplyId       string `json:"apply_id"`       //审核请求ID
		ApplyType     uint8  `json:"apply_type"`     //审核类型
		ApplyResponse string `json:"apply_response"` //审核请求的响应
	}
	AuditDataOption func(property *AuditData)
)

func NewAuditData(options ...AuditDataOption) (res *AuditData) {
	res = &AuditData{
	}
	for _, item := range options {
		item(res)
	}
	return
}

func AuditDataClient(client Client) AuditDataOption {
	return func(property *AuditData) {
		property.client = client
	}
}
func AuditCtx(ctx base.Context) AuditDataOption {
	return func(property *AuditData) {
		property.Ctx = ctx
	}
}
func AuditContext(context context.Context) AuditDataOption {
	return func(property *AuditData) {
		property.Context = context
	}
}
func AuditParameters(parameters []AuditParametersInterface) AuditDataOption {
	return func(property *AuditData) {
		property.Parameters = parameters
	}
}

func (r *AuditParametersVideoUrls) Audit(auditData *AuditData) (result *ApplyResult, err error) {
	var argByte []byte
	argByte = r.GetArg()
	if result, err = auditData.client.Do(argByte); err != nil {
		return
	}
	return
}

func (r *AuditParametersMusicUrls) Audit(auditData *AuditData) (result *ApplyResult, err error) {
	var argByte []byte
	argByte = r.GetArg()
	if result, err = auditData.client.Do(argByte); err != nil {
		return
	}
	return
}

func (r *AuditParametersImg) Audit(auditData *AuditData) (result *ApplyResult, err error) {
	var argByte []byte
	argByte = r.GetArg()
	if result, err = auditData.client.Do(argByte); err != nil {
		return
	}
	return
}

func (r *AuditParametersText) Audit(auditData *AuditData) (result *ApplyResult, err error) {
	var argByte []byte
	argByte = r.GetArg()
	if result, err = auditData.client.Do(argByte); err != nil {
		return
	}
	return
}

func (r *AuditData) Apply() (applyResult *ApplyResult, err error) {
	applyResult = &ApplyResult{}
	applyResult.Status = DataChatStatusOk
	var (
		result *ApplyResult
	)
	for _, item := range r.Parameters {
		if result, err = item.Audit(r); err != nil {
			return
		}
		switch result.Status {
		case DataChatStatusOk: //审核通过
		case DataChatStatusWaiting: //待审核
			applyResult = result
		case DataChatStatusFailure: //审核失败
			applyResult = result
		}
	}
	return
}

func (r *AuditParametersVideoUrls) GetArg() (res []byte) {
	if len(r.VideoUrls) == 0 {
		return
	}
	res, _ = json.Marshal(r)
	return
}

func (r *AuditParametersMusicUrls) GetArg() (res []byte) {
	if len(r.MusicUrls) == 0 {
		return
	}
	res, _ = json.Marshal(r)
	return
}

func (r *AuditParametersImg) GetArg() (res []byte) {
	if len(r.ImgUrls) == 0 {
		return
	}
	res, _ = json.Marshal(r)
	return
}

func (r *AuditParametersText) GetArg() (res []byte) {
	if len(r.Text) == 0 {
		return
	}
	res, _ = json.Marshal(r)
	return
}
