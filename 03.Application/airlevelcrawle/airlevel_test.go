package airlevelcrawle

import (
	"fmt"
	"testing"
)

func TestGetEnterprise(t *testing.T) {
	list := GetCity()
	for a, v := range list {
		fmt.Println(a, v)
	}
}
