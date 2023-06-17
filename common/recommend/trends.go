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

const (
	TrendContentShowYes uint8 = iota + 1 //可见
	TrendContentShowNo                   //不可见
)

var (
	SliceTrendContentShow = base.ModelItemOptions{
		{
			Value: TrendContentShowYes,
			Label: "可见",
		},
		{
			Value: TrendContentShowNo,
			Label: "不可见",
		},
	}
)

type (
	//动态信息
	TrendContent struct {
		UserHid        int64           `json:"user_hid"`             //用户信息
		DataType       string          `json:"data_type"`            //数据类型
		DataId         string          `json:"data_id"`              //数据ID
		Img            string          `json:"img,omitempty"`        //头图
		HaveVideo      bool            `json:"have_video,omitempty"` //是否有视频
		Video          string          `json:"video,omitempty"`      //视频信息
		UserShow       uint8           `json:"user_show"`            //用户是否可见 1-可见 2-不可见
		AttendUserShow uint8           `json:"attend_user_show"`     //关注用户是否可见 1-可件 2-不可见
		Title          string          `json:"title,omitempty"`      //动态标题
		Desc           string          `json:"desc,omitempty"`       //动态内容
		Time           base.TimeNormal `json:"time"`                 //时间
	}
	TrendContents struct {
		Data []*TrendContent `json:"data"`
	}
	ResultAddTrend struct {
		Result bool `json:"result"`
	}
)

func (r *TrendContent) Default() (res string) {
	if r.AttendUserShow == 0 {
		r.AttendUserShow = TrendContentShowYes
	}
	if r.UserShow == 0 {
		r.UserShow = TrendContentShowYes
	}
	return
}

func (r *TrendContent) ParseAttendUserShow() (res string) {
	mapShow, _ := SliceTrendContentShow.GetMapAsKeyUint8()
	if tmp, ok := mapShow[r.AttendUserShow]; ok {
		res = tmp
		return
	}
	res = fmt.Sprintf("未知类型(%v)", r.UserShow)
	return
}

func (r *TrendContent) ParseUserShow() (res string) {
	mapShow, _ := SliceTrendContentShow.GetMapAsKeyUint8()
	if tmp, ok := mapShow[r.UserShow]; ok {
		res = tmp
		return
	}
	res = fmt.Sprintf("未知类型(%v)", r.UserShow)
	return
}

func (r *TrendContents) GetUserHidAndMap() (userHIds []int64, dataMap map[string][]*TrendContent, err error) {
	var l = len(r.Data)
	userHIds = make([]int64, 0, l)
	dataMap = make(map[string][]*TrendContent, l)
	var kv string
	for _, item := range r.Data {
		kv = fmt.Sprintf("%v", item.UserHid)
		if _, ok := dataMap[kv]; !ok {
			dataMap[kv] = make([]*TrendContent, 0, l)
			userHIds = append(userHIds, item.UserHid)
			continue
		}
		dataMap[kv] = append(dataMap[kv], item)
	}
	return
}

func (r *TrendContents) GetJsonByte() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	return
}

func (r *TrendContents) Default(ctx *base.Context) (err error) {
	for _, item := range r.Data {
		item.Default()
	}
	return
}

//TODO 添加多条动态 ,当前直接调用接口写入数据库,后续建议使用MQ写入队列 解耦
func AddTrends(ctx *base.Context, data *TrendContents) (err error) {
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
	err = AddTrends(ctx, &TrendContents{
		Data: []*TrendContent{data},
	})
	return
}
