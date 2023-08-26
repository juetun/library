package mall_comment

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"time"
)

var (
	CacheParamConfig = map[string]*redis_pkg.CacheProperty{
		"CacheSpuComment": { //店铺变跟资质缓存的缓存Key
			Key:    "m:spu:comment:%v",
			Expire: 12 * time.Hour,
		},
	}
)

func GetCacheParamConfig(key string) (res *redis_pkg.CacheProperty, err error) {
	var ok bool
	if res, ok = CacheParamConfig[key]; ok {
		return
	}
	err = fmt.Errorf("您当前未配置缓存信息(%v)", key)
	return
}
