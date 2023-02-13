package recommend

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

const (
	SceneNameTendencies    = "tendencies"     //广告用户首页动态
	SceneNameHomeRecommend = "home_recommend" //广告首页数据推荐
)

const (
	AdDataDataTypeSpu               int8 = iota + 1 //广告商品信息
	AdDataDataTypeSocialIntercourse                 //广告社交动态信息
	AdDataStatusUserSet                             //用户手工设置
)

const (
	AdDataStatusCanUse  uint8 = iota + 1 //广告可用
	AdDataStatusOffLine                  //广告下架

)

var (
	SliceAdDataType = base.ModelItemOptions{ //
		{
			Label: "电商",
			Value: AdDataDataTypeSpu,
		},
		{
			Label: "社交",
			Value: AdDataDataTypeSocialIntercourse,
		},
		{
			Label: "手工指定",
			Value: AdDataStatusUserSet,
		},
	}
	SliceAdDataStatus = base.ModelItemOptions{
		{
			Label: "可用",
			Value: AdDataStatusCanUse,
		},
		{
			Label: "下架",
			Value: AdDataStatusOffLine,
		},
	}
)

type (
	ArgWriteRecommendData struct {
		Ctx      *base.Context `json:"-"`         //上下文信息
		UserHid  int64         `json:"user_hid"`  //用户
		DataType int8          `json:"data_type"` //数据类型
		DataId   string        `json:"data_id"`   //数据ID
		SceneKey string        `json:"scene_key"`
		Status   uint8         `json:"status"`
		Weight   int64         `json:"weight"`
	}
)

func (r *ArgWriteRecommendData) Default() (err error) {
	var (
		dataTypeMap map[int8]string
		ok          bool
	)
	if dataTypeMap, err = SliceAdDataType.GetMapAsKeyInt8(); err != nil {
		if _, ok = dataTypeMap[r.DataType]; !ok {
			err = fmt.Errorf("当前不支持你选择的数据类型")
			r.Ctx.Error(map[string]interface{}{
				"arg": r,
				"err": err.Error(),
			}, "ArgWriteRecommendDataDefault")
			return
		}
		return
	}
	switch r.SceneKey {
	case SceneNameTendencies: //如果是用户首页广告动态,则必须填写用户ID
		if r.UserHid == 0 {
			err = fmt.Errorf("请填写要推送的用户信息")
			return
		}
	}

	if r.Status == 0 { //默认不传为可用数据
		r.Status = AdDataStatusCanUse
	}
	return
}

//TODO 根据店铺ID获取店铺信息，
//当前采用接口直接发送，后续采用消息队列解耦重新实现
func WriteRecommendData(arg *ArgWriteRecommendData) (res bool, err error) {
	if err = arg.Default(); err != nil {
		return
	}
	var value = url.Values{}
	value.Set("user_hid", fmt.Sprintf("%d", arg.UserHid))
	value.Set("data_type", fmt.Sprintf("%d", arg.DataType))
	value.Set("data_id", arg.DataId)
	value.Set("scene_key", arg.SceneKey)
	value.Set("status", fmt.Sprintf("%d", arg.Status))
	value.Set("weight", fmt.Sprintf("%d", arg.Weight))
	ro := rpc.RequestOptions{
		Method:      http.MethodGet,
		AppName:     app_param.AppNameRecommend,
		URI:         "/recommend/write_data",
		Header:      http.Header{},
		Value:       value,
		Context:     arg.Ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                                   `json:"code"`
		Data struct{ Result bool `json:"result"` } `json:"data"`
		Msg  string                                `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	res = data.Data.Result
	return
}
