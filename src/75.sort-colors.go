package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			swap(nums, i, left)
			left++
			i++
		} else if nums[i] == 2 && i < right {
			swap(nums, i, right)
			right--
		} else {
			i++
		}
	}
}

func swap(nums []int, a int, b int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

// @lc code=end
