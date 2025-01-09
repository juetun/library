package app_param

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"strings"
)

const(
	PageNeedUserYes bool=true
)
type PageCommonParams struct {
	City             string          `json:"city" form:"city"`
	Cate             string          `json:"cate" form:"cate"`
	HeaderInfoString string          `json:"-" form:"-"`
	TimeNow          base.TimeNormal `json:"" form:"-"`
	common.HeaderInfo
	RequestUser
}

func (r *PageCommonParams) Default(c *base.Context, needUsers ...bool) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	_ = r.InitHeaderInfo(c.GinContext)
	if r.HeaderInfo.HTerminal == "" {
		r.HeaderInfo.HTerminal = TerminalWeb
	}
	if r.HeaderInfo.HApp == "" {
		r.HeaderInfo.HApp = HTerminal
	}
	switch len(needUsers) {
	case 1:
		r.City = r.GetCityCode(c.GinContext)
		if needUsers[0] {
			if err = r.InitRequestUser(c); err != nil {
				return
			}
		}

	case 2:
		r.City = r.GetCityCode(c.GinContext)
		if needUsers[0] {
			if err = r.InitRequestUser(c, needUsers[1]); err != nil {
				return
			}
		}
	default:
		r.City = r.GetCityCode(c.GinContext)
	}
	if r.HeaderInfoString, err = r.InitWebHeaderInfo(r.HeaderInfo, c.GinContext); err != nil {
		return
	}
	return
}

func (r *PageCommonParams) GetCityCode(c *gin.Context) (code string) {
	v, ok := c.Get(base.MiddleCityCode)
	if ok {
		code = v.(string)
	}
	return
}

func (r *PageCommonParams) InitWebHeaderInfo(headerInfo common.HeaderInfo, c *gin.Context) (headerInfoString string, err error) {
	var (
		headerInfoByte []byte
	)
	if headerInfoByte, err = json.Marshal(headerInfo); err != nil {
		return
	}
	if len(headerInfoByte) == 0 {
		return
	}
	base64Code := base64.StdEncoding.EncodeToString(headerInfoByte)
	aesObject := common.NewAes()
	if headerInfoString, err = aesObject.EncryptionCtr(base64Code, app_obj.TmpSignKey); err != nil {
		return
	}
	return
}

func (r *PageCommonParams) GetDetailParamByKey(c *gin.Context, key string, ext ...string) string {
	extName := ".html"
	if len(ext) > 0 {
		extName = ext[0]
	}
	data := c.Params.ByName(key)
	return strings.TrimSuffix(data, extName)
}
