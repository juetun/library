package app_user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	UserBrowser struct {
		UserHid   int64           `json:"-"`
		DataType  string          `json:"data_type"`
		DataId    string          `json:"data_id"`
		TimeStamp base.TimeNormal `json:"time_stamp"`
	}
)

const (
	UserBrowserCacheNameSpace = "user_browser" //用户浏览计数器缓存的namespace
)

//用户浏览缓存的KEY
func GetUserBrowser(userHid int64) (res string) {
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
func SetUserBrowser(ctx *base.Context, userHid int64, dataList []*UserBrowser, ctxs ...context.Context) (err error) {
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
		cacheClient, _ = app_obj.GetRedisClient(UserBrowserCacheNameSpace)
		dataItem       *redis.Z
		data           = make([]*redis.Z, 0, len(dataList))
		item           *UserBrowser
	)

	for _, item = range dataList {
		dataItem = &redis.Z{
			Score:  float64(item.TimeStamp.UnixNano()),
			Member: item,
		}
		data = append(data, dataItem)
	}

	var (
		ctxt = getCtx(ctxs...)
		key  = GetUserBrowser(userHid)
	)

	if len(data) > 0 {
		//添加数据
		if err = cacheClient.ZAdd(ctxt, key, data...).Err(); err != nil {
			return
		}
		//移除有序集合中300名以后的数据
		if err = cacheClient.ZRemRangeByRank(ctxt, key, 300, 0).Err(); err != nil {
			return
		}
	}

	return
}
