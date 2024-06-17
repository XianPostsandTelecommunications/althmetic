package main

//128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

func longestConsecutive(nums []int) int {
	numset := map[int]bool{}
	for _, num := range nums {
		numset[num] = true
	}
	longestStreak := 0
	for num := range numset {
		if !numset[num-1] {
			currentNum := num
			currentStreak := 1
			for numset[currentNum+1] {
				currentNum++

			}
		}
	}
}
