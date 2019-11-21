/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.05%)
 * Likes:    1188
 * Dislikes: 0
 * Total Accepted:    155.5K
 * Total Submissions: 388K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
package main

import "fmt"

func main() {
	fmt.Println("{[]}", isValid("{[]}"))
	fmt.Println("([)]", isValid("([)]"))
	fmt.Println("(]", isValid("(]"))
	fmt.Println("", isValid(""))
}

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	charMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	chars := []string{}
	for idx, c := range s {
		cStr := string(c)
		v, ok := charMap[cStr]
		if idx == 0 && ok {
			return false
		}
		if !ok {
			chars = append(chars, cStr)
		} else {
			l := len(chars)
			last := chars[l-1]
			if last == v {
				chars = chars[:l-1]
			} else {
				return false
			}
			if len(chars) == 0 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
