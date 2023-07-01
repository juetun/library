package upload_operate

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/base/cache_act"
	"time"
)

type (
	CacheProductPicAndVideoAction struct {
		cache_act.CacheActionBase
		arg                      *ArgUploadGetInfo
		argCommon                *base.GetDataTypeCommon
		GetByIdsFromDb           GetProductPicAndVideoByIdsFromDb
		HandlerGetUploadCacheKey HandlerGetUploadCacheKey
	}
	GetProductPicAndVideoByIdsFromDb    func(arg *ArgUploadGetInfo) (resData *ResultMapUploadInfo, err error)
	CacheProductPicAndVideoActionOption func(cacheFreightAction *CacheProductPicAndVideoAction)
	HandlerGetUploadCacheKey            func(id interface{}, Type string, expireTimeRands ...bool) (res string, timeExpire time.Duration)
)

func CacheProductPicAndVideoActionArg(arg *ArgUploadGetInfo, argCommon *base.GetDataTypeCommon) CacheProductPicAndVideoActionOption {
	return func(cacheFreightAction *CacheProductPicAndVideoAction) {
		cacheFreightAction.arg = arg
		cacheFreightAction.argCommon = argCommon
		return
	}
}

func CacheHandlerGetUploadCacheKey(arg HandlerGetUploadCacheKey) CacheProductPicAndVideoActionOption {
	return func(cacheFreightAction *CacheProductPicAndVideoAction) {
		cacheFreightAction.HandlerGetUploadCacheKey = arg
		return
	}
}

func CacheProductPicAndVideoActionGetByIdsFromDb(getByIdsFromDb GetProductPicAndVideoByIdsFromDb) CacheProductPicAndVideoActionOption {
	return func(cacheFreightAction *CacheProductPicAndVideoAction) {
		cacheFreightAction.GetByIdsFromDb = getByIdsFromDb
		return
	}
}

func NewCacheProductPicAndVideoAction(options ...CacheProductPicAndVideoActionOption) (res *CacheProductPicAndVideoAction) {
	res = &CacheProductPicAndVideoAction{CacheActionBase: cache_act.NewCacheActionBase()}
	for _, handler := range options {
		handler(res)
	}
	if res.Ctx == nil {
		res.Ctx = context.TODO()
	}
	return
}

func (r *CacheProductPicAndVideoAction) LoadBaseOption(options ...cache_act.CacheActionBaseOption) *CacheProductPicAndVideoAction {
	r.LoadBase(options...)
	return r
}

func (r *CacheProductPicAndVideoAction) SetToCacheNew(key string, duration time.Duration, data interface{}, expireTimeRand ...bool) (err error) {
	if err = r.Context.CacheClient.Set(r.Ctx, key, data, duration).Err(); err != nil {
		r.Context.Info(map[string]interface{}{
			"data":           data,
			"key":            key,
			"duration":       duration,
			"expireTimeRand": expireTimeRand,
		}, "CacheActionSetToCache")
		return
	}
	return
}

func (r *CacheProductPicAndVideoAction) saveCache(res *ResultMapUploadInfo) (err error) {
	var (
		key      string
		duration time.Duration
	)
	if len(res.Music) > 0 {
		for id, value := range res.Music {
			key, duration = r.HandlerGetUploadCacheKey(id, FileTypeMusic)
			if err = r.SetToCacheNew(key, duration, value); err != nil {
				return
			}
		}
	}
	//if len(res.Img) > 0 {
	//
	//	for id, value := range res.Img {
	//		key, duration = r.HandlerGetUploadCacheKey(id, FileTypePicture)
	//		if err = r.SetToCacheNew(key, duration, value); err != nil {
	//			return
	//		}
	//	}
	//}

	if len(res.Video) > 0 {

		for id, value := range res.Video {
			key, duration = r.HandlerGetUploadCacheKey(id, FileTypeVideo)
			if err = r.SetToCacheNew(key, duration, value); err != nil {
				return
			}
		}
	}
	//if len(res.Material) > 0 {
	//
	//	for id, value := range res.Material {
	//		key, duration = r.HandlerGetUploadCacheKey(id, FileTypeMaterial)
	//		if err = r.SetToCacheNew(key, duration, value); err != nil {
	//			return
	//		}
	//	}
	//}
	if len(res.File) > 0 {
		for id, value := range res.File {
			key, duration = r.HandlerGetUploadCacheKey(id, FileTypeFile)
			if err = r.SetToCacheNew(key, duration, value); err != nil {
				return
			}
		}
	}
	return
}

func (r *CacheProductPicAndVideoAction) Action() (res *ResultMapUploadInfo, err error) {
	if err = r.argCommon.Default(); err != nil {
		return
	}
	switch r.argCommon.GetType {
	case base.GetDataTypeFromDb: // 从数据库获取数据
		if res, err = r.GetByIdsFromDb(r.arg); err != nil {
			return
		}
		switch r.argCommon.RefreshCache {
		case base.RefreshCacheYes:
			if err = r.saveCache(res); err != nil {
				return
			}
		}
	case base.GetDataTypeFromCache: // 从缓存获取数据
		res, _, err = r.getByIdsFromCache(r.arg)
	case base.GetDataTypeFromAll: // 优先从缓存获取，如果没有数据，则从数据库获取
		res, err = r.getByIdsFromAll(r.arg)
	default:
		err = fmt.Errorf("当前不支持你选择的获取数据类型(%s)", r.argCommon.GetType)
	}
	return
}

func (r *CacheProductPicAndVideoAction) getFromCache(id interface{}, Type string, data interface{}) (err error) {
	defer func() {
		if err != nil && err != redis.Nil {
			r.Context.Info(map[string]interface{}{
				"id":   id,
				"Type": Type,
				"err":  err.Error(),
			}, "CacheProductPicAndVideoActionGetFromCache")
			return
		}
		err = base.NewErrorRuntime(err, base.ErrorRedisCode)
	}()
	key, _ := r.HandlerGetUploadCacheKey(id, Type)
	cmd := r.Context.CacheClient.Get(r.Ctx, key)
	if err = cmd.Err(); err != nil {
		return
	}

	if err = cmd.Scan(data); err != nil {
		return
	}
	return
}

func (r *CacheProductPicAndVideoAction) getByIdsFromCache(arg *ArgUploadGetInfo) (res *ResultMapUploadInfo, noCacheIds *ArgUploadGetInfo, err error) {
	var e error

	res = NewResultMapUploadInfo()

	//收集缓存中没有的数据ID，便于后边查询使用
	noCacheIds = NewArgUploadGetInfo()
	for _, it := range arg.VideoKeys {
		var data = &UploadVideo{}
		if e = r.getFromCache(it, FileTypeVideo, data); e != nil {
			if e != redis.Nil {
				err = e
				return
			}
			noCacheIds.VideoKeys = append(noCacheIds.VideoKeys, it)
			return
		}
		res.Video[it] = data
	}
	for _, it := range arg.MusicKey {
		var data = &UploadMusic{}
		if e = r.getFromCache(it, FileTypeMusic, data); e != nil {
			if e != redis.Nil {
				err = e
				return
			}
			noCacheIds.MusicKey = append(noCacheIds.MusicKey, it)
			return
		}
		res.Music[it] = data
	}
	for _, it := range arg.File {
		var data = &UploadFile{}
		if e = r.getFromCache(it, FileTypeFile, data); e != nil {
			if e != redis.Nil {
				err = e
				return
			}
			noCacheIds.File = append(noCacheIds.File, it)
			return
		}
		res.File[it] = data
	}

	return
}

func (r *CacheProductPicAndVideoAction) getByIdsFromAll(arg *ArgUploadGetInfo) (res *ResultMapUploadInfo, err error) {
	var (
		argCacheNotHave *ArgUploadGetInfo
		dt              *ResultMapUploadInfo
	)
	if res, argCacheNotHave, err = r.getByIdsFromCache(arg); err != nil {
		return
	}
	if len(argCacheNotHave.VideoKeys) == 0 && len(arg.MusicKey) == 0 {
		return
	}

	if dt, err = r.GetByIdsFromDb(argCacheNotHave); err != nil {
		return
	}

	for key, value := range dt.Video {
		res.Video[key] = value
	}
	for key, value := range dt.Music {
		res.Music[key] = value
	}

	for key, value := range dt.File {
		res.File[key] = value
	}

	return
}
