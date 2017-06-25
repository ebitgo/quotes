package thread

import (
	"sync"

	_ck "github.com/jojopoper/go-models/checker"
)

// CheckManagerInstance 全局唯一实例
var CheckManagerInstance *TManager

// NewCheckManager 获取实例
func NewCheckManager() *TManager {
	ret := new(TManager)
	return ret.Init()
}

// TManager 线程监控管理
type TManager struct {
	checker map[string]_ck.ICheckInterface
	lock    *sync.Mutex
}

// Init 初始化
func (ths *TManager) Init() *TManager {
	ths.lock = new(sync.Mutex)
	ths.regChecker()
	return ths
}

func (ths *TManager) regChecker() {
	ths.checker = make(map[string]_ck.ICheckInterface)

	b38 := new(Btc38Checker)
	b38.Init(20).RegistManager(ths.checker)
	poloniex := new(PoloniexChecker)
	poloniex.Init(20).RegistManager(ths.checker)
	yb := new(YuanbaoChecker)
	yb.Init(20).RegistManager(ths.checker)
	reg := new(PeriodReg)
	reg.Init(600).RegistManager(ths.checker)
}

// Check 启动检测
func (ths *TManager) Check() {
	for _, c := range ths.checker {
		if !c.IsRunning() {
			if c.IsBeginStart() {
				go c.Start()
			}
		}
	}
}
