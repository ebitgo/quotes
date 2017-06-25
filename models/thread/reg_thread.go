package thread

import (
	_L "github.com/ebitgo/quotes/models/log"
	_ck "github.com/jojopoper/go-models/checker"
	_r "github.com/jojopoper/rhttp"
)

// PeriodReg 周期上报自己IP
type PeriodReg struct {
	_ck.CheckBase
}

// Init 初始化
func (ths *PeriodReg) Init(interval int) _ck.ICheckInterface {
	defer _L.LoggerInstance.InfoPrint("Period Reg init complete\n")
	ths.CheckBase.Init(interval)
	ths.SetName("PeriodReg")
	ths.SetExeFunc(ths.exe)
	return ths
}

func (ths *PeriodReg) exe() {
	_L.LoggerInstance.InfoPrint("Period reg to ebitgo server ...\n")
	h := new(_r.CHttp)
	_, err := h.Get("http://wechat.ebitgo.com/v/Nalsu354lsnd10lsk", _r.ReturnString)
	if err != nil {
		_L.LoggerInstance.ErrorPrint("Connect http://ebitgo.com/quotes has error\n%+v\n", err)
	}
	_L.LoggerInstance.InfoPrint("Period reg to ebitgo server end.\n")
}
