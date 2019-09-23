package FunctionAndInterface

import "math"

//如果我们想要一个既能计算方形又能计算矩形的area函数，要怎么做呢，这里就要用到接口的方法了
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
