package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func check(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v \n", got, want)
	}
}

func Test(t *testing.T) {
	t.Run("两数之和", func(t *testing.T) {
		for _, tc := range SumsCase {
			got := twoSum(tc.nums, tc.target)
			check(t, got, tc.want)
		}
	})

	t.Run("计算质数", func(t *testing.T) {
		fmt.Println(countPrimes(10))
	})

	t.Run("学生出勤记录", func(t *testing.T) {
		for _, tc := range RecordCase {
			got := checkRecord(tc.Record)
			if got != tc.want {
				t.Errorf("got %v want %v \n", got, tc.want)
			}
		}
	})

}
