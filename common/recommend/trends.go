package recommend

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

type (
	//动态信息
	TrendContent struct {
		DataType  string          `json:"data_type"`            //数据类型
		DataId    string          `json:"data_id"`              //数据ID
		Img       string          `json:"img,omitempty"`        //头图
		HaveVideo bool            `json:"have_video,omitempty"` //是否有视频
		Video     string          `json:"video,omitempty"`      //视频信息
		Title     string          `json:"title,omitempty"`
		Desc      string          `json:"desc,omitempty"` //动态内容
		Time      base.TimeNormal `json:"time"`           //时间
	}
	TrendContents  []*TrendContent
	ResultAddTrend struct {
		Result bool `json:"result"`
	}
)

func (r *TrendContents) GetJsonByte() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	return
}

//TODO 添加多条动态 ,当前直接调用接口写入数据库,后续建议使用MQ写入队列 解耦
func AddTrends(ctx *base.Context, data TrendContents) (err error) {
	arg := url.Values{}
	params := rpc.RequestOptions{
		Context:     ctx,
		Method:      http.MethodPost,
		AppName:     app_param.AppNameComment,
		URI:         "/add_trends",
		Value:       arg,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}
	if ctx.GinContext != nil {
		params.Header.Set(app_obj.HttpHeaderInfo, ctx.GinContext.GetHeader(app_obj.HttpHeaderInfo))
	}
	if params.BodyJson, err = data.GetJsonByte(); err != nil {
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
		Code int            `json:"code"`
		Data ResultAddTrend `json:"data"`
		Msg  string         `json:"message"`
	}
	if err = json.Unmarshal(body, &resResult); err != nil {
		return
	}
	if resResult.Code > 0 {
		err = fmt.Errorf(resResult.Msg)
		return
	}
	return
}

//添加一条动态
func AddTrend(ctx *base.Context, data *TrendContent) (err error) {
	err = AddTrends(ctx, []*TrendContent{data})
	return
}
