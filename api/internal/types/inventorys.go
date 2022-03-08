package types

type ReqSetInv struct {
	StockId int32 `form:"stockid"`
	Num     int32 `form:"num"`
}

type ReqInvDetail struct {
	StockId int32 `form:"stockid"`
}

type RespSetInv struct {
	DetailMeta `json:"meta"`
}

type RespInvDetail struct {
	ReqSetInv  `json:"message"`
	DetailMeta `json:"meta"`
}
