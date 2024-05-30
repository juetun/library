package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/const_apply"
	"gorm.io/gorm"
	"time"
)

const (
	UserMainStatusInit    = const_apply.ApplyStatusInit     // 用户注册初始化(待审核)
	UserMainStatusOk      = const_apply.ApplyStatusOk       // 用户审核通过
	UserMainStatusInvalid = const_apply.ApplyStatusInvalid  // 已失效
	UserMainStatusFailure = const_apply.ApplyStatusFailure  // 用户审核失败
	UserMainStatusWaiting = const_apply.ApplyStatusAuditing // 待审核

	UserMainGenderMale   = 0 // 男性
	UserMainGenderFeMale = 1 // 女性

	UserMainPortraitStatusInit    = const_apply.ApplyStatusInit     // 头像初始化
	UserMainPortraitStatusOk      = const_apply.ApplyStatusOk       // 审核通过
	UserMainPortraitStatusInvalid = const_apply.ApplyStatusInvalid  // 已失效
	UserMainPortraitStatusFailure = const_apply.ApplyStatusFailure  // 审核失败
	UserMainPortraitStatusWaiting = const_apply.ApplyStatusAuditing // 待审核
)

const (
	UserMainAuthTypeGeneral uint8 = iota //个人用户
	UserMainAuthTypeCompany              //企业用户
)

var (
	SliceUserMainAuthType = base.ModelItemOptions{
		{Label: "个人", Value: UserMainAuthTypeGeneral},
		{Label: "企业", Value: UserMainAuthTypeCompany},
	}
	SliceUserMainStatus = base.ModelItemOptions{
		{Label: "未认证", Value: UserMainStatusInit},
		{Label: "已认证", Value: UserMainStatusOk},
		{Label: "已失效", Value: UserMainStatusInvalid},
		{Label: "审核失败", Value: UserMainStatusFailure},
		{Label: "待审核", Value: UserMainStatusWaiting},
	}
	SliceUserMainPortraitStatus = base.ModelItemOptions{
		{Label: "未审核", Value: UserMainPortraitStatusInit},
		{Label: "审核通过", Value: UserMainPortraitStatusOk},
		{Label: "已失效", Value: UserMainPortraitStatusInvalid},
		{Label: "审核失败", Value: UserMainPortraitStatusFailure},
		{Label: "待审核", Value: UserMainPortraitStatusWaiting},
	}
	SliceUserMainGender = base.ModelItemOptions{
		{Label: "男", Value: UserMainGenderMale},
		{Label: "女", Value: UserMainGenderFeMale},
	}
)
var (
	UserMainTableNumber int64 = 16
)

type (
	UserMain struct {
		Id               int              `gorm:"column:id;primary_key" json:"id"`
		UserHid          int64            `gorm:"uniqueIndex:idx_user_hid;column:user_hid;not null;default:0;type:bigint(20) COLLATE utf8mb4_bin" json:"user_hid"` // sql:"unique_index" 创建表时生成唯一索引
		Portrait         string           `gorm:"column:portrait;not null;type:varchar(100);default:'';comment:头图地址;" json:"portrait"`
		PortraitUrl      string           `json:"portrait_url" gorm:"-"`
		PortraitTmp      string           `gorm:"column:portrait_tmp;not null;type:varchar(100);default:'';comment:头图临时地址修改头像时使用;" json:"portrait_tmp"`
		PortraitTmpUrl   string           `json:"portrait_tmp_url" gorm:"-"`
		PortraitStatus   int8             `gorm:"column:portrait_status;not null;type:varchar(10);default:0;comment:用户审核状态从右向左每位依次昵称-头像;" json:"portrait_status"`
		NickName         string           `gorm:"column:nick_name;not null;type:varchar(30);default:'';comment:昵称" json:"nick_name"`
		UserName         string           `gorm:"column:user_name;not null;size:50;default:'';comment:用户名" json:"user_name" `
		Gender           uint8            `gorm:"column:gender;not null;type:tinyint(1);default:0;comment:性别 0-男 1-女" json:"gender"`
		Status           int8             `gorm:"column:status;not null;type:tinyint(1);default:0;comment:状态 0-可用 1-不可用" json:"status"`
		Score            int              `gorm:"column:score;not null;type:int(10);default:0;comment:用户积分" json:"score"`
		Balance          float64          `gorm:"column:balance;not null;type:decimal(10,2);default:0;comment:用户账户余额" json:"balance"`
		CurrentShopId    int64            `gorm:"column:current_shop_id;not null;default:0;comment:当前店铺ID" json:"current_shop_id"`
		Country          string           `gorm:"column:country;not null;type:varchar(30) COLLATE utf8mb4_general_ci;comment:国籍" json:"country"`
		ProvinceId       int              `gorm:"column:province_id;not null;type:int(10);default:0;comment:所在省份" json:"province_id"`
		CityId           int              `gorm:"column:city_id;not null;type:int(10);default:0;comment:所在城市" json:"city_id"`
		AreaId           int              `gorm:"column:area_id;not null;type:int(10);default:0;comment:所在区或县" json:"area_id"`
		OrgCode          string           `gorm:"column:org_code;not null;type:varchar(180) COLLATE utf8mb4_bin;comment:机构号" json:"org_code"`
		OrgRoot          string           `gorm:"column:org_root;not null;type:varchar(32) COLLATE utf8mb4_bin;comment:机构号" json:"org_root"`
		IsV              int              `gorm:"column:is_v;not null;type:tinyint(1);default:0;comment:用户头像加V 0-不加 1-加" json:"is_v" ` // 用户头像加V
		AuthType         uint8            `gorm:"column:auth_type;not null;type:tinyint(2);default:0;comment:认证类型 0-个人用户 1-企业用户" json:"auth_type"`
		AuthStatus       uint8            `gorm:"column:auth_status;not null;type:tinyint(2);default:2;comment:认证状态0-未审核 1-可用 2-不可用 3-待审核" json:"auth_status"`
		AuthDesc         string           `gorm:"column:auth_desc;not null;type:varchar(30);default:'';comment:认证描述" json:"auth_desc"` // 认证描述
		HaveDashboard    uint8            `gorm:"column:have_dashboard;not null;type:tinyint(1);default:2;comment:性别 1-有客服后台权限 2-无权限"  json:"have_dashboard"`
		Email            string           `gorm:"column:email;uniqueIndex:idx_email,priority:1;not null;type:varchar(100);default:'';comment:邮箱" json:"-"`
		EmailVerifiedAt  *base.TimeNormal `gorm:"column:email_verified_at;uniqueIndex:idx_email,priority:3;type:datetime;default:'2000-01-01 00:00:00';comment:认证时间;" json:"-"`
		CountryCode      string           `gorm:"column:country_code;uniqueIndex:idx_mobile,priority:2;type:varchar(15) COLLATE utf8mb4_bin;default:'';not null;comment:手机国别默认86" json:"country_code"`
		Mobile           string           `gorm:"column:mobile;not null;default:'';uniqueIndex:idx_mobile,priority:1;type:varchar(20) COLLATE utf8mb4_bin;comment:手机号" json:"-"`
		MobileVerifiedAt *base.TimeNormal `json:"mobile_verified_at" gorm:"column:mobile_verified_at;uniqueIndex:idx_mobile,priority:4;default:'2000-01-01 00:00:00'"`
		CreatedAt        base.TimeNormal  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" `
		UpdatedAt        base.TimeNormal  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at" `
		DeletedAt        *base.TimeNormal `gorm:"column:deleted_at;" json:"-"`
		DataName         string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
		TbName           string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
	}
)

func (r *UserMain) ParseTime(t interface{}, stringFormat ...string) (res base.TimeNormal, err error) {
	if t == nil {
		return
	}
	switch t.(type) {
	case base.TimeNormal:
		res = t.(base.TimeNormal)
	case *base.TimeNormal:
		t1 := t.(*base.TimeNormal)
		res = *t1
	case string:
		timeString := t.(string)
		f := "2006-01-02 15:04:05"
		if len(stringFormat) > 0 {
			f = stringFormat[0]
		}
		var t1 time.Time
		t1, err = time.ParseInLocation(f, timeString, time.Local)
		res = base.TimeNormal{Time: t1}
		return
	case time.Time:
		t1 := t.(time.Time)
		res = base.TimeNormal{Time: t1}
		return
	default:
		err = fmt.Errorf("数据格式不正确")
	}
	return
}

func (r *UserMain) BeforeCreate(db *gorm.DB) (err error) {

	return
}

func (r *UserMain) GetTableComment() (res string) {
	res = "用户信息主表"
	return
}
func (r *UserMain) ParseStatusString() (res string, err error) {
	var UserMainStatusMap map[int8]string
	if UserMainStatusMap, err = SliceUserMainStatus.GetMapAsKeyInt8(); err != nil {
		return
	}
	if dt, ok := UserMainStatusMap[r.Status]; ok {
		res = dt
		return
	}
	err = fmt.Errorf("当前用户审核状态不合法(%d)", r.Status)
	return
}

func (r *UserMain) ParseAuthType() (res string) {
	mapAuthType, _ := SliceUserMainAuthType.GetMapAsKeyUint8()
	var ok bool
	if res, ok = mapAuthType[r.AuthType]; ok {
		return
	}
	res = fmt.Sprintf("未知认证类型(%d)", r.AuthType)
	return
}

func (r *UserMain) ParseGender() (res string) {
	var UserMainGenderMap, _ = SliceUserMainGender.GetMapAsKeyUint8()
	if dt, ok := UserMainGenderMap[r.Gender]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.Gender)
	return
}

func (r *UserMain) ParsePortraitStatus() (res string) {
	var userMainPortraitStatusMap, _ = SliceUserMainPortraitStatus.GetMapAsKeyInt8()
	if dt, ok := userMainPortraitStatusMap[r.PortraitStatus]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.PortraitStatus)
	return
}

func (r *UserMain) Default() (err error) {
	if r.AuthStatus == 0 {
		r.AuthStatus = 2
	}
	return
}
