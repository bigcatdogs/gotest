package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("爬取开始--------------------------------")
	// 创建一个Collector
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
	})

	// 设置一个回调函数，该函数将在请求完成后调用
	//c.OnResponse(func(r *colly.Response) {
	//
	//	fmt.Println(string(r.Body))
	//})

	// 开始爬取
	c.Visit("http://www.baidu.com")
	//fmt.Println(string(c.Body))
}
