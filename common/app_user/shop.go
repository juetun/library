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

type (
	ArgUserReviewSubmit struct {
		ShopId   int64           `json:"shop_id" form:"shop_id"`
		UserHid  int64           `json:"user_hid" form:"user_hid"`
		BatchId  string          `json:"batch_id" form:"batch_id"`
		DataType string          `json:"data_type" form:"data_type"`
		Mark     string          `json:"mark" form:"mark"`
		Status   int8            `json:"status" form:"status"`
		TimeNow  base.TimeNormal `json:"time_now" form:"time_now"`
	}
)

//根据店铺ID获取店铺信息
func GetShopDataByUIds(ctx *base.Context, shopIds []int64, dataTypes ...string) (res map[int64]*mall.ShopData, err error) {
	res = map[int64]*mall.ShopData{}
	if len(shopIds) == 0 {
		return
	}
	var value = url.Values{}
	var shopID = make([]string, 0, len(shopIds))
	for _, value := range shopIds {
		shopID = append(shopID, fmt.Sprintf("%d", value))
	}
	value.Set("shop_ids", strings.Join(shopID, ","))
	if len(dataTypes) > 0 {
		value.Set("data_type", strings.Join(dataTypes, ","))
	}
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

func (r *ArgUserReviewSubmit) Default(ctx *base.Context) (err error) {
	if r.BatchId == "" || r.BatchId == "0" {
		err = fmt.Errorf("请选择您要审核信息")
		return
	}
	if err = r.validateDataType(); err != nil {
		return
	}
	switch r.Status {
	case UserApplyStatusFailure:
		if r.DataType != UpdateBatchTypeAuthType && r.Mark == "" {
			err = fmt.Errorf("请填写审核失败的备注")
			return
		}
	case UserApplyStatusUsing:
	default:
		err = fmt.Errorf("审核状态当前只支持审核成功和审核失败")
		return
	}
	if r.UserHid == 0 {
		err = fmt.Errorf("请选择数据所属用户")
		return
	}
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}

func (r *ArgUserReviewSubmit) validateDataType() (err error) {
	var mapDataType map[string]string
	mapDataType, _ = SliceUpdateBatchType.GetMapAsKeyString()
	if _, ok := mapDataType[r.DataType]; !ok {
		err = fmt.Errorf("当前暂不支持你审核的数据类型")
		return
	}
	return
}
