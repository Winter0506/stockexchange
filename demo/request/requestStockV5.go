package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// 正式代码 可以把切片改成结构体
func main() {
	client := &http.Client{}
	// 请求方法可以包括 http method  GET、POST、PUT、DELETE
	req, err := http.NewRequest("GET", "http://hq.sinajs.cn/?format=text&list=sh600519", nil)
	var resp *http.Response
	for {
		if err != nil {
			log.Println(err)
			return
		}
		req.Header.Add("If-None-Match", `W/"wyzzy"`)
		resp, err = client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		/*
			if resp.StatusCode != http.StatusOK {
				fmt.Println("Error: status code", resp.StatusCode)
				return
			}
		 */
		utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
		sitemap, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			log.Fatal(err)
			return
		}
		// 股票代码 股票名字 今日开盘价 昨日收盘价 当前时刻价格 今日最高价 今日最低价
		// 做成字符串切片传出
		// 模拟打印
		var stockSlice []string
		var stockNo, stockName string
		stockNo, stockName = byteToName(strings.Split(string(sitemap), ",")[0])
		/*todayOpeningPrice := strings.Split(string(sitemap), ",")[1]
		yesterdayClosingPrice := strings.Split(string(sitemap), ",")[2]
		currentPrice := strings.Split(string(sitemap), ",")[3]
		todayHighPrice := strings.Split(string(sitemap), ",")[4]
		todayLowPrice := strings.Split(string(sitemap), ",")[5]*/
		stockPriceSlice := strings.Split(string(sitemap), ",")[1:32]

		stockSlice = append(append(append(stockSlice, stockNo), stockName), stockPriceSlice...)
		fmt.Println(stockSlice)
		time.Sleep(60 * time.Second)
	}
	defer resp.Body.Close()
}

func byteToName(stockNoAndName string) (string, string) {
	stockNo := strings.Split(stockNoAndName, "=")[0]
	stockName := strings.Split(stockNoAndName, "=")[1]
	return stockNo, stockName
}