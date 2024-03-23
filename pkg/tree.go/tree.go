package tree

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{value: val, left: nil, right: nil}
}

func (t *TreeNode) GetValue() int {
	return t.value
}

func (t *TreeNode) GetLeft() *TreeNode {
	return t.left
}

func (t *TreeNode) GetRight() *TreeNode {
	return t.right
}

func (t *TreeNode) Insert(val int) {
	if val < t.value {
		if t.left == nil {
			t.left = NewTreeNode(val)
		} else {
			t.left.Insert(val)
		}
	} else {
		if t.right == nil {
			t.right = NewTreeNode(val)
		} else {
			t.right.Insert(val)
		}
	}
}

func (t *TreeNode) Search(val int) bool {
	if t == nil {
		return false
	}

	if val == t.value {
		fmt.Print("WE FOUND IT")
		return true
	} else if val < t.value {
		return t.left.Search(val)
	} else {
		return t.right.Search(val)
	}
}

func (t *TreeNode) InorderTraversal() []int {
	var result []int
	t.inorderTraversalHelper(&result)
	return result

}

func (t *TreeNode) inorderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	t.left.inorderTraversalHelper(result)
	*result = append(*result, t.value)
	t.right.inorderTraversalHelper(result)

}

func (t *TreeNode) PostorderTraversal() []int {
	var result []int
	t.postOrderTraversalHelper(&result)
	return result
}

func (t *TreeNode) postOrderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	t.left.postOrderTraversalHelper(result)
	t.right.postOrderTraversalHelper(result)
	*result = append(*result, t.value)

}

func (t *TreeNode) PreorderTraversal() []int {
	var result []int
	t.preOrderTraversalHelper(&result)
	return result
}

func (t *TreeNode) preOrderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	*result = append(*result, t.value)
	t.left.preOrderTraversalHelper(result)
	t.right.preOrderTraversalHelper(result)

}
