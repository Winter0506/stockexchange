package logic

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"stockexchange/rpc/stock/stock"
	"strconv"
	"time"
)

// 传入切片 传出 三个结构体
func buildStockStruct(stockPrice []string) (*stock.BasicInfo, *stock.FiveBuyInfo, *stock.FiveSellInfo) {

	todayOpenPrice, _ := strconv.ParseFloat(stockPrice[2], 10)
	lastClosePrice, _ := strconv.ParseFloat(stockPrice[3], 10)
	presentPrice, _ := strconv.ParseFloat(stockPrice[4], 10)
	highPrice, _ := strconv.ParseFloat(stockPrice[5], 10)
	lowPrice, _ := strconv.ParseFloat(stockPrice[6], 10)
	currentBuyPrice, _ := strconv.ParseFloat(stockPrice[7], 10)
	currentSellPrice, _ := strconv.ParseFloat(stockPrice[8], 10)
	transCount, _ := strconv.ParseInt(stockPrice[9], 10, 32)
	transAmount, _ := strconv.ParseFloat(stockPrice[10], 10)

	baseInfo := stock.BasicInfo{
		TodayOpenPrice:   float32(todayOpenPrice),
		LastClosePrice:   float32(lastClosePrice),
		PresentPrice:     float32(presentPrice),
		HighPrice:        float32(highPrice),
		LowPrice:         float32(lowPrice),
		CurrentBuyPrice:  float32(currentBuyPrice),
		CurrentSellPrice: float32(currentSellPrice),
		TransCount:       int32(transCount),
		TransAmount:      float32(transAmount),
	}

	buyOneCount, _ := strconv.ParseInt(stockPrice[11], 10, 32)
	buyOnePrice, _ := strconv.ParseFloat(stockPrice[12], 10)
	buyTwoCount, _ := strconv.ParseFloat(stockPrice[13], 10)
	buyTwoPrice, _ := strconv.ParseFloat(stockPrice[14], 10)
	buyThreeCount, _ := strconv.ParseFloat(stockPrice[15], 10)
	buyThreePrice, _ := strconv.ParseFloat(stockPrice[16], 10)
	buyFourCount, _ := strconv.ParseFloat(stockPrice[17], 10)
	buyFourPrice, _ := strconv.ParseFloat(stockPrice[18], 10)
	buyFiveCount, _ := strconv.ParseFloat(stockPrice[19], 10)
	buyFivePrice, _ := strconv.ParseFloat(stockPrice[20], 10)

	fiveBuyInfo := stock.FiveBuyInfo{
		BuyOneCount:   int32(buyOneCount),
		BuyOnePrice:   float32(buyOnePrice),
		BuyTwoCount:   int32(buyTwoCount),
		BuyTwoPrice:   float32(buyTwoPrice),
		BuyThreeCount: int32(buyThreeCount),
		BuyThreePrice: float32(buyThreePrice),
		BuyFourCount:  int32(buyFourCount),
		BuyFourPrice:  float32(buyFourPrice),
		BuyFiveCount:  int32(buyFiveCount),
		BuyFivePrice:  float32(buyFivePrice),
	}

	sellOneCount, _ := strconv.ParseInt(stockPrice[21], 10, 32)
	sellOnePrice, _ := strconv.ParseFloat(stockPrice[22], 10)
	sellTwoCount, _ := strconv.ParseFloat(stockPrice[23], 10)
	sellTwoPrice, _ := strconv.ParseFloat(stockPrice[24], 10)
	sellThreeCount, _ := strconv.ParseFloat(stockPrice[25], 10)
	sellThreePrice, _ := strconv.ParseFloat(stockPrice[26], 10)
	sellFourCount, _ := strconv.ParseFloat(stockPrice[27], 10)
	sellFourPrice, _ := strconv.ParseFloat(stockPrice[28], 10)
	sellFiveCount, _ := strconv.ParseFloat(stockPrice[29], 10)
	sellFivePrice, _ := strconv.ParseFloat(stockPrice[30], 10)

	fiveSellInfo := stock.FiveSellInfo{
		SellOneCount:   int32(sellOneCount),
		SellOnePrice:   float32(sellOnePrice),
		SellTwoCount:   int32(sellTwoCount),
		SellTwoPrice:   float32(sellTwoPrice),
		SellThreeCount: int32(sellThreeCount),
		SellThreePrice: float32(sellThreePrice),
		SellFourCount:  int32(sellFourCount),
		SellFourPrice:  float32(sellFourPrice),
		SellFiveCount:  int32(sellFiveCount),
		SellFivePrice:  float32(sellFivePrice),
	}

	return &baseInfo, &fiveBuyInfo, &fiveSellInfo
}

// TimestampProto 方法将 time.Time 类型转换为 *timestamppb.Timestamp 类型
func TimestampProto(priceTimeString1, priceTimeString2 string) (*timestamppb.Timestamp, error) {
	// 加载时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	priceTimeString := fmt.Sprintf("%s %s", priceTimeString1, priceTimeString2)
	fmt.Println(priceTimeString)
	priceTime, _ := time.ParseInLocation("2006-01-02 15:04:05", priceTimeString, loc)
	fmt.Println(priceTime)
	ts := timestamppb.New(priceTime)
	return ts, nil
}
