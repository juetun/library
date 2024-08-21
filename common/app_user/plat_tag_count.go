package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
)

const (
	PlatTagCountCacheNameSpace = "data_tag" //用户标签计数器缓存的namespace
)

type (
	PlatTagCount struct {
		TagKey string  `json:"tag_key"` //数量对应的key
		Count  float64 `json:"count"`   //数量
	}
)

func getPlatTagKey() (res string) {
	res = fmt.Sprintf("tag:plat")
	return
}

//设置用户的tag标签值
func SetPlatTagCount(ctx *base.Context, data []*PlatTagCount, ctxs ...context.Context) (err error) {
	otherData := make(map[string]interface{})
	defer func() {

		if err == nil || ctx == nil {
			return
		} else if err != nil {
			ctx.Error(map[string]interface{}{
				"data":      data,
				"otherData": otherData,
				"err":       err.Error(),
			}, "SetPlatTagCount")
		}
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()

	if len(data) == 0 {
		return
	}
	var ctxt = GetCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(PlatTagCountCacheNameSpace)

	var l = len(data)
	var items = make([]interface{}, 0, l)
	for _, item := range data {
		items = append(items, item.TagKey, item.Count)
	}
	var cacheKey string
	var e error
	//批量将数据写入redis
	if len(items) == 0 {
		return
	}
	cacheKey = getPlatTagKey()
	if e = cacheClient.HMSet(ctxt, cacheKey, items...).Err(); e != nil {
		otherData[fmt.Sprintf("%v", cacheKey)] = e.Error()
	}
	return
}

//获取用户的数量
func GetPlatTagCount(ctx *base.Context, tagKey string, ctxs ...context.Context) (count float64, err error) {
	if tagKey == "" {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"tagKey": tagKey,
			"err":    err.Error(),
		}, "GetPlatTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = GetCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(PlatTagCountCacheNameSpace)
	var e error
	if count, e = cacheClient.
		HGet(ctxt, getPlatTagKey(), tagKey).
		Float64(); e != nil && e != redis.Nil {
		err = e
		return
	}
	return
}

//获取多个店铺的多个标签
func GetPlatTagsCount(ctx *base.Context, tagKeys []string, ctxs ...context.Context) (resShopIdValue map[string]float64, err error) {
	l := len(tagKeys)
	resShopIdValue = make(map[string]float64, l)
	if l == 0 || len(tagKeys) == 0 {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"tagKeys": tagKeys,
			"err":     err.Error(),
		}, "GetPlatTagsCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = GetCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(PlatTagCountCacheNameSpace)

	var e error
	var result []interface{}
	if result, e = cacheClient.
		HMGet(ctxt, getPlatTagKey(), tagKeys...).Result(); e != nil && e != redis.Nil {
		err = e
		return
	}
	resShopIdValue = getPlatTagKeysValue(result, tagKeys)
	return
}

func getPlatTagKeysValue(result []interface{}, tagKeys []string) (res map[string]float64) {
	res = make(map[string]float64, len(result))
	for k, item := range result {
		if item != nil {
			switch item.(type) {
			case int64:
				res[tagKeys[k]] = float64(item.(int64))
			case float64:
				res[tagKeys[k]] = item.(float64)
			case string:
				res[tagKeys[k]], _ = strconv.ParseFloat(item.(string), 64)
			}
		}
	}
	return
}
