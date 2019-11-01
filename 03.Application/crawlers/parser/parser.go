package parser

import (
	"bytes"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"

	"learn-go/03.Application/crawlers/engine"
)

const cityListRe = `<a href="/air/([a-zA-Z]*)/">([一-龥]*)</a>`

func CityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// 去除重复的城市名称和拼音
	citylist := make(map[string]string)
	for _, v := range matches {
		citylist[string(v[1])] = string(v[2])
	}

	for a, v := range citylist {
		cityname := v
		result.Items = append(result.Items, cityname)
		result.Requests = append(result.Requests, engine.Request{
			Url: "http://air-level.com/air/" + a,
			ParserFunc: func(i []byte) engine.ParseResult {
				return CityAirInfo(i, cityname)
			},
		})
	}
	return result
}

func CityAirInfo(contents []byte, cityname string) engine.ParseResult {
	log.Printf("Parser city: %s", cityname)
	docs := bytes.NewReader(contents)
	doc, err := goquery.NewDocumentFromReader(docs)
	if err != nil {
		log.Fatalln(err)
	}
	result := engine.ParseResult{}
	doc.Find("table").Find("tr").Each(func(i int, ele *goquery.Selection) {
		cityairlevel := CityAirLevel{
			Addr:  ele.Find("td").Eq(0).Text(),
			Aqi:   ele.Find("td").Eq(1).Text(),
			Level: ele.Find("td").Eq(2).Text(),
			Pm25:  ele.Find("td").Eq(3).Text(),
			Pm10:  ele.Find("td").Eq(4).Text(),
		}
		result.Items = append(result.Items, cityairlevel)
	})
	if result.Items != nil {
		result.Items = result.Items[1:]
	}
	return result
}

// 城市空气质量等级结构化
type CityAirLevel struct {
	Addr  string `json:"addr"`
	Aqi   string `json:"aqi"`
	Level string `json:"level"`
	Pm25  string `json:"pm_25"`
	Pm10  string `json:"pm_10"`
}
