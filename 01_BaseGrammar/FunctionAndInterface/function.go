package FunctionAndInterface

//定义一个方形的结构
type Rectangle struct {
	Width  float64
	Height float64
}

//计算周长
func Perimeter(width, height float64) float64 {
	return (width + height) * 2
}

//计算面积
func Areas(width, height float64) float64 {
	return width * height
}

//使用struct的方式
func StructPerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func StructArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
