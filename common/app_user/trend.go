package app_user

import "github.com/juetun/base-wrapper/lib/base"

const (
	TrendTypeShopInfo     = "shop_info"     //店铺信息
	TrendTypeShopBrand    = "shop_brand"    //店铺品牌
	TrendTypeShopCate     = "shop_cate"     //店铺类目
	TrendTypeSpuEdit      = "product_edit"  //编辑商品
	TrendTypeSns          = "sns"           //圈子动态
	TrendTypeFishingSport = "fishing_spots" //钓点信息
	TrendTypeRegister     = "register"      //用户注册
	TrendTypeUsrApply     = "usr_apply"     //用户审核
)

var (
	SliceTrendType = base.ModelItemOptions{
		{
			Value: TrendTypeShopInfo,
			Label: "店铺信息变更",
		},
		{
			Value: TrendTypeShopBrand,
			Label: "店铺品牌变更",
		},
		{
			Value: TrendTypeShopCate,
			Label: "店铺类目变更",
		},
		{
			Value: TrendTypeSpuEdit,
			Label: "编辑商品",
		},
		{
			Value: TrendTypeSns,
			Label: "圈子动态",
		},
		{
			Value: TrendTypeFishingSport,
			Label: "钓点信息",
		},
		{
			Value: TrendTypeRegister,
			Label: "注册账号",
		},
		{
			Value: TrendTypeUsrApply,
			Label: "账号认证审核",
		},
	}
)
