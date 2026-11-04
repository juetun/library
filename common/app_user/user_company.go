package app_user

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common/response"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"strings"
)

type (
	ArgCompanyInfo struct {
		UserHid           int64                  `json:"user_hid" form:"user_hid"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"-" form:"-"`
	}
	ResultCompany struct {
		ID                    int64                 `json:"id"`
		TmpID                 int64                 `json:"tmp_id"`
		UpdateBatchId         string                `json:"update_batch_id"`
		CompanyName           string                `json:"company_name"`
		SocialCredit          string                `json:"social_credit"`
		ManagerUserHid        int64                 `json:"manager_user_hid"`
		LegalPerson           string                `json:"legal_person"`
		LegalCertificateType  uint8                 `json:"legal_certificate_type"`
		LegalNumber           string                `json:"legal_number"`
		BusinessLicense       string                `json:"business_license"`
		BusinessLicenseUrl    string                `json:"business_license_url"`
		LegalIdCardFront      string                `json:"legal_id_card_front"`
		LegalIdCardFrontUrl   string                `json:"legal_id_card_front_url"`
		LegalIdCardBack       string                `json:"legal_id_card_back"`
		LegalIdCardBackUrl    string                `json:"legal_id_card_back_url"`
		LegalDateExpiry       string                `json:"legal_date_expiry"`
		LegalDateOfIssue      string                `json:"legal_date_of_issue"`
		LegalIsNeverExpires   uint8                 `json:"legal_is_never_expires"`
		LicenseIsNeverExpires uint8                 `json:"license_is_never_expires"`
		LicenseDateExpiry     string                `json:"license_date_expiry"`
		LicenseDateOfIssue    string                `json:"license_date_of_issue"`
		Mobile                string                `json:"mobile"`
		Email                 string                `json:"email"`
		Address               string                `json:"address"`
		Desc                  string                `json:"desc"`
		IdTypes               base.ModelItemOptions `json:"id_types"`
		Url                   string                `json:"url"`
	}
	Company struct {
		ID                    int64            `gorm:"primary_key;column:id;" json:"id"`
		CompanyName           string           `gorm:"column:company_name;not null;default:'';type:varchar(100) COLLATE utf8mb4_general_ci;comment:公司名称" json:"company_name"`
		SocialCredit          string           `gorm:"column:social_credit;not null;default:'';type:varchar(100) COLLATE utf8mb4_general_ci;comment:社会信用代码" json:"social_credit"`
		ManagerUserHid        int64            `gorm:"column:manager_user_hid;uniqueIndex:idx_uqid,priority:1;default:0;not null;type:bigint(20)" json:"manager_user_hid"`
		LegalPerson           string           `gorm:"column:legal_person;not null;default:'';type:varchar(100) COLLATE utf8mb4_general_ci;comment:法定代表人" json:"legal_person"`
		LegalCertificateType  uint8            `gorm:"column:legal_certificate_type;not null;default:1;type:tinyint(2);comment:1-身份证 2-港澳通行证 3-台湾通行证 4-护照" json:"legal_certificate_type"`
		LegalNumber           string           `gorm:"column:legal_number;not null;default:'';type:varchar(50) COLLATE utf8mb4_general_ci;comment:法人证件号" json:"legal_number"`
		LegalIdCardFront      string           `gorm:"column:legal_id_card_front;not null;default:'';type:varchar(80) COLLATE utf8mb4_general_ci;comment:法人证件正面" json:"legal_id_card_front"`
		LegalIdCardFrontUrl   string           `json:"legal_id_card_front_url" gorm:"-"`
		LegalIdCardBack       string           `gorm:"column:legal_id_card_back;not null;default:'';type:varchar(80) COLLATE utf8mb4_general_ci;comment:法人证件反面" json:"legal_id_card_back"`
		LegalIdCardBackUrl    string           `json:"legal_id_card_back_url" gorm:"-"`
		LegalDateExpiry       base.TimeNormal  `gorm:"column:legal_date_expiry;not null;type:date;default:'2000-01-01';comment:有效期时间值" json:"legal_date_expiry"`
		LegalDateOfIssue      *base.TimeNormal `gorm:"column:legal_date_of_issue;type:date;default:null;comment:发证日期值" json:"legal_date_of_issue"`
		LegalIsNeverExpires   uint8            `gorm:"column:legal_is_never_expires;default:1;type:tinyint(2);not null;comment:是否永久有效 1-有过期时间 2-永久有效" json:"legal_is_never_expires"`
		BusinessLicense       string           `gorm:"column:business_license;not null;default:'';type:varchar(80) COLLATE utf8mb4_general_ci;comment:营业执照" json:"business_license"`
		BusinessLicenseUrl    string           `json:"business_license_url" gorm:"-"`
		LicenseDateExpiry     base.TimeNormal  `gorm:"column:license_date_expiry;not null;type:date;default:'2000-01-01';comment:有效期时间值" json:"license_date_expiry"`
		LicenseDateOfIssue    *base.TimeNormal `gorm:"column:license_date_of_issue;type:date;default:null;comment:发证日期值" json:"license_date_of_issue"`
		LicenseIsNeverExpires uint8            `gorm:"column:license_is_never_expires;default:1;type:tinyint(2);not null;comment:是否永久有效 1-有过期时间 2-永久有效" json:"license_is_never_expires"`
		ShopId                int64            `gorm:"column:shop_id;index:idx_shop_id,priority:1;default:0;type:bigint(20);not null;comment:店铺ID" json:"shop_id,omitempty"`
		Mobile                string           `gorm:"column:mobile;not null;default:'';type:varchar(20) COLLATE utf8mb4_bin;comment:手机号" json:"mobile"`
		Email                 string           `gorm:"column:email;not null;default:'';type:varchar(50) COLLATE utf8mb4_bin;comment:邮箱" json:"email"`
		Address               string           `gorm:"column:address;not null;default:'';type:varchar(80) COLLATE utf8mb4_general_ci;comment:公司地址" json:"address"`
		Desc                  string           `gorm:"column:desc;not null;default:'';type:varchar(1800) COLLATE utf8mb4_general_ci;comment:公司简介" json:"desc"`
		Url                   string           `gorm:"column:url;not null;default:'';type:varchar(400) COLLATE utf8mb4_general_ci;comment:公司网址"  json:"url"`
	}
	ArgCompanyFetch struct {
		response.PageQuery
		CompanyId    int64  `json:"company_id" form:"company_id"`
		CompanyName  string `json:"company_name" form:"company_name"`
		SocialCredit string `json:"social_credit" form:"social_credit"`
	}
	ResultCompanyFetch struct {
		List       []*Company `json:"list"`
		TotalCount int64      `json:"total_count,omitempty"`
		IsNext     bool       `json:"is_next,omitempty"` // [bool] 是否有下一页，true=有下一页；false=无下页，可关闭列表
		response.PagerParameter
	}
)

func (r *ArgCompanyFetch) Default(ctx *base.Context) (err error) {

	return
}

func (r *ArgCompanyInfo) Default(ctx *base.Context) (err error) {

	return
}

//根据获取公司信息列表
func GetCompanyList(ctx *base.Context, arg *ArgCompanyFetch) (res ResultCompanyFetch, err error) {
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUser,
		URI:         "/user/get_company_list",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	ro.BodyJson, _ = json.Marshal(arg)
	var data = struct {
		Code int                `json:"code"`
		Data ResultCompanyFetch `json:"data"`
		Msg  string             `json:"message"`
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

//根据公司ID获取用户信息
func GetCompanyByCompanyIds(ctx *base.Context, companyIds ...int64) (res map[int64]*Company, err error) {
	var l = len(companyIds)
	res = make(map[int64]*Company, l)
	if l == 0 {
		return
	}
	var value = url.Values{}
	var companyId = make([]string, 0, l)
	for _, value := range companyIds {
		companyId = append(companyId, fmt.Sprintf("%d", value))
	}
	uIds := strings.Join(companyId, ",")
	value.Set("company_ids", uIds)
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUser,
		URI:         "/user/get_company_by_company_ids",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                `json:"code"`
		Data map[int64]*Company `json:"data"`
		Msg  string             `json:"message"`
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
