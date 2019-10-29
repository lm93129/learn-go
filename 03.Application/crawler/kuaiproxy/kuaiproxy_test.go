package kuaiproxy

import (
	"fmt"
	"testing"
)

func TestCityAir(t *testing.T) {
	v := KuaiProxy("1")
	for _, m := range v {
		fmt.Println(m)
	}
}
