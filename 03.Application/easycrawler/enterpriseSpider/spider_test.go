package enterpriseSpider

import (
	"fmt"
	"testing"

	"github.com/tidwall/gjson" //获取json内容
)

func TestGetEnterprise(t *testing.T) {
	list := GetEnterpriseList("航天")
	for _, pid := range list {
		date := GetEnterpriseInfo(string(pid[1]))
		value := gjson.Get(string(date), "data.entName")
		fmt.Println(value.String())
	}
}
