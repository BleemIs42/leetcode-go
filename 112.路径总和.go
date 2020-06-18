/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (48.20%)
 * Likes:    317
 * Dislikes: 0
 * Total Accepted:    83.1K
 * Total Submissions: 166.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \      \
 * ⁠       7    2      1
 *
 *
 * 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
	res := hasPathSum(&tree, 3)
	fmt.Println(res)
}

// @lc code=start
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumQ(root, sum)

	if root == nil {
		return false
	}
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func hasPathSumQ(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queueNode := []*TreeNode{root}
	queueSum := []int{sum}
	for len(queueNode) > 0 {
		l := len(queueNode)
		for i := 0; i < l; i++ {
			node := queueNode[i]
			sum = queueSum[i] - node.Val
			if node.Left == nil && node.Right == nil && sum == 0 {
				return true
			}
			if node.Left != nil {
				queueNode = append(queueNode, node.Left)
				queueSum = append(queueSum, sum)
			}
			if node.Right != nil {
				queueNode = append(queueNode, node.Right)
				queueSum = append(queueSum, sum)
			}
		}
		queueNode = queueNode[l:]
		queueSum = queueSum[l:]
	}
	return false
}

// @lc code=end
