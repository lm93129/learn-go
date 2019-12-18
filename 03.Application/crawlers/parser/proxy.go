package parser

import (
	"bytes"
	"log"
	"strings"

	"learn-go/03.Application/crawlers/engine"

	"github.com/PuerkitoBio/goquery"
)

func KuaiProxy(contents []byte) engine.ParseResult {
	docs := bytes.NewReader(contents)
	doc, err := goquery.NewDocumentFromReader(docs)
	if err != nil {
		log.Fatalln(err)
	}
	result := engine.ParseResult{}
	// 使用GoQuery
	doc.Find("tr").Each(func(i int, ele *goquery.Selection) {
		proxy := ProxyList{
			Ip:    ele.Find("td").Eq(0).Text(),
			Port:  ele.Find("td").Eq(1).Text(),
			Types: strings.ToLower(ele.Find("td").Eq(3).Text()),
		}
		result.Items = append(result.Items, proxy)
	})
	// 如果数组为空则直接返回空数组避免错误
	if len(result.Items) != 0 {
		// 删除第一个空元素
		result.Items = result.Items[1:]
	}
	return result
}

type ProxyList struct {
	Ip    string `json:"ip"`
	Port  string `json:"port"`
	Types string `json:"types"`
}
