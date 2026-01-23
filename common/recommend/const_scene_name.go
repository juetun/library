package recommend

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	SceneNameTendencies           = "tendencies"             //广告用户首页动态
	SceneNameAttendUserTendencies = "attend_user_tendencies" //关注用户首页动态
	SceneNameHomeRecommend        = "home_recommend"         //广告首页数据推荐
)

const (
	TableAdDataCommon = "common" //推荐公共数据存储
)

const (
	AdDataDataTypeSpu               string = "1"  //商品信息
	AdDataDataTypeSocialIntercourse        = "2"  //广告社交动态信息
	AdDataStatusUserSet                    = "3"  //广告场景数据
	AdDataDataTypeUserShop                 = "4"  //店铺后台
	AdDataDataTypeSku                      = "5"  //sku信息
	AdDataDataTypeUser                     = "6"  //用户信息
	AdDataDataTypeSpuCategory              = "7"  //商品类目
	AdDataDataTypeFishingSport             = "8"  //钓点信息
	AdDataDataTypeShopNotice               = "9"  //店铺公告（上传图片时使用）
	AdDataDataTypeShopLogo                 = "10" //店铺LOGO（上传图片时使用）
	AdDataDataTypeShopBgImg                = "11" //店铺背景图（上传图片时使用）
	AdDataDataTypeShopBrandQuality         = "12" //店铺品牌资质（上传图片或文件时使用）
	AdDataDataTypeIdCard                   = "13" //用户身份证
	AdDataDataTypeUserAvatar               = "14" //用户头像
	AdDataDataTypeSuggestion               = "15" //投诉建议
	AdDataDataTypeOrderComment             = "16" //订单评论
	AdDataDataTypeOrderCompany             = "17" //公司资质
	AdDataDataTypePlatNotice               = "18" //平台公告
	AdDataDataTypeChat                     = "19" //聊天信息
	AdDataDataTypeHelpDocument             = "20" //帮助文档文件
	AdDataDataTypeRing                     = "21" //圈子
	AdDataDataTypeUserShopHome             = "22" //店铺主页 （前台）
	AdDataDataTypeRingComment              = "23" //评论 （圈子动态或钓点信息）
	AdDataDataTypeExpress                  = "24" //物流公司LOGO
	AdDataDataTypeFactory                  = "25" //厂家LOGO
	AdDataDataTypePlatActivity             = "26" //平台活动缩略图
	AdDataStatusDataScene                  = "27" //广告场景
	AdDataStatusOrderDetail                = "28" //订单详情
	AdDataStatusOrderSubDetail             = "29" //子订单详情

	AdDataDataTypeAppLogo    = "198" //手机APP logo
	AdDataDataTypeAppInstall = "199" //手机APP安装包
	AdDataDataTypeOther      = "200" //其他数据
	AdDataDataTypeAllSpu     = "-1"  //所有的商品

)

const AdDataDataTypeGetSnsData = "-2" //获取社交数据使用

const (
	AdDataStatusCanUse  uint8 = iota + 1 //广告可用
	AdDataStatusOffLine                  //广告下架
)

var (
	//业务数据
	SliceAdDataTypeBiz = base.ModelItemOptions{
		{
			Label: "电商",
			Value: AdDataDataTypeSpu,
		},
		{
			Label: "社交",
			Value: AdDataDataTypeSocialIntercourse,
		},
		{
			Label: "钓点信息",
			Value: AdDataDataTypeFishingSport,
		},
		{
			Label: "圈子",
			Value: AdDataDataTypeRing,
		},
		{
			Label: "用户信息",
			Value: AdDataDataTypeUser,
		},
		{
			Label: "店铺信息",
			Value: AdDataDataTypeUserShopHome,
		},
	}
	SliceAdDataType = base.ModelItemOptions{ //
		{
			Label: "广告场景",
			Value: AdDataStatusDataScene,
		},
		{
			Label: "电商_SPU",
			Value: AdDataDataTypeSpu,
		},
		{
			Label: "店铺",
			Value: AdDataDataTypeUserShopHome,
		},
		{
			Label: "用户",
			Value: AdDataDataTypeUser,
		},
		{
			Label: "平台活动",
			Value: AdDataDataTypePlatActivity,
		},
		{
			Label: "社交",
			Value: AdDataDataTypeSocialIntercourse,
		},
		{
			Label: "圈子",
			Value: AdDataDataTypeRing,
		},
		{
			Label: "钓点",
			Value: AdDataDataTypeFishingSport,
		},
		{
			Label: "电商类目",
			Value: AdDataDataTypeSpuCategory,
		},
		{
			Label: "电商_SKU",
			Value: AdDataDataTypeSku,
		},
		{
			Label: "客服手工指定",
			Value: AdDataStatusUserSet,
		},
		{
			Label: "客户店铺手工指定",
			Value: AdDataDataTypeUserShop,
		},
	}
	SliceAdDataStatus = base.ModelItemOptions{
		{
			Label: "可用",
			Value: AdDataStatusCanUse,
		},
		{
			Label: "下架",
			Value: AdDataStatusOffLine,
		},
	}
)

type (
	ArgWriteRecommendData struct {
		Ctx *base.Context `json:"-"` //上下文信息
		ArgWriteRecommendDataSingle
	}
	ArgWriteRecommendDataSingle struct {
		UserHid     int64            `json:"user_hid"`      //用户
		ToUserHid   int64            `json:"to_user_hid"`   //如果指定了固定用户展示值 默认0 适用个性化推荐数据
		DataType    string           `json:"data_type"`     //数据类型
		DataId      string           `json:"data_id"`       //数据ID
		ShopId      string           `json:"shop_id"`       //店铺ID
		SceneKey    string           `json:"scene_key"`     //场景KEY
		CategoryId  int64            `json:"category_id"`   //商品所属类目
		Status      uint8            `json:"status"`        //状态
		PullOnTime  *base.TimeNormal `json:"pull_on_time"`  //上架时间
		PullOffTime *base.TimeNormal `json:"pull_off_time"` //下架时间
		Weight      int64            `json:"weight"`        //权重
	}
	ArgWriteRecommendDataList struct {
		SceneKey   string                         `json:"scene_key"` //广告场景KEY
		CurrentUid int64                          `json:"current_uid"`
		Ctx        *base.Context                  `json:"-"` //上下文信息
		List       []*ArgWriteRecommendDataSingle `json:"list"`
	}

	ArgInitTrendDataImport struct {
		ActType         string          `json:"act_type" form:"act_type"`       //操作动作 枚举型 attend-关注 cancel-取关
		CurrentLoginUid int64           `json:"current_uid" form:"current_uid"` //当前登录用户
		UidString       string          `json:"user_hids" form:"user_hids"`
		UserHids        []int64         `json:"-" form:"-"` //多个用户的目的是批量关注数据导入
		TimeNow         base.TimeNormal `json:"-" form:"-"`
	}
	ArgAttendUserImport struct {
		AttendedUid int64           `json:"attended_uid"` //被关注用户Id
		TimeNow     base.TimeNormal `json:"-" form:"-"`
	}
)

func (r *ArgAttendUserImport) Default(c *base.Context) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}
func (r *ArgWriteRecommendData) Default() (err error) {
	var (
		dataTypeMap map[string]string
		ok          bool
	)
	if dataTypeMap, err = SliceAdDataType.GetMapAsKeyString(); err != nil {
		if _, ok = dataTypeMap[r.DataType]; !ok {
			err = fmt.Errorf("当前不支持你选择的数据类型")
			r.Ctx.Error(map[string]interface{}{
				"arg": r,
				"err": err.Error(),
			}, "ArgWriteRecommendDataDefault")
			return
		}
		return
	}
	switch r.SceneKey {
	case SceneNameTendencies: //如果是用户首页广告动态,则必须填写用户ID
		if r.UserHid == 0 {
			err = fmt.Errorf("请填写要推送的用户信息")
			return
		}
	}

	if r.Status == 0 { //默认不传为可用数据
		r.Status = AdDataStatusCanUse
	}
	return
}

func (r *ArgWriteRecommendData) GetJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

//TODO 根据店铺ID获取店铺信息，
//当前采用接口直接发送，后续采用消息队列解耦重新实现
func WriteRecommendData(arg *ArgWriteRecommendData) (res bool, err error) {
	if err = arg.Default(); err != nil {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameRecommend,
		URI:         "/recommend/write_data",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     arg.Ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int                                   `json:"code"`
		Data struct{ Result bool `json:"result"` } `json:"data"`
		Msg  string                                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res = data.Data.Result
	return
}

//当前采用接口直接发送，后续采用消息队列解耦重新实现
func WriteRecommendDataList(arg *ArgWriteRecommendDataList) (res bool, err error) {
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameRecommend,
		URI:         "/recommend/write_data_batch",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     arg.Ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var (
		data = struct {
			Code int                                   `json:"code"`
			Data struct{ Result bool `json:"result"` } `json:"data"`
			Msg  string                                `json:"message"`
		}{}
		body []byte
	)
	req := rpc.NewHttpRpc(&ro).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	if body = req.Body; len(body) == 0 {
		err = fmt.Errorf("系统异常,请重试")
		return
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return
	}
	if data.Code > 0 {
		err = fmt.Errorf(data.Msg)
		return
	}
	res = data.Data.Result
	return
}

//移除取消关注的数据
func RemoveAttendTrendData(arg *ArgInitTrendDataImport, ctx *base.Context) (res bool, err error) {
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameRecommend,
		URI:         "/recommend/remove_attend_trend_data",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int                                   `json:"code"`
		Data struct{ Result bool `json:"result"` } `json:"data"`
		Msg  string                                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res = data.Data.Result
	return
}

//关注当前用户的动态更新
func AttendTrendUserImport(arg *ArgInitTrendDataImport, ctx *base.Context) (res bool, err error) {
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameComment,
		URI:         "/trends/attend_user_import",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int                                   `json:"code"`
		Data struct{ Result bool `json:"result"` } `json:"data"`
		Msg  string                                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res = data.Data.Result
	return
}

//更新所有关注指定用户的动态
func AttendTrendedUserImport(arg *ArgAttendUserImport, ctx *base.Context) (res bool, err error) {
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameComment,
		URI:         "/trends/attend_user_trend_import",
		Header:      http.Header{},
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int                                   `json:"code"`
		Data struct{ Result bool `json:"result"` } `json:"data"`
		Msg  string                                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res = data.Data.Result
	return
}

func (r *ArgWriteRecommendDataList) Default(c *base.Context) (err error) {

	return
}

func (r *ArgWriteRecommendDataList) GetJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

func (r *ArgAttendUserImport) GetJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

func (r *ArgInitTrendDataImport) GetJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

func (r *ArgInitTrendDataImport) Default(ctx *base.Context) (err error) {
	if err = r.InitWithUidStr(); err != nil {
		return
	}
	r.TimeNow = base.GetNowTimeNormal()
	return
}

func (r *ArgInitTrendDataImport) InitWithUidStr() (err error) {
	if r.UidString == "" {
		return
	}
	uidSlice := strings.Split(r.UidString, ",")
	r.UserHids = make([]int64, 0, len(uidSlice))
	var uid int64
	for _, item := range uidSlice {
		if item == "" {
			continue
		}
		if uid, err = strconv.ParseInt(item, 10, 64); err != nil {
			return
		}
		r.UserHids = append(r.UserHids, uid)
	}
	return
}
