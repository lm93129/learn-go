package main

import (
	"learn-go/03.Application/crawlers/engine"
	"learn-go/03.Application/crawlers/parser"
)

func main() {
	// 从初始页面开始爬取
	engine.Run(engine.Request{
		Url:        "http://air-level.com/",
		ParserFunc: parser.CityList,
	})
}
