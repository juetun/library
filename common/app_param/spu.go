package app_param

import "github.com/juetun/base-wrapper/lib/base"

const (
	SpuDataBase                = "spu_data"         //商品信息
	SpuDataShop                = "spu_shop"         // 店铺信息
	SpuDataShopExt             = "spu_shop_ext"     // 店铺扩展信息
	SpuDataShopNotice          = "spu_shop_notice"  // 店铺公告信息
	SpuDataBrand               = "spu_brand"        // 品牌信息
	SpuDataTypeProductDesc     = "spu_product_desc" // 商品描述
	SpuDataTypeSKus            = "spu_sku"          // SKU信息
	SpuDataTypeSKusGift        = "spu_gifts"        // 赠品信息
	SpuDataTypeSKusStock       = "spu_sku_stock"    // sku库存
	SpuDataTypeSKusRelate      = "spu_sku_relate"   // sku关联属性
	SpuDataTypeSKusProperty    = "spu_sku_property" // sku属性
	SpuDataTypeFreightTemplate = "spu_freight"      // 运费模板
	SpuDataTypeComment         = "comment"          // 评论信息（推荐或最近评论的信息）
	/*****************************上传信息获取**********************************/
	//SpuUploadAll          = "spu_img"         // 商品图片获取所有图片 (包括SpuImg前缀的属性如:SpuUploadImgThumbnail,SpuUploadImgPic,SpuSkuImg,SpuUploadDescription）
	SpuUploadImgThumbnail = "spu_img_thumb"   // 商品缩略图
	SpuUploadImgPic       = "spu_img_li"      // 商品图片
	SpuUploadSkuImg       = "spu_sku_img"     // SKU图片
	SpuUploadBrandImg     = "spu_brand_img"   // 品牌logo
	SpuUploadDescription  = "spu_upload_desc" // 商品详情的信息
	SpuUploadVideo        = "spu_video"       // 上传视频视频地址
	SpuShopLogo           = "spu_shop_logo"   // 店铺LOGO
	SpuShopNoticeImg      = "spu_notice_img"  // 店铺公告图片
)

