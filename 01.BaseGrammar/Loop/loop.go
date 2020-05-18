package loop

import "fmt"

//go语言只有for循环，没有其他循环了
func Loop(n int) int {
	var m int
	for i := 0; i < n; i++ {
		m += i
	}
	return m
}

//其他简写的方法
func Loop2(n int) int {
	var i = 0
	for i <= n {
		i++
	}
	return i
}

//死循环
func Loop3() {
	i := 0
	for {
		i++
	}
}

//遍历数组
func RangeMap() {
	var amap = []int{1, 2, 3, 4, 5}
	for i, m := range amap {
		fmt.Printf("i: %d , v: %d", i, m)
	}
}
