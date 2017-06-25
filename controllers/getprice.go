package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	_m "github.com/ebitgo/quotes/models"
	_L "github.com/ebitgo/quotes/models/log"
	_V "github.com/jojopoper/go-models/validation"
)

// GetPriceController 获取当前价格
type GetPriceController struct {
	beego.Controller
}

// Get get function
func (ths *GetPriceController) Get() {
	ths.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With, Accept")
	ths.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ths.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	ths.Ctx.Output.Header("content-type", "application/json;charset=utf8")

	_L.LoggerInstance.InfoPrint("Get request quotes ...\n")
	op := &_m.OperationModel{}
	op.DecodeContext(&ths.Controller)
	ret := op.QueryExecute()
	if ret.CodeID == _V.NoError {
		ret = op.GetResultData()
	}

	if ret.CodeID > 0 {
		_L.LoggerInstance.ErrorPrint("[%s] Request quotes Error : \r\n\t%s\r\n", ths.Ctx.Input.IP(), ret.ErrorMsg)
	}
	data, _ := json.Marshal(ret)
	ths.Data["json"] = string(data)
	ths.ServeJSON(false)
}
