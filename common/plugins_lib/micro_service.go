package plugins_lib

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_start"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

//前端路径映射
var MicroService *MicroServiceConfig

type MicroServiceConfig struct {
	RegisterOpen   bool `json:"registeropen"`
	RegisterCenter bool `json:"registercenter"` //注册中心类型 ~~~  枚举 当前取值为 -consul
	Register       map[string]struct {
		RegisterAddr string `json:"registeraddr"` //consul主节点 "localhost:8500"多个用逗号隔开
	} `json:"register"`
}

func PluginMicroService(arg *app_start.PluginsOperate) (err error) {
	var syncLock sync.Mutex
	syncLock.Lock()
	defer syncLock.Unlock()
	io := base.NewSystemOut().SetInfoType(base.LogLevelInfo)
	io.SystemOutPrintf("Load MicroService start")
	configName := "MicroService"
	defer io.SetInfoType(base.LogLevelInfo).
		SystemOutPrintf(fmt.Sprintf("【%s】Load webMap config finished \n", configName))

	var yamlFile []byte
	var filePath = common.GetCommonConfigFilePath("micro_service.yml", true)
	if yamlFile, err = os.ReadFile(filePath); err != nil {
		io.SystemOutFatalf("【%s】yamlFile.Get MicroService err #%v \n", configName, err)
	}
	if err = yaml.Unmarshal(yamlFile, MicroService); err != nil {
		io.SystemOutFatalf("【%s】Load  MicroService config failure(%#v) \n", configName, err)
	}
	io.SetInfoType(base.LogLevelInfo).
		SystemOutPrintf("【%s】 %#v", configName, string(yamlFile))
	return

}
