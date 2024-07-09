package app_param

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"strings"
)

const IndexPrefix string = "jt"

func GetEsIndex(indexName string) (res string) {
	return gerPreString() + indexName
}

func GetEsRealIndex(indexName string) (res string) {
	return strings.TrimLeft(indexName, gerPreString())
}

func gerPreString() (res string) {
	res = fmt.Sprintf("%v%v_", IndexPrefix, app_obj.App.AppEnv)
	return
}
