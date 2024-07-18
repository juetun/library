package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// 获取用户信息对应的表
const (
	UserDataTypeIndex = "user_index"

	UserDataTypeMain        = "user_main"
	UserDataTypePortrait    = "user_portrait"     //用户头像
	UserDataTypePortraitTmp = "user_portrait_tmp" //用户临时头像
	UserDataTypeInfo        = "user_info"

	UserHidDivString      = "_" // 用户ID字符串切割
	UpdateColumnDivString = "." // 修改用户数据时传参的分割符
)
const (
	NeedValidateShopYes = true  //需要校验当前用户是否有店铺权限
	NeedValidateShopNo  = false //不需要校验当前用户是否有有店铺权限
)
const (
	UserHaveDashboardYes uint8 = iota + 1 //用户是否有管理后台权限 (有)
	UserHaveDashboardNo                   //用户是否有管理后台权限 (无)
)

var (
	SliceUserHaveDashboard = base.ModelItemOptions{
		{
			Value: UserHaveDashboardYes,
			Label: "有",
		},
		{
			Value: UserHaveDashboardNo,
			Label: "无",
		},
	}
)

const (
	RegChannelAdmin   = "admin"   // 客服后台注册的账号 渠道标记
	RegChannelAndroid = "android" // 客服后台注册的账号 安卓
	RegChannelIos     = "ios"     // 客服后台注册的账号 ios
	RegChannelH5      = "h5"      // 客服后台注册的账号 h5
	RegChannelWebSite = "website" // 客服后台注册的账号 h5
	RegChannelWeiXin  = "wei_xin" // 客服后台注册的账号 微信
	RegChannelWeiBo   = "wei_bo"  // 客服后台注册的账号 微博
	RegChannelAliPay  = "alipay"  // 客服后台注册的账号 支付宝
	RegChannelTikTok  = "tiktok"  // 客服后台注册的账号 抖音
	RegChannelBaiDu   = "baidu"   // 客服后台注册的账号 百度
)

const (
	DefaultBirthDay = "1970-01-01" //默认生日
)

var (
	IdCardSecret = "ABCDEFGHIJKLMNOP" // 加密KEY长度必须为16的倍数

	RegisterChannelMap = map[string]string{
		RegChannelAdmin:   "管理后台",
		RegChannelAndroid: "安卓",
		RegChannelIos:     "IOS",
		RegChannelH5:      "m站",
		RegChannelWebSite: "网站",
		RegChannelWeiXin:  "微信",
		RegChannelWeiBo:   "微博",
		RegChannelAliPay:  "支付宝",
		RegChannelTikTok:  "抖音",
		RegChannelBaiDu:   "百度",
	}
)

// 获取用户信息的响应参数结构
type (
	ResultUser struct {
		List map[int64]ResultUserItem `json:"list"`
	}
	ResultUserItem struct {
		UserHid           int64            `json:"user_hid,omitempty"`  // 用户ID
		Portrait          string           `json:"portrait,omitempty"`  // 头像
		PortraitUrl       string           `json:"portrait_url"`        //头像链接
		NickName          string           `json:"nick_name,omitempty"` // 昵称
		UserName          string           `json:"user_name,omitempty"` // 用户名
		RealName          string           `json:"real_name"`           // 真实姓名
		Gender            uint8            `json:"gender,omitempty"`    //
		Status            int8             `json:"status,omitempty"`    //
		Score             int              `json:"score,omitempty"`     //
		AuthStatus        uint8            `json:"auth_status,omitempty"`
		AuthDesc          string           `json:"auth_desc,omitempty"` // 认证描述
		AuthType          uint8            `json:"auth_type,omitempty"` // 认证类型
		IsV               int              `json:"is_v,omitempty"`      // 用户头像加V
		Remark            string           `json:"remark" `             // 个性签名
		Signature         string           `json:"signature,omitempty"`
		RegisterChannel   string           `json:"register_channel,omitempty"`
		CountryCode       string           `json:"country_code,omitempty"`
		Mobile            string           `json:"mobile,omitempty"`
		MobileVerifiedAt  *base.TimeNormal `json:"mobile_verified_at,omitempty"`
		Email             string           `json:"email,omitempty"`
		EmailVerifiedAt   *base.TimeNormal `json:"email_verified_at,omitempty"`
		ShopId            int64            `json:"shop_id"`
		UserMobileIndex   string           `json:"user_mobile_index"`
		UserEmailIndex    string           `json:"user_email_index"`
		RememberToken     string           `json:"remember_token"`
		MsgReadTimeCursor base.TimeNormal  `json:"msg_read_time_cursor"`
		HaveDashboard     uint8            `json:"have_dashboard"`
		IsMocking         bool             `json:"is_mocking"` //当前是否是模拟账号
	}
	RequestUser struct {
		UUserHid           int64           `json:"u_user_hid" form:"u_user_hid"`                         //用户
		UUserMobileIndex   string          `json:"u_user_mobile_index" form:"u_user_mobile_index"`       //手机数据存储位置
		UUserEmailIndex    string          `json:"u_user_email_index" form:"u_user_email_index"`         //email存储位置
		UPortrait          string          `json:"u_portrait" form:"u_portrait"`                         //头像
		UNickName          string          `json:"u_nick_name" form:"u_nick_name"`                       //昵称
		UUserName          string          `json:"u_user_name" form:"u_user_name"`                       //账号
		UGender            uint8           `json:"u_gender" form:"u_gender"`                             //性别
		UStatus            int8            `json:"u_status" form:"u_status"`                             //状态
		UAuthStatus        uint8           `json:"u_auth_status" form:"u_auth_status"`                   //认证审核状态
		UAuthType          uint8           `json:"u_auth_type" form:"u_auth_type"`                       //认证类型
		UScore             int             `json:"u_score" form:"u_score"`                               //积分
		URememberToken     string          `json:"u_remember_token" form:"u_remember_token"`             //是否记住密码
		UMsgReadTimeCursor base.TimeNormal `json:"u_msg_read_time_cursor" form:"u_msg_read_time_cursor"` //消息未读时刻节点
		UShopId            int64           `json:"u_shop_id" form:"u_shop_id"`                           //店铺ID
		UHaveDashboard     uint8           `json:"u_have_dashboard" form:"u_have_dashboard"`             //是否有客服后台权限
		UIsMocking         bool            `json:"uis_mocking" form:"uis_mocking"`                       //当前是否在模拟状态
		//UHeaderInfo *base.HeaderInfo `json:"u_header_info" form:"u_header_info"` //用户设备信息
	}

	User struct {
		UserHid   int64      `json:"user_hid"`
		UserIndex *UserIndex `json:"user_index,omitempty"`
		UserMain  *UserMain  `json:"user_main,omitempty"`
		UserInfo  *UserInfo  `json:"user_info,omitempty"`
	}
	UserIndex struct {
		ID            int64            `gorm:"column:id;primary_key" json:"id"`
		UserName      string           `gorm:"column:user_name;not null;type:varchar(50) COLLATE utf8mb4_bin;default:'';uniqueIndex;comment:用户名" json:"user_name" `
		TmpAccount    string           `gorm:"column:tmp_account;not null;type:varchar(200) COLLATE utf8mb4_bin;default:'';comment:注册时临时账号" json:"tmp_account" `
		IsUse         int              `json:"is_use" gorm:"column:is_use;not null;type:tinyint(1);default:0;comment:是否启用 0-未启用 大于0-已启用"`
		AttendNum     int64            `gorm:"column:attend_num;not null;default:0;type:bigint(20);comment:关注数 实时性不是高" json:"attend_num"`   // 关注数
		LoveNum       int64            `gorm:"column:love_num;not null;default:0;type:bigint(20);comment:点赞数 实时性不是高" json:"love_num"`       // 点赞数
		CommentNum    int64            `gorm:"column:comment_num;not null;default:0;type:bigint(20);comment:评论数 实时性不是高" json:"comment_num"` // 评论数
		ApplyFlag     string           `gorm:"column:apply_flag;not null;type:varchar(200) COLLATE utf8mb4_bin;default:'';comment:审核位管理" json:"apply_flag" `
		UpdateBatchId string           `gorm:"column:update_batch_id;not null;type:varchar(200) COLLATE utf8mb4_bin;default:'';comment:变更数据的ID" json:"update_batch_id" `
		AuthType      uint8            `gorm:"column:auth_type;not null;type:tinyint(2);default:0;comment:认证类型 0-个人用户 1-企业用户" json:"auth_type"`
		Status        int              `gorm:"column:status;not null;type:tinyint(1);default:0;comment:状态 0-可用 1-不可用" json:"status"`
		Email         string           `gorm:"column:email;not null;type:varchar(100);default:'';index:idx_email;comment:邮箱" json:"-"`
		CountryCode   string           `gorm:"column:country_code;type:varchar(15) COLLATE utf8mb4_bin;default:'';index:inx_mobile,priority:2;not null;comment:手机国别默认86" json:"country_code"`
		Mobile        string           `gorm:"column:mobile;not null;default:'';type:varchar(20) COLLATE utf8mb4_bin;index:inx_mobile,priority:1;comment:手机号" json:"-"`
		WeiXinToken   string           `gorm:"column:wei_xin_token;not null;default:'';type:varchar(255) COLLATE utf8mb4_bin;comment:微信登录token" json:"wei_xin_token"`
		AliPayToken   string           `gorm:"column:ali_pay_token;not null;default:'';type:varchar(255) COLLATE utf8mb4_bin;comment:支付宝授权登录token" json:"ali_pay_token"`
		TikTokToken   string           `gorm:"column:tik_tok_token;not null;default:'';type:varchar(255) COLLATE utf8mb4_bin;comment:抖音小程序授权登录token" json:"tik_tok_token"`
		BaiduToken    string           `gorm:"column:baidu_token;not null;default:'';type:varchar(255) COLLATE utf8mb4_bin;comment:百度小程序授权登录token" json:"baidu_token"`
		WeiboToken    string           `gorm:"column:weibo_token;not null;default:'';type:varchar(255) COLLATE utf8mb4_bin;comment:微博授权登录token" json:"weibo_token"`
		CreatedAt     base.TimeNormal  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;" `
		UpdatedAt     base.TimeNormal  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;" `
		DeletedAt     *base.TimeNormal `json:"-" gorm:"column:deleted_at;"`
		DataName      string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
		TbName        string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
	}
	UserInfo struct {
		ID                int              `gorm:"column:id;primary_key" json:"id"`
		RealName          string           `gorm:"column:real_name;type:varchar(60);not null;comment:真实姓名"  json:"real_name"`
		UserIndexHid      string           `gorm:"column:user_index_hid;type:varchar(60);not null;comment:user_main表位置" json:"user_index_hid"`
		UserHid           int64            `gorm:"column:user_hid;not null;uniqueIndex:idx_userhid;type:bigint(20) COLLATE utf8mb4_bin" json:"user_hid"`
		Remark            string           `json:"remark" gorm:"column:remark;not null;type:varchar(150);default:'';comment:个性签名"` // 个性签名
		RememberToken     string           `gorm:"column:remember_token;not null;default:'';size:500;comment:登录的token" json:"remember_token"`
		MsgReadTimeCursor base.TimeNormal  `gorm:"column:msg_read_time_cursor;not null;default:CURRENT_TIMESTAMP;comment:最近一次读取系统公告时间" json:"msg_read_time_cursor"`
		Level             uint8            `gorm:"column:level;not null;type:tinyint(2);default:0;comment:用户等级0-普通用户" json:"level"`
		Password          string           `gorm:"column:password;not null;type:varchar(256) COLLATE utf8mb4_general_ci;comment:密码" json:"password"`
		IdCard            string           `gorm:"column:id_card;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:身份证号加密串" json:"id_card"`
		IdCardSuffix      string           `gorm:"column:id_card_suffix;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:后缀后6位字符" json:"id_card_suffix"`
		QQ                string           `gorm:"column:qq;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:qq" json:"qq"`
		WeiXin            string           `gorm:"column:wei_xin;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:微信账号" json:"wei_xin"`
		BirthDay          base.TimeNormal  `gorm:"column:birth_day;not null;default:'1970-01-01';type:date;comment:出生日期" json:"birth_day"`
		BirCountry        string           `gorm:"column:bir_country;not null;type:varchar(30) COLLATE utf8mb4_general_ci;comment:国籍" json:"bir_country"`
		BirProvinceId     int              `gorm:"column:bir_province_id;not null;type:int(10);default:0;comment:出生省份" json:"bir_province_id"`
		BirCityId         int              `gorm:"column:bir_city_id;not null;type:int(10);default:0;comment:出生城市" json:"bir_city_id"`
		BirAreaId         int              `gorm:"column:bir_area_id;not null;type:int(10);default:0;comment:出生区或县" json:"bir_area_id"`
		DingDing          string           `gorm:"column:ding_ding;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:钉钉账号" json:"ding_ding"`
		WeiBo             string           `gorm:"column:wei_bo;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:微博账号" json:"wei_bo"`
		Signature         string           `gorm:"column:signature;not null;type:varchar(256) COLLATE utf8mb4_general_ci;comment:用户签名" json:"signature"`
		RegisterChannel   string           `gorm:"column:register_channel;not null;type:varchar(50) COLLATE utf8mb4_general_ci;comment:账号注册渠道" json:"register_channel"`
		MockAdminToken    string           `gorm:"column:mock_admin_token;not null;type:varchar(1000) COLLATE utf8mb4_general_ci;comment:mock用户token" json:"mock_admin_token,omitempty"`
		MockAdminUid      int64            `gorm:"column:mock_admin_uid;not null;type:bigint(20) COLLATE utf8mb4_bin;comment:mock用户ID" json:"mock_admin_uid,omitempty"`
		InviteCode        int64            `gorm:"column:invite_code;not null;default:0;type:int(10);comment:邀请码" json:"invite_code"`
		AttendNum         int64            `gorm:"column:attend_num;not null;default:0;type:bigint(20);comment:关注数 实时性不是高" json:"attend_num"`   // 关注数
		LoveNum           int64            `gorm:"column:love_num;not null;default:0;type:bigint(20);comment:点赞数 实时性不是高" json:"love_num"`       // 点赞数
		CommentNum        int64            `gorm:"column:comment_num;not null;default:0;type:bigint(20);comment:评论数 实时性不是高" json:"comment_num"` // 评论数
		MobileVerifiedAt  *base.TimeNormal `json:"mobile_verified_at" gorm:"column:mobile_verified_at;uniqueIndex:idx_mobile,priority:4;default:'2000-01-01 00:00:00'"`
		EmailVerifiedAt   *base.TimeNormal `gorm:"column:email_verified_at;uniqueIndex:idx_email,priority:3;type:datetime;default:'2000-01-01 00:00:00';comment:认证时间;" json:"-"`
		DataName          string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
		TbName            string           `json:"-" gorm:"-"` // 数据所在分库分表的位置
	}

	UserEmail struct {
		ID              int              `gorm:"column:id;primary_key" json:"-"`
		UserHid         int64            `json:"user_hid" gorm:"column:user_hid;uniqueIndex:idx_userhid,priority:1;type:bigint(20);default:0;not null;"`
		UserIndexHid    string           `json:"user_index_hid" gorm:"column:user_index_hid;type:varchar(60);not null;comment:user_main表位置"`
		Email           string           `gorm:"column:email;uniqueIndex:idx_email,priority:1;not null;type:varchar(100);default:0;comment:邮箱" json:"-"`
		EmailVerifiedAt *base.TimeNormal `gorm:"column:email_verified_at;not null;uniqueIndex:idx_email,priority:3;type:datetime;default:'2000-01-01 00:00:00';comment:认证时间;" json:"-"`
		IsDel           int              `json:"is_del" gorm:"column:is_del;uniqueIndex:idx_userhid,priority:2;uniqueIndex:idx_email,priority:2;type:tinyint(2);default:0;comment:是否删除0-未删除 大于0-已删除"`
	}
	UserMobile struct {
		ID               int              `gorm:"column:id;primary_key" json:"-"`
		UserHid          int64            `json:"user_hid" gorm:"column:user_hid;uniqueIndex:inx_userhid,priority:1;not null;default:0;type:bigint(20) COLLATE utf8mb4_bin;"`
		UserIndexHid     string           `json:"user_index_hid" gorm:"column:user_index_hid;not null;default:'';type:varchar(60) COLLATE utf8mb4_bin;comment:user_main表位置"`
		CountryCode      string           `gorm:"column:country_code;uniqueIndex:idx_mobile,priority:2;type:varchar(15) COLLATE utf8mb4_bin;not null;comment:手机国别默认86" json:"country_code"`
		Mobile           string           `gorm:"column:mobile;not null;default:'';uniqueIndex:idx_mobile,priority:1;type:varchar(20) COLLATE utf8mb4_bin;comment:手机号" json:"-"`
		MobileVerifiedAt *base.TimeNormal `json:"mobile_verified_at" gorm:"column:mobile_verified_at;not null;uniqueIndex:idx_mobile,priority:4;default:'2000-01-01 00:00:00'"`
		IsDel            int              `json:"is_del" gorm:"column:is_del;type:tinyint(2);uniqueIndex:inx_userhid,priority:2;uniqueIndex:idx_mobile,priority:3;not null;idx_mobile,priority:1;default:0;comment:是否删除0-未删除 大于0-已删除"`
	}
)

//根据token 获取当前登录用户信息
func GetUserInfoByXAuthToken(xAuthToken string, ctx *base.Context) (requestUser *RequestUser, err error) {
	requestUser = &RequestUser{}
	var (
		jwtUser = base.JwtUser{}
		user    *ResultUser
	)
	if err = base.ParseJwtKey(xAuthToken, ctx, &jwtUser); err != nil { // 如果解析token失败
		err = fmt.Errorf("用户信息异常")
		return
	}
	if jwtUser.UserId > 0 {
		requestUser.UUserHid = jwtUser.UserId
		if user, err = GetResultUserByUid(fmt.Sprintf("%d", jwtUser.UserId), ctx); err != nil {
			return
		}
		requestUser.SetResultUser(user)
	}
	return
}

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
		r.AuthStatus = item.UserMain.AuthStatus
		r.AuthType = item.UserMain.AuthType
		r.AuthDesc = item.UserMain.AuthDesc
		r.Portrait = item.UserMain.Portrait
		r.PortraitUrl = item.UserMain.PortraitUrl
		if r.PortraitUrl == "" {
			r.PortraitUrl = item.UserMain.PortraitTmpUrl
		}
		r.NickName = item.UserMain.NickName
		r.Gender = item.UserMain.Gender
		r.Status = item.UserMain.Status
		r.Score = item.UserMain.Score
		r.IsV = item.UserMain.IsV
		r.Mobile = item.UserMain.Mobile
		r.MobileVerifiedAt = item.UserMain.MobileVerifiedAt
		r.Email = item.UserMain.Email
		r.EmailVerifiedAt = item.UserMain.EmailVerifiedAt
		r.HaveDashboard = item.UserMain.HaveDashboard
		r.ShopId = item.UserMain.CurrentShopId
	}

	//如果当前账号有管理员权限

	if item.UserInfo != nil {
		r.Signature = item.UserInfo.Signature
		r.Remark = item.UserInfo.Remark
		r.RegisterChannel = item.UserInfo.RegisterChannel
		r.RealName = item.UserInfo.RealName
		r.RememberToken = item.UserInfo.RememberToken
		r.MsgReadTimeCursor = item.UserInfo.MsgReadTimeCursor
		if item.UserInfo.MockAdminToken != "" {
			r.IsMocking = true
		}
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
		if len(needValidateShop) > 0 {
			if needValidateShop[0] && r.UShopId == 0 {
				err = base.NewErrorRuntime(fmt.Errorf("对不起,您当前的账号没有店铺管理权限"), base.ErrorHasNotPermit)
				return
			}
		}
		if r.UShopId > 0 {
			return
		}
		if app_obj.App.UseDefaultShopId { //店铺ID默认值，调试数据使用（测试环境）
			shopId := ctx.GinContext.GetHeader(app_obj.HttpShopId)
			r.UShopId, _ = strconv.ParseInt(shopId, 10, 64)
		}
	}()
	//获取用户信息
	if r.UUserHid == 0 {
		jwtUser, exit := common.TokenValidate(ctx, true)
		if exit {
			err = base.NewErrorRuntime(fmt.Errorf("请登录账号"), base.ErrorNotLogin)
			return
		}
		r.UUserHid = jwtUser.UserId
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
	if user, err = GetResultUserByUid(fmt.Sprintf("%d", r.UUserHid), ctx); err != nil {
		return
	}
	r.SetResultUser(user)

	return
}

func (r *RequestUser) SetResultUser(user *ResultUser) {
	if user == nil {
		return
	}
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
	r.UAuthStatus = userInfo.AuthStatus
	r.UShopId = userInfo.ShopId
	r.UHaveDashboard = userInfo.HaveDashboard
	r.UUserMobileIndex = userInfo.UserMobileIndex
	r.UUserEmailIndex = userInfo.UserEmailIndex
	r.URememberToken = userInfo.RememberToken
	r.UMsgReadTimeCursor = userInfo.MsgReadTimeCursor
	r.UHaveDashboard = userInfo.HaveDashboard
	r.UIsMocking = userInfo.IsMocking
}

func GetResultUserByUid(userId string, ctx *base.Context, dataTypes ...string) (res *ResultUser, err error) {
	res = &ResultUser{List: make(map[int64]ResultUserItem, )}
	var value = url.Values{}

	value.Set("user_hid", userId)
	if len(dataTypes) == 0 {
		dataTypes = []string{UserDataTypeMain, UserDataTypeInfo}
	}
	value.Set("data_type", strings.Join(dataTypes, ","))
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
