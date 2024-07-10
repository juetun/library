package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"strings"
)

const (
	FullTextIndexPrefix string = "jt"
	FullTextSpu                = "spu"       //商品
	FullTextFishSport          = "fish_spot" //钓点
	FullTextSns                = "sns"       //动态
	FullTextShop               = "shop"      //店铺
	FullTextUser               = "user"      //用户
)

func GetEsIndex(indexName string) (res string) {
	return gerPreString() + indexName
}

func GetEsRealIndex(indexName string) (res string) {
	return strings.TrimLeft(indexName, gerPreString())
}

func gerPreString() (res string) {
	res = fmt.Sprintf("%v%v_", FullTextIndexPrefix, app_obj.App.AppEnv)
	return
}
