package sort

/**
希尔排序：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序
参考：https://www.runoob.com/w3cnote/shell-sort.html
「O(nLogn)~O(n^2)」
分组用的增量的选择：第一次gap=  len(arr) / 2 :相当于2个为一组
				 第二次 gap = gap/2 相当于分组数量减半，4个为一组
						。。。
				 最后一次：gap = 1 ,相当于最后最有一个大组了，
*/
func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}

	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}
