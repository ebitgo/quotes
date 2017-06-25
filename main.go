package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	_AC "github.com/ebitgo/quotes/models/appconf"
	_C "github.com/ebitgo/quotes/models/coin"
	_L "github.com/ebitgo/quotes/models/log"
	_T "github.com/ebitgo/quotes/models/thread"
	_ "github.com/ebitgo/quotes/routers"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_L.LoggerInstance = _L.NewLoggerInstance(fmt.Sprintf("logs/coinreader.%s", time.Now().Format("2006-01-02_15.04.05.000")))
	_L.LoggerInstance.OpenDebug = true
	_L.LoggerInstance.SetLogFunCallDepth(4)

	_L.LoggerInstance.InfoPrint(" > Init Coin manager instance...\r\n")
	_C.CoinMangerInstance = _C.NewCoinManager()

	_L.LoggerInstance.InfoPrint(" > Init Appconfig instance...\r\n")
	_AC.ConfigInstance = _AC.NewConfigController()

	_L.LoggerInstance.InfoPrint(" > Init Thread manager instance...\r\n")
	_T.CheckManagerInstance = _T.NewCheckManager()
	_T.CheckManagerInstance.Check()
	beego.Run()
}
