package utils

import "sort"

// Sort sorts values (in-place) with respect to the given comparator.
//
// Uses Go's sort (hybrid of quicksort for large and then insertion sort for smaller slices).
// 对values进行从小到大排序
func Sort(values []interface{}, comparator Comparator) {
	sort.Sort(sortable{values, comparator})
}

//实现了sort接口，定义了值和它的比较方法，对于不同的类型定义不同的比较方法即可
type sortable struct {
	values     []interface{}
	comparator Comparator
}

func (s sortable) Len() int {
	return len(s.values)
}
func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
