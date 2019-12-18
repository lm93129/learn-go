package sort

import (
	"fmt"
	"reflect"
	"testing"
)

//检查函数
func checkSort(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

//冒泡排序
func TestBubbleSort(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		for _, m := range SortCase {
			got := BubbleASort(m.values)
			fmt.Println(got)
			checkSort(t, got, m.want)
		}

	})

	t.Run("use interface", func(t *testing.T) {
		for _, m := range SortInterfaceCase {
			BubbleBSort(m.Vector)
			fmt.Println(m.Vector)
			checkSort(t, m.Vector, m.want)
		}
	})

}

//插入排序
func TestInsertSort(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		for _, m := range SortCase {
			fmt.Println(m.values)
			got := InsertASort(m.values)
			fmt.Println(got)
			checkSort(t, got, m.want)
		}
	})
	t.Run("B", func(t *testing.T) {
		for _, m := range SortInterfaceCase {
			fmt.Println(m.Vector)
			InsertBSort(m.Vector)
			fmt.Println(m.Vector)
			checkSort(t, m.Vector, m.want)
		}
	})
}

//翻转数据
func TestReverse(t *testing.T) {
	var want = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	t.Run("A", func(t *testing.T) {
		got := ReverseA([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		checkSort(t, got, want)
	})
	t.Run("B", func(t *testing.T) {
		Vector := element{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		ReverseB(Vector)
		fmt.Println(Vector)
		checkSort(t, Vector, want)
	})
}
