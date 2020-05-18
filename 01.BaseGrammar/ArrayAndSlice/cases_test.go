package ArrayAndSlice

var ArrayCase = []struct {
	array [5]int
	want  int
}{
	{
		[5]int{1, 2, 3, 4, 5},
		15,
	},
	{
		[5]int{1, 0, 3, 4, 5},
		13,
	},
}
