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
	PlatMallManager  = "plat_mall_manager"  //电商管理
	PlatBadgeURefund = "plat_badge_urefund" //电商管理-退款单
	PlatOrderBadge   = "plat_order_badge"   //电商管理-订单管理
	PlatBadgeUOrder  = "plat_badge_uorder"  //电商管理-订单管理-用户订单

	PlatUserBadge        = "plat_user_badge"     //用户信息
	PlatUserManagerBadge = "plat_umanager_badge" //用户信息 -用户管理
	PlatWaitingReview    = "plat_waiting_review" //用户信息 -用户管理 - 资料审核

	PlatMarketing  = "plat_marketing"   //营销管理
	PlatSuggestion = "plat_suggestion"  //营销管理-投诉与建议
	PlatDataStatic = "plat_data_static" //营销管理-统计

	PlatSNSApprove        = "plat_sns_approve"         //社交与钓点
	PlatArticleApprove    = "plat_article_approve"     //社交与钓点- 圈子动态
	PlatArticleFishSports = "plat_article_fish_sports" //社交与钓点- 钓点信息
)

const (
	ShopTagRefundAndSaleAfterCount = "shop_refund_sale_after" //退款和售后
	ShopTagOrderListCount          = "shop_order_list"        //订单列表

	ShopTagUnReadChatCount       = "shop_chat_msg"         //未读聊天信息
	ShopTagOrderWaitingSend      = "shop_send_goods"       //待发货订单数
	ShopTagOrderWaitingSure      = "shop_refund_sure"      //退款确认
	ShopTagOrderWaitingComplaint = "shop_refund_complaint" //订单申诉
)

//
//
//user_main
const (
	UserTagUnReadMsgCount    = "user_message"    //未读消息(当前包括聊天未读信息和通知信息)
	UserTagUnReadChatCount   = "user_chat_msg"   //未读聊天消息
	UserTagUnReadNoticeCount = "user_notice_msg" //未读通知信息
	UserTagCartNumCount      = "user_cart_num"   //购物车数量
	UserTagAttendedCount     = "user_attend_me"  //粉丝数
	UserTagMyAttendCount     = "user_my_attend"  //我关注数量
	UserTagLoveCount         = "user_love"       //我的点赞数
	UserTagSeeCount          = "user_browser"    //我的喜欢(浏览)
	UserCollectCount         = "user_collect"    //我的收藏
	UserCanUseCouponCount    = "user_coupon_num" //可用优惠券数量
	UserTagScoreCount        = "user_score"      //积分数
	UserTagExportCount       = "user_export"     //导出数据未下载的数量
	UserTagHomeFlag          = "user_home"       //用户主页（我的）是否点显示 0-不显示 1-显示
	UserTagMainFlag          = "user_main"       //首页(打开APP第一个页面)是否点显示 0-不显示 1-显示

	UserTagOrderBadge = "user_order_badge" //当前未处理的订单数（网站个人中心左侧菜单 “我的订单” 上使用）

	UserTagOrderNotPay          = "user_waiting_pay"        //未付款订单数
	UserTagOrderWaitingSendGood = "user_receive_send_goods" //待发货订单数
	UserTagOrderWaitingComment  = "user_order_comment"      //待评价
	UserTagOrderNotOverRefund   = "user_not_over_refund"    //未完成退款单
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

//设置用户的tag标签值自增
func UpdateUseTagAddCount(ctx *base.Context, userHid int64, key string, value float64, ctxs ...context.Context) (err error) {
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
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	err = cacheClient.HIncrByFloat(ctxt, getUserTagKeyByUid(userHid), key, value).Err()
	return
}

//设置用户的tag标签值
func SetUseTagCount(ctx *base.Context, data []*UserTagCount, ctxs ...context.Context) (err error) {
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
	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
	var l = len(data)
	var dataListMap = make(map[int64][]interface{}, l)
	for _, item := range data {
		if _, ok := dataListMap[item.UserHid]; !ok {
			dataListMap[item.UserHid] = make([]interface{}, 0, l*2)
		}
		dataListMap[item.UserHid] = append(dataListMap[item.UserHid], item.TagKey, item.Count)
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

	cacheClient, _ := app_obj.GetRedisClient(UserTagCountCacheNameSpace)
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
