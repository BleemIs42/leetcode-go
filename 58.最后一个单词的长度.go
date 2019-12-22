/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (31.50%)
 * Likes:    155
 * Dislikes: 0
 * Total Accepted:    57.2K
 * Total Submissions: 179.9K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */

package main

import "fmt"

func main() {
	s := "Hello World "
	fmt.Println(s, lengthOfLastWord(s))
}

// @lc code=start
func lengthOfLastWord(s string) int {
	count := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if count > 0 {
				return count
			}
		} else {
			count++
		}
	}
	return count
}

// @lc code=end
