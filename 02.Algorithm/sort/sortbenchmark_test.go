package sort

import (
	"math/rand"
	"testing"
	"time"
)

var values = generateRandomNumber(1, 9999, 1000)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("A", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleASort(values)
		}
	})

	Vector := element{}
	Vector = generateRandomNumber(1, 9999, 1000)
	b.Run("B", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleBSort(Vector)
		}
	})
}

func BenchmarkInsertSort(b *testing.B) {

	b.Run("A", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InsertASort(values)
		}
	})

	Vector := element{}
	Vector = generateRandomNumber(1, 9999, 1000)
	b.Run("B", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InsertBSort(Vector)
		}
	})
}

//生成随机不重复的数组用于性能测试
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
