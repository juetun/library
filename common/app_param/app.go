package app_param

import (
	"github.com/juetun/base-wrapper/lib/base"
	"sync"
)

const (
	AppNameAdmin             = "admin-main"
	AppNameUpload            = "api-upload"
	AppNameExport            = "api-export"
	AppNameUser              = "api-user"
	AppNameTag               = "api-tag"
	AppNameNotice            = "api-notice"
	AppNameComment           = "api-comment"      //社交评论
	AppNameChat              = "api-chat"         //聊天
	AppNameCar               = "api-car"          //汽车
	AppNameMall              = "api-mall"         //电商
	AppNameMallOrder         = "api-order"        //订单
	AppNameMallOrderComment  = "api-ordercomment" //订单评论
	AppNameMallActivity      = "api-activity"     //电商活动
	AppNameRecommend         = "api-recommend"    //推荐
	AppNameSocialIntercourse = "api-sns"          //社交
)

const (
	TerminalWeb       = "web" //web站
	TerminalMina      = "mina"
	TerminalH5        = "h5"
	TerminalAndroid   = "android"
	TerminalWINPHONE  = "winphone"
	TerminalIos       = "ios"
	TerminalHarmonyOS = "harmonyOS"
)

var (
	SliceTerminal = base.ModelItemOptions{
		{Label: "网站", Value: TerminalWeb},
		{Label: "小程序", Value: TerminalMina},
		{Label: "h5", Value: TerminalH5},
		{Label: "安卓", Value: TerminalAndroid},
		{Label: "IOS", Value: TerminalIos},
		{Label: "华为鸿蒙", Value: TerminalHarmonyOS},
	}
)

//标签类型定义
const (
	DataPapersGroupCategoryTag          = "user_tag"           // 用户标签
	DataPapersGroupCategoryMallCategory = "mall_category"      // 电商类目
	DataPapersGroupCategoryMallBrand    = "mall_brand_quality" // 电商品牌类目
)

var MapDataPapersGroupCategory = map[string]string{
	DataPapersGroupCategoryTag:          "用户标签",
	DataPapersGroupCategoryMallCategory: "电商类目",
	DataPapersGroupCategoryMallBrand:    "电商品牌",
}
var (
	//当前动作服务支持的动动作 key 和key对应的描述映射
	//如: map[string]string{"user_reg":"用户注册信息"}
	TrendsTypes = make(map[string]string, 150)

	SliceAppNames = base.ModelItemOptions{
		{
			Label: "客服后台",
			Value: AppNameAdmin,
		},
		{
			Label: "上传",
			Value: AppNameUpload,
		},
		{
			Label: "导入导出",
			Value: AppNameExport,
		},
		{
			Label: "用户",
			Value: AppNameUser,
		},
		{
			Label: "标签",
			Value: AppNameTag,
		},
		{
			Label: "评论",
			Value: AppNameComment,
		},
		{
			Label: "私信",
			Value: AppNameChat,
		},
		{
			Label: "汽车资讯",
			Value: AppNameCar,
		},
		{
			Label: "电商",
			Value: AppNameMall,
		},
		{
			Label: "订单",
			Value: AppNameMallOrder,
		},
		{
			Label: "电商评论",
			Value: AppNameMallOrderComment,
		},
		{
			Label: "电商活动",
			Value: AppNameMallActivity,
		},
		{
			Label: "广告推荐",
			Value: AppNameRecommend,
		},
		{
			Label: "社交",
			Value: AppNameSocialIntercourse,
		},
	}
)

//添加动态类型 map格式
func AppendTrendsTypesAsMap(mapTrendTypes map[string]string) {
	if len(mapTrendTypes) == 0 {
		return
	}
	var (
		syncLock sync.RWMutex
	)
	syncLock.Lock()
	defer syncLock.Unlock()
	for key, item := range mapTrendTypes {
		TrendsTypes[key] = item
	}
	return
}

//添加动态类型 ModelItemOptions格式
func AppendTrendsTypesAsModelItemOptions(trendsTypes base.ModelItemOptions) {
	var (
		mapTrendTypes, _ = trendsTypes.GetMapAsKeyString()
	)
	AppendTrendsTypesAsMap(mapTrendTypes)
	return
}
