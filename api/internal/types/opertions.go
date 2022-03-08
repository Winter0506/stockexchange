package types

type ReqUserFav struct {
	UserId  int32 `form:"userid"`
	StockId int32 `form:"stockid"`
}

type ReqStockFav struct {
	StockId int32 `form:"stockid"`
}

type ReqStockFavList struct {
	StockId int32 `form:"stockid"`
}

type ReqUserFavList struct {
	UserId int32 `form:"userid"`
}

type RespFavDetail struct {
	DetailMeta `json:"meta"`
}

type FavMessage struct {
	Total   int      `json:"total"`
	FavList []string `json:"favlist"`
}

type RespFavList struct {
	FavMessage `json:"message"`
	DetailMeta `json:"meta"`
}
