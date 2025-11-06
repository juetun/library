package app_param

import (
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"time"
)

const (
	QueueConsumes       = "QueueConsumesInfo"
	QueueCacheNameSpace = "library"
)

var (
	CacheParamConfig = map[string]*redis_pkg.CacheProperty{
		QueueConsumes: {Key: "p:MsgConsumeData:%v", Expire: 7 * 24 * time.Hour,}, //消息队列消费者缓存
	}
)

