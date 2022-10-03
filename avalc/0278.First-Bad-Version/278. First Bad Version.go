package _278_First_Bad_Version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//左闭右开的写法
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了22.70%的用户
*/
func firstBadVersion(n int) int {
	l, r := 1, n+1
	for l < r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m
			continue
		}
		l = m + 1
	}
	//跳出前最后一次：l=m=r-1
	//跳出for的原因，是l=r，如果是r=m导致的，r=m=l=BadVersion；如果是l=m+1导致的，那原来的m是okVersion，m+1=l为第一个badVersion
	//可以直接return l的原因是，肯定至少有一个错误版本 ，所以，如果只有1个值，那也一定是bad version，然后r=m=l,跳出循转
	//这里return right和return left一样
	return l
}

//左闭右闭的写法:
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了5.04%的用户
*/
func firstBadVersion2(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m - 1
			continue
		}
		l = m + 1
	}
	//跳出前最后一次：l==r的时候，判断的是[ok,?,bad],对应 的下标是[l-1,l=r=m,r+1]
	//这里跳出循环的条件是 l>r，如果是从if跳出，m=l=badVersion，r=okVersion ,否则，l=m+1=badVersion，r=okVersion
	return l
}
func isBadVersion(version int) bool {
	return true
}
