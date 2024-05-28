package plugins_lib

import (
	"fmt"
	"os"
	"sync"

	"github.com/juetun/base-wrapper/lib/app/app_start"

	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"gopkg.in/yaml.v2"
)

//前端路径映射
var WebMap = map[string]string{}

func PluginWebMap(arg *app_start.PluginsOperate) (err error) {
	var syncLock sync.Mutex
	syncLock.Lock()
	defer syncLock.Unlock()
	io := base.NewSystemOut().SetInfoType(base.LogLevelInfo)
	io.SystemOutPrintf("Load webMap start")
	defer io.SetInfoType(base.LogLevelInfo).
		SystemOutPrintf(fmt.Sprintf("Load webMap config finished \n"))

	var yamlFile []byte
	var filePath = common.GetCommonConfigFilePath("webmap.yaml")
	if yamlFile, err = os.ReadFile(filePath); err != nil {
		io.SystemOutFatalf("yamlFile.Get err #%v \n", err)
	}
	if err = yaml.Unmarshal(yamlFile, &WebMap); err != nil {
		io.SystemOutFatalf("Load  webMap config failure(%#v) \n", err)
	}
	for key, value := range WebMap {
		io.SetInfoType(base.LogLevelInfo).
			SystemOutPrintf("【%s】 %#v", key, value)
	}
	return

}
