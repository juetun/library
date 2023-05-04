package recommend

import (
	"fmt"
	"github.com/juetun/library/common/plugins_lib"
	"net/url"
	"strings"
)

const (
	//商品前台界面
	PageNameSpu  = "spu"
	PageNameShop = "shop"
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

var MapDataTypeBiz = map[string]string{
	AdDataDataTypeSpu:               PageNameSpu, //商品信息
	AdDataDataTypeSocialIntercourse: PageNameSns, //广告社交动态信息
}

var (
	MapPageMallName = map[string]string{
		PageNameSpu:  "/#/pages/mall/detail/index",
		PageNameShop: "/#/pages/mall/shop/home/index",
	}
	MapPageSNsName  = map[string]string{PageNameSns: "/#/pages/sns/detail/index",}
	MapPageUserShop = map[string]string{
		UserShopInfo: "/shop/info",
		UserShopHome: "/",
	}
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

//获取页面链接
func GetPageLink(urlValue *url.Values, dataType string, pageNames ...string) (res string, err error) {

	switch dataType {
	case AdDataDataTypeUserShop:
		var (
			stringValue  string
			urlValue1    = url.Values{}
			paramsDivide = "?"
		)

		if tmp, ok := MapDataTypeBiz[dataType]; ok {
			dataType = tmp
		}
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s%s", tmp, getPageUserShopPathByPageName(pageNames...))
		} else {
			stringValue = fmt.Sprintf("//localhost:3000%s", getPageUserShopPathByPageName(pageNames...))
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
			for key, value := range urlValue1 {
				urlValue.Set(key, strings.Join(value, ""))
			}
		}
		res = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, urlValue.Encode())
		return
	case AdDataDataTypeSpu: //广告商品信息
		var (
			stringValue  string
			urlValue1    = url.Values{}
			paramsDivide = "?"
		)
		if tmp, ok := MapDataTypeBiz[dataType]; ok {
			dataType = tmp
		}
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s%s", tmp, getPageSpuPathByPageName(pageNames...))
		} else {
			stringValue = fmt.Sprintf("//127.0.0.1:10086%s", getPageSpuPathByPageName(pageNames...))
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
			for key, value := range urlValue1 {
				urlValue.Set(key, strings.Join(value, ""))
			}
		}
		res = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, urlValue.Encode())
		return
	case AdDataDataTypeSocialIntercourse: //广告商品信息
		var (
			stringValue  string
			urlValue1    = url.Values{}
			paramsDivide = "?"
		)
		if tmp, ok := MapDataTypeBiz[dataType]; ok {
			dataType = tmp
		}
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s%s", tmp, getPageSNSPathByPageName(pageNames...))
		} else {
			stringValue = fmt.Sprintf("//127.0.0.1:10086%s", getPageSNSPathByPageName(pageNames...))
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
			for key, value := range urlValue1 {
				urlValue.Set(key, strings.Join(value, ""))
			}
		}
		res = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, urlValue.Encode())
		return
	default:
		err = fmt.Errorf("")
	}
	return
}
