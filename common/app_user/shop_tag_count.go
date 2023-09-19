package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	ShopTagUnReadChatMsgCount = "unread_chat"  //未读私信
	ShopTagAttendNum          = "attend"       //关注店铺的人数
	ShopTagLevel              = "score"        //店铺综合评分
	ShopTagComment            = "comment"      //店铺商品评分
	ShopTagFreight            = "ems_comment"  //物流评价评分
	ShopTagSaleComment        = "sale_comment" //售后评价评分
)

const (
	ShopTagCountCacheNameSpace = "shop_tag" //用户标签计数器缓存的namespace
)

type (
	ShopTagCount struct {
		ShopId int64   `json:"shop_id"` //用户ID
		TagKey string  `json:"tag_key"` //数量对应的key
		Count  float64 `json:"count"`   //数量
	}
)

func getShopTagKeyByUid(userHid int64) (res string) {
	res = fmt.Sprintf("shop:tag:%v", userHid)
	return
}

//设置用户的tag标签值
func SetShopTagCount(ctx *base.Context, data []*ShopTagCount, ctxs ...context.Context) (err error) {
	otherData := make(map[string]interface{})
	defer func() {

		if err == nil || ctx == nil {
			return
		} else if err != nil {
			ctx.Error(map[string]interface{}{
				"data":      data,
				"otherData": otherData,
				"err":       err.Error(),
			}, "SetShopTagCount")
		}
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()

 	if len(data) == 0 {
		return
	}
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(ShopTagCountCacheNameSpace)

	var l = len(data)
	var dataListMap = make(map[int64][]interface{}, l)
	for _, item := range data {
		if _, ok := dataListMap[item.ShopId]; !ok {
			dataListMap[item.ShopId] = make([]interface{}, 0, l*2)
		}
		dataListMap[item.ShopId] = append(dataListMap[item.ShopId], item.TagKey, item.Count)
	}
	var cacheKey string
	var e error
	//批量将数据写入redis
	for shopId, items := range dataListMap {
		if len(items) == 0 {
			continue
		}
		cacheKey = getShopTagKeyByUid(shopId)
		if e = cacheClient.HMSet(ctxt, cacheKey, items...).Err(); e != nil {
			otherData[fmt.Sprintf("%v", shopId)] = e.Error()
		}
	}
	return
}

func getCtxWithMany(ctxs ...context.Context) (ctxt context.Context) {
	if len(ctxs) == 0 {
		ctxt = context.TODO()
	} else {
		ctxt = ctxs[0]
	}
	return
}

//获取用户的数量
func GetShopTagCount(ctx *base.Context, shopId int64, tagKey string, ctxs ...context.Context) (count float64, err error) {
	if shopId == 0 || tagKey == "" {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"shopId": shopId,
			"tagKey": tagKey,
			"err":    err.Error(),
		}, "GetShopTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(ShopTagCountCacheNameSpace)
	var e error
	if count, e = cacheClient.
		HGet(ctxt, getShopTagKeyByUid(shopId), tagKey).
		Float64(); e != nil && e != redis.Nil {
		err = e
		return
	}
	return
}

//获取多个店铺的多个标签
func GetShopsTagsCount(ctx *base.Context, shopIds []int64, tagKeys []string, ctxs ...context.Context) (resShopIdValue map[int64]map[string]float64, err error) {
	lShop := len(shopIds)
	resShopIdValue = make(map[int64]map[string]float64, lShop)
	if lShop == 0 || len(tagKeys) == 0 {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"shopIds": shopIds,
			"tagKeys": tagKeys,
			"err":     err.Error(),
		}, "GetShopTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(ShopTagCountCacheNameSpace)

	for _, shopId := range shopIds {
		var e error
		var result []interface{}
		if result, e = cacheClient.
			HMGet(ctxt, getShopTagKeyByUid(shopId), tagKeys...).Result(); e != nil && e != redis.Nil {
			err = e
			return
		}
		resShopIdValue[shopId] = getShopTagKeysValue(result, tagKeys)
	}

	return
}

func getShopTagKeysValue(result []interface{}, tagKeys []string) (res map[string]float64) {
	for k, item := range result {
		if item != nil {
			switch item.(type) {
			case int64:
				res[tagKeys[k]] = float64(item.(int64))
			case float64:
				res[tagKeys[k]] = item.(float64)
			}
		}
	}
	return
}
