package _154_Find_Minimum_in_Rotated_Sorted_Array_II

import "fmt"

//by ava
func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>2
		//if nums[m] > nums[m+1] {
		//	fmt.Println("ava m res :",m+1)
		//	return nums[m+1]
		//}
		fmt.Println("ava--M:", m)
		if nums[m] > nums[0] {
			l = m + 1
		} else if nums[m] < nums[0] {
			r = m
		} else {
			l++
		}
		fmt.Println(fmt.Printf("ava--L=%d,R=%d ", l, r))
	}
	fmt.Println("ava  res:", l)
	return nums[l]
}
func findMin2(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>2
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r--
		}
	}
	fmt.Println("ava  res:", l)
	return nums[l]
}
