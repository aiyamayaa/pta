package _034_Find_First_and_Last_Position_of_Element_in_Sorted_Array

/**
by ava self
执行用时：4 ms, 在所有 Go 提交中击败了95.77%的用户
内存消耗：3.8 MB, 在所有 Go 提交中击败了99.97%的用户
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	var ret []int
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return []int{-1, -1}
	} else {
		ret = append(ret, l)
		for l < len(nums) && nums[l] == target {
			l++
		}
		ret = append(ret, l-1)
	}
	return ret
}

/**
github 题解
*/
func searchRange2(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}

}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 用两个边界方法
/**
执行用时：4 ms, 在所有 Go 提交中击败了95.77%的用户
内存消耗：3.8 MB, 在所有 Go 提交中击败了58.14%的用户
*/
func searchRange3(nums []int, target int) []int {
	// 目标值开始位置：为 target 的左侧边界
	start := findLeftBound(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 目标值结束位置：为 target 的右侧边界
	end := findRightBound(nums, target)
	return []int{start, end}
}

// 寻找左侧边界的二分查找
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 返回左侧边界
	return left // note
}

// 寻找右侧边界的二分查找
func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 返回右侧边界
	return right
}
