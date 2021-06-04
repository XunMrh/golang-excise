package main

import "strconv"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (root *treeNode) AppendValue(value []int) []int {
	if root == nil {
		return []int{}
	}
	value = root.left.AppendValue(value)
	value = append(value, root.value)
	value = root.right.AppendValue(value)
	return value
}

func (root *treeNode) Add(value int) *treeNode {
	if root == nil {
		root = new(treeNode)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = root.left.Add(value)
	} else {
		root.right = root.right.Add(value)
	}
	return root
}

func (root *treeNode) Strings() string {
	var str string
	if root == nil {
		return str
	}
	str += root.left.Strings()
	str += strconv.Itoa(root.value)
	str += root.left.Strings()
	return str
}
