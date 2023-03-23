package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strings"
)

// 0-普通店 1-网站官方自营店 2-官方店 3-旗舰店 4-授权店
const (
	ShopEntryTypeGeneral         uint8 = iota + 1 // 普通店
	ShopEntryTypeSite                             // 网站官方自营店
	ShopEntryTypeOfficial                         // 品牌官方自营店
	ShopEntryTypeFlagshipStore                    // 旗舰店
	ShopEntryTypeAuthorizedStore                  // 授权店
)

const (
	FlagTesterNo  uint8 = iota + 1 // 为测试数据
	FlagTesterYes                  // 不为测试数据
)

// 店铺入驻状态
const (
	ShopStatusInit    uint8 = iota + 1 // 入驻状态初始化
	ShopStatusOk                       // 入驻状态审核通过
	ShopStatusFailure                  // 入驻状态审核失败

)
const (
	ShopTypePerson     uint8 = iota + 1 // 个人店
	ShopTypeBussiness                   // 企业店
	ShopTypeGeneralYes                  // 公募
	ShopTypeGeneralNo                   // 非公募
)

const (
	ShopNeedVerifyStatusOk       uint8 = iota + 1 //审核通过
	ShopNeedVerifyStatusUpdating                  //审核中
	ShopNeedVerifyStatusFailure                   //审核失败
	ShopNeedVerifyStatusExpire                    // 店铺资质已过期
)

//修改店铺信息支持的字段
const (
	ShopCanUpdateColumnName          = "name"
	ShopCanUpdateColumnShopType      = "shop_type"
	ShopCanUpdateColumnShopEntryType = "shop_entry_type"
	ShopCanUpdateColumnIcon          = "icon"     //修改logo
	ShopCanUpdateColumnBgImage       = "bg_image" //修改背景图
)

// 0-个人店 1-企业店 2-公募 3-公募
var (
	SliceFlagTester = base.ModelItemOptions{
		{
			Value: FlagTesterNo,
			Label: "是",
		},
		{
			Value: FlagTesterYes,
			Label: "否",
		},
	}
	SliceShopType = base.ModelItemOptions{
		{
			Value: ShopTypePerson,
			Label: "个人店",
		},
		{
			Value: ShopTypeBussiness,
			Label: "企业店",
		},
		{
			Value: ShopTypeGeneralYes,
			Label: "公募",
		},
		{
			Value: ShopTypeGeneralNo,
			Label: "非公募",
		},
	}
	SliceShopEntryType = base.ModelItemOptions{
		{
			Value: ShopEntryTypeGeneral,
			Label: "普通店",
		},
		{
			Value: ShopEntryTypeSite,
			Label: "官方自营店",
		},
		{
			Value: ShopEntryTypeOfficial,
			Label: "自营店",
		},
		{
			Value: ShopEntryTypeFlagshipStore,
			Label: "旗舰店",
		},
		{
			Value: ShopEntryTypeAuthorizedStore,
			Label: "授权店",
		},
	}
	SliceShopStatus = base.ModelItemOptions{
		{
			Value: ShopStatusInit,
			Label: "初始化",
		},
		{
			Value: ShopStatusOk,
			Label: "审核通过",
		},
		{
			Value: ShopStatusFailure,
			Label: "审核失败",
		},
	}

	//修改店铺资料时使用
	SliceShopNeedVerifyStatus = base.ModelItemOptions{
		{
			Value: ShopNeedVerifyStatusOk,
			Label: "审核通过",
		},
		{
			Value: ShopNeedVerifyStatusUpdating,
			Label: "审核中",
		},
		{
			Value: ShopNeedVerifyStatusFailure,
			Label: "审核失败",
		},
		{
			Value: ShopNeedVerifyStatusExpire,
			Label: "资质已过期",
		},
	}
)

const (
	ShopSliceVerifyStatusValue   = "0" // 初始化
	ShopSliceVerifyStatusOk      = "1" // 审核成功
	ShopSliceVerifyStatusIng     = "2" // 审核中
	ShopSliceVerifyStatusFailure = "3" // 审核失败
)

const (
	ShopSliceVerifyStatusIcon = "shop_icon"
)

var (
	//注意:此数据只能在后边添加,否则会影响数据结构
	SliceVerifyStatus = base.ModelItemOptions{
		{
			Value: ShopSliceVerifyStatusIcon,
			Label: "头像",
		},
	}
)

type (
	Shop struct {
		ShopID           int64            `gorm:"column:shop_id;primary_key" json:"shop_id"`
		Name             string           `gorm:"column:name;type:varchar(255);not null;default:'';comment:店铺名称" json:"name"`
		Logo             string           `gorm:"column:icon;type:varchar(255);not null;default:'';comment:店铺logo" json:"icon"`
		LogoUrl          string           `gorm:"-" json:"logo_url" `
		BgImageUrl       string           `json:"bg_image_url" gorm:"-"`
		BgImage          string           `gorm:"column:bg_image;type:varchar(255);not null;default:'';comment:店铺背景图" json:"bg_image"`
		ShopType         uint8            `gorm:"column:shop_type;type:tinyint(2);not null;default:1;comment:店铺类型 1-个人店 2-企业店 3-公募 4-非公募" json:"shop_type"`
		ShopEntryType    uint8            `gorm:"column:shop_entry_type;type:tinyint(2);not null;default:1;comment:店铺入驻类型 1-普通店 2-本站官方自营店 3-官方店 4-旗舰店 5-授权店" json:"shop_entry_type"`
		Status           uint8            `gorm:"column:status;type:tinyint(2);not null;default:1;comment:店铺审核状态1-审核通过 3-待审核 2-审核失败" json:"status"`
		FlagTester       uint8            `gorm:"column:flag_tester;not null;type: tinyint(2);default:2;comment:是否为测试数据店铺 1-是 2-不是"  json:"flag_tester"`
		AdminUserHid     int64            `gorm:"column:admin_user_hid;default:0;index:idx_userHid,priority:1;type:bigint(20);not null;comment:管理管理员账号" json:"admin_user_hid"`
		NeedVerifyStatus uint8            `gorm:"column:need_verify_status;type:tinyint(2);not null;default:1;comment:需要审核状态 1-审核通过 2-待审核 3-审核失败" json:"need_verify_status"`
		VerifyStatus     string           `gorm:"column:verify_status;type:varchar(20);not null;default:'';comment:审核数据状态" json:"verify_status"`
		CreatedAt        base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt        base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
		DeletedAt        *base.TimeNormal `gorm:"column:deleted_at;" json:"deleted_at"`
	}
)

func (r *Shop) TableName() string {
	return fmt.Sprintf("%sshop", TablePrefix)
}

func (r *Shop) Default() {
	r.ShopType = 1
	r.ShopEntryType = 1
	r.Status = 1
	r.FlagTester = 2
	r.NeedVerifyStatus = 1
	return
}

func (r *Shop) GetShopName() (res string) {
	switch r.ShopEntryType {
	case ShopEntryTypeGeneral: //普通店
		res = r.Name
	case ShopEntryTypeSite: //官方自营店
		res = fmt.Sprintf("%s官方自营店", r.Name)
	case ShopEntryTypeOfficial: //官方自营店
		res = fmt.Sprintf("%s官方旗舰店", r.Name)
	case ShopEntryTypeFlagshipStore: //旗舰店
		res = fmt.Sprintf("%s旗舰店", r.Name)
	case ShopEntryTypeAuthorizedStore: //旗舰店
		res = fmt.Sprintf("%s授权店", r.Name)
	default:
		res = fmt.Sprintf("未知店铺类型(%d)", r.ShopEntryType)
	}
	return
}

// ParseShopType 店铺状态
func (r *Shop) ParseShopType() (res string) {
	shopTypeMap, _ := SliceShopType.GetMapAsKeyUint8()
	if dt, ok := shopTypeMap[r.ShopType]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.ShopType)
	return
}

func (r *Shop) ParseFlagTester() (res string) {
	sliceFlagTester, _ := SliceShopType.GetMapAsKeyUint8()
	if dt, ok := sliceFlagTester[r.FlagTester]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.FlagTester)
	return
}

func (r *Shop) ParseShopNeedVerifyStatus() (res string) {
	mapKey, _ := SliceShopNeedVerifyStatus.GetMapAsKeyUint8()
	if dt, ok := mapKey[r.NeedVerifyStatus]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.NeedVerifyStatus)
	return
}

// ParseShopEntryType 店铺名称格式
func (r *Shop) ParseShopEntryType() (res string) {
	shopEntryTypeMap, _ := SliceShopEntryType.GetMapAsKeyUint8()
	if dt, ok := shopEntryTypeMap[r.ShopEntryType]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.ShopEntryType)
	return
}

// ParseStatus 店铺状态
func (r *Shop) ParseStatus() (res string) {
	mapShopStatus, _ := SliceShopStatus.GetMapAsKeyUint8()
	if dt, ok := mapShopStatus[r.Status]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知状态(%d)", r.Status)
	return
}
func (r *Shop) GetTableComment() (res string) {
	res = "店铺表"
	return
}

func (r *Shop) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *Shop) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}

func (r *Shop) GetHid() (res int64) {
	res = r.ShopID
	return
}

func (r *Shop) GetShopTypeOptions() base.ModelItemOptions {
	return SliceShopType
}
func (r *Shop) GetDefaultIcon() string {
	if r.Logo == "" {
		r.Logo = DefaultImageShow
	}
	return r.Logo
}

func (r *Shop) SetVerifyStatus(tp string, v string) {
	s := []byte(r.VerifyStatus)
	mapValue := r.verifyStatusFormat(s)
	mapValue[tp] = v
	sValue := make([]string, 0, len(SliceVerifyStatus))
	for _, item := range SliceVerifyStatus {
		if dt, ok := mapValue[item.Value.(string)]; ok {
			sValue = append(sValue, dt)
		}
	}
	r.VerifyStatus = strings.Join(sValue, "")
	return
}

func (r *Shop) GetVerifyStatus(tp string) (res string) {
	s := []byte(r.VerifyStatus)
	mapValue := r.verifyStatusFormat(s)
	if tmp, ok := mapValue[tp]; ok {
		res = tmp
		return
	}
	res = fmt.Sprintf("当前不支持你输入的KEY:%s校验", tp)
	return
}

func (r *Shop) verifyStatusFormat(runes []byte) (mapValue map[string]string) {
	l := len(runes)
	mapValue = make(map[string]string, l)
	for k, item := range SliceVerifyStatus {
		if k < l {
			mapValue[item.Value.(string)] = string(runes[k])
		} else {
			mapValue[item.Value.(string)] = ShopSliceVerifyStatusValue
		}

	}
	return
}
