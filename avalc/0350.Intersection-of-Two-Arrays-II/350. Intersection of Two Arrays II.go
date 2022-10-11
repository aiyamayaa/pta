package _350_Intersection_of_Two_Arrays_II

/**
执行用时：4 ms, 在所有 Go 提交中击败了50.73%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了33.77%的用户
*/
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了33.77%的用户
*/
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}
