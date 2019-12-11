/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (58.25%)
 * Likes:    714
 * Dislikes: 0
 * Total Accepted:    140.5K
 * Total Submissions: 240.8K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * 	Val  int
 * 	Next *ListNode
 * }
 */

// 需要注释多余代码才能测试成功

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := newLinkList([]int{1, 2, 4})
	l2 := newLinkList([]int{1, 3, 4})
	fmt.Println(walker(l1, []int{}))
	fmt.Println(walker(l2, []int{}))
	ll := mergeTwoLists(l1, l2)
	fmt.Println(walker(ll, []int{}))
}

func newLinkList(list []int) *ListNode {
	var head *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		head = &ListNode{
			Val:  list[i],
			Next: head,
		}
	}
	return head
}

func walker(ll *ListNode, values []int) []int {
	if ll == nil {
		return values
	}
	values = append(values, ll.Val)
	return walker(ll.Next, values)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// @lc code=end
