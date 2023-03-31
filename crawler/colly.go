package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	c := colly.NewCollector()
	// Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// c.OnHTML("li", func(e *colly.HTMLElement) {
	// 	//var b = make([]byte, 1000)
	// 	//n, err := e.Request.Body.Read(b)
	// 	e.Attr("")

	// })

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
    if err != nil {
        log.Fatal(err)
    }
    divNodes := htmlquery.Find(doc, `/html/body/div[@id='outer-wrapper']/div[@id='wrap2']/div[@id='content-wrapper']/div[@id='main-wrapper']/div[@id='main']/div[@id='Blog1']/div[@class='blog-posts hfeed']/div[@class='post hentry uncustomized-post-template']`)
    for _, node := range divNodes {
        url := htmlquery.FindOne(node, "./h1[@class='post-title entry-title']/a/@href")
		
		fmt.Println(htmlquery.InnerText(url))
    }
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})


	
    if p, err := proxy.RoundRobinProxySwitcher(
        "http://192.168.1.110:7890",
    ); err == nil {
        c.SetProxyFunc(p)
    }


	c.Visit("https://program-think.blogspot.com/")
}
