package thread

import (
	"sync"

	_A "github.com/ebitgo/quotes/models/appconf"
	_C "github.com/ebitgo/quotes/models/coin"
	_L "github.com/ebitgo/quotes/models/log"
	_ck "github.com/jojopoper/go-models/checker"
)

// Btc38Checker 线程检查
type Btc38Checker struct {
	_ck.CheckBase
	btc38Reader []*_C.Btc38Reader
}

// Init 初始化
func (ths *Btc38Checker) Init(interval int) _ck.ICheckInterface {
	defer _L.LoggerInstance.InfoPrint("Btc38Checker checker init complete\n")
	ths.CheckBase.Init(interval)
	ths.SetName("Btc38OrderReader")
	ths.SetExeFunc(ths.exe)
	ths.btc38Reader = make([]*_C.Btc38Reader, 0)

	ret := _A.ConfigInstance.GetBtc38Config()
	for _, itm := range ret {
		rd := new(_C.Btc38Reader)
		rd.Init(itm[0], itm[1])
		ths.btc38Reader = append(ths.btc38Reader, rd)
	}
	return ths
}

func (ths *Btc38Checker) exe() {
	_L.LoggerInstance.InfoPrint("Btc38Checker exe...\n")
	wg := new(sync.WaitGroup)
	for _, reader := range ths.btc38Reader {
		wg.Add(1)
		go func(w *sync.WaitGroup, r *_C.Btc38Reader) {
			defer w.Done()
			err := r.ReadDatas()
			if err != nil {
				_L.LoggerInstance.ErrorPrint(" <Btc38.com> checking[%s - %s] has error \n%+v\n",
					r.MonetaryName, r.CoinName, err)
			}
		}(wg, reader)
	}
	wg.Wait()
	_L.LoggerInstance.InfoPrint("Btc38Checker exe end.\n")
	// _L.LoggerInstance.DebugPrint("%+v\n", _C.CoinMangerInstance.Get("btc38"))
}
