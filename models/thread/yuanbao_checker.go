package thread

import (
	"sync"

	_A "github.com/ebitgo/quotes/models/appconf"
	_C "github.com/ebitgo/quotes/models/coin"
	_L "github.com/ebitgo/quotes/models/log"
	_ck "github.com/jojopoper/go-models/checker"
)

// YuanbaoChecker 线程检查
type YuanbaoChecker struct {
	_ck.CheckBase
	ybReader []*_C.YuanbaoReader
	waiting  *sync.WaitGroup
}

// Init 初始化
func (ths *YuanbaoChecker) Init(interval int) _ck.ICheckInterface {
	defer _L.LoggerInstance.InfoPrint("YuanbaoChecker checker init complete\n")
	ths.CheckBase.Init(interval)
	ths.SetName("YuanbaoOrderReader")
	ths.SetExeFunc(ths.exe)
	ths.ybReader = make([]*_C.YuanbaoReader, 0)
	ths.waiting = new(sync.WaitGroup)

	ret := _A.ConfigInstance.GetYuanbaoConfig()
	for _, itm := range ret {
		rd := new(_C.YuanbaoReader)
		rd.Init(itm[0], itm[1])
		ths.ybReader = append(ths.ybReader, rd)
	}
	return ths
}

func (ths *YuanbaoChecker) exe() {
	_L.LoggerInstance.InfoPrint("YuanbaoChecker exe...\n")
	for _, reader := range ths.ybReader {
		ths.waiting.Add(1)
		go func(w *sync.WaitGroup, r *_C.YuanbaoReader) {
			defer w.Done()
			err := r.ReadDatas()
			if err != nil {
				_L.LoggerInstance.ErrorPrint(" <www.yuanbao.com> checking[%s - %s] has error \n%+v\n",
					r.MonetaryName, r.CoinName, err)
			}
		}(ths.waiting, reader)
	}
	ths.waiting.Wait()
	_L.LoggerInstance.InfoPrint("YuanbaoChecker exe end.\n")
	// _L.LoggerInstance.DebugPrint(" == Yuanbao datas \n%+v\n", _C.CoinMangerInstance.Get("poloniex"))
}
