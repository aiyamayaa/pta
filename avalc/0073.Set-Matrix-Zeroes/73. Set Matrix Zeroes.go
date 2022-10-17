package _073_Set_Matrix_Zeroes

func setZeroes(matrix [][]int) {
	cowX_have_0 := false
	for j, m := range matrix[0] {
		if m == 0 {
			cowX_have_0 = true
			matrix[0][j] = 0
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
			if matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	//注意后面的先后顺序
	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if cowX_have_0 {
		matrix[0] = make([]int, len(matrix[0]))
	}
}
