package main

import (
	"learn-go/03.Application/crawlers/engine"
	"learn-go/03.Application/crawlers/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://air-level.com/",
		ParserFunc: parser.CityList,
	})
}
