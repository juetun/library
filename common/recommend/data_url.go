package recommend

import (
	"fmt"
	"github.com/juetun/library/common/plugins_lib"
	"net/url"
	"strings"
)

//获取页面链接
func GetPageLink(urlValue *url.Values, dataType string) (res string, err error) {

	switch dataType {
	case AdDataDataTypeSpu: //广告商品信息
		var (
			stringValue  string
			urlValue1    = url.Values{}
			paramsDivide = "?"
		)
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s/#/pages/mall/detail/index", tmp)
		} else {
			stringValue = fmt.Sprintf("//127.0.0.1:10086/#/pages/mall/detail/index")
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
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s/#/pages/sns/detail/index", tmp)
		} else {
			stringValue = fmt.Sprintf("//127.0.0.1:10086/#/pages/sns/detail/index")
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
