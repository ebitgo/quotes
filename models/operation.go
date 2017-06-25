package models

import (
	"fmt"

	"github.com/astaxie/beego"
	_V "github.com/jojopoper/go-models/validation"
)

// OperationModel 操作模块定义
type OperationModel struct {
	operation _V.IOperation
}

// GetResultData 按照接口要求的获取结果
func (ths *OperationModel) GetResultData() *_V.OperationResult {
	if ths.operation != nil {
		return ths.operation.GetResultData()
	}
	return &_V.OperationResult{
		ErrorMsg: "Undefined parameter format[2]",
		CodeID:   _V.UnknownFormatError,
	}
}

// QueryExecute 按照接口要求执行过程
func (ths *OperationModel) QueryExecute() *_V.OperationResult {
	if ths.operation != nil {
		return ths.operation.QueryExecute()
	}
	return &_V.OperationResult{
		ErrorMsg: "Can not find out result object",
		CodeID:   _V.CommomError,
	}
}

// DecodeContext 按照接口要求解码获取参数
func (ths *OperationModel) DecodeContext(ctl *beego.Controller) {
	paramType := ctl.Input().Get("t")

	switch paramType {
	case "all":
		ths.operation = &AllOrdersOp{}
	default:
		fmt.Println(paramType)
		return
	}
	ths.operation.DecodeContext(ctl)
}
