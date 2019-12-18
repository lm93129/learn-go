package sort

var SortCase = []struct {
	values []int
	want   []int
}{
	{
		[]int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12},
		[]int{4, 12, 27, 37, 80, 81, 84, 85, 93, 93},
	},
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

var SortInterfaceCase = []struct {
	Vector element
	want   []int
}{
	{
		element{4, 93, 84, 85, 80, 37, 81, 93, 27, 12},
		[]int{4, 12, 27, 37, 80, 81, 84, 85, 93, 93},
	},
	{
		element{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}
