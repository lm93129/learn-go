package airlevelcrawle

import (
	"fmt"
	"testing"
)

func TestGetEnterprise(t *testing.T) {
	list := GetCity()
	for a, v := range list {
		fmt.Println(a, v)
		//v := CityAir(a)
		//for _, m := range v {
		//	fmt.Println(m)
		//}
	}
}

func TestCityAir(t *testing.T) {
	v := CityAir("hefei")
	for _, m := range v {
		fmt.Println(m)
	}
}
