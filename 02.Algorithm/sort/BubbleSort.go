package sort

//直接简单的写法
func BubbleASort(element []int) []int {
	for i := 0; i < len(element)-1; i++ {
		for j := i + 1; j < len(element); j++ {
			if element[i] > element[j] {
				element[i], element[j] = element[j], element[i]
			}
		}
	}
	return element
}

//运用方法、结构体和接口的思想
type element []int

type Vector interface {
	Len() int
	//如果第i个元素大于第j个元素则返回True : False
	Greater(i, j int) bool
	//如果第i个元素小于第j个元素则返回True : False
	Less(i, j int) bool
	//交换元素
	Swap(i, j int)
}

func (e element) Greater(i, j int) bool {
	return e[i] > e[j]
}

func (e element) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e element) Len() int {
	return len(e)
}

func (e element) Less(i, j int) bool  {
	return e[i] < e[j]
}

func BubbleBSort(vector Vector) Vector {
	l := vector.Len()
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if vector.Greater(i, j) {
				vector.Swap(i, j)
			}
		}
	}
	return vector
}
