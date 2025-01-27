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
	"strings"
	"sync"
)

type (
	GetBizData struct {
		Context           *base.Context
		GetDataTypeCommon base.GetDataTypeCommon `json:"common"`
		DataTypes         string                 `json:"data_types"`
		UserHID           int64                  `json:"user_hid"`
		HeaderInfoString  string                 `json:"header_info_string"`
	}
	ArgumentGetBizDataItem struct {
		DataTypes string   `json:"data_types"`
		DataIds   []string `json:"data_ids"`
	}
	DataStructArguments struct {
		AppName   string `json:"app_name"`
		URI       string `json:"uri"`
		Method    string `json:"method"`
		OrgParams OrgGetDataParamsHandler
	}
	OrgGetDataParamsHandler func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte)
)

//获取业务数据的参数配置
func GetMapDataGetHandler(dataTypes ...string) (res map[string]DataStructArguments) {
	var dataType string
	if len(dataTypes) > 0 {
		dataType = dataTypes[0]
	}
	return map[string]DataStructArguments{
		AdDataDataTypeUserShop: {
			AppName: app_param.AppNameMall,
			URI:     "/shop/get_for_recomm",
			Method:  http.MethodPost, OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		}, //配置获取电商数据映射
		AdDataDataTypeSpu: {
			AppName: app_param.AppNameMall,
			URI:     "/product/get_spu_by_ids",
			Method:  http.MethodPost,
			OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		}, //配置获取电商数据映射
		AdDataDataTypeSpuCategory: {
			AppName: app_param.AppNameMall,
			URI:     "/category/get_for_recomm",
			Method:  http.MethodPost, OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		}, //配置获取电商数据映射
		AdDataDataTypeSocialIntercourse: {
			AppName: app_param.AppNameSocialIntercourse,
			URI:     "/data/get_article_by_ids",
			Method:  http.MethodPost,
			OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		},
		AdDataDataTypeFishingSport: {
			AppName: app_param.AppNameSocialIntercourse,
			URI:     "/data/get_fishing_spots_by_ids",
			Method:  http.MethodPost,
			OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		},
		AdDataDataTypeRing: { //圈子数据
			AppName: app_param.AppNameSocialIntercourse,
			URI:     "/data/get_ring_by_ids",
			Method:  http.MethodPost,
			OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		},
		AdDataDataTypeGetSnsData: { //获取社交和钓点数据可使用此参数集中获取
			AppName: app_param.AppNameSocialIntercourse,
			URI:     "/data/base_data_by_ids",
			Method:  http.MethodPost,
			OrgParams: func(argItem *ArgumentGetBizDataItem) (urlValue url.Values, requestBody []byte) {
				urlValue = url.Values{}
				urlValue.Set("data_types", dataType)
				urlValue.Set("ids", strings.Join(argItem.DataIds, ","))
				return
			},
		},
	}
}

func (r *GetBizData) SyncGetData(groupMapDataId map[string]*ArgumentGetBizDataItem, l int) (res map[string]*DataItem, err error) {
	res = make(map[string]*DataItem, l)

	var (
		dataMul           sync.WaitGroup
		lock              sync.Mutex
		ok                bool
		handler           DataStructArguments
		MapDataGetHandler = GetMapDataGetHandler(r.DataTypes)
	)

	for key, argumentItem := range groupMapDataId {
		if handler, ok = MapDataGetHandler[key]; !ok {
			err = fmt.Errorf("当前不支持您选择的商品数据类型(%s)", key)
			return
		}
		if len(argumentItem.DataIds) == 0 {
			continue
		}
		dataMul.Add(1)

		//并行获取商品数据详情
		go func(bizCode string, argumentIt *ArgumentGetBizDataItem, handlerOp DataStructArguments) {

			defer dataMul.Done()
			var (
				e       error
				resData map[string]*DataItem
			)

			//发送请求获取数据
			if resData, e = r.GetFromApplication(handlerOp, argumentIt); e != nil {
				return
			}

			lock.Lock()
			defer lock.Unlock()
			for _, value := range resData {
				if value.DataType == "" {
					value.DataType = bizCode
				}
				value.Default()
				res[GetUniqueKey(value.DataType, value.DataId)] = value
			}
		}(key, argumentItem, handler)
	}
	dataMul.Wait()
	return
}

func (r *GetBizData) GetFromApplication(handlerOp DataStructArguments, argumentIt *ArgumentGetBizDataItem) (res map[string]*DataItem, err error) {
	//args url.Values, appName, URI string, method string
	//.OrgParams(), handlerOp.AppName, handlerOp.URI, handlerOp.Method
	res = map[string]*DataItem{}
	if handlerOp.AppName == "" {
		err = fmt.Errorf("请选择查询数据的应用")
		return
	}
	var (
		urlValue, requestBody = handlerOp.OrgParams(argumentIt)
		params                = rpc.RequestOptions{
			Context:     r.Context,
			Method:      handlerOp.Method,
			AppName:     handlerOp.AppName,
			URI:         handlerOp.URI,
			Value:       urlValue,
			PathVersion: app_obj.App.AppRouterPrefix.Intranet,
			Header:      http.Header{},
			BodyJson:    requestBody,
		}
	)
	params.Header.Set(app_obj.HttpHeaderInfo, r.HeaderInfoString)
	params.Header.Set(app_obj.HttpUserHid, fmt.Sprintf("%v", r.UserHID))
	httpRpc := rpc.NewHttpRpc(&params)
	req := httpRpc.Send()

	if err = req.Error; err != nil {
		return
	}
	var body []byte
	if body = req.GetBody().Body; len(body) == 0 {
		return
	}
	if err = req.Error; err != nil {
		return
	}
	var resResult struct {
		Code int                  `json:"code"`
		Data map[string]*DataItem `json:"data"`
		Msg  string               `json:"message"`
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

//根据场景Key获取数据
func GetRecommendDataByScenes(arg *ArgGetDataByScenes, ctx *base.Context) (res ResultGetDataByScenes, err error) {

	var header = http.Header{}
	header.Set(app_obj.HttpHeaderInfo, arg.HeaderInfoString)
	header.Set(app_obj.HttpUserHid, fmt.Sprintf("%v", arg.UUserHid))
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameRecommend,
		URI:         "/recommend/get_data_by_scenes",
		Header:      header,
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}

	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int                   `json:"code"`
		Data ResultGetDataByScenes `json:"data"`
		Msg  string                `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res = data.Data
	return
}
