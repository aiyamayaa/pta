package sort

/*
	归并排序：整个序列一分为二，保证左边有序和右边有序，然后合并保证整体有序
	参考：https://www.runoob.com/w3cnote/merge-sort.html
	思路：通过递归实现
*/
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle] //包含arr[0]到arr[middle-1]
	right := arr[middle:] //包含arr[middle]到最后
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	//这里使用：
	result = append(result, left...)
	result = append(result, right...)
	//for len(left) != 0 {
	//	result = append(result, left[0])
	//	left = left[1:]
	//}
	//
	//for len(right) != 0 {
	//	result = append(result, right[0])
	//	right = right[1:]
	//}

	return result
}
