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
)

const (
	SceneNameTendencies    = "tendencies"     //广告用户首页动态
	SceneNameHomeRecommend = "home_recommend" //广告首页数据推荐
)

const (
	TableAdDataCommon = "common" //推荐公共数据存储
)

const (
	AdDataDataTypeSpu               string = "1"   //商品信息
	AdDataDataTypeSku                      = "5"   //sku信息
	AdDataDataTypeSpuCategory              = "7"   //商品类目
	AdDataDataTypeSocialIntercourse        = "2"   //广告社交动态信息
	AdDataStatusUserSet                    = "3"   //用户手工设置
	AdDataDataTypeUserShop                 = "4"   //店铺后台
	AdDataDataTypeFishingSport             = "8"   //钓点信息
	AdDataDataTypeShopNotice               = "9"   //店铺公告（上传图片时使用）
	AdDataDataTypeShopLogo                 = "10"  //店铺LOGO（上传图片时使用）
	AdDataDataTypeShopBgImg                = "11"  //店铺背景图（上传图片时使用）
	AdDataDataTypeShopBrandQuality         = "12"  //店铺品牌资质（上传图片或文件时使用）
	AdDataDataTypeUser                     = "6"   //用户信息
	AdDataDataTypeIdCard                   = "13"  //用户身份证
	AdDataDataTypeUserAvatar               = "14"  //用户头像
	AdDataDataTypeSuggestion               = "15"  //投诉建议
	AdDataDataTypeOrderComment             = "16"  //订单评论
	AdDataDataTypeOrderCompany             = "17"  //公司资质
	AdDataDataTypePlatNotice               = "18"  //平台公告
	AdDataDataTypeChat                     = "19"  //聊天信息
	AdDataDataTypeHelpDocument             = "20"  //帮助文档文件
	AdDataDataTypeRing                     = "21"  //圈子
	AdDataDataTypeUserShopHome             = "22"  //店铺主页 （前台）

	AdDataDataTypeOther                    = "200" //其他数据
	AdDataDataTypeAllSpu                   = "-1"  //所有的商品

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
	}
	SliceAdDataType = base.ModelItemOptions{ //
		{
			Label: "电商_SPU",
			Value: AdDataDataTypeSpu,
		},
		{
			Label: "社交",
			Value: AdDataDataTypeSocialIntercourse,
		},
		{
			Label: "钓点",
			Value: AdDataDataTypeFishingSport,
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
		Ctx         *base.Context    `json:"-"`             //上下文信息
		UserHid     int64            `json:"user_hid"`      //用户
		DataType    string           `json:"data_type"`     //数据类型
		DataId      string           `json:"data_id"`       //数据ID
		SceneKey    string           `json:"scene_key"`     //场景KEY
		Status      uint8            `json:"status"`        //状态
		PullOnTime  *base.TimeNormal `json:"pull_on_time"`  //上架时间
		PullOffTime *base.TimeNormal `json:"pull_off_time"` //下架时间
		Weight      int64            `json:"weight"`        //权重
	}
)

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
