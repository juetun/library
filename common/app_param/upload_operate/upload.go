package upload_operate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/base/cache_act"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"time"
)

type (
	DaoUploadImpl struct {
		base.ServiceDao
		ctx context.Context
	}
	DaoUpload interface {
		GetUploadByKeys(arg *ArgUploadGetInfo, argCommon *base.GetDataTypeCommon) (res *ResultMapUploadInfo, err error)
	}
)

func (r *DaoUploadImpl) getDataByUserIdsFromUploadServer(arg *ArgUploadGetInfo) (resData *ResultMapUploadInfo, err error) {
	resData = NewResultMapUploadInfo()
	var value = url.Values{}
	var bodyByte []byte

	//判断参数是否为空
	if arg == nil || arg.IsNull() {
		return
	}

	if bodyByte, err = json.Marshal(arg); err != nil {
		return
	}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUpload,
		URI:         "/upload/get_upload_address",
		Header:      http.Header{},
		Value:       value,
		BodyJson:    bodyByte,
		Context:     r.Context,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                  `json:"code"`
		Data *ResultMapUploadInfo `json:"data"`
		Msg  string               `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	if data.Data != nil {
		resData = data.Data
	}

	return
}

func (r *DaoUploadImpl) getUploadCacheKey(id interface{}, Type string, expireTimeRands ...bool) (res string, timeExpire time.Duration) {
	res = fmt.Sprintf(CacheUploadCache.Key, Type, id)
	timeExpire = CacheUploadCache.Expire
	var expireTimeRand bool
	if len(expireTimeRands) > 0 {
		expireTimeRand = expireTimeRands[0]
	}
	if expireTimeRand {
		randNumber, _ := r.RealRandomNumber(60) //打乱缓存时长，防止缓存同一时间过期导致数据库访问压力变大
		timeExpire = timeExpire + time.Duration(randNumber)*time.Second
	}
	return
}

//从
func (r *DaoUploadImpl) GetUploadByKeys(arg *ArgUploadGetInfo, argCommon *base.GetDataTypeCommon) (res *ResultMapUploadInfo, err error) {

	res, err = NewCacheProductPicAndVideoAction(
		CacheProductPicAndVideoActionArg(arg, argCommon),
		CacheHandlerGetUploadCacheKey(r.getUploadCacheKey),
		CacheProductPicAndVideoActionGetByIdsFromDb(r.getDataByUserIdsFromUploadServer),
	).LoadBaseOption(
		cache_act.CacheActionBaseContext(r.Context),
		cache_act.CacheActionBaseCtx(r.ctx), ).
		Action()
	return
}

//移除接口
func (r *DaoUploadImpl) RemoveFile() (res *ResultUploadRemove, err error) {

	return
}
func NewDaoUpload(ctx ...*base.Context) DaoUpload {
	p := &DaoUploadImpl{}
	p.SetContext(ctx...)
	if p.ctx == nil {
		p.ctx = context.TODO()
	}
	return p
}
