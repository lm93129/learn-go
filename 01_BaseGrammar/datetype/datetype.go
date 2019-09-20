package datetype

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包全局变量数据定义的不同方式
var a = 1 //定义为整数
var (
	b int = 2    //定义为整数
	d     = true //定义为布尔类型
)

//常量的定义
const c string = "hello" //定义为字符串

//整数
func AddInt(x, y int) int {
	return x + y
}

func AddFloat(x, y float32) float32 {
	return x + y
}

func TypeBool() bool {
	e := "World" // 使用:=的方式声明变量只能在函数内使用
	fmt.Println(a, b, d, c, e)
	return true
}

//函数名首字母大写才能在别的文件调用
func Euler() {
	euler := cmplx.Pow(math.E, 1i*math.Pi) + 1 //欧拉公式： e^i*pi
	euler = cmplx.Exp(1i*math.Pi) + 1          //e^i*pi 写法不一样，但是结果是一样的
	fmt.Printf("%f", euler)
}

//强制数据类型转换
func Trig() int {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //实现int转float64和float64转int
	fmt.Println(c)
	return c
}
