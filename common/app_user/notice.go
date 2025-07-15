package app_user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

type (
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
	NoticeTemplateSelectTypeOption func(noticeTemplateSelectType *NoticeTemplateSelectType)
	ResultNoticeTemplate           map[string]NoticeTemplateItem

	ArgNoticeTemplate struct {
		Keys    string   `json:"keys" form:"keys"`
		KeyList []string `json:"-" form:"-"`
	}
)

func (r *ArgNoticeTemplate) Default(c *gin.Context) (err error) {
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
