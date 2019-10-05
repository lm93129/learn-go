package easy

var SumsCase = []struct {
	nums   []int
	target int
	want   []int
}{
	{
		[]int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12},
		16,
		[]int{0, 9},
	},
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		3,
		[]int{7, 8},
	},
}

var RecordCase = []struct {
	Record string
	want   bool
}{
	{
		"PPALLL",
		false,
	},
	{
		"PPALLP",
		true,
	},
}
