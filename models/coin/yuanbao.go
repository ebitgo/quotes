package coin

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	_r "github.com/jojopoper/rhttp"
)

const (
	YuanBaoKeyName = "yuanbao"
)

// MarketTradeData market交易数据定义
type MarketTradeData struct {
	CoinFrom string `json:"coin_from"`
	CoinTo   string `json:"coin_to"`
	Current  string `json:"current"`
}

// YuanbaoMarketsData 元宝Data信息定义
type YuanbaoMarketsData struct {
	Cny []*MarketTradeData
}

// YuanbaoTradeHistory 历史信息定义
type YuanbaoTradeHistory struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []*YuanbaoMarketsData
}

// YuanbaoOrderDef 挂单定义
type YuanbaoOrderDef struct {
	BuyOrder  [][]float64 `json:"bids"`
	SellOrder [][]float64 `json:"asks"`
	// Trades    []*YuanbaoTradeHistory
}

// PoloniexReader 时代网站读取数据
type YuanbaoReader struct {
	BaseReader
}

// Init 初始化
func (ths *YuanbaoReader) Init(m, c string) {
	ths.BaseUrl = "https://www.yuanbao.com"
	ths.decodeFunc = ths.decoderOrder
	ths.BaseReader.Init(strings.ToLower(m), strings.ToLower(c))
	rep := new(ReportDatas)
	rep.Init()
	CoinMangerInstance.Add(YuanBaoKeyName, rep)
}

// ReadDatas 读取网站数据
func (ths *YuanbaoReader) ReadDatas() error {
	curRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ths.UpdateDecodeFunc(ths.decoderOrder)
	orders, err := ths.readOrder(curRand)
	if err == nil {
		// ths.UpdateDecodeFunc(ths.decoderHistory)
		// historys, err := ths.readHistory(curRand)

		// if err == nil {
		ybData := orders.(*YuanbaoOrderDef)
		// 	poloData.Trades = historys.([]*YuanbaoTradeHistory)
		ths.formatData(ybData)
		// }
	}
	return err
}

func (ths *YuanbaoReader) formatData(pdata *YuanbaoOrderDef) {
	report := CoinMangerInstance.Get(YuanBaoKeyName)
	report.NewData()
	lensell := len(pdata.SellOrder)
	orderList := &Btc38OrderDef{
		BuyOrder:  pdata.BuyOrder,
		SellOrder: make([][]float64, lensell),
		Trades:    make([][]interface{}, 0),
	}

	index := lensell - 1
	for i := 0; i < lensell; i++ {
		orderList.SellOrder[index] = pdata.SellOrder[i]
		index--
	}

	// for _, itm := range pdata.Trades {
	// 	hists := make([]interface{}, 3)
	// 	hists[0] = itm.Date
	// 	hists[1], _ = strconv.ParseFloat(itm.Price, 64)
	// 	hists[2], _ = strconv.ParseFloat(itm.Amount, 64)
	// 	orderList.Trades = append(orderList.Trades, hists)
	// }
	report.Set(orderList)
}

func (ths *YuanbaoReader) readOrder(r *rand.Rand) (interface{}, error) {
	addr := fmt.Sprintf("%s/market/depths?depth=%s2%s&v=0.%d",
		ths.BaseUrl, ths.CoinName, ths.MonetaryName, r.Int31())

	return ths.chttp.Get(addr, _r.ReturnCustomType)
}

func (ths *YuanbaoReader) readHistory(r *rand.Rand) (interface{}, error) {
	addr := fmt.Sprintf("%s/coins/markets?t=0.%d",
		ths.BaseUrl, r.Int31())

	return ths.chttp.Get(addr, _r.ReturnCustomType)
}

func (ths *YuanbaoReader) decoderOrder(b []byte) (interface{}, error) {
	orderList := &YuanbaoOrderDef{}
	err := json.Unmarshal(b, orderList)
	if err != nil {
		return nil, err
	}
	// fmt.Printf(" poloniex orders =====\n%+v\n", orderList)
	return orderList, err
}

func (ths *YuanbaoReader) decoderHistory(b []byte) (interface{}, error) {
	Trades := make([]*YuanbaoTradeHistory, 0)
	// fmt.Printf(" history orig data = \n%s\n", string(b))
	err := json.Unmarshal(b, &Trades)
	if err != nil {
		// fmt.Printf("Umarshal has error \n%+v\n", err)
		return nil, err
	}
	// fmt.Printf(" poloniex trades =====\n%+v\n", Trades)
	return Trades, err
}
