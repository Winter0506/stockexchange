package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 这种写法有错
	res, err := http.Get("http://hq.sinajs.cn/?format=text&list=sh600519")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	stockInfo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", stockInfo)
}
