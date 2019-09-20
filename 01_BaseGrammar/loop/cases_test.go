package loop

var LoopCase = []struct {
	m    int
	want int
}{
	{
		1000,
		499500,
	},
	{
		10,
		55,
	},
	{
		100,
		5050,
	},
}

var Loop2Case = []struct {
	m    int
	want int
}{
	{
		1000,
		1001,
	},
	{
		10,
		11,
	},
	{
		100,
		101,
	},
}
