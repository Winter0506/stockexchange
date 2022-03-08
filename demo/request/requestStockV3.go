package main

import (
	"fmt"
	"github.com/syhlion/greq"
	"github.com/syhlion/requestwork.v2"
	"log"
	"time"
)

func main() {

	// 因應業務需求，需要有大量可以存取各式 rest api 並且又要控制好併發數量，而開發出來的套件，使用方式如下
	// 套件上已經有在內部做釋放資源 Close 的動作，所以不需要像原生的套件一樣，需要特地把 Body.Close()
	//need import https://github.com/syhlion/requestwork.v2
	// 這是最高同時併發的控制，是 greq 的核心套件
	worker := requestwork.New(50)

	client := greq.New(worker, 5*time.Second, false)

	//GET
	data, _, err := client.Get("http://hq.sinajs.cn/?format=text&list=sh600519", nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("data:%s", data)

	//POST
	//v := url.Values{}

	//post form data
	/*v.Add("data", string(data))
	data, httpstatus, err := client.Post("https://tw.yahoo.com", bytes.NewBufferString(v.Encode()))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("data:%s ,httpstatus:%d", data, httpstatus)*/

}
