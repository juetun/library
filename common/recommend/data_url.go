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
		var stringValue string
		if tmp, ok := plugins_lib.WebMap[dataType]; ok {
			stringValue = fmt.Sprintf("//%s/#/pages/mall/detail/index", tmp)
		} else {
			stringValue = fmt.Sprintf("//127.0.0.1/#/pages/mall/detail/index")
		}
		var (
			urlValue1    = url.Values{}
			paramsDivide = "?"
			dataSlice    = strings.Split(stringValue, paramsDivide)
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
		stringValue = fmt.Sprintf("%s%s%s", stringValue, paramsDivide, urlValue.Encode())
		return
	case AdDataDataTypeSocialIntercourse: //广告商品信息
	default:
		err = fmt.Errorf("")
	}
	return
}
