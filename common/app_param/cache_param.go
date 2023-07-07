package app_param

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"net/http"
	"net/url"
)

type (
	ArgClearCacheByKeyPrefix struct {
		MicroApp  string `json:"micro_app" form:"micro_app"`
		KeyPrefix string `json:"key_prefix" form:"key_prefix"`
	}

	ResultClearCacheByKeyPrefix struct {
		Resutlt bool `json:"resutlt"`
	}

	ArgGetCacheParamConfig struct {
		MicroApp string `json:"micro_app" form:"micro_app"`
	}

	ResultGetCacheParamConfig map[string]*redis_pkg.CacheProperty
)

func (r *ArgClearCacheByKeyPrefix) Default(c *base.Context) (err error) {

	return
}

func (r *ArgGetCacheParamConfig) Default(c *base.Context) (err error) {

	return
}

//
func ReloadAppCacheConfig(ctx *base.Context, argGetCacheParamConfig *ArgGetCacheParamConfig) (res ResultGetCacheParamConfig, err error) {
	res = ResultGetCacheParamConfig{}
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodGet,
		AppName:     argGetCacheParamConfig.MicroApp,
		URI:         "/cache/get_cache_param_config",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}
	if params.BodyJson, err = json.Marshal(argGetCacheParamConfig); err != nil {
		return
	}
	req := rpc.NewHttpRpc(&params).
		Send().GetBody()
	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.Body; len(body) == 0 {
		return
	}

	var resResult struct {
		Code int                       `json:"code"`
		Data ResultGetCacheParamConfig `json:"data"`
		Msg  string                    `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}
	res = resResult.Data

	return
}
