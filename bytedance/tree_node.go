package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. 二叉树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	right := LowestCommonAncestor(root.Right, p, q)
	left := LowestCommonAncestor(root.Left, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

//  700. 二叉搜索树中的搜索
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}
