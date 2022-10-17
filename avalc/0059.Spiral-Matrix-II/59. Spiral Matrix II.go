package _059_Spiral_Matrix_II

import "fmt"

func generateMatrix(n int) [][]int {
	res, k := make([][]int, n), 1
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for m := 0; m < n/2; m++ {
		//初始坐标：(m,m),(m,n-m-1),(n-m-1,n-m-1),(n-m-1,m)
		fmt.Println("m:", m)
		for j := m; j < n-m-1; j++ {
			fmt.Println(fmt.Sprintf("res[%d][%d]=%d", m, j, k))
			res[m][j] = k
			k++
		}
		for i := m; i < n-m-1; i++ {
			fmt.Println(fmt.Sprintf("res[%d][%d]=%d", i, n-m-1, k))
			res[i][n-m-1] = k
			k++
		}
		for j := n - m - 1; j > m; j-- {
			fmt.Println(fmt.Sprintf("res[%d][%d]=%d", n-m-1, j, k))
			res[n-m-1][j] = k
			k++
		}
		for i := n - m - 1; i > m; i-- {
			fmt.Println(fmt.Sprintf("res[%d][%d]=%d", i, m, k))
			res[i][m] = k
			k++
		}
	}
	if n%2 == 1 {
		res[n/2][n/2] = k
	}
	return res
}
func generateMatrix1(n int) [][]int {
	res, k := make([][]int, n), 1
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for m := 0; m < n/2; m++ {
		//初始坐标：(m,m),(m,n-m*2-1),(n-m*2-1,n-m*2-1),(n-m*2-1,m)
		for j := m; j < n-m-1; j++ {
			res[m][j] = k
			k++
		}
		for i := m; i < n-m-1; i++ {
			res[i][n-m-1] = k
			k++
		}
		for j := n - m - 1; j > m; j-- {
			res[n-m-1][j] = k
			k++
		}
		for i := n - m - 1; i > m; i-- {
			res[i][m] = k
			k++
		}
	}
	if n%2 == 1 {
		res[n/2][n/2] = k
	}
	return res
}
