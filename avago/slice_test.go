package avago

import "testing"

/**
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
一个slice由三个部分构成：指针、长度和容量
	指针:指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。
	长度:对应slice中元素的数目；长度不能超过容量.len函数返回slice的长度
	容量:对应从slice的开始位置到底层数据的结尾位置。cap函数返回slice的容量。
*/
func TestSlice(t *testing.T) {

}
