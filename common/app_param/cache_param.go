package app_param

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
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
