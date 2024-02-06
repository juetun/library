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
		Context          *base.Context
		DataTypes        string `json:"data_types"`
		HeaderInfoString string `json:"header_info_string"`
	}
	dataStruct struct {
		AppName    string     `json:"app_name"`
		URI        string     `json:"uri"`
		Method     string     `json:"method"`
		Parameters url.Values `json:"parameters"`
	}
)

func (r *GetBizData) SyncGetData(groupMapDataId map[string][]string, l int) (res map[string]*DataItem, err error) {
	res = make(map[string]*DataItem, l)

	var (
		//TODO 获取数据在此配置实现方法
		MapDataGetHandler = map[string]dataStruct{
			AdDataDataTypeUserShop: {
				AppName: app_param.AppNameMall,
				URI:     "/shop/get_for_recomm",
				Method:  http.MethodPost, Parameters: url.Values{},
			}, //配置获取电商数据映射
			AdDataDataTypeSpu: {
				AppName: app_param.AppNameMall,
				URI:     "/product/get_spu_by_ids",
				Method:  http.MethodPost,
				Parameters: func() (urlValue url.Values) {
					urlValue = url.Values{}
					urlValue.Set("data_types", r.DataTypes)
					return
				}(),
			}, //配置获取电商数据映射
			AdDataDataTypeSocialIntercourse: {
				AppName: app_param.AppNameSocialIntercourse,
				URI:     "/social_intercourse/get_data_by_ids",
				Method:  http.MethodGet, Parameters: url.Values{},
			},
		}
		dataMul sync.WaitGroup
		lock    sync.Mutex
		ok      bool
		handler dataStruct
	)

	for key, ids := range groupMapDataId {
		if handler, ok = MapDataGetHandler[key]; !ok {
			err = fmt.Errorf("当前不支持您选择的商品数据类型(%s)", key)
			return
		}

		dataMul.Add(1)

		//并行获取商品数据详情
		go func(bizCode string, idString []string, handlerOp dataStruct) {

			defer dataMul.Done()

			var (
				err     error
				resData map[string]*DataItem
			)

			handlerOp.Parameters.Set("ids", strings.Join(ids, ","))
			//发送请求获取数据
			if resData, err = r.GetFromApplication(handlerOp.Parameters, handlerOp.AppName, handlerOp.URI, handlerOp.Method); err != nil {
				return
			}

			lock.Lock()
			defer lock.Unlock()
			for _, value := range resData {
				value.Default()
				res[GetUniqueKey(bizCode, value.DataId)] = value
			}
		}(key, ids, handler)
	}
	dataMul.Wait()
	return
}

func (r *GetBizData) GetFromApplication(args url.Values, appName, URI string, method string) (res map[string]*DataItem, err error) {
	res = map[string]*DataItem{}
	if args == nil {
		return
	}
	if appName == "" {
		err = fmt.Errorf("请选择查询数据的应用")
		return
	}

	params := rpc.RequestOptions{
		Context:     r.Context,
		Method:      method,
		AppName:     appName,
		URI:         URI,
		Value:       args,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		Header:      http.Header{},
	}
	params.Header.Set(app_obj.HttpHeaderInfo, r.HeaderInfoString)
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
