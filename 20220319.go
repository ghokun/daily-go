package dailygo

type _20220319_node struct {
	Value string
	Left  *_20220319_node
	Right *_20220319_node
}

/*
A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

	  0
	 / \
	1   0
	   / \
	  1   0
	 / \
	1   1
*/
func _20220319(binaryTree *_20220319_node) int {
	if binaryTree == nil {
		return 0
	}
	left := binaryTree.Left
	right := binaryTree.Right

	if left != nil && right != nil {
		count := 0
		if left.Value == right.Value {
			count = 1
		}
		return count + _20220319(left) + _20220319(right)
	}
	if left != nil {
		return _20220319(left)
	}
	if right != nil {
		return _20220319(right)
	}
	return 1 // both nil
}
