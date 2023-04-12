package app_user

import "github.com/juetun/base-wrapper/lib/base"

const (
	UpdateBatchTypeInfo        = "shop_info"    //店铺信息
	UpdateBatchTypeBrand       = "shop_brand"   //店铺品牌
	UpdateBatchTypeCate        = "shop_cate"    //店铺类目
	UpdateBatchTypeAuthType    = "auth_type"    //店铺入驻审核
	UpdateBatchTypeAvatar      = "avatar"       //用户头像审核
	UpdateBatchTypeUserProfile = "user_profile" //用户资料
)
const (
	UserApplyStatusUsing   uint8 = iota + 1 //使用中
	UserApplyStatusInvalid                  //失效
	UserApplyStatusFailure                  //审核失败
	UserApplyStatusInit                     //编辑中
	UserApplyStatusSubmit                   //提交审核
)

var (
	SliceUpdateBatchType = base.ModelItemOptions{
		{
			Value: UpdateBatchTypeInfo,
			Label: "店铺信息",
		},
		{
			Value: UpdateBatchTypeBrand,
			Label: "店铺品牌",
		},
		{
			Value: UpdateBatchTypeCate,
			Label: "店铺类目",
		},
		{
			Value: UpdateBatchTypeUserProfile,
			Label: "用户资料",
		},
		{
			Value: UpdateBatchTypeAuthType,
			Label: "店铺入驻",
		},
		{
			Value: UpdateBatchTypeAvatar,
			Label: "用户头像",
		},
	}
)

var (
	//用户资料审核状态
	SliceUserApplyStatus = base.ModelItemOptions{
		{
			Label: "使用中...",
			Value: UserApplyStatusUsing,
		},
		{
			Label: "已失效",
			Value: UserApplyStatusInvalid,
		},
		{
			Label: "审核失败",
			Value: UserApplyStatusFailure,
		},
		{
			Label: "编辑中...",
			Value: UserApplyStatusInit,
		},
		{
			Label: "审核中...",
			Value: UserApplyStatusSubmit,
		},
	}
)
