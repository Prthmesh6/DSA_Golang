package linkedlist

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode

	Tail *ListNode
}

func NewLinkedListHead(Data int) *LinkedList {
	LinkedList := &LinkedList{Head: &ListNode{
		Data: Data,
		Next: nil,
	}}
	LinkedList.Tail = LinkedList.Head
	return LinkedList
}

func (ll *LinkedList) InsertFront(Data int) {
	ll.Head = &ListNode{Data: Data, Next: ll.Head}
}

func (ll *LinkedList) InsertBack(Data int) {
	ll.Tail.Next = &ListNode{Data: Data, Next: nil}
	ll.Tail = ll.Tail.Next
}

func (ll *LinkedList) PrintElements() {
	originalHead := ll.Head
	for {
		if ll.Head == nil {
			break
		}
		fmt.Println(ll.Head.Data)
		ll.Head = ll.Head.Next
	}

	ll.Head = originalHead
}
