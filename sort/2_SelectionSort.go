package sort

/**
选择排序：每次挑选最小的，放在最左边
参考：https://www.runoob.com/w3cnote/selection-sort.html
*/
func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
