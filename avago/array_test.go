package avago

import (
	"fmt"
	"testing"
)

/**
golang中各种类型数据的初始化
*/

/**
new函数：表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
*/
func TestArray0(t *testing.T) {
	//（1）默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // 0 print the first element
	fmt.Println(a[len(a)-1]) // 0 print the last element, a[2]
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func TestArray(t *testing.T) {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)    //[1 2 3]
	fmt.Println(r)    //[1 2 0]
	fmt.Println(r[2]) // "0"
	//(2-1)如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	m := [...]int{1, 2, 3}
	fmt.Printf("%T\n", m) // "[3]int"

	//(2-2-1)symbol := [...]string{0: "$", 1: "€", 2: "￡", 3: "￥"}
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(symbol)           //[$ € ￡ ￥]
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	//(2-2-2)symbol2 := [...]string{0: "$", 1: "€",  3: "￥",2: "￡"}
	symbol2 := [...]string{USD: "$", EUR: "€", RMB: "￥", GBP: "￡"}
	fmt.Println(symbol2)           //[$ € ￡ ￥]
	fmt.Println(RMB, symbol2[RMB]) // "3 ￥"

	//(2-3下面表示第10个数字(下标为9)为-1
	rr := [...]int{9: -1}
	fmt.Println(rr)
	fmt.Printf("%T\n", rr) // "[10]int"
}

func TestArrayCompare(t *testing.T) {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}
func TestArrayOther(t *testing.T) {
	//清空数组

}
func zeroArray(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func zeroArray2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func TestMap(t *testing.T) {

}
