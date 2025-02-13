package recommend

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/common/response"
	"github.com/juetun/base-wrapper/lib/utils"
	"github.com/juetun/library/common/app_param"
)

type (
	ArgGetDataByScenes struct {
		NowItemId         string                 `json:"now_id,omitempty" form:"now_id"`                   // 当前数据的ID
		NowItemType       int8                   `json:"now_type,omitempty" form:"now_type"`               // 当前数据的类型
		HttpRequestId     string                 `json:"http_request_id,omitempty" form:"http_request_id"` // 每次请求的数据的ID
		DeviceId          string                 `json:"device_id,omitempty" form:"device_id"`             // 用户设备号
		ClientType        string                 `json:"client_type,omitempty" form:"client_type"`         // 终端类型 "m ,android,iso,weixin,alipay"
		Channel           string                 `json:"channel,omitempty" form:"channel"`                 // APP发布的渠道
		AppVersion        string                 `json:"app_version,omitempty" form:"app_version"`         // APP版本号
		App               string                 `json:"app,omitempty" form:"app"`                         // APP名称
		GetDataTypeCommon base.GetDataTypeCommon `json:"common,omitempty" form:"common"`
		Scenes            []*ArgDataBySceneItem  `json:"scenes,omitempty"`

		HeaderInfoString string          `json:"header_info_str" form:"-"`
		TimeNow          base.TimeNormal `json:"-" form:"-"`
		app_param.RequestUser
		common.HeaderInfo
	}
	ArgDataBySceneItem struct {
		Scene      string                 `json:"scene,omitempty" form:"scene"`
		SceneId    int64                  `json:"scene_id,omitempty" form:"scene_id"`
		NowType    string                 `json:"now_type,omitempty" form:"now_type"`
		NowId      string                 `json:"now_id,omitempty" form:"now_id"`
		Parameters map[string]interface{} `json:"parameters" form:"parameters"`
		response.PageQuery
	}
	ResultGetDataByScenes map[string]*PagerRecommend

	PagerRecommend struct {
		List          []DataItem     `json:"list"`
		TotalCount    int64          `json:"total_count,omitempty"`
		IsNext        bool           `json:"is_next,omitempty"` // [bool] 是否有下一页，true=有下一页；false=无下页，可关闭列表
		SceneProperty *SceneProperty `json:"scene_property"`
		AdCount       int            `json:"ad_count,omitempty"` //广告推流数量
		response.PageQuery
	}
	SceneProperty struct {
		LimitType uint8  `json:"limit_type,omitempty" form:"limit_type"`
		PicW      string `json:"pic_w,omitempty" form:"pic_w"`
		PicH      string `json:"pic_h,omitempty" form:"pic_h"`
		RatioW    string `json:"ratio_w,omitempty" form:"ratio_w"`
		RatioH    string `json:"ratio_h,omitempty" form:"ratio_h"`
	}
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
	_ = r.InitHeaderInfo(ctx.GinContext)
	_ = r.InitRequestUser(ctx)
	r.HeaderInfoString = ctx.GinContext.Request.Header.Get(app_obj.HttpHeaderInfo)
	if r.HttpRequestId == "" {
		if r.HttpRequestId = ctx.GinContext.GetHeader(app_obj.HttpTraceId); r.HttpRequestId == "" {
			r.HttpRequestId = utils.Guid("")
		}
	}
	return
}

func (r *ArgGetDataByScenes) GetJson() (res []byte, err error) {
	if r == nil {
		return
	}
	res, err = json.Marshal(r)
	return
}
