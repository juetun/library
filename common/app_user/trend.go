package app_user

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/common/response"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/recommend"
	"net/http"
	"net/url"
)

const (
	TrendTypeShopInfo     = "shop_info"     //店铺信息
	TrendTypeShopBrand    = "shop_brand"    //店铺品牌
	TrendTypeShopCate     = "shop_cate"     //店铺类目
	TrendTypeSpuEdit      = "product_edit"  //编辑商品
	TrendTypeSns          = "sns"           //圈子动态
	TrendTypeFishingSport = "fishing_spots" //钓点信息
	TrendTypeRegister     = "register"      //用户注册
	TrendTypeUsrApply     = "usr_apply"     //用户审核
)

var (
	SliceTrendType = base.ModelItemOptions{
		{
			Value: TrendTypeShopInfo,
			Label: "店铺信息变更",
		},
		{
			Value: TrendTypeShopBrand,
			Label: "店铺品牌变更",
		},
		{
			Value: TrendTypeShopCate,
			Label: "店铺类目变更",
		},
		{
			Value: TrendTypeSpuEdit,
			Label: "编辑商品",
		},
		{
			Value: TrendTypeSns,
			Label: "圈子动态",
		},
		{
			Value: TrendTypeFishingSport,
			Label: "钓点信息",
		},
		{
			Value: TrendTypeRegister,
			Label: "注册账号",
		},
		{
			Value: TrendTypeUsrApply,
			Label: "账号认证审核",
		},
	}
)

type (
	ArgRecommendAttendUser struct {
		app_param.RequestUser
		common.HeaderInfo
		response.PageQuery
		TimeNow           base.TimeNormal        `json:"-" form:"-"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"-" form:"-"`
	}
)

//根据场景Key获取数据
func GetUserRecommendData(arg *ArgRecommendAttendUser, headerInfoString string, ctx *base.Context) (res *recommend.PagerRecommend, err error) {

	var header = http.Header{}
	res = &recommend.PagerRecommend{}
	header.Set(app_obj.HttpHeaderInfo, headerInfoString)
	header.Set(app_obj.HttpUserHid, fmt.Sprintf("%v", arg.UUserHid))
	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameUser,
		URI:         "/user/recommend_attend_user",
		Header:      header,
		Value:       url.Values{},
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}

	if ro.BodyJson, err = arg.GetJson(); err != nil {
		return
	}
	var data = struct {
		Code int `json:"code"`
		Data struct {
			Pager struct {
				List          []recommend.DataItem     `json:"list"`
				TotalCount    int64                    `json:"total_count,omitempty"`
				IsNext        bool                     `json:"is_next,omitempty"` // [bool] 是否有下一页，true=有下一页；false=无下页，可关闭列表
				SceneProperty *recommend.SceneProperty `json:"scene_property"`
				AdCount       int                      `json:"ad_count,omitempty"` //广告推流数量
				response.PageQuery
			} `pager`
		} `json:"data"`
		Msg string `json:"message"`
	}{}

	if err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error; err != nil {
		return
	}
	res.List = data.Data.Pager.List
	return
}

func (r *ArgRecommendAttendUser) Default(ctx *base.Context) (err error) {
	_ = r.InitRequestUser(ctx)
	_ = r.InitHeaderInfo(ctx.GinContext)
	r.TimeNow = base.GetNowTimeNormal()
	return
}

func (r *ArgRecommendAttendUser) GetJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}
