package HelloWorld

import "testing"

//单元测试
func TestHelloWorld(t *testing.T) {
	for _, tc := range HelloCases {
		got := LanguageSwitch(tc.name, tc.language)
		if got != tc.want {
			t.Errorf("got is %s , want is %s", got, tc.want)
		}
	}
}

//基准测试
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageSwitch("li", "spanish")
	}
}
