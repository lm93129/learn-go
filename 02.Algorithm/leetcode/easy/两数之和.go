package easy

import "fmt"

func twoSum(nums []int, target int) []int {
	// hash 一次遍历 空间：O(n) 时间O(n)
	info := make(map[int]int)
	for i, num := range nums {
		//如果target-num存在值，说明该num为对应的值，然后输出该num的index和target的index
		if index, ok := info[target-num]; ok {
			return []int{index, i}
		} else if _, ok := info[num]; !ok {
			info[num] = i
			fmt.Println(info)
		}
	}
	return []int{}
}
