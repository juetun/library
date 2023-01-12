package app_user

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/app_param/mall"
	"net/http"
	"net/url"
	"strings"
)

//根据店铺ID获取店铺信息
func GetShopDataByUIds(ctx *base.Context, shopIds []int64, dataTypes ...string) (res map[int64]*mall.ShopData, err error) {
	var value = url.Values{}
	var shopID = make([]string, 0, len(shopIds))
	for _, value := range shopIds {
		shopID = append(shopID, fmt.Sprintf("%d", value))
	}
	uIds := strings.Join(shopID, ",")
	value.Set("user_hid", uIds)
	value.Set("data_type", strings.Join(dataTypes, ","))
	ro := rpc.RequestOptions{
		Method:      http.MethodGet,
		AppName:     app_param.AppNameMall,
		URI:         "/shop/get_by_shop_ids",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                      `json:"code"`
		Data map[int64]*mall.ShopData `json:"data"`
		Msg  string                   `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	res = data.Data
	return
}
