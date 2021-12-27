package types

type ReqStockId struct {
	Id int32 `path:"id"`
}

type DetailMessage struct {
	Id               int32   `json:"id"`
	StockName        string  `json:"stockName"`
	StockCode        string  `json:"stockCode"`
	TodayOpenPrice   float32 `json:"todayOpenPrice"`
	LastClosePrice   float32 `json:"lastClosePrice"`
	PresentPrice     float32 `json:"presentPrice"`
	HighPrice        float32 `json:"highPrice"`
	LowPrice         float32 `json:"lowPrice"`
	CurrentBuyPrice  float32 `json:"currentBuyPrice"`
	CurrentSellPrice float32 `json:"currentSellPrice"`
	TransCount       int32   `json:"transCount"`
	TransAmount      float32 `json:"transAmount"`
	BuyOneCount      int32   `json:"buyOneCount"`
	BuyOnePrice      float32 `json:"buyOnePrice"`
	BuyTwoCount      int32   `json:"buyTwoCount"`
	BuyTwoPrice      float32 `json:"buyTwoPrice"`
	BuyThreeCount    int32   `json:"buyThreeCount"`
	BuyThreePrice    float32 `json:"buyThreePrice"`
	BuyFourCount     int32   `json:"buyFourCount"`
	BuyFourPrice     float32 `json:"buyFourPrice"`
	BuyFiveCount     int32   `json:"buyFiveCount"`
	BuyFivePrice     float32 `json:"buyFivePrice"`
	SellOneCount     int32   `json:"sellOneCount"`
	SellOnePrice     float32 `json:"sellOnePrice"`
	SellTwoCount     int32   `json:"sellTwoCount"`
	SellTwoPrice     float32 `json:"sellTwoPrice"`
	SellThreeCount   int32   `json:"sellThreeCount"`
	SellThreePrice   float32 `json:"sellThreePrice"`
	SellFourCount    int32   `json:"sellFourCount"`
	SellFourPrice    float32 `json:"sellFourPrice"`
	SellFiveCount    int32   `json:"sellFiveCount"`
	SellFivePrice    float32 `json:"sellFivePrice"`
	CurrentTime      string  `json:"currentTime"`
}

type DetailMeta struct {
	Msg    string `json:"msg""`
	Status int16  `json:"status"`
}

type RespStockDetail struct {
	DetailMessage `json:"message"`
	DetailMeta    `json:"meta"`
}

type ReqStockCreate struct {
	StockName string `json:"stockName"`
	StockCode string `json:"stockCode"`
}
