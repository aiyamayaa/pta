package _977_Squares_of_a_Sorted_Array

import "sort"

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	for l, r, i := 0, len(nums)-1, len(nums)-1; l <= r; i-- {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			ret[i] = nums[l] * nums[l]
			l++
		} else {
			ret[i] = nums[r] * nums[r]
			r--
		}
	}
	return ret
}

// 解法二
func sortedSquares1(A []int) []int {
	for i, value := range A {
		A[i] = value * value
	}
	sort.Ints(A)
	return A
}
