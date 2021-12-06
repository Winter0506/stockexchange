package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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
		utf8Reader := transform.NewReader(resp.Body,
			simplifiedchinese.GBK.NewDecoder())
		sitemap, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s\n", sitemap)
		time.Sleep(5 * time.Second)
	}
	defer resp.Body.Close()
}
