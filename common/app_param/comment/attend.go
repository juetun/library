package comment

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"strconv"
)

type (
	ArgHasAttendStatus struct {
		CurrentUid    int64                           `json:"current_uid" form:"current_uid"`
		TargetUid     []*ArgHasAttendStatusTargetItem `json:"target_uid" form:"target_uid"` //多个用逗号隔开
		TargetUids    []int64                         `json:"-" form:"-"`
		TargetShopIds []int64                         `json:"-" form:"-"`
	}
	ArgHasAttendStatusTargetItem struct {
		Type string `json:"type" form:"type"`
		Id   string `json:"id" form:"id"`
	}
	ResultHasAttendStatus map[string]map[int64]bool
)

func (r *ArgHasAttendStatus) Default(ctx *base.Context) (err error) {
	r.TargetUids = make([]int64, 0)
	var (
		id int64
		l  = len(r.TargetUid)
	)
	if len(r.TargetUid) > 0 {
		r.TargetUids = make([]int64, 0, l)
		for _, item := range r.TargetUid {
			if id, err = strconv.ParseInt(item.Id, 10, 64); err != nil {
				err = fmt.Errorf("target_uid参数格式错误")
				return
			}
			if id == 0 {
				continue
			}
			switch item.Type {
			case AttendDataTypeShop:
				r.TargetShopIds = append(r.TargetShopIds, id)
			case AttendDataTypeUser:
				r.TargetUids = append(r.TargetUids, id)
			default:

			}
		}
	}

	return
}

func (r *ArgHasAttendStatus) ToJson() (res []byte) {
	res, _ = json.Marshal(r)
	return
}

//根据用户ID获取用户信息
func FetchHasAttendStatus(ctx *base.Context, arg *ArgHasAttendStatus) (res map[string]map[int64]bool, err error) {
	var (
		l = len(arg.TargetUid)
	)
	res = make(map[string]map[int64]bool, l)
	if l == 0 {
		return
	}
	var value = url.Values{}
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameComment,
		URI:         "/act/has_attend",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		BodyJson:    arg.ToJson(),
	}
	var data = struct {
		Code int                       `json:"code"`
		Data map[string]map[int64]bool `json:"data"`
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
