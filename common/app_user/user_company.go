package app_user

import "github.com/juetun/base-wrapper/lib/base"

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
)

func (r *ArgCompanyInfo) Default(ctx *base.Context) (err error) {

	return
}
