package _349_Intersection_of_Two_Arrays

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			res = append(res, n)
			delete(m, n)
		}
	}
	return res
}
