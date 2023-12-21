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

func printInOrderTraversal(Node *TreeNode) {
	if Node == nil {
		return
	}
	printInOrderTraversal(Node.LeftNode)
	fmt.Println(Node.Data)
	printInOrderTraversal(Node.RightNode)
}

func printPreOrderTraversal(Node *TreeNode) {
	if Node == nil {
		return
	}
	fmt.Println(Node.Data)
	printInOrderTraversal(Node.LeftNode)
	printInOrderTraversal(Node.RightNode)
}

func printPostOrderTraversal(Node *TreeNode) {
	if Node == nil {
		return
	}
	printInOrderTraversal(Node.LeftNode)
	printInOrderTraversal(Node.RightNode)
	fmt.Println(Node.Data)
}

func (bst *BinaryTree) PrintTree() {
	root := bst.Root
	// root1 := root
	// root2 := root
	fmt.Println("Printing InOrder Traversal")
	printInOrderTraversal(root)
	fmt.Println("Printing PostOrder Traversal")
	printPostOrderTraversal(root)
	fmt.Println("Printing PreOrder Traversal")
	printPreOrderTraversal(root)
}
