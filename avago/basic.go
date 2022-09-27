package avago

import (
	"crypto/sha256"
	"fmt"
)

//基本的赋值：元组赋值
func fuzhi() {
	x := 1
	y := 2
	x, y = y, x
	//赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值
	a := make([]int, 5)
	a[1], a[2] = a[2], a[1]
	//元组赋值也可以使一系列琐碎赋值更加紧凑
	i, j, k := 2, 3, 5
	fmt.Println(i, j, k)
}

//计算两个整数值的的最大公约数（GCD），辗转相除法
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//计算斐波纳契数列（Fibonacci）的第N个数（斐波那契：后面的数等于其前面两个数的和）
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//Sha crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要
func Sha() {
	//消息摘要有256bit大小，因此对应[32]byte数组类型。如果两个消息摘要是相同的，那么可以认为两个消息本身也是相同
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	//%x :它用于指定以十六进制的格式打印数组或slice全部的元素
	//%t :用于打印布尔型数据，
	//%T :用于显示一个值对应的数据类型。
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}
