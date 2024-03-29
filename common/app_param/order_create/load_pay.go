package order_create

import (
	"github.com/juetun/base-wrapper/lib/app/app_start"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

var (
	// 数据库配置信息存储对象
	ConfigPay SystemPayConfig
)

type (
	SystemPayConfig struct {
		Pay SystemPayConfigPay `json:"pay" yaml:"pay"`
	}
	SystemPayConfigPay struct {
		AliPay    SystemPayConfigAliPay    `json:"alipay" yaml:"alipay"`
		WeiXinPay SystemPayConfigWeiXinPay `json:"weixin" yaml:"weixin"`
	}
	SystemPayConfigAliPay struct {
		AppId     string `json:"app_id" yaml:"app_id"`
		Secret    string `json:"secret" yaml:"secret"`
		FlatRabat string `json:"flat_rabat" yaml:"flat_rabat"` //支付平台手续费率
	}
	SystemPayConfigWeiXinPay struct {
		AppId     string `json:"app_id" yaml:"app_id"`
		Secret    string `json:"secret" yaml:"secret"`
		FlatRabat string `json:"flat_rabat" yaml:"flat_rabat"` //支付平台手续费率
	}
)

var io = base.NewSystemOut().
	SetInfoType(base.LogLevelInfo)

func LoadPay(arg *app_start.PluginsOperate) (err error) {

	var syncLock sync.Mutex
	syncLock.Lock()
	defer syncLock.Unlock()
	io.SystemOutPrintln("init pay parameters config")
	defer io.SystemOutPrintln("Init  pay parameters finished")

	var yamlFile []byte
	if yamlFile, err = os.ReadFile(common.GetConfigFilePath("pay.yaml")); err != nil {
		io.SetInfoType(base.LogLevelFatal).SystemOutFatalf("yamlFile.Get err   #%v \n", err)
	}
	if err = yaml.Unmarshal(yamlFile, &ConfigPay); err != nil {
		io.SetInfoType(base.LogLevelFatal).SystemOutFatalf("Fatal error pay parameters file(%#v) \n", err)
	}
	io.SetInfoType(base.LogLevelInfo).SystemOutPrintf("Load pay parameters config is : '%#v' ", ConfigPay)
	return
}
