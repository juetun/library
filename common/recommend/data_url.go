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
	//商品前台界面
	PageNameSpu  = "spu"
	PageNameShop = "shop"
	PageNameUsr  = "user"

	PageNameOther = "other" //无实际意义拼接数据使用
)

const (
	//社交前台界面
	PageNameSns = "sns"
)

const (
	//店铺后台界面名称
	UserShopInfo = "shop_info"
	UserShopHome = "shop_home"
)

var (
	MapDataTypeBiz = map[string]string{
		AdDataDataTypeSpu:               PageNameSpu,  //商品信息（spu）
		AdDataDataTypeSku:               PageNameSpu,  //商品(sku)信息
		AdDataDataTypeUserShop:          PageNameShop, //店铺信息
		AdDataDataTypeUser:              PageNameUsr,  //用户信息
		AdDataDataTypeSocialIntercourse: PageNameSns,  //广告社交动态信息

		AdDataDataTypeOther: PageNameOther, //其他信息
	}
	MapPageMallName = map[string]string{
		PageNameSpu:                     "/#/pages/mall/detail/index",
		PageNameShop:                    "/#/pages/mall/shop/home/index",
		PageNameUsr:                     "/#/pages/user/view/index",
		AdDataDataTypeSocialIntercourse: "/#/pages/sns/detail/index",
	}
	MapPageSNsName  = map[string]string{PageNameSns: "/#/pages/sns/detail/index",}
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
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s%s", tmp, getPagePathHandler(pageNames...))
		} else {
			stringValue = fmt.Sprintf("//localhost:3000%s", getPagePathHandler(pageNames...))
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
			AdDataDataTypeUser:              getPageSpuPathByPageName,
			AdDataDataTypeSocialIntercourse: getPageSNSPathByPageName,
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
func getPageLinkMina(urlValue *url.Values, dataType string, pageNames ...string) (res DataItemLinkMina, err error) {
	var (
		pageName string
		ok       bool
	)

	if pageName, ok = MapDataTypeBiz[dataType]; !ok {
		err = fmt.Errorf("系统暂不支持您选中的数据类型(%v)链接生成", dataType)
		return
	}

	if len(pageNames) > 0 {
		pageName = pageNames[0]
	}
	res.PageName = pageName
	res.Query = make(map[string]interface{}, 10)
	if urlValue != nil {
		for key := range *urlValue {
			res.Query[key] = urlValue.Get(key)
		}
	}

	return
}

//获取页面链接
func GetPageLink(headerInfo *common.HeaderInfo, urlValue *url.Values, dataType string, pageNames ...string) (res interface{}, err error) {
	switch headerInfo.HTerminal {
	case app_param.TerminalMina: //小程序
		res, err = getPageLinkMina(urlValue, dataType, pageNames...)
	default:
		//TerminalMina    = "mina"
		//TerminalH5      = "h5"
		//TerminalAndroid = "android"
		//TerminalIos     = "ios"
		res, err = getPageLinkDefault(urlValue, dataType, pageNames...)
	}

	return
}
