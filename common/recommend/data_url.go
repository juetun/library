package recommend

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/plugins_lib"
	"net/url"
	"strings"
)

const (
	PageNameHome = "home" //首页

	//商品前台界面
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
	MapPageMallName = map[string]string{
		PageNameSpu:                     "/#/pages/mall/detail/index",
		PageNameShop:                    "/#/pages/mall/shop/home/index",
		PageNameUsr:                     "/#/pages/user/view/index",
		AdDataDataTypeSocialIntercourse: "/#/pages/sns/detail/index",
		AdDataDataTypeRing:              "/#/pages/sns/ring_article/index", //圈子界面
		AdDataDataTypeFishingSport:      "/#/pages/fishingsport/detail/index",
		PageNamePayFinish:               "/#/pages/mall/order/pay_finish/index", //支付结果界面
		PageNamePayPreview:              "/#/pages/mall/order/preview/index",    //支付预付界面
	}
	MapPageSNsName = map[string]string{
		PageNameHome:         "/#/pages/home/index/index",
		PageNameSns:          "/#/pages/sns/detail/index",
		PageNameFishingSport: "/#/pages/fishingsport/detail/index",
		PageNameRing:         "/#/pages/ring_article/detail/index",
	}
	MapPageUserShop = map[string]string{
		UserShopInfo: "/shop/info",
		UserShopHome: "/",
	}
)

type (
	GetPagePathHandler func(pageNames ...string) (res string)
)

func getPageSpuPathByPageName(pageNames ...string) (res string) {
	var pageName = PageNameSpu
	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}

	if tmp, ok := MapPageMallName[pageName]; ok {
		res = tmp
		return
	}
	return
}

func getPageUserShopPathByPageName(pageNames ...string) (res string) {
	var pageName = UserShopInfo
	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}

	if tmp, ok := MapPageUserShop[pageName]; ok {
		res = tmp
		return
	}
	return
}
func getPageShopPathByPageName(pageNames ...string) (res string) {
	var pageName = UserShopHome
	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}

	if tmp, ok := MapPageUserShop[pageName]; ok {
		res = tmp
		return
	}
	return
}

func getPageSNSPathByPageName(pageNames ...string) (res string) {
	var pageName = PageNameSns
	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}

	if tmp, ok := MapPageSNsName[pageName]; ok {
		res = tmp
		return
	}
	return
}

func getPageFishingSpotsPathByPageName(pageNames ...string) (res string) {
	var pageName = PageNameFishingSport
	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}

	if tmp, ok := MapPageSNsName[pageName]; ok {
		res = tmp
		return
	}
	return
}

func getPageLink(getPagePathHandler GetPagePathHandler, urlValue *url.Values, dataType string, pageNames ...string) (res string, err error) {
	var (
		stringValue  string
		urlValue1    = url.Values{}
		paramsDivide = "?"
	)

	if tmp, ok := MapDataTypeBiz[dataType]; ok {
		dataType = tmp
	}
	if getPagePathHandler != nil {
		suffix := getPagePathHandler(pageNames...)
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s%s", tmp, suffix)
		} else {
			stringValue = fmt.Sprintf("//localhost:3000%s", suffix)
		}
	}
	var (
		dataSlice = strings.Split(stringValue, paramsDivide)
	)
	var l = len(dataSlice)
	if l > 1 {
		if urlValue1, err = url.ParseQuery(dataSlice[l-1]); err != nil {
			return
		}
		preUrl := dataSlice[0 : l-1]
		stringValue = strings.Join(preUrl, paramsDivide)
		if urlValue1 != nil {
			for key, value := range urlValue1 {
				urlValue.Set(key, strings.Join(value, ""))
			}
		}

	}
	res = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, urlValue.Encode())
	return
}

func getPageLinkDefault(urlValue *url.Values, dataType string, pageNames ...string) (res string, err error) {
	var (
		mapGetPagePath = map[string]GetPagePathHandler{
			AdDataDataTypeSpu:               getPageSpuPathByPageName,
			AdDataDataTypeSku:               getPageSpuPathByPageName,
			AdDataDataTypeUserShop:          getPageUserShopPathByPageName,
			AdDataDataTypeUserShopHome:      getPageShopPathByPageName,
			AdDataDataTypeUser:              getPageSpuPathByPageName,
			AdDataDataTypeSocialIntercourse: getPageSNSPathByPageName,
			AdDataDataTypeSpuCategory:       getPageSNSPathByPageName,
			AdDataDataTypeFishingSport:      getPageFishingSpotsPathByPageName,
			AdDataDataTypeOther:             getPageSpuPathByPageName,
		}

		ok      bool
		handler GetPagePathHandler
	)
	if dataType != AdDataDataTypeOther {
		if handler, ok = mapGetPagePath[dataType]; !ok {
			err = fmt.Errorf("对不起,系统当前暂不支持生成您的数据类型(%s)", dataType)
			return
		}
	}
	if res, err = getPageLink(handler, urlValue, dataType, pageNames...); err != nil {
		return
	}
	return
}

//小程序参数生成
func getPageLinkMina(argument *LinkArgument) (res DataItemLinkMina, err error) {
	var (
		pageName string
		ok       bool
	)

	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
		return
	}

	if argument.PageName != "" {
		pageName = argument.PageName
	}
	res.PageName = pageName
	res.Query = make(map[string]interface{}, 10)
	if argument.NeedHeaderInfo {
		if argument.HeaderInfo.HApp != "" {
			res.Query["h_app"] = argument.HeaderInfo.HApp
		}
		if argument.HeaderInfo.HTerminal != "" {
			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
		}
		if argument.HeaderInfo.HChannel != "" {
			res.Query["h_channel"] = argument.HeaderInfo.HChannel
		}
		if argument.HeaderInfo.HVersion != "" {
			res.Query["h_version"] = argument.HeaderInfo.HVersion
		}
	}
	if argument.UrlValue != nil {
		for key := range * argument.UrlValue {
			res.Query[key] = argument.UrlValue.Get(key)
		}
	}

	return
}

func getPageH5Mina(argument *LinkArgument) (res DataItemLinkMina, err error) {
	var (
		pageName string
		ok       bool
	)

	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
		return
	}

	if argument.PageName != "" {
		pageName = argument.PageName
	}
	res.PageName = pageName
	res.Query = make(map[string]interface{}, 10)
	if argument.NeedHeaderInfo {
		if argument.HeaderInfo.HApp != "" {
			res.Query["h_app"] = argument.HeaderInfo.HApp
		}
		if argument.HeaderInfo.HTerminal != "" {
			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
		}
		if argument.HeaderInfo.HChannel != "" {
			res.Query["h_channel"] = argument.HeaderInfo.HChannel
		}
		if argument.HeaderInfo.HVersion != "" {
			res.Query["h_version"] = argument.HeaderInfo.HVersion
		}
	}
	if argument.UrlValue != nil {
		for key := range * argument.UrlValue {
			res.Query[key] = argument.UrlValue.Get(key)
		}
	}

	return
}

func getDefault(argument *LinkArgument) (res interface{}, err error) {
	if argument.NeedHeaderInfo {
		if argument.UrlValue == nil {
			argument.UrlValue = &url.Values{}
		}
		if argument.HeaderInfo.HApp != "" {
			argument.UrlValue.Set("h_app", argument.HeaderInfo.HApp)
		}
		if argument.HeaderInfo.HTerminal != "" {
			argument.UrlValue.Set("h_terminal", argument.HeaderInfo.HTerminal)
		}
		if argument.HeaderInfo.HChannel != "" {
			argument.UrlValue.Set("h_channel", argument.HeaderInfo.HChannel)
		}
		if argument.HeaderInfo.HVersion != "" {
			argument.UrlValue.Set("h_version", argument.HeaderInfo.HVersion)
		}
	}

	res, err = getPageLinkDefault(argument.UrlValue, argument.DataType, argument.PageName)
	return
}

//小程序参数生成
func getPageLinkApp(argument *LinkArgument) (res DataItemLinkMina, err error) {
	var (
		pageName string
		ok       bool
	)

	if pageName, ok = MapDataTypeBiz[argument.DataType]; !ok {
		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", argument.DataType)
		return
	}

	if len(argument.PageName) > 0 {
		pageName = argument.PageName
	}
	res.PageName = pageName
	res.Query = make(map[string]interface{}, 10)
	if argument.NeedHeaderInfo {
		if argument.HeaderInfo.HApp != "" {
			res.Query["h_app"] = argument.HeaderInfo.HApp
		}
		if argument.HeaderInfo.HTerminal != "" {
			res.Query["h_terminal"] = argument.HeaderInfo.HTerminal
		}
		if argument.HeaderInfo.HChannel != "" {
			res.Query["h_channel"] = argument.HeaderInfo.HChannel
		}
		if argument.HeaderInfo.HVersion != "" {
			res.Query["h_version"] = argument.HeaderInfo.HVersion
		}
	}
	if argument.UrlValue != nil {
		for key := range *argument.UrlValue {
			res.Query[key] = argument.UrlValue.Get(key)
		}
	}
	return
}

type LinkArgument struct {
	HeaderInfo     *common.HeaderInfo
	UrlValue       *url.Values
	DataType       string
	PageName       string
	NeedHeaderInfo bool `json:"need_header_info"` //拼接参数时，带上header_info数据
	LinkTypIsURL   bool `json:"link_typ_is_url"`  //返回的链接地址是字符串//
}

//获取页面链接
//headerInfo *common.HeaderInfo, urlValue *url.Values, dataType string, pageNames ...string
func GetPageLink(argument *LinkArgument) (res interface{}, err error) {
	//如果返回的为url连接地址
	if argument.LinkTypIsURL {
		res, err = getDefault(argument)
		return
	}
	type GetPageLinkHandler = func(argument *LinkArgument) (res DataItemLinkMina, err error)
	var getLinkMap = map[string]GetPageLinkHandler{
		app_param.TerminalMina:    getPageLinkMina, //小程序
		app_param.TerminalH5:      getPageH5Mina,   //H5页面操作使用
		app_param.TerminalAndroid: getPageLinkApp,  //安卓
		app_param.TerminalIos:     getPageLinkApp,  //IOS
	}

	if handler, ok := getLinkMap[argument.HeaderInfo.HTerminal]; ok {
		res, err = handler(argument)
		return
	}
	res, err = getDefault(argument)
	return
}

func GetUserHref(info *common.HeaderInfo, urlV url.Values) (link interface{}, err error) {
	//urlV := url.Values{}
	//urlV.Set("uid", fmt.Sprintf("%v", r.ToUserHid))
	link, err = GetPageLink(&LinkArgument{
		HeaderInfo: info,
		UrlValue:   &urlV,
		DataType:   AdDataDataTypeUser,
		PageName:   PageNameUsr,
	})
	return
}
