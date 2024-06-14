package recommend

import (
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/common/response"
	"github.com/juetun/base-wrapper/lib/utils"
	"github.com/juetun/library/common/app_param"
)

type (
	ArgGetDataByScenes struct {
		NowItemId         string                 `json:"now_id" form:"now_id"`           // 当前数据的ID
		NowItemType       int8                   `json:"now_type" form:"now_type"`       // 当前数据的类型
		RequestId         string                 `json:"request_id" form:"request_id"`   // 每次请求的数据的ID
		DeviceId          string                 `json:"device_id" form:"device_id"`     // 用户设备号
		ClientType        string                 `json:"client_type" form:"client_type"` // 终端类型 "m ,android,iso,weixin,alipay"
		Channel           string                 `json:"channel" form:"channel"`         // APP发布的渠道
		AppVersion        string                 `json:"app_version" form:"app_version"` // APP版本号
		App               string                 `json:"app" form:"app"`                 // APP名称
		GetDataTypeCommon base.GetDataTypeCommon `json:"common" form:"common"`
		Scenes            []*ArgDataBySceneItem  `json:"scenes"`

		HeaderInfoString string          `json:"header_info_str" form:"-"`
		TimeNow          base.TimeNormal `json:"-" form:"-"`
		app_param.RequestUser
		common.HeaderInfo
	}
	ArgDataBySceneItem struct {
		Scene   string `json:"scene" form:"scene"`
		SceneId int64  `json:"scene_id" form:"scene_id"`
		response.PageQuery
	}
	ResultGetDataByScenes map[string]*response.Pager
)

func (r *ArgGetDataByScenes) GetSceneKeys() (res []string) {
	var l = len(r.Scenes)
	res = make([]string, 0, l)
	var mapScenes = make(map[string]bool, l)
	for _, item := range r.Scenes {
		if item.Scene == "" {
			continue
		}
		if _, ok := mapScenes[item.Scene]; !ok {
			mapScenes[item.Scene] = true
			res = append(res, item.Scene)
		}
	}
	return
}

func (r *ArgGetDataByScenes) Default(ctx *base.Context) (err error) {
	if err = r.InitHeaderInfo(ctx.GinContext); err != nil {
		return
	}
	_ = r.InitRequestUser(ctx)
	r.HeaderInfoString = ctx.GinContext.Request.Header.Get(app_obj.HttpHeaderInfo)
	if r.RequestId == "" {
		if r.RequestId = ctx.GinContext.GetHeader(app_obj.HttpTraceId); r.RequestId == "" {
			r.RequestId = utils.Guid("")
		}
	}
	return
}
