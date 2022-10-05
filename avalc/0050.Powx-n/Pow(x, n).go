package _050_Powx_n

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, n)
}
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	r := quickMul(x, n/2)
	if n%2 == 0 {
		return r * r
	}
	return r * r * x
}
