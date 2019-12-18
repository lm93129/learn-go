package sort

func InsertASort(element []int) []int {
	l := len(element)
	for i := 1; i < l; i++ {
		// 每一趟不满足条件就选择i为哨兵保存，将哨兵插入0~i-1有序序列（0~i-1始终是有序的）
		for j := i; j > 0 && element[j] < element[j-1]; j-- {
			element[j], element[j-1] = element[j-1], element[j]
		}
	}
	return element
}

//插入排序
func InsertBSort(vector Vector) Vector {
	l := vector.Len()
	for i := 1; i < l; i++ {
		// 每一趟不满足条件就选择i为哨兵保存，将哨兵插入0~i-1有序序列（0~i-1始终是有序的）
		for j := i; j > 0 && vector.Less(j, j-1); j-- {
			vector.Swap(j, j-1)
		}
	}
	return vector
}
