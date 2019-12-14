/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (39.03%)
 * Likes:    298
 * Dislikes: 0
 * Total Accepted:    100.6K
 * Total Submissions: 257.8K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("mississippi", "issi"))
}

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	nl := len(needle)
	hl := len(haystack)
	max := hl - nl

	if max < 0 {
		return -1
	}

	// for i := 0; i <= max; i++ {
	// 	if haystack[i:i+nl] == needle {
	// 		return i
	// 	}
	// }

	// KMP
	for i := 0; i <= max; {
		if haystack[i:i+nl] == needle {
			return i
		} else {
			if i+nl < hl {
				idx := strings.LastIndex(needle, string(haystack[i+nl]))
				i += nl - idx
			} else {
				return -1
			}
		}
	}
	return -1
}

// @lc code=end
