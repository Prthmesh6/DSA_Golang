package binarytree

import "fmt"

type TreeNode struct {
	Data      int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(Data int) *BinaryTree {
	return &BinaryTree{
		Root: &TreeNode{
			Data: Data,
		},
	}
}

func (bst *BinaryTree) InsertNode(Data int) {
	bst.Root = insertNode(bst.Root, Data)
}

func insertNode(Node *TreeNode, Data int) *TreeNode {
	if Node == nil {
		return &TreeNode{Data: Data, LeftNode: nil, RightNode: nil}
	}

	if Node.Data > Data {
		Node.LeftNode = insertNode(Node.LeftNode, Data)
	}
	if Node.Data < Data {
		Node.RightNode = insertNode(Node.RightNode, Data)
	}

	return Node
}

func printBst(Node *TreeNode) {
	if Node == nil {
		return
	}
	printBst(Node.LeftNode)
	fmt.Println(Node.Data)
	printBst(Node.RightNode)
}

func (bst *BinaryTree) PrintTree() {
	printBst(bst.Root)
}
