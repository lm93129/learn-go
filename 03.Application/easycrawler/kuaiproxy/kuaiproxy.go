package kuaiproxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func KuaiProxy(num string) (Proxylist []ProxyList) {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://61.164.39.66:53281")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Timeout: 20 * time.Second, Transport: transport}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://www.kuaidaili.com/free/inha/%s", num), nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// 使用GoQuery
	doc.Find("tr").Each(func(i int, ele *goquery.Selection) {
		proxy := ProxyList{
			Ip:    ele.Find("td").Eq(0).Text(),
			Port:  ele.Find("td").Eq(1).Text(),
			Types: strings.ToLower(ele.Find("td").Eq(3).Text()),
		}
		Proxylist = append(Proxylist, proxy)
	})
	// 如果数组为空则直接返回空数组避免错误
	if len(Proxylist) == 0 {
		return Proxylist
	}
	// 删除第一个空元素
	return Proxylist[1:]
}

type ProxyList struct {
	Ip    string `json:"ip"`
	Port  string `json:"port"`
	Types string `json:"types"`
}
