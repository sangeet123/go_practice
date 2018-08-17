package main

import (
	"fmt"
)

type node struct {
	val int
	l   *node
	r   *node
}

type BST struct {
	root *node
}

func getNode(val int) *node {
	b := new(node)
	b.l = nil
	b.r = nil
	b.val = val
	return b
}

func (n *node) add(val int) *node {
	if n == nil {
		n = getNode(val)
		return n
	}

	if n.val > val {
		n.l = n.l.add(val)
	} else {
		n.r = n.r.add(val)
	}

	return n
}

func (b *BST) Add(val int) {
	if b.root == nil {
		b.root = getNode(val)
		return
	}
	b.root.add(val)
}

func (n *node) append(values []int) []int {
	if n != nil {
		values = n.l.append(values)
		values = append(values, n.val)
		values = n.r.append(values)
	}
	return values
}

func (b *BST) Sort() []int {
	return b.root.append(nil)
}

func main() {
	var b BST

	b.Add(4)
	b.Add(3)
	b.Add(1)
	b.Add(5)
	b.Add(10)
	b.Add(2)

	fmt.Println(b.Sort())
}
