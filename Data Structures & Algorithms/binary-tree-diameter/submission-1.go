/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {

	var dfs func(*TreeNode)int

	res := 0

	dfs = func(root *TreeNode)int{
		if root == nil{
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		res = max(res, left + right)

		return 1 + max(left, right)
	}

	dfs(root)
	return res
    
}
