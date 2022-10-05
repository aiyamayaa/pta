package _367_Valid_Perfect_Square

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了8.50%的用户
*/
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		m := l + (r-l)>>2
		if m*m == num {
			return true
		} else if m*m < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

/**
牛顿迭代法
*/
func isPerfectSquare2(num int) bool {
	for i := 1; num > 0; i += 2 {
		num -= i
	}
	return num == 0
}
