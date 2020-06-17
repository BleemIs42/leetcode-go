/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (40.56%)
 * Likes:    271
 * Dislikes: 0
 * Total Accepted:    82.5K
 * Total Submissions: 193.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
 *
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
		},
		Right: &TreeNode{
			Val:  1,
			Left: nil,
		},
	}
	res := minDepth(&tree)
	fmt.Println(res)
}

// @lc code=start
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDeep := minDepth(root.Left)
	rightDeep := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return 1 + leftDeep + rightDeep
	}
	return 1 + minInt(leftDeep, rightDeep)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
