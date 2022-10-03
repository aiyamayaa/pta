package avago

import (
	"fmt"
	"math"
	"testing"
)

func TestTest(t *testing.T) {

}

/** 运算符
&      位运算 AND   与
|      位运算 OR    或
^      位运算 XOR   异或 ：不同为1，相同为0
&^     位清空（AND NOT）:c = a &^ b: 把b中1对应位置在c位置上清空。举例：a: 00001100，b：00000100，c:：00001000
<<     左移 算术上，一个x<<n左移运算等价于乘以2^n （乘以2的n次方）左移运算用零填充右边空缺的bit位，
>>     右移 算术上，一个x>>n右移运算等价于除以2^n （除以2的n次方）
				无符号数的右移运算也是用0填充左边空缺的bit位，但是有符号数的右移运算会用符号位的值填充左边空缺的bit位
*/
func TestInt(t *testing.T) {

	fmt.Printf("%08b\n", 1)         //00000001
	fmt.Printf("%08b\n", 1<<5)      //00100000
	fmt.Printf("%08b\n", 1<<1|1<<5) //00100010
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

}

/**
  常量math.MaxFloat32表示float32能表示的最大数值，大约是 3.4e38；
  常量math.MaxFloat64大约是1.8e308。
  float32和float64它们分别能表示的最小值近似为1.4e-45和4.9e-324。

*/
func TestFloat(t *testing.T) {
	//(1)因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	//用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印.
	//但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适
	//打印精度是小数点后三个小数精度和8个字符宽度：
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	//+Inf -Inf :正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	//函数math.IsNaN用于测试一个数是否是非数NaN，math.NaN则返回非数对应的值。
	//注意：NaN和任何数都是不相等
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

/**
复数：
（1）复数的表示：创建一个复数，获取实部、虚部
（2）复数的运算：求复数的平方根函数和求幂函数。。。
（3）复数的比较：可以用==和!=进行相等比较
*/
func TestComplex(t *testing.T) {
	//（1）内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"  实部
	fmt.Println(imag(x * y))         // "10"  虚部
	//参考：https://books.studygolang.com/gopl-zh/ch3/ch3-03.html

}

func TestBool(t *testing.T) {

}

func TestString(t *testing.T) {
	ages := make(map[int]int)
	ages[1] = 1
	ages[2] = 2
	val := ages[3]
	fmt.Printf("%+T", val)
}

func TestConst(t *testing.T) {

}
