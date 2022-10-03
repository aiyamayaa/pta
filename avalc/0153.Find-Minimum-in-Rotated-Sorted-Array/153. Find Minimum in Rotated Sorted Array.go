package _153_Find_Minimum_in_Rotated_Sorted_Array

/**
执行用时：4 ms, 在所有 Go 提交中击败了17.87%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了99.92%的用户
*/
func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	for l < r {
		m := l + (r-l)>>1
		//这里判断比较耗时，这个方法不好
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[m] >= nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了58.89%的用户
  左闭右开
*/
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
