package models

import (
	"github.com/astaxie/beego"
	_c "github.com/ebitgo/quotes/models/coin"
	_V "github.com/jojopoper/go-models/validation"
)

// OperRetData 返回数据结构定义
type OperRetData struct {
	Btc38Rpt *_c.ReportDatas `json:"btc38_data"`
	PolonRpt *_c.ReportDatas `json:"polon_data"`
	YuanbRpt *_c.ReportDatas `json:"yuanbao_data"`
}

// AllOrdersOp 获取所有交易所挂单信息
type AllOrdersOp struct {
	_V.CheckParamValid
	data *OperRetData
}

// GetResultData 返回处理结果
func (ths *AllOrdersOp) GetResultData() *_V.OperationResult {
	return ths.OperationRetData
}

// QueryExecute 处理过程
func (ths *AllOrdersOp) QueryExecute() *_V.OperationResult {
	ths.data.Btc38Rpt = _c.CoinMangerInstance.Get(_c.Btc38KeyName)
	ths.data.PolonRpt = _c.CoinMangerInstance.Get(_c.PoloniexKeyName)
	ths.data.YuanbRpt = _c.CoinMangerInstance.Get(_c.YuanBaoKeyName)
	ths.OperationRetData.ResultData = ths.data
	return ths.OperationRetData
}

// DecodeContext 解码获取参数
func (ths *AllOrdersOp) DecodeContext(ctl *beego.Controller) {
	ths.OperationRetData = new(_V.OperationResult)
	ths.data = new(OperRetData)
}
