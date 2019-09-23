package FunctionAndInterface

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("use Perimeter", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want) //这里的 f 对应 float64，.2 表示输出 2 位小数
		}
	})
	//使用结构体
	t.Run("use StructPerimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := StructPerimeter(rectangle)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want) //这里的 f 对应 float64，.2 表示输出 2 位小数
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("use Area", func(t *testing.T) {
		got := Areas(10.0, 40.0)
		want := 400.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want) //这里的 f 对应 float64，.2 表示输出 2 位小数
		}
	})
	//使用结构体
	t.Run("user structArea", func(t *testing.T) {
		rectangle := Rectangle{10.0, 40.0}
		got := StructArea(rectangle)
		want := 400.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want) //这里的 f 对应 float64，.2 表示输出 2 位小数
		}
	})
}

func TestArea2(t *testing.T) {
	t.Run("test Rectangle", func(t *testing.T) {
		for _, tc := range areaTests {
			got := tc.shape.Area()
			if got != tc.want {
				t.Errorf("got %.2f want %.2f", got, tc.want)
			}
		}
	})
}
