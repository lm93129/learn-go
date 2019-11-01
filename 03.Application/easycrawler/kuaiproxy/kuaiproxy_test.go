package kuaiproxy

import (
	"fmt"
	"testing"
)

func TestCityAir(t *testing.T) {
	v := KuaiProxy("2")
	fmt.Println(len(v))
	for _, m := range v {
		fmt.Println(m)
	}
}
