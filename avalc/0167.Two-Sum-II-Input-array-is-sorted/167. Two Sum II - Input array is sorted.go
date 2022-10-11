package _167_Two_Sum_II_Input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	var res []int
	for i, n := range numbers {
		if index, ok := m[target-n]; !ok {
			m[n] = i + 1
		} else {
			return []int{index, i + 1}
		}
	}
	return res
}
func twoSum2(numbers []int, target int) []int {
	res := make([]int, 2)
	left, right := 0, len(numbers)-1
	for {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}
	}
}
