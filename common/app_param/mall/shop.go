package mall

import "github.com/juetun/library/common/app_param/mall/models"

const (
	ShopExt        = "shop_ext"      //店铺信息
	ShopImgIcon    = "shop_img_icon" //店铺logo
	ShopImgbgImage = "shop_bg_img"   //店铺背景图
	ShopNotice     = "shop_notice"   //店铺公告
)
const ManagerGroupIdAdmin int64 = - 1 //托管到客服后台的聊天信息
type (
	ShopData struct {
		Shop       models.Shop        `json:"shop"`        // 店铺基本信息
		ShopExt    *models.ShopExt    `json:"shop_ext"`    // 店铺扩展信息
		ShopNotice *models.ShopNotice `json:"shop_notice"` // 店铺公告
	}
	MapShopSingle struct {
		MapShopIdExt    map[int64]*models.ShopExt
		MapShopIdNotice map[int64]*models.ShopNotice
	}
	BrandData struct {
		Brand models.Brand `json:"brand"` // 店铺基本信息
	}
	MapBrandSingle struct {
	}
)

func (r *ShopData) SetShopExt(ext *models.ShopExt) {
	r.ShopExt = ext
}

func (r *ShopData) SetShopNotice(notice *models.ShopNotice) {
	r.ShopNotice = notice
}
