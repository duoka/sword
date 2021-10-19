package array

import "sort"

// 题目描述
// 在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字是重复的，
// 也不知道每个数字重复几次。请找出数组中任意一个重复的数字
// Input: {2, 3, 1, 0, 2, 5}
// Output: 2

// findRepeatNumberMap .
// 判断map key
func findRepeatNumberMap(nums []int) int {
	temp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := temp[num]; ok {
			return num
		}
		temp[num] = struct{}{}
	}
	return -1
}

// findRepeatNumberIndex .
// 交换下标索引，需要被交换的元素相等即可得出结果
func findRepeatNumberIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			// 需要交换的元素相等
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// findRepeatNumberSort .
func findRepeatNumberSort(nums []int) int {
	// 排序
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
