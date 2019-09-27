package easy

import (
	"reflect"
	"testing"
)

func checkSort(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v \n", got, want)
	}
}

func Test(t *testing.T) {
	t.Run("sums", func(t *testing.T) {
		for _, tc := range SumsCase {
			got := twoSum(tc.nums, tc.target)
			checkSort(t, got, tc.want)
		}
	})
}
