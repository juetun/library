package app_user

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
)

//根据用户ID获取用户信息
func GetUserByUIds(ctx *base.Context, userId []int64, dataTypes ...string) (res map[int64]*app_param.User, err error) {
	var (
		l = len(userId)
	)
	res = make(map[int64]*app_param.User, l)
	if l == 0 {
		return
	}
	var value = url.Values{}
	var userHID = make([]string, 0, l)
	for _, value := range userId {
		userHID = append(userHID, fmt.Sprintf("%d", value))
	}
	uIds := strings.Join(userHID, ",")
	value.Set("user_hid", uIds)
	value.Set("data_type", strings.Join(dataTypes, ","))
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUser,
		URI:         "/user/get_intact_user_by_uid",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	var data = struct {
		Code int                       `json:"code"`
		Data map[int64]*app_param.User `json:"data"`
		Msg  string                    `json:"message"`
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

//根据用户ID获取用户信息
func GetResultUserByUid(userId string, ctx *base.Context) (res *app_param.ResultUser, err error) {
	res, err = app_param.GetResultUserByUid(userId, ctx)
	return
}

type (
	ArgUpdateByUid struct {
		Uid  int64                  `json:"uid" form:"uid"`
		Data map[string]interface{} `json:"data" form:"data"`
	}
	ArgUpdateByUIds struct {
		UserInfos []ArgUpdateByUid `json:"user_infos"`
	}
	ResultUpdateByUIds struct {
		Result bool `json:"result"`
	}
)

//根据用户ID获取用户信息
func UpdateUserByUserHIds(ctx *base.Context, args *ArgUpdateByUIds) (res *ResultUpdateByUIds, err error) {
	res = &ResultUpdateByUIds{}
	var value = url.Values{}

	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUser,
		URI:         "/user/update_by_uid_list",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = json.Marshal(args); err != nil {
		return
	}
	var data = struct {
		Code int                 `json:"code"`
		Data *ResultUpdateByUIds `json:"data"`
		Msg  string              `json:"message"`
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
