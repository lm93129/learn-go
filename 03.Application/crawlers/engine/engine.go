package engine

import (
	"log"

	"learn-go/03.Application/crawlers/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	// 将种子页面传入
	for _, r := range seeds {
		requests = append(requests, r)
	}
	// 循环爬取和解析
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetcher url: %s", r.Url)
		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Printf("fetcher err : %s", err)
			continue
		}
		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			log.Printf("Got item: %s", item)
		}
	}
}
