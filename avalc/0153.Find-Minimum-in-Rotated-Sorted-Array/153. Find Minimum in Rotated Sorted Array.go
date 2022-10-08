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

//leetcode官方解答-返回的值是正确的，但是下标有问题
func findMin3(nums []int) int {
	lens := len(nums)
	left, right := 0, lens-1
	for left < right { //需要mid与右边的邻居比较才能判断mid是否满足条件，使用模板2
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { //右边有序，则mid已经是右边min,向左下坡
			right = mid
		} else { //向右下坡
			left = mid + 1
		}
	}
	return nums[left] // 此时left与right重合
}
