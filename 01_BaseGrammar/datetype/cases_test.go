package datetype

//构造测试用的数据
var AddIntCases = []struct {
	x    int
	y    int
	want int
}{
	{
		1,
		2,
		3,
	},
	{
		10,
		30,
		40,
	},
}

var AddFloatCases = []struct {
	x    float32
	y    float32
	want float32
}{
	{
		1.1,
		2.0,
		3.1,
	},
	{
		10.1,
		30.01,
		40.11,
	},
}
