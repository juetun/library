package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	UseTagUnReadMsgCount = "unread_msg"         //未读消息
	UseTagCartNumCount   = "cart_num_msg"       //购物车数量
	UseTagAttendedCount  = "attended_num"       //粉丝数
	UseTagMyAttendCount  = "my_attend"          //我关注数量
	UseTagLoveCount      = "my_love"            //我的点赞数
	UseTagSeeCount       = "my_see"             //我的喜欢(浏览)
	UseCollectCount      = "my_collect"         //我的收藏
	UseCanUseCouponCount = "can_use_coupon_num" //可用优惠券数量
	UseTagScoreCount     = "score_num"          //积分数
	UseTagExportCount    = "export_num"         //导出数据未下载的数量
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
	res = fmt.Sprintf("u:tag:%v", userHid)
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
	var ctxt context.Context
	if len(ctxs) == 0 {
		ctxt = context.TODO()
	} else {
		ctxt = ctxs[0]
	}
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	for _, item := range data {
		if err = cacheClient.HSet(ctxt, getUserTagKeyByUid(item.UserHid), item.TagKey, item.Count).Err(); err != nil {
			return
		}
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
	var ctxt context.Context
	if len(ctxs) == 0 {
		ctxt = context.TODO()
	} else {
		ctxt = ctxs[0]
	}
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	var e error
	if count, e = cacheClient.
		HGet(ctxt, getUserTagKeyByUid(useHid), tagKey).
		Float64(); e != nil && e != redis.Nil {
		err = e
		return
	}
	return
}
