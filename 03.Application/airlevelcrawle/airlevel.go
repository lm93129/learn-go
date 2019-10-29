package airlevelcrawle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{Timeout: 20 * time.Second}

func GetCity() map[string]string {

	req, _ := http.NewRequest("GET", "http://air-level.com/", nil)
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

func CityAir(city string) (cityinfos []AirLevel) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://air-level.com/air/%s", city), nil)
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

	doc.Find("table").Find("tr").Each(func(i int, ele *goquery.Selection) {

		cityinfo := AirLevel{
			Addr:  ele.Find("td").Eq(0).Text(),
			Aqi:   ele.Find("td").Eq(1).Text(),
			Level: ele.Find("td").Eq(2).Text(),
			Pm25:  ele.Find("td").Eq(3).Text(),
			Pm10:  ele.Find("td").Eq(4).Text(),
		}
		cityinfos = append(cityinfos, cityinfo)
	})

	return cityinfos
}

type AirLevel struct {
	Addr  string
	Aqi   string
	Level string
	Pm25  string
	Pm10  string
}
