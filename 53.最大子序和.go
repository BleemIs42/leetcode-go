/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (48.00%)
 * Likes:    1421
 * Dislikes: 0
 * Total Accepted:    132.6K
 * Total Submissions: 274.3K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret1 := maxSubArray(nums)
	ret2 := maxSubArrayDP(nums)
	ret3 := maxSubArrayDAC(nums)
	fmt.Println(nums, ret1, ret2, ret3)
}

// @lc code=start
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	max := nums[0]
	sum := max
	for _, v := range nums[1:] {
		sum = maxInt(v, v+sum)
		max = maxInt(max, sum)
	}
	return max
}

// @lc code=end

func maxSubArrayDP(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := 1; i <= len(nums[1:]); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		max = maxInt(max, nums[i])

	}
	return max
}

func maxSubArrayDAC(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	return helper(nums, 0, l-1)
}

func helper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2

	leftSum := helper(nums, left, mid)
	rightSum := helper(nums, mid+1, right)
	crossSum := crossSum(nums, left, right, mid)
	return maxInt(maxInt(leftSum, rightSum), crossSum)
}
func crossSum(nums []int, left, right, mid int) int {
	if left == right {
		return nums[left]
	}

	halfMin := ^(int(^uint(0) >> 1)) / 2

	currNum := halfMin
	maxLeft := currNum
	for i := mid; i >= left; i-- {
		currNum += nums[i]
		maxLeft = maxInt(currNum, maxLeft)
	}

	currNum = halfMin
	maxRight := currNum
	for i := mid + 1; i <= right; i++ {
		currNum += nums[i]
		maxRight = maxInt(currNum, maxRight)
	}

	return maxLeft + maxRight
}
