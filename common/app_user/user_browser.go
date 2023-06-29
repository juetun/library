package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
)

type (
	UserBrowser struct {
		HeaderInfo *common.HeaderInfo `json:"-"`
		UserHid    int64              `json:"-"`
		DataType   string             `json:"data_type"`
		DataId     string             `json:"data_id"`
		TimeStamp  base.TimeNormal    `json:"time_stamp"`
	}
)

const (
	UserBrowserCacheNameSpace = "user_browser" //用户浏览计数器缓存的namespace
	BrowserMaxCount           = 300
)

//用户浏览缓存的KEY
func GetUserBrowserCacheKey(userHid int64) (res string) {
	res = fmt.Sprintf("u:bw:%v", userHid)
	return
}
func getCtx(ctxs ...context.Context) (ctx context.Context) {
	if len(ctxs) == 0 {
		ctx = context.TODO()
	} else {
		ctx = ctxs[0]
	}
	return
}

//设置用户的tag标签值
func SetUserBrowser(ctx *base.Context, dataList []*UserBrowser, ctxs ...context.Context) (err error) {
	defer func() {
		if err == nil || ctx == nil {
			return
		}
		ctx.Error(map[string]interface{}{
			"data": dataList,
			"err":  err.Error(),
		}, "SetUseTagCount")
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	if len(dataList) == 0 {
		return
	}

	var (
		l              = len(dataList)
		cacheClient, _ = app_obj.GetRedisClient(UserBrowserCacheNameSpace)
		dataItem       *redis.Z
		item           *UserBrowser
		groupData      = make(map[int64][]*redis.Z, l)
	)

	for _, item = range dataList {
		if _, ok := groupData[item.UserHid]; !ok {
			groupData[item.UserHid] = make([]*redis.Z, 0, )
		}
		dataItem = &redis.Z{
			Score:  float64(item.TimeStamp.UnixNano()),
			Member: item,
		}
		groupData[item.UserHid] = append(groupData[item.UserHid], dataItem)
	}

	var (
		ctxt     = getCtx(ctxs...)
		cacheKey string
	)

	if len(groupData) > 0 {
		for userHid, data := range groupData {
			cacheKey = GetUserBrowserCacheKey(userHid)
			//添加数据
			if err = cacheClient.ZAdd(ctxt, cacheKey, data...).Err(); err != nil {
				return
			}

			//移除有序集合中300名以后的数据
			if err = cacheClient.ZRemRangeByRank(ctxt, cacheKey, BrowserMaxCount, 0).Err(); err != nil {
				return
			}
		}

	}

	return
}
