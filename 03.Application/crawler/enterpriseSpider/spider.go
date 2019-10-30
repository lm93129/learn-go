package enterpriseSpider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// 输入关键字获取百度企业信用PID
func GetEnterpriseList(keyword string) [][][]byte {
	client := &http.Client{Timeout: 20 * time.Second}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://xin.baidu.com/s?q=%s&t=1", keyword), nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var all []byte
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		all, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	}
	re := regexp.MustCompile(`<a class="zx-list-item-url" target="_blank" href="/detail/compinfo.pid=([a-zA-z0-9-]*)"`)
	matches := re.FindAllSubmatch(all, -1)
	return matches
}

// 通过PID获取企业信息JSON
func GetEnterpriseInfo(pid string) []byte {
	resp, err := http.Get(fmt.Sprintf("https://xin.baidu.com/detail/basicAjax?pid=%s", pid))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var body []byte
	if resp.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	}
	return body
}
