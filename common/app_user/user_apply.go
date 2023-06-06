package app_user

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/mall/models"
	"github.com/juetun/library/common/const_apply"
)

const (
	UpdateBatchTypeInfo        = "shop_info"    //店铺信息
	UpdateBatchTypeBrand       = "shop_brand"   //店铺品牌
	UpdateBatchTypeCompany     = "company"      //企业信息
	UpdateBatchTypeManager     = "manager"      //代理人信息
	UpdateBatchTypeCate        = "shop_cate"    //店铺类目
	UpdateBatchTypeAuthType    = "auth_type"    //店铺入驻审核
	UpdateBatchTypeAvatar      = "avatar"       //用户头像审核
	UpdateBatchTypeUserProfile = "user_profile" //用户资料
)
const (
	//审核状态用店铺审核状态值同步
	UserApplyStatusUsing   = models.ShopStatusOk       //使用中
	UserApplyStatusInvalid = models.ShopStatusInvalid  //失效
	UserApplyStatusFailure = models.ShopStatusFailure  //审核失败
	UserApplyStatusInit    = models.ShopStatusInit     //编辑中
	UserApplyStatusSubmit  = models.ShopStatusAuditing //提交审核

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
		{
			Value: UpdateBatchTypeCompany,
			Label: "企业信息",
		},
		{
			Value: UpdateBatchTypeManager,
			Label: "代理人信息",
		},
	}
)

var (
	//用户资料审核状态
	SliceUserApplyStatus = base.ModelItemOptions{
		{
			Label: "审核中...",
			Value: UserApplyStatusSubmit,
		},
		{
			Label: "审核失败",
			Value: UserApplyStatusFailure,
		},
		{
			Label: "已失效",
			Value: UserApplyStatusInvalid,
		},
		{
			Label: "审核通过",
			Value: UserApplyStatusUsing,
		},
		{
			Label: "初始化...",
			Value: UserApplyStatusInit,
		},
		{
			Label: "编辑中...",
			Value: const_apply.ApplyStatusTimeEditing,
		},
	}
)
