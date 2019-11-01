package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("Citylist_test.html")
	if err != nil {
		t.Error(err)
	}

	result := CityList(contents)
	fmt.Println(len(result.Items))
}

func TestParserCityAirInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("Citylist_test.html")
	if err != nil {
		t.Error(err)
	}

	result := CityAirInfo(contents, "铜陵")
	fmt.Println(result.Items)
}
