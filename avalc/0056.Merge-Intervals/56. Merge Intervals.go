package _056_Merge_Intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	index := 0
	var ret [][]int
	ret = append(ret, intervals[0])
	for _, interval := range intervals {
		if interval[0] > ret[index][1] {
			index++
			ret = append(ret, interval)
		} else {
			ret[index][1] = max(interval[1], ret[index][1])
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//利用贪心算法实现
func merge2(intervals [][]int) [][]int {
	//先从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//再弄重复的
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1]) //赋值最大值
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}
