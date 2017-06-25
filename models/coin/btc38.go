package coin

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	_r "github.com/jojopoper/rhttp"
)

const (
	Btc38KeyName = "btc38"
)

// Btc38OrderDef 挂单定义
type Btc38OrderDef struct {
	BuyOrder  [][]float64     `json:"buyStr"`
	SellOrder [][]float64     `json:"sellStr"`
	Trades    [][]interface{} `json:"tradeStr"`
}

// Btc38Reader 时代网站读取数据
type Btc38Reader struct {
	BaseReader
}

// Init 初始化
func (ths *Btc38Reader) Init(m, c string) {
	ths.BaseUrl = "http://www.btc38.com"
	ths.decodeFunc = ths.decoder
	ths.BaseReader.Init(m, c)
	rep := new(ReportDatas)
	rep.Init()
	CoinMangerInstance.Add(Btc38KeyName, rep)
}

// ReadDatas 读取网站数据
func (ths *Btc38Reader) ReadDatas() error {
	curRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := fmt.Sprintf("%s/trade/getTradeList30.php?mk_type=%s&coinname=%s&n=0.00%d1",
		ths.BaseUrl, ths.MonetaryName, ths.CoinName, curRand.Int31())

	_, err := ths.chttp.Get(addr, _r.ReturnCustomType)
	return err
}

func (ths *Btc38Reader) decoder(b []byte) (interface{}, error) {
	orderList := &Btc38OrderDef{}
	err := json.Unmarshal(b, orderList)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("Order list ===========\n%+v\n", orderList)
	report := CoinMangerInstance.Get(Btc38KeyName)
	report.NewData()
	report.Set(orderList)
	return report, err
}
