package HelloWorld

import "testing"

//string拼接
func TestHelloName(t *testing.T) {
	for _, tc := range HelloCases {
		got := HelloName(tc.name)
		if got != tc.want {
			t.Errorf("got is %s , want is %s", got, tc.want)
		}
	}
}

//单元测试
func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello World!"
	if got != want {
		t.Errorf("got is %s , want is %s ", got, want)
	}
}

//基准测试
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}
