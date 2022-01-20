package ds

import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (n *treeNode) add(val int) {

	switch {
	case val < n.value:
		if n.left == nil {
			//log.Printf("Adding val:=%d to left of Node:=%d\n", val, n.value)
			n.left = &treeNode{value: val, left: nil, right: nil}
		} else {
			n.left.add(val)
		}

	case val > n.value:
		if n.right == nil {
			//log.Printf("Adding val:=%d to right of Node:=%d\n", val, n.value)
			n.right = &treeNode{value: val, left: nil, right: nil}
		} else {
			n.right.add(val)
		}

	case val == n.value:
		fmt.Printf("%d value already exists\n", val)
	}

}

func (n *treeNode) delete(val int) {

	switch {
	case val < n.value:
		if n.left == nil {
			return
		} else {
			n.left.delete(val)
		}

	case val > n.value:
		if n.right == nil {
			return
		} else {
			n.right.delete(val)
		}

	case val == n.value:
		//TODO: Add remove logic
	}
}

type Btree struct {
	root *treeNode
}

func NewBTree() Btree {
	return Btree{root: nil}
}

func (bt *Btree) Add(val int) {

	if bt.root == nil {
		bt.root = &treeNode{value: val,
			left:  nil,
			right: nil}
	} else {
		bt.root.add(val)
	}
}
func (bt *Btree) Display() {
	bt.root.Show()
}
func (n *treeNode) Show() {

	if n.left == nil && n.right == nil {
		fmt.Println(n.value)
		return
	}

	if n.left != nil {
		n.left.Show()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.Show()
	}

}
