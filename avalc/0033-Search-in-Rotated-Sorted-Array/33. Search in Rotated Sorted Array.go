package _033_Search_in_Rotated_Sorted_Array

//正确案例：二分法左闭右闭
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		}
		//nums[m]=nums[0]的情况，此时l=m,r>m，有2个值，
		//[3,1]，这种情况，
		//如果把3算作左边的序列处理，1为右边的序列，左边递增右边递增没问题
		//如果把3算作右边的序列处理，左边序列为空，右边序列递减，不对
		if nums[m] >= nums[0] {
			if nums[m] > target && nums[l] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

//错误案例：左闭右开，容易造成越界，不好
func search2(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		}
		if nums[m] >= nums[0] {
			if nums[m] > target && nums[l] <= target {
				r = m
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[0] {
			//这里一般需要怕断右边和target的值，选择左闭右开容易造成越界， nums[r-1]也容易造成越界
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
