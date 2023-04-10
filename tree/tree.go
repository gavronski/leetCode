package tree

import (
	"log"
	"strings"
)

type Node struct {
	Data     string
	Children []*Node
}

type Tree struct {
	Root *Node
}

func (tree *Tree) AddChild(parent *Node, data string) {
	var node *Node = &Node{Data: data}

	if tree.Root == nil {
		tree.Root = node
		return
	}

	parent.Children = append(parent.Children, node)
}

func (tree *Tree) Traverse(node *Node, lvl int) {
	if tree.Root == nil {
		return
	}

	log.Println(strings.Repeat("-", lvl), node.Data)

	if len(node.Children) == 0 {
		return
	}
	lvl++

	for _, child := range node.Children {
		tree.Traverse(child, lvl)
	}
}

func (tree *Tree) Delete(parent *Node, data string) {
	if tree.Root == nil {
		return
	}
	if data == tree.Root.Data {
		tree.Root = nil
		return
	}

	if len(parent.Children) == 0 {
		return
	}

	for i, child := range parent.Children {
		if child.Data == data {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return
		}
		tree.Delete(child, data)
	}
}
