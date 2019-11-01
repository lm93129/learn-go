package kuaiproxy

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCityAir(t *testing.T) {

	v := KuaiProxy("1")

	for _, m := range v {
		fmt.Println(m.Ip)
		content := strings.Join([]string{m.Types, "://", m.Ip, ":", m.Port, "\n"}, "")
		wiretestring(content)
	}

}

func wiretestring(str string) {
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(str)
	fd.Write(buf)
	fd.Close()
}
