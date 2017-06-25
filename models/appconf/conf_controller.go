package appconf

import (
	"fmt"

	_c "github.com/astaxie/beego/config"
	_L "github.com/ebitgo/quotes/models/log"
)

// ConfigInstance 配置文件唯一实例
var ConfigInstance *ConfigController

// ConfigController 配置文件读取控制器
type ConfigController struct {
	appConfig _c.Configer
}

// NewConfigController new ConfigController
func NewConfigController() *ConfigController {
	ret := new(ConfigController)
	return ret.Init()
}

// Init 初始化
func (ths *ConfigController) Init() *ConfigController {
	var err error
	ths.appConfig, err = _c.NewConfig("ini", "conf/coin_app.conf")
	if err != nil {
		_L.LoggerInstance.ErrorPrint("Read 'coin_app.conf' has error : \r\n%v\r\n", err)
		panic(err.Error())
	}
	return ths
}

// GetBtc38Config 获取btc38基础货币和虚拟币配置
func (ths *ConfigController) GetBtc38Config() (ret [][]string) {
	count := ths.appConfig.DefaultInt("btc38::count", 0)
	ret = make([][]string, 0)
	for i := 0; i < count; i++ {
		itm := make([]string, 2)
		itm[0] = ths.appConfig.String(fmt.Sprintf("btc38::monetary_%d", i+1))
		itm[1] = ths.appConfig.String(fmt.Sprintf("btc38::coin_%d", i+1))
		ret = append(ret, itm)
	}
	return
}

// GetPoloniexConfig 获取poloniex基础货币和虚拟币配置
func (ths *ConfigController) GetPoloniexConfig() (ret [][]string) {
	count := ths.appConfig.DefaultInt("poloniex::count", 0)
	ret = make([][]string, 0)
	for i := 0; i < count; i++ {
		itm := make([]string, 2)
		itm[0] = ths.appConfig.String(fmt.Sprintf("poloniex::monetary_%d", i+1))
		itm[1] = ths.appConfig.String(fmt.Sprintf("poloniex::coin_%d", i+1))
		ret = append(ret, itm)
	}
	return
}

// GetYuanbaoConfig 获取yuanbao基础货币和虚拟币配置
func (ths *ConfigController) GetYuanbaoConfig() (ret [][]string) {
	count := ths.appConfig.DefaultInt("yuanbao::count", 0)
	ret = make([][]string, 0)
	for i := 0; i < count; i++ {
		itm := make([]string, 2)
		itm[0] = ths.appConfig.String(fmt.Sprintf("yuanbao::monetary_%d", i+1))
		itm[1] = ths.appConfig.String(fmt.Sprintf("yuanbao::coin_%d", i+1))
		ret = append(ret, itm)
	}
	return
}

// GetProxyIP 获取代理IP配置
func (ths *ConfigController) GetProxyIP() string {
	return ths.appConfig.DefaultString("proxy::ip", "")
}

// GetProxyPort 获取代理Port配置
func (ths *ConfigController) GetProxyPort() string {
	return ths.appConfig.DefaultString("proxy::port", "")
}
