package types

type ReqUserFav struct {
	UserId  string `json:"userid"`
	StockId string `json:"stockid"`
}

type ReqStockFav struct {
	UserId  string `json:"userid"`
	StockId string `json:"stockid"`
}

type RespFavDetail struct {
	DetailMeta `json:"meta"`
}

type FavMessage struct {
	Total   int      `json:"total""`
	FavList []string `json:"favlist"`
}

type RespFavList struct {
	FavMessage `json:"message"`
	DetailMeta `json:"meta"`
}
