package coin

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_AC "github.com/ebitgo/quotes/models/appconf"
	_r "github.com/jojopoper/rhttp"
)

const (
	PoloniexKeyName = "poloniex"
)

// PoloniexTradeHistory 历史信息定义
type PoloniexTradeHistory struct {
	Date   string `json:"date"`
	Price  string `json:"rate"`
	Amount string `json:"amount"`
}

// PoloniexOrderDef 挂单定义
type PoloniexOrderDef struct {
	BuyOrder  [][]interface{} `json:"bids"`
	SellOrder [][]interface{} `json:"asks"`
	Trades    []*PoloniexTradeHistory
}

// PoloniexReader 时代网站读取数据
type PoloniexReader struct {
	BaseReader
	OrderDepth int
	proxyIP    string
	proxyPort  string
}

// Init 初始化
func (ths *PoloniexReader) Init(m, c string) {
	ths.BaseUrl = "https://poloniex.com"
	ths.decodeFunc = ths.decoderOrder
	ths.OrderDepth = 10
	ths.proxyIP = _AC.ConfigInstance.GetProxyIP()
	ths.proxyPort = _AC.ConfigInstance.GetProxyPort()
	ths.BaseReader.Init(strings.ToUpper(m), strings.ToUpper(c))
	if len(ths.proxyIP) > 0 {
		client, _ := ths.chttp.GetProxyClient(30, ths.proxyIP, ths.proxyPort)
		ths.chttp.SetClient(client)
	}
	rep := new(ReportDatas)
	rep.Init()
	CoinMangerInstance.Add(PoloniexKeyName, rep)
}

// ReadDatas 读取网站数据
func (ths *PoloniexReader) ReadDatas() error {
	curRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ths.UpdateDecodeFunc(ths.decoderOrder)
	orders, err := ths.readOrder(curRand)
	if err == nil {
		ths.UpdateDecodeFunc(ths.decoderHistory)
		historys, err := ths.readHistory(curRand)

		if err == nil {
			poloData := orders.(*PoloniexOrderDef)
			poloData.Trades = historys.([]*PoloniexTradeHistory)
			ths.formatData(poloData)
		}
	}
	return err
}

func (ths *PoloniexReader) formatData(pdata *PoloniexOrderDef) {
	report := CoinMangerInstance.Get(PoloniexKeyName)
	report.NewData()
	orderList := &Btc38OrderDef{
		BuyOrder:  make([][]float64, 0),
		SellOrder: make([][]float64, 0),
		Trades:    make([][]interface{}, 0),
	}

	for _, itm := range pdata.BuyOrder {
		buyod := make([]float64, 2)
		buyod[0], _ = strconv.ParseFloat(itm[0].(string), 64)
		buyod[1], _ = itm[1].(float64)
		orderList.BuyOrder = append(orderList.BuyOrder, buyod)
	}

	for _, itm := range pdata.SellOrder {
		sellod := make([]float64, 2)
		sellod[0], _ = strconv.ParseFloat(itm[0].(string), 64)
		sellod[1], _ = itm[1].(float64)
		orderList.SellOrder = append(orderList.SellOrder, sellod)
	}

	for _, itm := range pdata.Trades {
		hists := make([]interface{}, 3)
		hists[0] = itm.Date
		hists[1], _ = strconv.ParseFloat(itm.Price, 64)
		hists[2], _ = strconv.ParseFloat(itm.Amount, 64)
		orderList.Trades = append(orderList.Trades, hists)
	}
	report.Set(orderList)
}

func (ths *PoloniexReader) readOrder(r *rand.Rand) (interface{}, error) {
	addr := fmt.Sprintf("%s/public?command=returnOrderBook&depth=%d&currencyPair=%s_%s&_=0.%d",
		ths.BaseUrl, ths.OrderDepth, ths.MonetaryName, ths.CoinName, r.Int31())

	return ths.chttp.ClientGet(addr, _r.ReturnCustomType)
}

func (ths *PoloniexReader) readHistory(r *rand.Rand) (interface{}, error) {
	addr := fmt.Sprintf("%s/public?command=returnTradeHistory&currencyPair=%s_%s&_=0.%d",
		ths.BaseUrl, ths.MonetaryName, ths.CoinName, r.Int31())

	return ths.chttp.ClientGet(addr, _r.ReturnCustomType)
}

func (ths *PoloniexReader) decoderOrder(b []byte) (interface{}, error) {
	orderList := &PoloniexOrderDef{}
	err := json.Unmarshal(b, orderList)
	if err != nil {
		return nil, err
	}
	// fmt.Printf(" poloniex orders =====\n%+v\n", orderList)
	return orderList, err
}

func (ths *PoloniexReader) decoderHistory(b []byte) (interface{}, error) {
	// orderList := &PoloniexOrderDef{
	Trades := make([]*PoloniexTradeHistory, 0)
	// fmt.Printf(" history orig data = \n%s\n", string(b))
	err := json.Unmarshal(b, &Trades)
	if err != nil {
		// fmt.Printf("Umarshal has error \n%+v\n", err)
		return nil, err
	}
	// fmt.Printf(" poloniex trades =====\n%+v\n", Trades)
	return Trades, err
}
