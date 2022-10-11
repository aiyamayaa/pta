package simple_sort_2

import "sort"

func main() {
	var intervals = [][]int{{2, 3}, {1, 2}, {4, 5}, {3, 4}}

	//对该数组以每行第一个元素升序排列的顺序进行排序

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}
