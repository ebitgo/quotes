package coin

import (
	"encoding/json"
	"sync"
	"time"
)

// PriceDef 价格定义
type PriceDef struct {
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
	DateTime string  `json:"date_time"`
}

// ReportDatas 报告需要的数据定义
type ReportDatas struct {
	BuyList    []*PriceDef `json:"buy"`
	SellList   []*PriceDef `json:"sell"`
	TradeList  []*PriceDef `json:"trade"`
	locker     *sync.Mutex
	UpdateTime time.Time `json:"update_time"`
}

// Init 初始化
func (ths *ReportDatas) Init() {
	ths.locker = new(sync.Mutex)
}

// NewData 新的数据源
func (ths *ReportDatas) NewData() {
	ths.BuyList = make([]*PriceDef, 0)
	ths.SellList = make([]*PriceDef, 0)
	ths.TradeList = make([]*PriceDef, 0)
}

// Set 设置数据源
func (ths *ReportDatas) Set(src *Btc38OrderDef) {
	ths.locker.Lock()
	defer ths.locker.Unlock()

	ths.BuyList = ths.setOrderList(src.BuyOrder)
	ths.SellList = ths.setOrderList(src.SellOrder)
	ths.TradeList = ths.setTradeList(src.Trades)
	ths.UpdateTime = time.Now()
}

func (ths *ReportDatas) setOrderList(src [][]float64) []*PriceDef {
	if src == nil {
		return nil
	}
	dst := make([]*PriceDef, 0)
	for _, itm := range src {
		if itm != nil {
			prc := &PriceDef{
				Price:  itm[0],
				Amount: itm[1],
			}
			dst = append(dst, prc)
		}
	}
	return dst
}

func (ths *ReportDatas) setTradeList(src [][]interface{}) []*PriceDef {
	if src == nil {
		return nil
	}

	dst := make([]*PriceDef, 0)
	for _, itm := range src {
		if itm != nil {
			prc := &PriceDef{
				Price:    itm[1].(float64),
				Amount:   itm[2].(float64),
				DateTime: itm[0].(string),
			}
			dst = append(dst, prc)
		}
	}
	return dst
}

// 读取数据源json格式
func (ths *ReportDatas) Read() ([]byte, error) {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	return json.Marshal(ths)
}
