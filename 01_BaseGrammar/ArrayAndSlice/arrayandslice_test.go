package ArrayAndSlice

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	for _, tc := range ArrayCase {
		if Array(tc.array) != tc.want {
			t.Errorf("want: %d", tc.want)
		}
	}
}

func TestSumSlice(t *testing.T) {

	//封装一个检查函数，reflect.DeepEqual传入的值是string和slice也会返回true，所以还可以增加代码的类型安全性
	//这样如果传入的值是个string或者其他类型，还会报错
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	//t.Run可以在一个单元测试函数里面加多个子单元测试
	t.Run("sum slice", func(t *testing.T) {
		got := sumall([]int{1, 2, 3, 4, 5})
		want := 15
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("add 2 slice", func(t *testing.T) {
		got := SumSlice([]int{1, 2, 3, 4, 5}, []int{1, 3, 4, 5})
		want := []int{15, 13}
		// reflect.DeepEqual可以对比两个数组，当然也可以循环遍历对比。
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumEmptySlice([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSumSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSlice([]int{1, 2, 3, 4, 5}, []int{1, 3, 4, 5})
	}
}
