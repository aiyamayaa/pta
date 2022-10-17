package _073_Set_Matrix_Zeroes

import (
	"fmt"
	"testing"
)

type question73 struct {
	para73
	ans73
}

// para 是参数
// one 代表第一个参数
type para73 struct {
	matrix [][]int
}

// ans 是答案
// one 代表第一个答案
type ans73 struct {
	matrix [][]int
}

func Test_Problem73(t *testing.T) {

	qs := []question73{
		{
			para73{[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			ans73{[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			}},
		},
		{
			para73{[][]int{
				{-4, -2147483648, 6, -7, 0},
				{-8, 6, -8, -6, 0},
				{2147483647, 2, -9, -6, -10},
			}},
			ans73{[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{2147483647, 2, -9, -6, 0},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 73------------------------\n")

	for _, q := range qs {
		c, p := q.ans73, q.para73
		fmt.Printf("【input  】:%v       \n", p)
		setZeroes(p.matrix)
		fmt.Printf("【output 】:%v\n", p)
		fmt.Printf("【correct】:%v\n", c)
	}
	fmt.Printf("\n\n\n")
}
