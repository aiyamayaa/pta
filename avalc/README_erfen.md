#二分法

二分法对于数字的处理，一般是有两种定义
- 左闭右闭`[left，right]`：下标为0，1，2，3，4，对这5个数搜索
- 左闭右开`[left，right)`：下标为0，1，2，3，4，对前4个数搜索

##模版1-左闭右闭
```go
//左闭右闭，表示right指向的值需要被处理
left,right:=0,len(nums)-1
for left<=right{
	middle := left + (right-left)>>2
	// 更新右边界
	if nums[middle]>target{
		//因为right需要被搜索，但是middle已经不是要找的值了
        right = middle -1
	}else if nums[middle]<target{
		left = middle + 1
	}else{
		return middle
	}
}

```

##模版2-左闭右开
```go
//左闭右开，表示right指向的值不需要被处理
left,right:=0,len(nums)
//rigth不需要被搜索，left需要被搜索，left不要等于right
for (left<right){
	middle := left + (right-left)>>2 
	// 更新右边界 
	if nums[middle]>target{
		//因为right不需要被搜索，相当于middle已经不是要找的值了，虽然rigth = midlle,但是不搜索它 
		right = middle
	}else if nums[middle]<target{
		left = middle + 1
	}else{
		return middle
	}
}

```
##模版3-左开右闭
```go
//左闭右闭，表示right指向的值需要被处理
left,right:=0,len(nums)-1
for left+1<right{
	middle := left + (right-left)>>2
	// 更新右边界
	if nums[middle]<target{
		//因为right需要被搜索，但是middle已经不是要找的值了
        left = middle 
	}else {
		right = middle 
	}
}

```

- 实现二分查找的另一种方法。 
- 搜索条件需要访问元素的直接左右邻居。 
- 使用元素的邻居来确定它是向右还是向左。 
- 保证查找空间在每个步骤中至少有 3 个元素。 
- 需要进行后处理。 当剩下 2 个元素时，循环 / 递归结束。 需要评估其余元素是否符合条件。
