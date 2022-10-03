#33. Search in Rotated Sorted Array
## 题目
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
```azure
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

```
Example 2:
```azure
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```
## 题目
### 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标` k（0 <= k < nums.length）`上进行了 旋转，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 从 0 开始 计数）。
例如， `[0,1,2,4,5,6,7] `在下标 3 处经旋转后可能变为 `[4,5,6,7,0,1,2] `。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 `O(log n)` 的算法解决此问题。
示例 1：
```azure
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
```
示例 2：
```azure
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
```
示例 3：
```azure
输入：nums = [1], target = 0
输出：-1
```

##题目大意
在一个旋转的数组中，找出数字target的下标。没有返回-1。（数组没有重复）

##思路
可以用二分法，但是首先要判断，middle是在旋转后的数组的左边还是右边
- 左边：`nums[middle]>=nums[0]`
- 右边：`nums[middle]<nums[0]`

然后判断`target`与`nums[middle]`的大小关系
- 左边：
  - 如果`nums[0]<target<nums[middle]，right=m`
  - 否则，`left=m+1`
- 右边
  - 如果`nums[middle] < target < nums[len(nums)-1]，left = m+1`
  - 否则，`right=m`