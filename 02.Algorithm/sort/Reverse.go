package sort

//翻转数据
func ReverseA(element []int) []int {
	l := len(element) - 1
	r := len(element) / 2
	for i := 0; i < r; i++ {
		element[i], element[l-i] = element[l-i], element[i]
	}
	return element
}

func ReverseB(vector Vector) Vector {
	l := vector.Len() - 1
	r := vector.Len() / 2
	for i := 0; i < r; i++ {
		vector.Swap(i, l-i)
	}
	return vector
}
