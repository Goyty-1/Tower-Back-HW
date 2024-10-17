package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Add(num int) *Node {
	if n == nil {
		return &Node{Value: num}
	}
	if num < n.Value {
		n.Left = n.Left.Add(num)
	} else if num > n.Value {
		n.Right = n.Right.Add(num)
	}
	return n
}

func (n *Node) IsExist(num int) bool {
	if n == nil {
		return false
	}
	if num == n.Value {
		return true
	} else if num < n.Value {
		return n.Left.IsExist(num)
	} else {
		return n.Right.IsExist(num)
	}
}

func (n *Node) Delete(num int) *Node {
	if n == nil {
		return nil
	}

	if num < n.Value {
		n.Left = n.Left.Delete(num)

	} else if num > n.Value {
		n.Right = n.Right.Delete(num)

	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		}
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
	}
	return n
}

func main() {
	root := &Node{}

	root.Add(77)

	fmt.Println(root.IsExist(77))
	fmt.Println(root.IsExist(25))

	root.Delete(77)

	fmt.Println(root.IsExist(77))
}
