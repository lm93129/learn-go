package kuaiproxy

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{Timeout: 20 * time.Second}

func KuaiProxy(num string) (Proxylist []ProxyList) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://www.kuaidaili.com/free/inha/%s", num), nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	doc.Find("tr").Each(func(i int, ele *goquery.Selection) {
		proxy := ProxyList{
			Ip:    ele.Find("td").Eq(0).Text(),
			Port:  ele.Find("td").Eq(1).Text(),
			Types: ele.Find("td").Eq(3).Text(),
		}
		Proxylist = append(Proxylist, proxy)
	})

	// 删除第一个空元素
	return Proxylist[1:]
}

type ProxyList struct {
	Ip    string `json:"ip"`
	Port  string `json:"port"`
	Types string `json:"types"`
}
