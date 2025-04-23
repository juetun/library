package cache_key_mall

import (
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"time"
)

var (
	MallCacheParamConfig = map[string]*redis_pkg.CacheProperty{
		"CacheKeySpuList": { //在线商品有序集合
			Key:    "m:spu_set",
			Expire: 24 * time.Hour,
		},
		"CacheKeyShopSaleTop": { //店铺商品销量排行
			Key:    "m:shop_sale:%v",
			Expire: 24 * time.Hour,
		},
		"CacheKeySpuSaleTop": { //商品总销量排行
			Key:    "m:spu_sale:%v",
			Expire: 24 * time.Hour,
		},
	}
)
