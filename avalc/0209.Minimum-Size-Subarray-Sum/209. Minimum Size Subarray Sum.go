package _209_Minimum_Size_Subarray_Sum

func minSubArrayLen(target int, nums []int) int {
	pointL, sum := 0, 0
	length := len(nums) + 1
	for i, n := range nums {
		sum += n
		for sum >= target {
			length = min(length, i-pointL+1)
			sum -= nums[pointL]
			pointL += 1
		}
	}
	if length == len(nums)+1 {
		return 0
	}
	return length
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
