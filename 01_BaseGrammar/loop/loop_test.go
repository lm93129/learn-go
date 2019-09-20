package loop

import (
	"testing"
)

func TestLoop(t *testing.T) {
	for _, tc := range LoopCase {
		if Loop(tc.m) != tc.want {
			t.Errorf("got %d , want %d \n", Loop(tc.m), tc.want)
		}
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(1000)
	}
}
func TestLoop3(t *testing.T) {
	Loop3()
}

func TestLoop2(t *testing.T) {
	for _, tc := range Loop2Case {
		if Loop2(tc.m) != tc.want {
			t.Errorf("got %d , want %d \n", Loop2(tc.m), tc.want)
		}
	}
}

func TestRangeMap(t *testing.T) {
	RangeMap()
}
