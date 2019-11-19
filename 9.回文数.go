/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (56.66%)
 * Likes:    815
 * Dislikes: 0
 * Total Accepted:    199.4K
 * Total Submissions: 352K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */

package main

import "fmt"

func main() {
	list := []int{101, 10, -101}
	for _, v := range list {
		fmt.Println(v, isPalindrome(v))
	}
}

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ret := 0
	temp := x
	for temp != 0 {
		ret = ret*10 + temp%10
		temp /= 10
	}
	return ret == x
}

// @lc code=end
