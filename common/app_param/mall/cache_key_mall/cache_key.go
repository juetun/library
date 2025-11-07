package cache_key_mall

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"time"
)

var (
	MallCacheParamConfig = map[string]*redis_pkg.CacheProperty{
		"CacheKeySpuList": { //在线商品有序集合
			Key:           "m:online:spu",
			Expire:        24 * time.Hour,
			CacheDataType: redis_pkg.CacheDataTypeHashSortSet, //有序集合
		},
		"CacheKeyShopSaleTop": { //店铺商品销量排行
			Key:           "m:sale:shop_%v",
			Expire:        24 * time.Hour,
			CacheDataType: redis_pkg.CacheDataTypeHashSortSet, //有序集合
		},
		"CacheKeySpuSaleTop": { //商品总销量排行
			Key:           "m:sale:spu",
			Expire:        24 * time.Hour,
			CacheDataType: redis_pkg.CacheDataTypeHashSortSet, //有序集合
		},
	}
)

func GetCacheParamConfig(key string) (res *redis_pkg.CacheProperty, err error) {
	var ok bool
	if res, ok = MallCacheParamConfig[key]; ok {
		return
	}
	err = fmt.Errorf("您当前未配置缓存信息(%v)", key)
	return
}
