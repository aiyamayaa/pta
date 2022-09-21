package sort

/**
冒泡排序：每一次把最大的泡冒到最右边
参考：https://www.runoob.com/w3cnote/bubble-sort.html
*/
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
