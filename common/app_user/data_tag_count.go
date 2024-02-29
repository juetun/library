package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
)

//
//
//商品（或其他数据的访问记录数）
const (
	BrowserSpuCount = "spu_bro" //商品浏览数
)

const (
	DataTagCountCacheNameSpace = "data_tag" //用户标签计数器缓存的namespace
)

type (
	DataTagCount struct {
		DataId string  `json:"data_id"` //用户ID
		TagKey string  `json:"tag_key"` //数量对应的key
		Count  float64 `json:"count"`   //数量
	}
)

func getDataTagKeyByUid(dataId string) (res string) {
	res = fmt.Sprintf("tag:data:%v", dataId)
	return
}

//设置用户的tag标签值自增
func UpdateDataTagAddCount(ctx *base.Context, dataId string, key string, value float64, ctxs ...context.Context) (err error) {
	defer func() {
		if err == nil || ctx == nil {
			return
		} else if err != nil {
			ctx.Error(map[string]interface{}{
				"key":   key,
				"value": value,
				"err":   err.Error(),
			}, "AddUseTagCount")
		}

		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	if value == 0 {
		return
	}
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(DataTagCountCacheNameSpace)
	err = cacheClient.HIncrByFloat(ctxt, getDataTagKeyByUid(dataId), key, value).Err()
	return
}

//设置用户的tag标签值
func SetDataTagCount(ctx *base.Context, data []*DataTagCount, ctxs ...context.Context) (err error) {
	defer func() {
		if err == nil || ctx == nil {
			return
		} else if err != nil {
			ctx.Error(map[string]interface{}{
				"data": data,
				"err":  err.Error(),
			}, "SetUseTagCount")
		}

		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	if len(data) == 0 {
		return
	}
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(DataTagCountCacheNameSpace)
	var l = len(data)
	var dataListMap = make(map[string][]interface{}, l)
	for _, item := range data {
		if _, ok := dataListMap[item.DataId]; !ok {
			dataListMap[item.DataId] = make([]interface{}, 0, l*2)
		}
		dataListMap[item.DataId] = append(dataListMap[item.DataId], item.TagKey, item.Count)
	}
	var cacheKey string
	//批量将数据写入redis
	for userHid, items := range dataListMap {
		if len(items) == 0 {
			continue
		}
		cacheKey = getDataTagKeyByUid(userHid)
		_ = cacheClient.HMSet(ctxt, cacheKey, items...).Err()
	}

	return
}

//获取用户的数量
func GetDataTagCount(ctx *base.Context, dataId string, tagKey string, ctxs ...context.Context) (count float64, err error) {
	if dataId == "" || tagKey == "" {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"dataId": dataId,
			"tagKey": tagKey,
			"err":    err.Error(),
		}, "GetUseTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(DataTagCountCacheNameSpace)
	var e error
	var cacheKey = getDataTagKeyByUid(dataId)
	if count, e = cacheClient.
		HGet(ctxt, cacheKey, tagKey).
		Float64(); e != nil && e != redis.Nil {
		err = e
		return
	}
	return
}

//获取多个店铺的多个标签
func GetDataTagsCount(ctx *base.Context, dataIds []string, tagKeys []string, ctxs ...context.Context) (resUserIdValue map[string]map[string]float64, err error) {
	lUserHId := len(dataIds)
	resUserIdValue = make(map[string]map[string]float64, lUserHId)
	if lUserHId == 0 || len(tagKeys) == 0 {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"dataIds": dataIds,
			"tagKeys": tagKeys,
			"err":     err.Error(),
		}, "GetDataTagsCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(DataTagCountCacheNameSpace)
	var cacheKey string
	for _, dataId := range dataIds {
		var e error
		var result []interface{}
		cacheKey = getDataTagKeyByUid(dataId)
		if result, e = cacheClient.
			HMGet(ctxt, cacheKey, tagKeys...).Result(); e != nil && e != redis.Nil {
			err = e
			return
		}
		resUserIdValue[dataId] = getDataTagKeysValue(result, tagKeys)
	}

	return
}

func getDataTagKeysValue(result []interface{}, tagKeys []string) (res map[string]float64) {
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
