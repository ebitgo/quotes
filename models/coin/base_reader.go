package coin

import _r "github.com/jojopoper/rhttp"

// BaseReader 基础读取定义
type BaseReader struct {
	BaseUrl      string
	MonetaryName string
	CoinName     string
	chttp        *_r.CHttp
	decodeFunc   _r.DecodeFunction
}

// Init 初始化
func (ths *BaseReader) Init(m, c string) {
	ths.MonetaryName = m
	ths.CoinName = c
	ths.chttp = new(_r.CHttp)
	if ths.decodeFunc == nil {
		panic("Have to set decodeFun before call BaseReader.Init()")
	}
	ths.chttp.SetDecodeFunc(ths.decodeFunc)
}

// UpdateDecodeFunc 更新解码函数
func (ths *BaseReader) UpdateDecodeFunc(f _r.DecodeFunction) {
	ths.chttp.SetDecodeFunc(f)
}
