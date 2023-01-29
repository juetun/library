package app_param

import "github.com/juetun/base-wrapper/lib/base"

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
