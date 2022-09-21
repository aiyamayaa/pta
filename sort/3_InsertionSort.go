package sort

/**
插入排序：从左至右，每获取一个数，把它放到左边正确的位置，保证左边有序
参考：https://www.runoob.com/w3cnote/insertion-sort.html
优点：插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率；
缺点：但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位；
*/
func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}
