package datetype

import (
	"fmt"
	"testing"
)

//断言函数封装
func assertCorrectMessage(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("sum '%d' want '%d'", got, want)
	}
}

func TestAddInt(t *testing.T) {
	for _, tc := range AddIntCases {
		sum := AddInt(tc.x, tc.y)
		fmt.Println(sum)
		assertCorrectMessage(t, sum, tc.want)
	}
}

func TestAddFloat(t *testing.T) {
	for _, tc := range AddFloatCases {
		sum := AddFloat(tc.x, tc.y)
		fmt.Println(sum)
		assertCorrectMessage(t, sum, tc.want)
	}
}

//示例函数，可以在godoc中生成示例
func ExampleTypeBool() {
	out := TypeBool()
	fmt.Println(out)
	// Output: true
}

func ExampleAddInt() {
	sum := AddInt(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkAddFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddFloat(1.0, 2.3)
	}
}

func BenchmarkEuler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Euler()
	}
}

func TestTrig(t *testing.T) {
	if Trig() != 5 {
		t.Errorf("Trig is %d", Trig())
	}
}
