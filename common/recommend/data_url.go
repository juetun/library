package recommend

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
)

const (
	PageNameHome = "home" //首页

	//商品前台界面
	PageNameOrder        = "order"
	PageNameSpu          = "spu"
	PageNameShop         = "shop"
	PageNameUsr          = "user"
	PageNameSns          = "article"       //社交前台界面
	PageNameFishingSport = "fishing_spots" //钓点信息
	PageNameRing         = "ring_article"  //圈子信息

	PageNameOther      = "other"      //无实际意义拼接数据使用
	PageNamePayFinish  = "pay_finish" //支付结果界面
	PageNamePayPreview = "pay_view"   //预付款界面
)

const (
	//店铺后台界面名称
	UserShopInfo = "shop_info"
	UserShopHome = "shop_home"
)

var (
	MapDataTypeBiz = map[string]string{
		AdDataDataTypeSpu:               PageNameSpu,          //商品信息（spu）
		AdDataDataTypeSku:               PageNameSpu,          //商品(sku)信息
		AdDataDataTypeUserShop:          PageNameShop,         //店铺信息（后台）
		AdDataDataTypeUserShopHome:      PageNameShop,         //店铺信息（前台）
		AdDataDataTypeUser:              PageNameUsr,          //用户信息
		AdDataDataTypeSpuCategory:       PageNameHome,         //类目页面 -首页- 电商 tab
		AdDataDataTypeSocialIntercourse: PageNameSns,          //广告社交动态信息
		AdDataDataTypeFishingSport:      PageNameFishingSport, //钓点信息
		AdDataDataTypeRing:              PageNameRing,         //圈子信息
		AdDataDataTypeOther:             PageNameOther,        //其他信息
	}
	MapPageMallName = map[string]PageUrl{
		PageNameOrder: {
			H5:  "/#/order/pages/detail/index",
			Web: "/order_detail.html",
		},
		PageNameSpu: {
			H5:  "/#/mall/pages/detail/index",
			Web: "/spu_{{.spu_id}}.html",
		},
		PageNameShop: {
			H5:  "/#/shop/pages/home/index",
			Web: "/shop_{{.shop_id}}.html",
		},
		PageNameUsr: {
			H5:  "/#/user/pages/view/index",
			Web: "/usr_{{.user_id}}.html",
		},
		AdDataDataTypeSocialIntercourse: {
			H5:  "/#/sns/pages/detail/index",
			Web: "/article_{{.id}}.html",
		},
		AdDataDataTypeRing: {
			H5:  "/#/sns/pages/ring_article/index",
			Web: "/ring_{{.id}}.html",
		}, //圈子界面
		AdDataDataTypeFishingSport: {
			H5:  "/#/fishingsport/pages/detail/index",
			Web: "/fish_sports_{{.id}}.html",
		},
		PageNamePayFinish: {
			H5:  "/#/order/pages/pay_finish/index",
			Web: "/page_finish.html",
		}, //支付结果界面
		PageNamePayPreview: {
			H5:  "/#/order/pages/preview/index",
			Web: "/pay.html",
		}, //支付预付界面
	}
	MapPageSNsName = map[string]PageUrl{
		PageNameHome: {
			H5:  "/#/home/index/index",
			Web: "/home.html",
		},
		PageNameSns: {
			H5:  "/#/sns/pages/detail/index",
			Web: "/article_detail_{{.id}}.html",
		},
		PageNameFishingSport: {
			H5:  "/#/fishingsport/pages/detail/index",
			Web: "/fish_sports_detail_{{.id}}.html",
		},
		PageNameRing: {
			H5:  "/#/sns/pages/ring_article/detail/index",
			Web: "/ring_{{.id}}.html",
		},
	}
	MapPageUserShop = map[string]PageUrl{
		UserShopInfo: {
			H5: "/shop/info",
		},
		UserShopHome: {
			H5: "/",
		},
	}
)

type (
	PageUrl struct {
		H5  string `json:"h_5"`
		Web string `json:"web"`
	}
	GetPagePathHandler func(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string)

	ArgGetLinks struct {
		List  []*ArgGetLinksItem `json:"list"`
		MapPk map[string]bool    `json:"-"` //用于PK 去重之用
	}
	ArgGetLinksItem struct {
		Pk             string                 `json:"pk" form:"pk"`
		HeaderInfo     *common.HeaderInfo     `json:"header_info,omitempty"`
		UrlValue       *url.Values            `json:"url_value,omitempty" form:"url_value"`
		UrlLinkVal     map[string]interface{} `json:"url_link_val,omitempty"` //url链接中的参数
		DataType       string                 `json:"data_type,omitempty" form:"data_type"`
		PageName       string                 `json:"page_name,omitempty" form:"page_name"`
		NeedHeaderInfo bool                   `json:"need_header_info,omitempty"` //拼接参数时，带上header_info数据
		OpenTarget     string                 `json:"open_target,omitempty"`      //打开界面的方式
		LinkTypIsURL   bool                   `json:"link_typ_is_url,omitempty"`  //返回的链接地址是字符串//
	}

	ResultGetLinks map[string]interface{}
)

//初始化一个获取链接的参数
func NewArgGetLinksItem(headerInfo *common.HeaderInfo, pageName, pk string) (res *ArgGetLinksItem) {
	res = &ArgGetLinksItem{
		Pk:         pk,
		HeaderInfo: headerInfo,
		UrlValue:   &url.Values{},
		UrlLinkVal: make(map[string]interface{}, 3),
		PageName:   pageName,
	}
	return
}

func (r *ArgGetLinks) Default(ctx *base.Context) (err error) {

	return
}

func (r *ArgGetLinks) AppendLinksItem(item *ArgGetLinksItem) (err error) {
	if _, ok := r.MapPk[item.Pk]; ok {
		return
	}
	r.MapPk[item.Pk] = true
	r.List = append(r.List, item)
	return
}

//获取链接地址
func GetLinks(args ArgGetLinks, ctx *base.Context) (res ResultGetLinks, err error) {
	res = make(map[string]interface{}, len(args.List))
	var (
		ro = rpc.RequestOptions{
			Method:      http.MethodPost,
			AppName:     app_param.AppNameAdmin,
			URI:         "/get_links",
			Header:      http.Header{},
			Value:       url.Values{},
			Context:     ctx,
			PathVersion: app_obj.App.AppRouterPrefix.Intranet,
		}
		data = struct {
			Code int            `json:"code"`
			Data ResultGetLinks `json:"data"`
			Msg  string         `json:"message"`
		}{}
	)
	ro.BodyJson, _ = json.Marshal(args)
	if err = rpc.NewHttpRpc(&ro).Send().GetBody().Bind(&data).Error; err != nil {
		return
	}
	res = data.Data
	return
}

//
//func getPageSpuPathByPageName(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string) {
//	var pageName = PageNameSpu
//	if len(pageNames) > 0 {
//		pageName = pageNames[0]
//	}
//
//	if tmp, ok := MapPageMallName[pageName]; ok {
//		res = getLinkValue(terminal, urlLinkVal, tmp)
//		return
//	}
//	return
//}

//func getLinkValue(terminal string, urlLinkVal map[string]interface{}, pageUrl PageUrl) (res string) {
//	switch terminal {
//	case app_param.TerminalWeb:
//		res = pageUrl.Web
//	case app_param.TerminalMina:
//		res = pageUrl.H5
//	default:
//		res = pageUrl.H5
//	}
//	if res != "" && urlLinkVal != nil {
//		tmpl, _ := template.New("test").Parse(res)
//		var result bytes.Buffer
//		_ = tmpl.Execute(&result, urlLinkVal)
//		res = result.String()
//	}
//	return
//}

//func getPageUserShopPathByPageName(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string) {
//	var pageName = UserShopInfo
//	if len(pageNames) > 0 {
//		pageName = pageNames[0]
//	}
//
//	if tmp, ok := MapPageUserShop[pageName]; ok {
//		res = getLinkValue(terminal, urlLinkVal, tmp)
//		return
//	}
//	return
//}
//func getPageShopPathByPageName(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string) {
//	var pageName = UserShopHome
//	if len(pageNames) > 0 {
//		pageName = pageNames[0]
//	}
//
//	if tmp, ok := MapPageUserShop[pageName]; ok {
//		res = getLinkValue(terminal, urlLinkVal, tmp)
//		return
//	}
//	return
//}

//func getPageSNSPathByPageName(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string) {
//	var pageName = PageNameSns
//	if len(pageNames) > 0 {
//		pageName = pageNames[0]
//	}
//
//	if tmp, ok := MapPageSNsName[pageName]; ok {
//		res = getLinkValue(terminal, urlLinkVal, tmp)
//		return
//	}
//	return
//}
//
//func getPageFishingSpotsPathByPageName(terminal string, urlLinkVal map[string]interface{}, pageNames ...string) (res string) {
//	var pageName = PageNameFishingSport
//	if len(pageNames) > 0 {
//		pageName = pageNames[0]
//	}
//
//	if tmp, ok := MapPageSNsName[pageName]; ok {
//		res = getLinkValue(terminal, urlLinkVal, tmp)
//		return
//	}
//	return
//}

//argument.UrlValue, argument.DataType, argument.PageName...
//func getPageLink(getPagePathHandler GetPagePathHandler, argument *LinkArgument) (res string, err error) {
//	var (
//		stringValue  string
//		urlValue1    = url.Values{}
//		paramsDivide = "?"
//		dtp          string
//	)
//
//	if tmp, ok := MapDataTypeBiz[argument.DataType]; ok {
//		argument.DataType = tmp
//	}
//	if getPagePathHandler != nil {
//		suffix := getPagePathHandler(argument.HeaderInfo.HTerminal, argument.UrlLinkVal, argument.PageName)
//		if terminalInfo, ok := plugins_lib.WebMap[argument.HeaderInfo.HTerminal]; ok {
//			if tmp, ok := terminalInfo[argument.DataType]; ok {
//				stringValue = fmt.Sprintf("//%s%s", tmp, suffix)
//			} else {
//				stringValue = fmt.Sprintf("//localhost:3000%s", suffix)
//				dtp = "1"
//			}
//		} else {
//			stringValue = fmt.Sprintf("//localhost:3000%s", suffix)
//			dtp = "2"
//		}
//	}
//	dataSlice := strings.Split(stringValue, paramsDivide)
//	l := len(dataSlice)
//	if l > 1 {
//		if urlValue1, err = url.ParseQuery(dataSlice[l-1]); err != nil {
//			return
//		}
//		preUrl := dataSlice[0 : l-1]
//		stringValue = strings.Join(preUrl, paramsDivide)
//		if urlValue1 != nil {
//			for key, value := range urlValue1 {
//				argument.UrlValue.Set(key, strings.Join(value, ""))
//			}
//		}
//
//	}
//	if dtp != "" {
//		argument.UrlValue.Set("ldtp", dtp)
//	}
//	if len(*argument.UrlValue) != 0 {
//		res = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, argument.UrlValue.Encode())
//	} else {
//		res = stringValue
//	}
//	return
//}
//
//func getPageLinkDefault(argument *LinkArgument) (res string, err error) {
//
//	var (
//		mapGetPagePath = map[string]GetPagePathHandler{
//			AdDataDataTypeSpu:               getPageSpuPathByPageName,
//			AdDataDataTypeSku:               getPageSpuPathByPageName,
//			AdDataDataTypeUserShop:          getPageUserShopPathByPageName,
//			AdDataDataTypeUserShopHome:      getPageShopPathByPageName,
//			AdDataDataTypeUser:              getPageSpuPathByPageName,
//			AdDataDataTypeSocialIntercourse: getPageSNSPathByPageName,
//			AdDataDataTypeSpuCategory:       getPageSNSPathByPageName,
//			AdDataDataTypeFishingSport:      getPageFishingSpotsPathByPageName,
//			AdDataDataTypeOther:             getPageSpuPathByPageName,
//		}
//
//		ok      bool
//		handler GetPagePathHandler
//	)
//	if argument.DataType != AdDataDataTypeOther {
//		if handler, ok = mapGetPagePath[argument.DataType]; !ok {
//			err = fmt.Errorf("对不起,系统当前暂不支持生成您的数据类型(%s)", argument.DataType)
//			return
//		}
//	}
//	if res, err = getPageLink(handler, argument); err != nil {
//		return
//	}
//	return
//}
//
////小程序参数生成
//func getPageLinkMina(argument *LinkArgument) (res DataItemLinkMina, err error) {
//	var (
//		pageName string
//		ok       bool
//	)
//
//	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
//		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
//		return
//	}
//
//	if argument.PageName != "" {
//		pageName = argument.PageName
//	}
//	res.PageName = pageName
//	res.Query = make(map[string]interface{}, 10)
//	if argument.NeedHeaderInfo {
//		if argument.HeaderInfo.HApp != "" {
//			res.Query["h_app"] = argument.HeaderInfo.HApp
//		}
//		if argument.HeaderInfo.HTerminal != "" {
//			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
//		}
//		if argument.HeaderInfo.HChannel != "" {
//			res.Query["h_channel"] = argument.HeaderInfo.HChannel
//		}
//		if argument.HeaderInfo.HVersion != "" {
//			res.Query["h_version"] = argument.HeaderInfo.HVersion
//		}
//	}
//	if argument.UrlValue != nil {
//		for key := range * argument.UrlValue {
//			res.Query[key] = argument.UrlValue.Get(key)
//		}
//	}
//
//	return
//}
//
//func getPageH5Mina(argument *LinkArgument) (res DataItemLinkMina, err error) {
//	var (
//		pageName string
//		ok       bool
//	)
//
//	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
//		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
//		return
//	}
//
//	if argument.PageName != "" {
//		pageName = argument.PageName
//	}
//	res.PageName = pageName
//	res.Query = make(map[string]interface{}, 10)
//	if argument.NeedHeaderInfo {
//		if argument.HeaderInfo.HApp != "" {
//			res.Query["h_app"] = argument.HeaderInfo.HApp
//		}
//		if argument.HeaderInfo.HTerminal != "" {
//			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
//		}
//		if argument.HeaderInfo.HChannel != "" {
//			res.Query["h_channel"] = argument.HeaderInfo.HChannel
//		}
//		if argument.HeaderInfo.HVersion != "" {
//			res.Query["h_version"] = argument.HeaderInfo.HVersion
//		}
//	}
//	if argument.UrlValue != nil {
//		for key := range * argument.UrlValue {
//			res.Query[key] = argument.UrlValue.Get(key)
//		}
//	}
//
//	return
//}
//
////网站链接
//func getPageLinkWeb(argument *LinkArgument) (res interface{}, err error) {
//	if argument.NeedHeaderInfo {
//		if argument.UrlValue == nil {
//			argument.UrlValue = &url.Values{}
//		}
//		if argument.HeaderInfo.HApp != "" {
//			argument.UrlValue.Set("h_app", argument.HeaderInfo.HApp)
//		}
//		if argument.HeaderInfo.HTerminal != "" {
//			argument.UrlValue.Set("h_terminal", argument.HeaderInfo.HTerminal)
//		}
//		if argument.HeaderInfo.HChannel != "" {
//			argument.UrlValue.Set("h_channel", argument.HeaderInfo.HChannel)
//		}
//		if argument.HeaderInfo.HVersion != "" {
//			argument.UrlValue.Set("h_version", argument.HeaderInfo.HVersion)
//		}
//	}
//
//	res, err = getPageWebLinkDefault(argument)
//	return
//}
//
////网站链接
//func getPageWebLinkDefault(argument *LinkArgument) (res string, err error) {
//	var (
//		mapGetPagePath = map[string]GetPagePathHandler{
//			AdDataDataTypeSpu:               getPageSpuPathByPageName,
//			AdDataDataTypeSku:               getPageSpuPathByPageName,
//			AdDataDataTypeUserShop:          getPageUserShopPathByPageName,
//			AdDataDataTypeUserShopHome:      getPageShopPathByPageName,
//			AdDataDataTypeUser:              getPageSpuPathByPageName,
//			AdDataDataTypeSocialIntercourse: getPageSNSPathByPageName,
//			AdDataDataTypeSpuCategory:       getPageSNSPathByPageName,
//			AdDataDataTypeFishingSport:      getPageFishingSpotsPathByPageName,
//			AdDataDataTypeRing:              getPageSNSPathByPageName,
//			AdDataDataTypeOther:             getPageSpuPathByPageName,
//		}
//
//		ok      bool
//		handler GetPagePathHandler
//	)
//	if argument.DataType != AdDataDataTypeOther {
//		if handler, ok = mapGetPagePath[argument.DataType]; !ok {
//			err = fmt.Errorf("对不起,系统当前暂不支持生成您的数据类型(%s)", argument.DataType)
//			return
//		}
//	}
//	if res, err = getPageLink(handler, argument); err != nil {
//		return
//	}
//	return
//}
//
//func getDefault(argument *LinkArgument) (res interface{}, err error) {
//	if argument.NeedHeaderInfo {
//		if argument.UrlValue == nil {
//			argument.UrlValue = &url.Values{}
//		}
//		if argument.HeaderInfo.HApp != "" {
//			argument.UrlValue.Set("h_app", argument.HeaderInfo.HApp)
//		}
//		if argument.HeaderInfo.HTerminal != "" {
//			argument.UrlValue.Set("h_terminal", argument.HeaderInfo.HTerminal)
//		}
//		if argument.HeaderInfo.HChannel != "" {
//			argument.UrlValue.Set("h_channel", argument.HeaderInfo.HChannel)
//		}
//		if argument.HeaderInfo.HVersion != "" {
//			argument.UrlValue.Set("h_version", argument.HeaderInfo.HVersion)
//		}
//	}
//
//	res, err = getPageLinkDefault(argument)
//	return
//}
//
////小程序参数生成
//func getPageLinkApp(argument *LinkArgument) (res DataItemLinkMina, err error) {
//	var (
//		pageName string
//		ok       bool
//	)
//
//	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
//		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
//		return
//	}
//
//	if len(argument.PageName) > 0 {
//		pageName = argument.PageName
//	}
//	res.PageName = pageName
//	res.Query = make(map[string]interface{}, 10)
//	if argument.NeedHeaderInfo {
//		if argument.HeaderInfo.HApp != "" {
//			res.Query["h_app"] = argument.HeaderInfo.HApp
//		}
//		if argument.HeaderInfo.HTerminal != "" {
//			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
//		}
//		if argument.HeaderInfo.HChannel != "" {
//			res.Query["h_channel"] = argument.HeaderInfo.HChannel
//		}
//		if argument.HeaderInfo.HVersion != "" {
//			res.Query["h_version"] = argument.HeaderInfo.HVersion
//		}
//	}
//	if argument.UrlValue != nil {
//		for key := range *argument.UrlValue {
//			res.Query[key] = argument.UrlValue.Get(key)
//		}
//	}
//	return
//}

//获取页面链接
//headerInfo *common.HeaderInfo, urlValue *url.Values, dataType string, pageNames ...string
//func GetPageLink(argument *LinkArgument) (res interface{}, err error) {
//
//	type GetPageLinkHandler = func(argument *LinkArgument) (res DataItemLinkMina, err error)
//	switch argument.HeaderInfo.HTerminal {
//	case app_param.TerminalMina, app_param.TerminalH5, app_param.TerminalAndroid, app_param.TerminalIos:
//		//如果返回的为url连接地址
//		if argument.LinkTypIsURL {
//			res, err = getDefault(argument)
//			return
//		}
//		var getLinkMap = map[string]GetPageLinkHandler{
//			app_param.TerminalMina:    getPageLinkMina, //小程序
//			app_param.TerminalH5:      getPageH5Mina,   //H5页面操作使用
//			app_param.TerminalAndroid: getPageLinkApp,  //安卓
//			app_param.TerminalIos:     getPageLinkApp,  //IOS
//		}
//		if handler, ok := getLinkMap[argument.HeaderInfo.HTerminal]; ok {
//			res, err = handler(argument)
//			return
//		}
//	case app_param.TerminalWeb: //网站链接
//		res, err = getPageLinkWeb(argument)
//	default:
//		//如果返回的为url连接地址
//		if argument.LinkTypIsURL {
//			res, err = getDefault(argument)
//			return
//		}
//		res, err = getDefault(argument)
//	}
//	return
//}
//
//func GetUserHref(info *common.HeaderInfo, urlV url.Values) (link interface{}, err error) {
//
//	var (
//		linkArgument *LinkArgument
//	)
//
//	switch info.HTerminal {
//	case app_param.TerminalWeb:
//		uid := urlV.Get("uid")
//		linkArgument = &LinkArgument{
//			HeaderInfo: info,
//			UrlValue:   &urlV,
//			DataType:   AdDataDataTypeUser,
//			PageName:   PageNameUsr,
//		}
//		urlV.Del("uid")
//		if linkArgument.UrlLinkVal == nil {
//			linkArgument.UrlLinkVal = make(map[string]interface{}, 5)
//		}
//		linkArgument.UrlLinkVal["user_id"] = uid
//	default:
//		linkArgument = &LinkArgument{
//			HeaderInfo: info,
//			UrlValue:   &urlV,
//			DataType:   AdDataDataTypeUser,
//			PageName:   PageNameUsr,
//		}
//	}
//	link, err = GetPageLink(linkArgument)
//	return
//}
