package main

import (
	binarytree "github.com/Prthmesh6/DSA_Golang/BinaryTree"
	linkedlist "github.com/Prthmesh6/DSA_Golang/LinkedList"
)

func main() {
	checkLinkedList()
	checkTree()
}

func checkTree() {
	BinarySearchTree := binarytree.NewBinaryTree(1)
	BinarySearchTree.InsertNode(2)
	BinarySearchTree.InsertNode(4)
	BinarySearchTree.InsertNode(3)
	BinarySearchTree.InsertNode(-1)

	BinarySearchTree.PrintTree()
}

func checkLinkedList() {
	LinkedList := linkedlist.NewLinkedListHead(1)

	LinkedList.InsertFront(0)
	LinkedList.InsertFront(-1)
	LinkedList.InsertFront(-2)
	LinkedList.InsertBack(2)
	LinkedList.InsertBack(3)

	LinkedList.PrintElements()
}
