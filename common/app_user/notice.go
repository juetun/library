package app_user

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"strings"
)

const (
	NoticeKeyRegisterAndLogin = "register_user"  //注册登录短信验证码
	NoticeKeyUpdateMobile     = "update_mobile"  //修改手机号验证码
	NoticeKeyUpdateEmail      = "update_email"   //修改电子邮件
	NoticeKeyUpdateIdCard     = "update_id_card" //修改身份证
)
const (
	NoticeTemplateSelectTypeStart       = iota // 起始站位符，无实际意义
	NoticeTemplateSelectTypeSiteIn             //站内信
	NoticeTemplateSelectTypeShortMsg           //短信
	NoticeTemplateSelectTypeEmail              //邮件
	NoticeTemplateSelectTypePush               //app push
	NoticeTemplateSelectTypePersonalMsg        //私信
)
const (
	NoticeTemplateSelectTypeStartDesc       = ""
	NoticeTemplateSelectTypeSiteInDesc      = "select_type_site_in"
	NoticeTemplateSelectTypeShortMsgDesc    = "select_type_short_msg"
	NoticeTemplateSelectTypeEmailDesc       = "select_type_email"
	NoticeTemplateSelectTypePushDesc        = "select_type_push"
	NoticeTemplateSelectTypePersonalMsgDesc = "select_type_personal_msg"
)
const (
	NoticeTemplateSelectTypeNo  uint8 = iota //未启用
	NoticeTemplateSelectTypeYes              //启用
)

type (
	NoticeTemplateItemOptions []NoticeTemplateItemOption
	NoticeTemplateItemOption  struct {
		base.ModelItemOption
		Key string `json:"key"`
	}
	NoticeTemplateItem struct {
		NoticeTemplateSelectType
		NoticeTemplateSelectTypeDesc
	}
	NoticeTemplateSelectType struct {
		SelectTypeStart       uint8  `json:"-" form:"-"`                                     //占位符,
		SelectTypeSiteIn      uint8  `json:"select_type_site_in" form:"select_type_site_in"` //站内信
		SiteInTitle           string `json:"site_in_title" form:"site_in_title"`
		SelectTypeShortMsg    uint8  `json:"select_type_short_msg" form:"select_type_short_msg"` //短信
		SelectTypeEmail       uint8  `json:"select_type_email" form:"select_type_email"`         //邮件
		EmailTitle            string `json:"email_title" form:"email_title"`
		SelectTypePush        uint8  `json:"select_type_push" form:"select_type_push"`                 //app_push
		SelectTypePersonalMsg uint8  `json:"select_type_personal_msg" form:"select_type_personal_msg"` //私信
	}
	NoticeTemplateSelectTypeDesc struct {
		SelectTypeSiteInDesc      string `json:"select_type_site_in_desc" form:"select_type_site_in_desc"`           //站内信
		SelectTypeShortMsgDesc    string `json:"select_type_short_msg_desc" form:"select_type_short_msg_desc"`       //短信
		SelectTypeEmailDesc       string `json:"select_type_email_desc" form:"select_type_email_desc"`               //邮件
		SelectTypePushDesc        string `json:"select_type_push_desc" form:"select_type_push_desc"`                 //app_push
		SelectTypePersonalMsgDesc string `json:"select_type_personal_msg_desc" form:"select_type_personal_msg_desc"` //私信
	}
	NoticeTemplateSelectTypeName struct {
		SelectTypeStartName       string `json:"-"`                             //占位符,
		SelectTypeSiteInName      string `json:"select_type_site_in_name"`      //站内信
		SelectTypeShortMsgName    string `json:"select_type_short_msg_name"`    //短信
		SelectTypeEmailName       string `json:"select_type_email_name"`        //邮件
		SelectTypePushName        string `json:"select_type_push_name"`         //app_push
		SelectTypePersonalMsgName string `json:"select_type_personal_msg_name"` //私信
	}
	NoticeTemplateSelectTypeOption func(noticeTemplateSelectType *NoticeTemplateSelectType)
	ResultNoticeTemplate           map[string]NoticeTemplateItem

	ArgNoticeTemplate struct {
		Keys    string   `json:"keys" form:"keys"`
		KeyList []string `json:"-" form:"-"`
	}
	SelectLabel struct {
		Label string `json:"label"`
		Value uint8  `json:"value"`
	}
)

var (
	SliceNoticeTemplateSelectTypeValue = base.ModelItemOptions{
		{
			Label: "启用",
			Value: NoticeTemplateSelectTypeYes,
		},
		{
			Label: "未启用",
			Value: NoticeTemplateSelectTypeNo,
		},
	}
	SliceNoticeTemplateSelectType = NoticeTemplateItemOptions{
		{
			ModelItemOption: base.ModelItemOption{
				Label: "占位符",
				Value: NoticeTemplateSelectTypeStart,
			},
			Key: NoticeTemplateSelectTypeStartDesc,
		},
		{
			ModelItemOption: base.ModelItemOption{
				Label: "站内信",
				Value: NoticeTemplateSelectTypeSiteIn,
			},
			Key: NoticeTemplateSelectTypeSiteInDesc,
		},
		{
			ModelItemOption: base.ModelItemOption{
				Label: "短信",
				Value: NoticeTemplateSelectTypeShortMsg,
			},
			Key: NoticeTemplateSelectTypeShortMsgDesc,
		},
		{
			ModelItemOption: base.ModelItemOption{
				Label: "邮件",
				Value: NoticeTemplateSelectTypeEmail,
			},
			Key: NoticeTemplateSelectTypeEmailDesc,
		},
		{
			ModelItemOption: base.ModelItemOption{
				Label: "app push",
				Value: NoticeTemplateSelectTypePush,
			},
			Key: NoticeTemplateSelectTypePushDesc,
		},
		{
			ModelItemOption: base.ModelItemOption{
				Label: "私信",
				Value: NoticeTemplateSelectTypePersonalMsg,
			},
			Key: NoticeTemplateSelectTypePersonalMsgDesc,
		},
	}
)

func (r *NoticeTemplateItemOptions) GetMapAsKeyUint8() (res map[uint8]string, err error) {

	res = make(map[uint8]string, len(*r))
	for _, item := range *r {
		res[item.Value.(uint8)] = item.Label
	}
	return
}
func (r *ArgNoticeTemplate) Default(c *base.Context) (err error) {
	r.Keys = strings.TrimSpace(r.Keys)
	r.KeyList = make([]string, 0)
	if r.Keys != "" {
		keyList := strings.Split(r.Keys, ",")
		r.KeyList = make([]string, 0, len(keyList))
		for _, item := range keyList {
			if item == "" {
				continue
			}
			r.KeyList = append(r.KeyList, item)
		}
	}
	return
}

func (r *NoticeTemplateSelectTypeDesc) Default() (err error) {
	if r.SelectTypeSiteInDesc == "" {
		r.SelectTypeSiteInDesc = "{{.Content}}"
	}
	if r.SelectTypeShortMsgDesc == "" {
		r.SelectTypeShortMsgDesc = "{{.Content}}"
	}
	if r.SelectTypeEmailDesc == "" {
		r.SelectTypeEmailDesc = "{{.Content}}"
	}
	if r.SelectTypePushDesc == "" {
		r.SelectTypePushDesc = "{{.Content}}"
	}
	if r.SelectTypePersonalMsgDesc == "" {
		r.SelectTypePersonalMsgDesc = "{{.Content}}"
	}
	return
}

//根据通知KEY获取通知模版
func GetNoticeTemplateByTemplateKeys(ctx *base.Context, templateKeys *base.ArgGetByStringIds) (res ResultNoticeTemplate, err error) {
	res = make(map[string]NoticeTemplateItem, len(templateKeys.Ids))
	var value = url.Values{}

	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameNotice,
		URI:         "/announce/notice_template_by_keys",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = json.Marshal(templateKeys); err != nil {
		return
	}
	var data = struct {
		Code int                  `json:"code"`
		Data ResultNoticeTemplate `json:"data"`
		Msg  string               `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	res = data.Data
	return
}

func NewNoticeTemplateSelectType(options ...NoticeTemplateSelectTypeOption) (res *NoticeTemplateSelectType) {
	res = &NoticeTemplateSelectType{SelectTypeStart: 1}
	for _, handler := range options {
		handler(res)
	}
	return
}
func NoticeTemplateSelectTypeOptionPersonalMsg(personalMsg uint8) (res NoticeTemplateSelectTypeOption) {
	return func(noticeTemplateSelectType *NoticeTemplateSelectType) {
		noticeTemplateSelectType.SelectTypePersonalMsg = personalMsg
	}
}

func NoticeTemplateSelectTypeOptionPush(push uint8) (res NoticeTemplateSelectTypeOption) {
	return func(noticeTemplateSelectType *NoticeTemplateSelectType) {
		noticeTemplateSelectType.SelectTypePush = push
	}
}

func NoticeTemplateSelectTypeOptionEmail(email uint8) (res NoticeTemplateSelectTypeOption) {
	return func(noticeTemplateSelectType *NoticeTemplateSelectType) {
		noticeTemplateSelectType.SelectTypeEmail = email
	}
}

func NoticeTemplateSelectTypeOptionShortMsg(ShortMsg uint8) (res NoticeTemplateSelectTypeOption) {
	return func(noticeTemplateSelectType *NoticeTemplateSelectType) {
		noticeTemplateSelectType.SelectTypeShortMsg = ShortMsg
	}
}

func NoticeTemplateSelectTypeOptionSiteIn(siteIn uint8) (res NoticeTemplateSelectTypeOption) {
	return func(noticeTemplateSelectType *NoticeTemplateSelectType) {
		noticeTemplateSelectType.SelectTypeSiteIn = siteIn
	}
}

func (r *NoticeTemplateSelectType) UnSerialized(selectType string) (dt []string) {
	bt := []byte(selectType)
	dt = make([]string, len(SliceNoticeTemplateSelectType))
	for k, value := range bt {
		dt[k] = string(value)
	}
	return
}

func (r *NoticeTemplateSelectType) Serialized() (res string) {
	var data = make([]string, len(SliceNoticeTemplateSelectType))
	data[NoticeTemplateSelectTypeStart] = fmt.Sprintf("%d", r.SelectTypeStart)
	data[NoticeTemplateSelectTypeSiteIn] = fmt.Sprintf("%d", r.SelectTypeSiteIn)
	data[NoticeTemplateSelectTypeShortMsg] = fmt.Sprintf("%d", r.SelectTypeShortMsg)
	data[NoticeTemplateSelectTypeEmail] = fmt.Sprintf("%d", r.SelectTypeEmail)
	data[NoticeTemplateSelectTypePush] = fmt.Sprintf("%d", r.SelectTypePush)
	data[NoticeTemplateSelectTypePersonalMsg] = fmt.Sprintf("%d", r.SelectTypePersonalMsg)
	res = strings.Join(data, "")
	return
}

func (r *NoticeTemplateSelectType) SetValue(column string, value uint8) (res *NoticeTemplateSelectType, err error) {
	res = r
	switch column {
	case NoticeTemplateSelectTypeSiteInDesc: //站内信
		res.SelectTypeSiteIn = value
	case NoticeTemplateSelectTypeShortMsgDesc: //短信
		res.SelectTypeShortMsg = value
	case NoticeTemplateSelectTypeEmailDesc: //邮件
		res.SelectTypeEmail = value
	case NoticeTemplateSelectTypePushDesc: //app_push
		res.SelectTypePush = value
	case NoticeTemplateSelectTypePersonalMsgDesc: //私信
		res.SelectTypePersonalMsg = value
	default:
		err = fmt.Errorf("当前不支持你选择的字段修改(%s)", column)
	}

	return
}

func (r *NoticeTemplateSelectType) GetNoticeTemplateSelectTypeLabel() (res []SelectLabel) {
	res = []SelectLabel{}
	for key, value := range SliceNoticeTemplateSelectType {
		if key == NoticeTemplateSelectTypeStart { //起始位直接忽略
			continue
		}
		res = append(res, SelectLabel{
			Label: value.Label,
			Value: value.Value.(uint8),
		})
	}
	return
}
func (r *NoticeTemplateSelectType) ParseValueName(typeValue uint8) (res string) {
	mapV, _ := SliceNoticeTemplateSelectTypeValue.GetMapAsKeyUint8()
	if dt, ok := mapV[typeValue]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知状态(%d)", typeValue)
	return
}

func (r *NoticeTemplateSelectType) ParseName() (res *NoticeTemplateSelectTypeName) {
	res = &NoticeTemplateSelectTypeName{
		SelectTypeStartName:       r.ParseValueName(r.SelectTypeStart),
		SelectTypeSiteInName:      r.ParseValueName(r.SelectTypeSiteIn),
		SelectTypeShortMsgName:    r.ParseValueName(r.SelectTypeShortMsg),
		SelectTypeEmailName:       r.ParseValueName(r.SelectTypeEmail),
		SelectTypePushName:        r.ParseValueName(r.SelectTypePush),
		SelectTypePersonalMsgName: r.ParseValueName(r.SelectTypePersonalMsg),
	}

	return
}
