package ArrayAndSlice

//常用的是切片类型而不是数组类型
//数组类型为下面这种写法：
//[5]int{1,2,3,4,5}
//切片类型为：
//[]int{1,2,3,4,5}

func Array(number [5]int) int {
	sum := 0
	for _, v := range number {
		sum += v
	}
	return sum
}

func sumall(number []int) int {
	sum := 0
	for _, v := range number {
		sum += v
	}
	return sum
}

func SumSlice(number ...[]int) (sums []int) {
	for _, v := range number {
		sums = append(sums, sumall(v))
	}

	return sums
}

func SumEmptySlice(number ...[]int) (sums []int) {
	for _, numbers := range number {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] //截取部分切片，取到从索引 1 到最后一个元素
			sums = append(sums, sumall(tail))
		}
	}

	return sums
}
