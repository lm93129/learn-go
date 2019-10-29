package airlevelcrawle

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func GetCity() map[string]string {
	client := &http.Client{Timeout: 20 * time.Second}
	req, _ := http.NewRequest("GET", "http://www.air-level.com/", nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
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

	// 匹配城市名称和拼音
	re := regexp.MustCompile(`<a href="/air/([a-zA-Z]*)/">([一-龥]*)</a>`)
	matches := re.FindAllSubmatch(all, -1)

	// 去除重复的城市名称和拼音
	citylist := make(map[string]string)
	for _, v := range matches {
		citylist[string(v[1])] = string(v[2])
	}
	return citylist
}

func GetAirLevel(m map[string]string) {

}
