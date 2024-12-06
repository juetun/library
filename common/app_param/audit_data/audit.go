package audit_data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"sync"
)

//审核工具类型
const (
	ApplyToolTypeDefault uint8 = iota //默认不审核
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
const (
	IsSynchronousYes uint8 = iota + 1 //同步数据
	IsSynchronousNo
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
			Label: "不审核",
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
	AuditClient interface {
		Do(item AuditParametersInterface) (result *ApplyResult, err error)
	}
	AuditParametersInterface interface {
		GetArg() (res []byte) //获取参数
		GetApplyType() (res uint8)
		GetIsSynchronous() (isSynchronous uint8)
		Default()
	}
	AuditData struct {
		DefaultClient                                                   AuditClient                `json:"-"`
		PrivateClient                                                   AuditClient                `json:"-"`
		BaiDuClient                                                     AuditClient                `json:"-"`
		ShuMeiClient                                                    AuditClient                `json:"-"`
		ToolClientClient                                                AuditClient                `json:"-"`
		Ctx                                                             *base.Context              `json:"-"`
		Context                                                         context.Context            `json:"-"`
		onceDefault, oncePrivate, onceBaiDu, onceShuMei, onceToolClient sync.Once                  `json:"-"`
		Parameters                                                      []AuditParametersInterface `json:"arg"`
	}
	AuditParametersText struct {
		MsgId         string   `json:"msg_id"`
		Text          []string `json:"text"`           //审核的图片列表
		IsSynchronous uint8    `json:"is_synchronous"` //是否同步返回
		ApplyType     uint8    `json:"apply_type"`
	}
	AuditParametersImg struct {
		MsgId         string   `json:"msg_id"`
		ImgUrls       []string `json:"img_urls"`       //审核的图片列表
		IsSynchronous uint8    `json:"is_synchronous"` //是否同步返回
		ApplyType     uint8    `json:"apply_type"`
	}
	AuditParametersVideoUrls struct {
		MsgId         string   `json:"msg_id"`
		VideoUrls     []string `json:"video_urls"`     //审核的视频链接
		IsSynchronous uint8    `json:"is_synchronous"` //是否同步返回
		ApplyType     uint8    `json:"apply_type"`
	}

	AuditParametersMusicUrls struct {
		MsgId         string   `json:"msg_id"`
		MusicUrls     []string `json:"music_urls"`     //审核的音频链接
		IsSynchronous uint8    `json:"is_synchronous"` //是否同步返回
		ApplyType     uint8    `json:"apply_type"`
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

func AuditCtx(ctx *base.Context) AuditDataOption {
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

func (r *AuditData) InitAuditDefault() {
	r.onceDefault.Do(func() {
		r.DefaultClient = NewDefaultAudit(r.Ctx, r.Context)
	})
	return
}

func (r *AuditData) InitPrivateDefault() {
	r.oncePrivate.Do(func() {
		r.PrivateClient = NewPrivateAudit(r.Ctx, r.Context)
	})
	return
}

func (r *AuditData) InitBaiDuDefault() {
	r.onceBaiDu.Do(func() {
		r.BaiDuClient = NewBaiDuAudit(r.Ctx, r.Context)
	})
	return
}

func (r *AuditData) InitShuMeiDefault() {
	r.onceShuMei.Do(func() {
		r.ShuMeiClient = NewShuMeiAudit(r.Ctx, r.Context)
	})
	return
}

func (r *AuditData) InitToolClientDefault() {
	r.onceToolClient.Do(func() {
		r.ToolClientClient = NewToolClientAudit(r.Ctx, r.Context)
	})
	return
}

func (r *AuditData) Audit(item AuditParametersInterface) (applyResult *ApplyResult, err error) {
	applyResult = &ApplyResult{}
	switch item.GetApplyType() {
	case ApplyToolTypeDefault: //默认为审核
		r.InitAuditDefault()
		applyResult, err = r.DefaultClient.Do(item)
	case ApplyToolTypePrivate: //平台自审程序
		r.InitPrivateDefault()
		applyResult, err = r.PrivateClient.Do(item)
	case ApplyToolTypeBaiDu: //百度审核
		r.InitBaiDuDefault()
		applyResult, err = r.BaiDuClient.Do(item)
	case ApplyToolTypeShuMei: //数美审核
		r.InitShuMeiDefault()
		applyResult, err = r.ShuMeiClient.Do(item)
	case ApplyToolTypeClient: //平台人工审核
		r.InitToolClientDefault()
		applyResult, err = r.ToolClientClient.Do(item)
	default:
		err = fmt.Errorf("暂不支持你选择的审核类型")
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
		item.Default()
		if item.GetIsSynchronous() != IsSynchronousYes { //如果不是同步返回
			continue
		}
		if result, err = r.Audit(item); err != nil {
			return
		}
		switch result.Status {
		case DataChatStatusOk: //审核通过
		case DataChatStatusWaiting: //待审核
			applyResult = result
		case DataChatStatusFailure: //审核失败 (如果一条数据审核失败，那么这次审核都失败)
			applyResult = result
			return
		}
	}
	for _, item := range r.Parameters {
		if item.GetIsSynchronous() != IsSynchronousNo { //上边已经同步验证过
			continue
		}
		if result, err = r.Audit(item); err != nil {
			return
		}
		applyResult.Status = DataChatStatusWaiting
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

func (r *AuditParametersVideoUrls) GetApplyType() (res uint8) {
	return r.ApplyType
}

func (r *AuditParametersMusicUrls) GetApplyType() (res uint8) {
	return r.ApplyType
}

func (r *AuditParametersImg) GetApplyType() (res uint8) {
	return r.ApplyType
}

func (r *AuditParametersText) GetApplyType() (res uint8) {
	return r.ApplyType
}

func (r *AuditParametersVideoUrls) GetIsSynchronous() (res uint8) {
	return r.IsSynchronous
}

func (r *AuditParametersMusicUrls) GetIsSynchronous() (res uint8) {
	return r.IsSynchronous
}

func (r *AuditParametersImg) GetIsSynchronous() (res uint8) {
	return r.IsSynchronous
}

func (r *AuditParametersText) GetIsSynchronous() (res uint8) {
	return r.IsSynchronous
}

func (r *AuditParametersVideoUrls) Default() {
	if r.IsSynchronous == 0 {
		r.IsSynchronous = IsSynchronousNo
	}
	return
}

func (r *AuditParametersMusicUrls) Default() {
	if r.IsSynchronous == 0 {
		r.IsSynchronous = IsSynchronousNo
	}
	return
}

func (r *AuditParametersImg) Default() {

	if r.IsSynchronous == 0 {
		r.IsSynchronous = IsSynchronousNo
	}
	return
}

func (r *AuditParametersText) Default() {
	if r.IsSynchronous == 0 {
		r.IsSynchronous = IsSynchronousYes
	}

	return
}
