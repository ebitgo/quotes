package thread

import (
	"sync"

	_A "github.com/ebitgo/quotes/models/appconf"
	_C "github.com/ebitgo/quotes/models/coin"
	_L "github.com/ebitgo/quotes/models/log"
	_ck "github.com/jojopoper/go-models/checker"
)

// PoloniexChecker 线程检查
type PoloniexChecker struct {
	_ck.CheckBase
	poloReader []*_C.PoloniexReader
	waiting    *sync.WaitGroup
}

// Init 初始化
func (ths *PoloniexChecker) Init(interval int) _ck.ICheckInterface {
	defer _L.LoggerInstance.InfoPrint("PoloniexChecker checker init complete\n")
	ths.CheckBase.Init(interval)
	ths.SetName("PoloniexOrderReader")
	ths.SetExeFunc(ths.exe)
	ths.poloReader = make([]*_C.PoloniexReader, 0)
	ths.waiting = new(sync.WaitGroup)

	ret := _A.ConfigInstance.GetPoloniexConfig()
	for _, itm := range ret {
		rd := new(_C.PoloniexReader)
		rd.Init(itm[0], itm[1])
		ths.poloReader = append(ths.poloReader, rd)
	}
	return ths
}

func (ths *PoloniexChecker) exe() {
	_L.LoggerInstance.InfoPrint("PoloniexChecker exe...\n")
	for _, reader := range ths.poloReader {
		ths.waiting.Add(1)
		go func(w *sync.WaitGroup, r *_C.PoloniexReader) {
			defer w.Done()
			err := r.ReadDatas()
			if err != nil {
				_L.LoggerInstance.ErrorPrint(" <Poloniex.com> checking[%s - %s] has error \n%+v\n",
					r.MonetaryName, r.CoinName, err)
			}
		}(ths.waiting, reader)
	}
	ths.waiting.Wait()
	_L.LoggerInstance.InfoPrint("PoloniexChecker exe end.\n")
	// _L.LoggerInstance.DebugPrint(" == Poloniex datas \n%+v\n", _C.CoinMangerInstance.Get("poloniex"))
}
