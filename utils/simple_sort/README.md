#使用golang原生的sort

如果使用golang的sort对特定的结构类型进行排序，需要：
- 定义结构类型
- 实现三个方法
  - `Len()`：获取该数组的长度
  - `Swap(i,j int)`：对i,j下标的元素进行交换
  - `Less(i,j int) bool`：对i,j下标的元素进行比较，i在j前面返回true,否则返回false
    
当某结构体数组全部实现上述方法后，即可通过sort.Sort方法对该数组进行排序，排序顺序取决于Less方法的返回值。
```go

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

```