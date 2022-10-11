#数组排序的简单方法-sort.Slice

1.通常以实现sort.Interface接口中3个方法的形式实现二维数组的排序是较为繁琐的，可以通过sort.Slice方法快速实现二维数组的排序。

2.sort.Slice方法有两个参数：
- 待排序的二维数组，
- 类似Less方法的方法（比较器）： `func(i,j int) bool`，如果索引`i`对应的值在索引`j`对应的值前面，返回`true`，否则返回`false`

如：
```go
//定义二维数组

var intervals = [][]int{{2,3},{1,2},{4,5},{3,4}}

//对该数组以每行第一个元素升序排列的顺序进行排序

sort.Slice(intervals, func(i,j int) bool {

return intervals[i][0] < intervals[j][0]

})
```


3.使用sort.Slice方法不必实现上述三种方法，使用简便。