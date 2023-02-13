package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

// 获取用户信息对应的表
const (
	UserDataTypeIndex = "user_index"

	UserDataTypeEmail    = "user_email"
	UserDataTypeMain     = "user_main"
	UserDataTypePortrait = "user_portrait" //用户头像
	UserDataTypeInfo     = "user_info"
	UserDataTypeMobile   = "user_mobile"

	UserHidDivString      = "_" // 用户ID字符串切割
	UpdateColumnDivString = "." // 修改用户数据时传参的分割符
)
const (
	NeedValidateShopYes = true  //需要校验当前用户是否有店铺权限
	NeedValidateShopNo  = false //不需要校验当前用户是否有有店铺权限
)

// 获取用户信息的响应参数结构
type (
	ResultUser struct {
		List map[int64]ResultUserItem `json:"list"`
	}
	ResultUserItem struct {
		UserHid          int64      `json:"user_hid,omitempty"`  // 用户ID
		Portrait         string     `json:"portrait,omitempty"`  // 头像
		PortraitUrl      string     `json:"portrait_url"`        //头像链接
		NickName         string     `json:"nick_name,omitempty"` // 昵称
		UserName         string     `json:"user_name,omitempty"` // 用户名
		RealName         string     `json:"real_name"`           // 真实姓名
		Gender           int        `json:"gender,omitempty"`    //
		Status           int        `json:"status,omitempty"`    //
		Score            int        `json:"score,omitempty"`     //
		AuthDesc         string     `json:"auth_desc,omitempty"` // 认证描述
		IsV              int        `json:"is_v,omitempty"`      // 用户头像加V
		Remark           string     `json:"remark" `             // 个性签名
		Signature        string     `json:"signature,omitempty"`
		RegisterChannel  string     `json:"register_channel,omitempty"`
		CountryCode      string     `json:"country_code,omitempty"`
		Mobile           string     `json:"mobile,omitempty"`
		MobileVerifiedAt *time.Time `json:"mobile_verified_at,omitempty"`
		Email            string     `json:"email,omitempty"`
		EmailVerifiedAt  *time.Time `json:"email_verified_at,omitempty"`
		ShopId           int64      `json:"shop_id"`

		UserMobileIndex   string          `json:"user_mobile_index"`
		UserEmailIndex    string          `json:"user_email_index"`
		RememberToken     string          `json:"remember_token"`
		MsgReadTimeCursor base.TimeNormal `json:"msg_read_time_cursor"`
		HaveDashboard     uint8           `json:"have_dashboard"`
	}
	RequestUser struct {
		UUserHid           int64           `json:"u_user_hid" form:"u_user_hid"`                         //用户
		UUserMobileIndex   string          `json:"u_user_mobile_index" form:"u_user_mobile_index"`       //手机数据存储位置
		UUserEmailIndex    string          `json:"u_user_email_index" form:"u_user_email_index"`         //email存储位置
		UPortrait          string          `json:"u_portrait" form:"u_portrait"`                         //头像
		UNickName          string          `json:"u_nick_name" form:"u_nick_name"`                       //昵称
		UUserName          string          `json:"u_user_name" form:"u_user_name"`                       //账号
		UGender            int             `json:"u_gender" form:"u_gender"`                             //性别
		UStatus            int             `json:"u_status" form:"u_status"`                             //状态
		UScore             int             `json:"u_score" form:"u_score"`                               //积分
		URememberToken     string          `json:"u_remember_token" form:"u_remember_token"`             //是否记住密码
		UMsgReadTimeCursor base.TimeNormal `json:"u_msg_read_time_cursor" form:"u_msg_read_time_cursor"` //消息未读时刻节点
		UShopId            int64           `json:"u_shop_id" form:"u_shop_id"`                           //店铺ID
		UHaveDashboard     uint8           `json:"u_have_dashboard" form:"u_have_dashboard"`             //是否有客服后台权限

		UHeaderInfo *base.HeaderInfo `json:"u_header_info" form:"u_header_info"` //用户设备信息
	}

	User struct {
		UserHid    int64       `json:"user_hid"`
		UserIndex  *UserIndex  `json:"user_index,omitempty"`
		UserMain   *UserMain   `json:"user_main,omitempty"`
		UserEmail  *UserEmail  `json:"user_email,omitempty"`
		UserInfo   *UserInfo   `json:"user_info,omitempty"`
		UserMobile *UserMobile `json:"user_mobile,omitempty"`
	}
	UserIndex struct {
		ID         int64            `gorm:"column:id;primary_key" json:"-"`
		UserName   string           `gorm:"column:user_name;not null;type:varchar(50) COLLATE utf8mb4_bin;uniqueIndex;comment:用户名" json:"user_name" `
		TmpAccount string           `gorm:"column:tmp_account;not null;type:varchar(200) COLLATE utf8mb4_bin;comment:注册时临时账号" json:"tmp_account" `
		IsUse      int              `json:"is_use" gorm:"column:is_use;type:tinyint(1);default:0;comment:是否启用 0-启用 大于0-已启用"`
		CreatedAt  base.TimeNormal  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;" `
		UpdatedAt  base.TimeNormal  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;" `
		DeletedAt  *base.TimeNormal `json:"deleted_at" gorm:"column:deleted_at;" `
	}
	UserInfo struct {
		ID                int             `gorm:"column:id;primary_key" json:"id"`
		RealName          string          `gorm:"column:real_name;type:varchar(60);not null;comment:真实姓名"  json:"real_name"`
		UserIndexHid      string          `gorm:"column:user_index_hid;type:varchar(60);not null;comment:user_main表位置" json:"user_index_hid"`
		UserHid           int64           `gorm:"column:user_hid;not null;uniqueIndex:idx_userhid;default:0;type:bigint(20) COLLATE utf8mb4_bin" json:"user_hid"`
		RememberToken     string          `gorm:"column:remember_token;not null;default:'';size:500;comment:登录的token" json:"remember_token"`
		MsgReadTimeCursor base.TimeNormal `gorm:"column:msg_read_time_cursor;not null;default:CURRENT_TIMESTAMP;comment:最近一次读取系统公告时间" json:"msg_read_time_cursor"`
		Level             string          `gorm:"column:level;not null;type:tinyint(2);default:0;comment:用户等级0-普通用户" json:"level"`
		Remark            string          `json:"remark" gorm:"column:remark;not null;type:varchar(150);default:'';comment:个性签名"` // 个性签名
		Password          string          `gorm:"column:password;not null;type:varchar(256) COLLATE utf8mb4_general_ci;comment:密码" json:"password"`
		IdCard            string          `gorm:"column:id_card;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:身份证号加密串" json:"id_card"`
		IdCardSuffix      string          `gorm:"column:id_card_suffix;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:后缀后6位字符" json:"id_card_suffix"`
		QQ                string          `gorm:"column:qq;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:qq" json:"qq"`
		WeiXin            string          `gorm:"column:wei_xin;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:微信账号" json:"wei_xin"`
		DingDing          string          `gorm:"column:ding_ding;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:钉钉账号" json:"ding_ding"`
		WeiBo             string          `gorm:"column:wei_bo;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:微博账号" json:"wei_bo"`
		Signature         string          `gorm:"column:signature;not null;type:varchar(256) COLLATE utf8mb4_general_ci;comment:用户签名" json:"signature"`
		RegisterChannel   string          `gorm:"column:register_channel;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:账号注册渠道" json:"register_channel"`
		InviteCode        int64           `gorm:"column:invite_code;not null;default:0;type:int(10);comment:邀请码" json:"invite_code"`
	}

	UserEmail struct {
		ID              int        `gorm:"column:id;primary_key" json:"-"`
		UserHid         int64      `json:"user_hid" gorm:"column:user_hid;uniqueIndex:idx_userhid,priority:1;type:bigint(20);default:0;not null;"`
		UserIndexHid    string     `json:"user_index_hid" gorm:"column:user_index_hid;type:varchar(60);not null;comment:user_main表位置"`
		Email           string     `gorm:"column:email;uniqueIndex:idx_email,priority:1;not null;type:varchar(100);default:0;comment:邮箱" json:"-"`
		EmailVerifiedAt *time.Time `gorm:"column:email_verified_at;not null;uniqueIndex:idx_email,priority:3;type:datetime;default:'2000-01-01 00:00:00';comment:认证时间;" json:"-"`
		IsDel           int        `json:"is_del" gorm:"column:is_del;uniqueIndex:idx_userhid,priority:2;uniqueIndex:idx_email,priority:2;type:tinyint(2);default:0;comment:是否删除0-未删除 大于0-已删除"`
	}
	UserMobile struct {
		ID               int        `gorm:"column:id;primary_key" json:"-"`
		UserHid          int64      `json:"user_hid" gorm:"column:user_hid;uniqueIndex:inx_userhid,priority:1;not null;default:0;type:bigint(20) COLLATE utf8mb4_bin;"`
		UserIndexHid     string     `json:"user_index_hid" gorm:"column:user_index_hid;not null;default:'';type:varchar(60) COLLATE utf8mb4_bin;comment:user_main表位置"`
		CountryCode      string     `gorm:"column:country_code;uniqueIndex:idx_mobile,priority:2;type:varchar(15) COLLATE utf8mb4_bin;not null;comment:手机国别默认86" json:"country_code"`
		Mobile           string     `gorm:"column:mobile;not null;default:'';uniqueIndex:idx_mobile,priority:1;type:varchar(20) COLLATE utf8mb4_bin;comment:手机号" json:"-"`
		MobileVerifiedAt *time.Time `json:"mobile_verified_at" gorm:"column:mobile_verified_at;not null;uniqueIndex:idx_mobile,priority:4;default:'2000-01-01 00:00:00'"`
		IsDel            int        `json:"is_del" gorm:"column:is_del;type:tinyint(2);uniqueIndex:inx_userhid,priority:2;uniqueIndex:idx_mobile,priority:3;not null;idx_mobile,priority:1;default:0;comment:是否删除0-未删除 大于0-已删除"`
	}
)

// GetRealName 获取用户的真实姓名
func (r *ResultUserItem) GetRealName(nilDefaultValue ...string) (res string) {
	if r.RealName != "" {
		res = r.RealName
		return
	}
	if len(nilDefaultValue) > 0 {
		res = nilDefaultValue[0]
	}
	return
}

func (r *ResultUserItem) InitData(item *User) {
	r.UserHid = item.UserHid
	if item.UserMain != nil {
		r.AuthDesc = item.UserMain.AuthDesc
		r.Portrait = item.UserMain.Portrait
		r.PortraitUrl = item.UserMain.PortraitUrl
		r.NickName = item.UserMain.NickName
		r.Gender = item.UserMain.Gender
		r.Status = item.UserMain.Status
		r.Score = item.UserMain.Score
		r.IsV = item.UserMain.IsV
		r.UserEmailIndex = item.UserMain.UserEmailIndex
		r.UserMobileIndex = item.UserMain.UserMobileIndex
		r.HaveDashboard = item.UserMain.HaveDashboard
		r.ShopId = item.UserMain.CurrentShopId
	}
	if item.UserInfo != nil {
		r.Signature = item.UserInfo.Signature
		r.Remark = item.UserInfo.Remark
		r.RegisterChannel = item.UserInfo.RegisterChannel
		r.RealName = item.UserInfo.RealName
		r.RememberToken = item.UserInfo.RememberToken
		r.MsgReadTimeCursor = item.UserInfo.MsgReadTimeCursor

	}

	if item.UserEmail != nil {
		r.Email = item.UserEmail.Email
		r.EmailVerifiedAt = item.UserEmail.EmailVerifiedAt
	}
	if item.UserMobile != nil {
		r.Mobile = item.UserMobile.Mobile
		r.CountryCode = item.UserMobile.CountryCode
	}
	return
}

func (r *RequestUser) HaveShop() (res bool, err error) {
	if r.UShopId == 0 {
		err = fmt.Errorf("您当前的账号没有用户权限")
		return
	}
	res = true
	return
}

//   needValidateShop  //NeedValidateShopYes = true,需要校验当前用户是否有店铺权限  NeedValidateShopNo = false,不需要校验当前用户是否有店铺权限
func (r *RequestUser) InitRequestUser(ctx *base.Context, needValidateShop ...bool) (err error) {

	defer func() {
		if r.UShopId > 0 {
			return
		}
		if app_obj.App.UseDefaultShopId { //店铺ID默认值，调试数据使用（测试环境）
			shopId := ctx.GinContext.GetHeader(app_obj.HttpShopId)
			r.UShopId, _ = strconv.ParseInt(shopId, 10, 64)
		}
	}()
	r.UHeaderInfo = base.NewHeaderInfo()

	if err = r.UHeaderInfo.ParseFromString(ctx, ctx.GinContext.GetHeader(app_obj.HttpHeaderInfo)); err != nil {
		err = base.NewErrorRuntime(fmt.Errorf("系统异常,解析header失败"), base.ErrorNotLogin)
		return
	}
	if r.UUserHid > 0 {
		return
	}
	var uidString string
	if uidString = ctx.GinContext.GetHeader(app_obj.HttpUserHid); uidString == "" || uidString == "null" {
		err = base.NewErrorRuntime(fmt.Errorf("请先登录系统"), base.ErrorNotLogin)
		return
	}
	if r.UUserHid, err = strconv.ParseInt(uidString, 10, 64); err != nil {
		err = base.NewErrorRuntime(fmt.Errorf("用户信息参数格式不正确(uid:%s)", uidString), base.ErrorNotLogin)
		return
	}
	var user *ResultUser
	uidString = fmt.Sprintf("%d", r.UUserHid)
	if user, err = GetResultUserByUid(uidString, ctx); err != nil {
		return
	}
	r.SetResultUser(user)
	if len(needValidateShop) > 0 {
		if needValidateShop[0] && r.UShopId == 0 {
			err = base.NewErrorRuntime(fmt.Errorf("对不起,您当前的账号没有店铺管理权限"), base.ErrorHasNotPermit)

			return
		}
	}
	return
}

func (r *RequestUser) SetResultUser(user *ResultUser) {
	var userInfo ResultUserItem
	var ok bool
	if userInfo, ok = user.List[r.UUserHid]; !ok {
		return
	}

	r.UPortrait = userInfo.Portrait
	r.UNickName = userInfo.NickName
	r.UUserName = userInfo.UserName
	r.UGender = userInfo.Gender
	r.UStatus = userInfo.Status
	r.UShopId = userInfo.ShopId
	r.UHaveDashboard = userInfo.HaveDashboard
	r.UUserMobileIndex = userInfo.UserMobileIndex
	r.UUserEmailIndex = userInfo.UserEmailIndex
	r.URememberToken = userInfo.RememberToken
	r.UMsgReadTimeCursor = userInfo.MsgReadTimeCursor
	r.UHaveDashboard = userInfo.HaveDashboard
}

func GetResultUserByUid(userId string, ctx *base.Context) (res *ResultUser, err error) {
	var value = url.Values{}

	value.Set("user_hid", userId)
	value.Set("data_type", strings.Join([]string{UserDataTypeMain, UserDataTypeInfo, UserDataTypeEmail, UserDataTypeMobile}, ","))
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     AppNameUser,
		URI:         "/user/get_by_uid",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int         `json:"code"`
		Data *ResultUser `json:"data"`
		Msg  string      `json:"message"`
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
