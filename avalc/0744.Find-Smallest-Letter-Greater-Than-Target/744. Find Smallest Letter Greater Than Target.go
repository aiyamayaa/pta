package _744_Find_Smallest_Letter_Greater_Than_Target

import "sort"

/**
by ava
执行用时：4 ms, 在所有 Go 提交中击败了11.40%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了99.56%的用户
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		m := l + (r-l)>>1
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l >= len(letters) {
		return letters[0]
	}
	return letters[l]

}

/**
左闭右开
执行用时：4 ms, 在所有 Go 提交中击败了11.40%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了99.56%的用户
*/
func nextGreatestLetter2(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		m := l + (r-l)>>1
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return letters[l%len(letters)]
}

/* 左闭右闭
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了55.26%的用户
**/
func nextGreatestLetter2_2(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l <= r {
		m := l + (r-l)>>1
		if letters[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return letters[l%len(letters)]
}

func nextGreatestLetter3(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	i := sort.Search(len(letters)-1, func(i int) bool { return letters[i] > target })
	return letters[i]
}
