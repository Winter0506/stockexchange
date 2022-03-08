package main

import (
	"fmt"
	"github.com/syhlion/greq"
	"github.com/syhlion/requestwork.v2"
	"log"
	"time"
)

func main() {
	for {
		// 在V4基础上进行改进
		worker := requestwork.New(50)

		client := greq.New(worker, 5*time.Second, false)

		//GET
		stockdata, _, err := client.Get("http://hq.sinajs.cn/?format=text&list=sh600519", nil)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("data:%s", stockdata)
		time.Sleep(5 * time.Second)
	}
}
