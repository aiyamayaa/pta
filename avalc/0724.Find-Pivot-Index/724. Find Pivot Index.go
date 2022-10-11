package _724_Find_Pivot_Index

//思路1，by ava
func pivotIndex(nums []int) int {
	rSum, lSum := 0, 0
	for i := 1; i < len(nums); i++ {
		rSum += nums[i]
	}
	if rSum == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		rSum -= nums[i]
		lSum += nums[i-1]
		if lSum == rSum {
			return i
		}
	}
	return -1
}

//by carlsun
func pivotIndex1(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0  // 中心索引左半和
	rightSum := 0 // 中心索引右半和
	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		rightSum = sum - leftSum + nums[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

//思路2：前缀和的思路
func pivotIndex3(nums []int) int {
	var total, sum int
	for _, v := range nums {
		total += v
	}

	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
