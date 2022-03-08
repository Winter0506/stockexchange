package types

type ReqUserAccountId struct {
	Id int32 `path:"id"`
}

type AccountDetailMessage struct {
	UserId         int32   `json:"userId"`
	Account        float32 `json:"accout"`
	MarketValue    float32 `json:"marketvalue"`
	Available      float32 `json:"available"`
	ProfiltAndLoss float32 `json:"profiltAndLoss"`
}

type RespUserAccountDetail struct {
	AccountDetailMessage `json:"message"`
	DetailMeta           `json:"meta"`
}

type ReqUserAccountCreate struct {
	UserId  int32   `json:"userId"`
	Account float32 `json:"account"`
}

type ReqUserAccountUpdate struct {
	UserId  int32   `json:"userId"`
	Account float32 `json:"account"`
}

type HoldPositionListMeta struct {
	Msg    string `json:"msg""`
	Status int16  `json:"status"`
}

type RespHoldPositionList struct {
	HoldPositionList     []string `json:"holdPositionList"`
	HoldPositionListMeta `json:"meta"`
}

type ReqCreateTrustItem struct {
	UserId    int32 `json:"userId"`
	StockId   int32 `json:"stockId"`
	Num       int32 `json:"num"`
	Direction uint  `json:"direction"` // 1是买入 2是卖出
}

type TrustInfoResponse struct {
	Id         int32   `json:"id"`
	User       int32   `json:"userId"`
	Stock      int32   `json:"stockId"`
	Num        int32   `json:"num"`
	Cost       float32 `json:"cost"`
	Direction  uint    `json:"direction"`
	DealNumber int32   `json:"dealNumber"`
	DealCost   float32 `json:"dealCost"`
	Status     string  `json:"status"`
	TrustSn    string  `json:"trustSn"`
}

type RespTrustInfoResponse struct {
	TrustInfoResponse `json:"message"`
	DetailMeta        `json:"meta"`
}
