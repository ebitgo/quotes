package coin

// CoinMangerInstance 交易所数据唯一实例
var CoinMangerInstance *CManager

// NewCoinManager 创建新的实例
func NewCoinManager() *CManager {
	ret := new(CManager)
	return ret.Init()
}

// CManager 交易所数据管理
type CManager struct {
	datas map[string]*ReportDatas
}

// Init 初始化
func (ths *CManager) Init() *CManager {
	ths.datas = make(map[string]*ReportDatas)
	return ths
}

// Get 读取
func (ths *CManager) Get(key string) *ReportDatas {
	return ths.datas[key]
}

// Add 添加
func (ths *CManager) Add(key string, r *ReportDatas) {
	ths.datas[key] = r
}
