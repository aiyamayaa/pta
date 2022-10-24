#206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：

```go
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
```

示例 2：

```go
输入：head = [1,2]
输出：[2,1]
```

示例 3：
```go
输入：head = []
输出：[]
```

提示：
```go
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
```



进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

##思路
###思路1 双指针的方式实现
定义一个pre指针和一个cur指针，cur指针的Next指向pre指针，然后两个指针向后移动

###思路2 递归的方式实现