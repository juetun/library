package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	UserTagUnReadMsgCount = "unread_msg"     //未读消息
	UserTagCartNumCount   = "cart_num"       //购物车数量
	UserTagAttendedCount  = "att_num"        //粉丝数
	UserTagMyAttendCount  = "att"            //我关注数量
	UserTagLoveCount      = "love"           //我的点赞数
	UserTagSeeCount       = "see"            //我的喜欢(浏览)
	UserCollectCount      = "collect"        //我的收藏
	UserCanUseCouponCount = "can_use_coupon" //可用优惠券数量
	UserTagScoreCount     = "score"          //积分数
	UserTagExportCount    = "export"         //导出数据未下载的数量
	UserTagHomeFlag       = "uhome"          //用户主页是否点显示 0-不显示 1-显示

	UserTagOrderNotPay          = "not_pay"      //未付款订单数
	UserTagOrderWaitingSendGood = "wait_send"    //待发货订单数
	UserTagOrderWaitingComment  = "wait_comment" //待评价
	UserTagOrderRefund          = "not_over_rf"  //未完成退款单
)

const (
	UserTagCountCacheNameSpace = "user_tag" //用户标签计数器缓存的namespace
)

type (
	UserTagCount struct {
		UserHid int64  `json:"user_hid"` //用户ID
		TagKey  string `json:"tag_key"`  //数量对应的key
		Count   int64  `json:"count"`    //数量
	}
)

func getUserTagKeyByUid(userHid int64) (res string) {
	res = fmt.Sprintf("usr:tag:%v", userHid)
	return
}

//设置用户的tag标签值
func SetUseTagCount(ctx *base.Context, data []*UserTagCount, ctxs ...context.Context) (err error) {
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"data": data,
			"err":  err.Error(),
		}, "SetUseTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	if len(data) == 0 {
		return
	}
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	var l = len(data)
	var dataListMap = make(map[int64][]interface{}, l)
	for _, item := range data {
		if _, ok := dataListMap[item.UserHid]; !ok {
			dataListMap[item.UserHid] = make([]interface{}, 0, l*2)
		}
		dataListMap[item.UserHid] = append(dataListMap[item.UserHid], item.Count)
	}
	var cacheKey string
	//批量将数据写入redis
	for userHid, items := range dataListMap {
		if len(items) == 0 {
			continue
		}
		cacheKey = getUserTagKeyByUid(userHid)
		_ = cacheClient.HMSet(ctxt, cacheKey, items...).Err()
	}

	return
}

//获取用户的数量
func GetUseTagCount(ctx *base.Context, useHid int64, tagKey string, ctxs ...context.Context) (count float64, err error) {
	if useHid == 0 || tagKey == "" {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"useHid": useHid,
			"tagKey": tagKey,
			"err":    err.Error(),
		}, "GetUseTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	var e error
	var cacheKey = getUserTagKeyByUid(useHid)
	if count, e = cacheClient.
		HGet(ctxt, cacheKey, tagKey).
		Float64(); e != nil && e != redis.Nil {
		err = e
		return
	}
	return
}

//获取多个店铺的多个标签
func GetUsersTagsCount(ctx *base.Context, userHIds []int64, tagKeys []string, ctxs ...context.Context) (resUserIdValue map[int64]map[string]float64, err error) {
	lUserHId := len(userHIds)
	resUserIdValue = make(map[int64]map[string]float64, lUserHId)
	if lUserHId == 0 || len(tagKeys) == 0 {
		return
	}
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"userHIds": userHIds,
			"tagKeys":  tagKeys,
			"err":      err.Error(),
		}, "GetUsersTagsCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	var ctxt = getCtxWithMany(ctxs...)

	cacheClient, _ := app_obj.GetRedisClient(ShopTagCountCacheNameSpace)
	var cacheKey string
	for _, userHId := range userHIds {
		var e error
		var result []interface{}
		cacheKey = getUserTagKeyByUid(userHId)
		if result, e = cacheClient.
			HMGet(ctxt, cacheKey, tagKeys...).Result(); e != nil && e != redis.Nil {
			err = e
			return
		}
		resUserIdValue[userHId] = getUserTagKeysValue(result, tagKeys)
	}

	return
}

func getUserTagKeysValue(result []interface{}, tagKeys []string) (res map[string]float64) {
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
