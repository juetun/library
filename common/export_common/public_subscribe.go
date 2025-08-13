package export_common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	RedisMqTopicTmp = "export:mq:%v" //
)

const (
	RedisExportNameSpace = "export"
)

//获取聊天系统监听的队列参数,分布式系统使用
func GetRedisMqTopicTmp(topicNames ...string) (res string) {
	var topicName string

	//TODO 当前共用一个，后续可在此实现分桶逻辑
	//if len(topicNames) > 0 {
	//	topicName = topicNames[0]
	//}
	res = fmt.Sprintf(RedisMqTopicTmp, topicName)
	return
}

type DaoExportCommon struct {
	base.ServiceDao
	Ctx context.Context
}

func (r *DaoExportCommon) Init(ctx ...*base.Context) {
	r.SetContext(ctx...)
	if r.Ctx == nil {
		r.Ctx = context.TODO()
	}
	return
}

//各个后台的Websocket链接 缓存队列的连接客户端
func (r *DaoExportCommon) GetExportCacheClient() (cacheClient *redis.Client) {
	cacheClient, _ = app_obj.GetRedisClient(RedisExportNameSpace)
	return
}
