package cache_act_comment

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/base/cache_act"
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"github.com/juetun/library/common/app_param/mall/cache_key_mall"
	"github.com/juetun/library/common/app_param/mall_comment"
	"time"
)

func GetSpuCacheKey(productId interface{}, expireTimeRands ...bool) (res string, timeExpire time.Duration, err error) {
	var CacheKeyProductDescWithProductId *redis_pkg.CacheProperty
	if CacheKeyProductDescWithProductId, err = cache_key_mall.GetCacheParamConfig("CacheSpuComment"); err != nil {
		return
	}

	res = fmt.Sprintf(CacheKeyProductDescWithProductId.Key, productId)
	timeExpire = CacheKeyProductDescWithProductId.Expire
	var expireTimeRand bool
	if len(expireTimeRands) > 0 {
		expireTimeRand = expireTimeRands[0]
	}
	if expireTimeRand {
		randNumber, _ := base.RealRandomNumber(60) //打乱缓存时长，防止缓存同一时间过期导致数据库访问压力变大
		timeExpire = timeExpire + time.Duration(randNumber)*time.Second
	}
	return
}

type (
	CacheSpuCommentAction struct {
		cache_act.CacheActionBase
		arg            *base.ArgGetByStringIds
		GetByIdsFromDb GetSpuDescByIdsFromDb
	}
	GetSpuDescByIdsFromDb       func(id ...string) (resData map[string]*mall_comment.OrderComment, err error)
	CacheSpuCommentActionOption func(cacheFreightAction *CacheSpuCommentAction)
)

func CacheSpuCommentActionArg(arg *base.ArgGetByStringIds) CacheSpuCommentActionOption {
	return func(cacheFreightAction *CacheSpuCommentAction) {
		cacheFreightAction.arg = arg
		return
	}
}

func CacheSpuCommentActionGetByIdsFromDb(getByIdsFromDb GetSpuDescByIdsFromDb) CacheSpuCommentActionOption {
	return func(cacheFreightAction *CacheSpuCommentAction) {
		cacheFreightAction.GetByIdsFromDb = getByIdsFromDb
		return
	}
}

func NewCacheSpuCommentAction(options ...CacheSpuCommentActionOption) (res *CacheSpuCommentAction) {
	res = &CacheSpuCommentAction{CacheActionBase: cache_act.NewCacheActionBase()}
	for _, handler := range options {
		handler(res)
	}
	if res.Ctx == nil {
		res.Ctx = context.TODO()
	}
	return
}

func (r *CacheSpuCommentAction) LoadBaseOption(options ...cache_act.CacheActionBaseOption) *CacheSpuCommentAction {
	r.LoadBase(options...)
	return r
}

func (r *CacheSpuCommentAction) Action() (res map[string]*mall_comment.OrderComment, err error) {
	res = map[string]*mall_comment.OrderComment{}
	if len(r.arg.Ids) == 0 {
		return
	}
	// 初始化 type默认值
	if err = r.arg.Default(); err != nil {
		return
	}
	switch r.arg.GetType {
	case base.GetDataTypeFromDb: // 从数据库获取数据
		if res, err = r.GetByIdsFromDb(r.arg.Ids...); err != nil {
			return
		}
		switch r.arg.RefreshCache {
		case base.RefreshCacheYes:
			_ = r.RemoveCacheByStringId(r.arg.Ids...)
			for id, value := range res {
				if err = r.SetToCache(id, value); err != nil {
					return
				}
			}
		}
	case base.GetDataTypeFromCache: // 从缓存获取数据
		res, _, err = r.getByIdsFromCache(r.arg.Ids...)
	case base.GetDataTypeFromAll: // 优先从缓存获取，如果没有数据，则从数据库获取
		res, err = r.getByIdsFromAll(r.arg.Ids...)
	default:
		err = fmt.Errorf("当前不支持你选择的获取数据类型(%s)", r.arg.GetType)
	}
	return
}

func (r *CacheSpuCommentAction) getFromCache(id interface{}) (data *mall_comment.OrderComment, err error) {
	data = &mall_comment.OrderComment{}
	defer func() {
		if err != nil && err != redis.Nil {
			r.Context.Info(map[string]interface{}{
				"productId": id,
				"err":       err.Error(),
			}, "CacheSpuCommentActionGetFromCache")
			err = base.NewErrorRuntime(err, base.ErrorRedisCode)
			return
		}
	}()
	var key string
	if key, _, err = r.GetCacheKey(id); err != nil {
		return
	}
	cmd := r.Context.CacheClient.Get(r.Ctx, key)
	if err = cmd.Err(); err != nil {
		return
	}
	if err = cmd.Scan(data); err != nil {
		return
	}

	return
}

func (r *CacheSpuCommentAction) getByIdsFromCache(ids ...string) (res map[string]*mall_comment.OrderComment, noCacheIds []string, err error) {
	var e error

	res = map[string]*mall_comment.OrderComment{}

	//收集缓存中没有的数据ID，便于后边查询使用
	noCacheIds = make([]string, 0, len(ids))

	for _, id := range ids {
		if res[id], e = r.getFromCache(id); e != nil {
			if e != redis.Nil {
				err = e
				return
			}
			noCacheIds = append(noCacheIds, id)
			res[id] = mall_comment.NewOrderComment()
		}
	}
	return
}

func (r *CacheSpuCommentAction) getByIdsFromAll(ids ...string) (res map[string]*mall_comment.OrderComment, err error) {
	var pIds []string
	if res, pIds, err = r.getByIdsFromCache(ids...); err != nil {
		return
	}
	if len(pIds) == 0 {
		return
	}

	var dt map[string]*mall_comment.OrderComment
	if dt, err = r.GetByIdsFromDb(pIds...); err != nil {
		return
	}

	for _, pid := range pIds {
		if dta, ok := dt[pid]; ok {
			res[pid] = dta
			_ = r.SetToCache(pid, dta)
			continue
		}
		_ = r.SetToCache(pid, &mall_comment.OrderComment{})
	}

	return
}
