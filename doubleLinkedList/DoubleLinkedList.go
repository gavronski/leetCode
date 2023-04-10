package doubleLinkedList

import (
	"log"
)

type Node struct {
	Next *Node
	Prev *Node
	Data string
}

type DoubleLinkedList struct {
	First *Node
	Last  *Node
	Total int
}

func (ll *DoubleLinkedList) InsertFirst(data string) {
	var node *Node = &Node{Data: data}

	// check if first element exists
	if ll.First == nil {
		ll.First = node
		ll.Last = node
	} else {
		first := ll.First
		first.Prev = node
		node.Next = first
		ll.First = node
	}
	ll.Total++
}

func (ll *DoubleLinkedList) InsertBefore(data, query string) {
	var node *Node = &Node{Data: data}

	if ll.First == nil {
		return
	}

	prev := &Node{}
	current := ll.First

	for current != nil {
		if current.Data == query {
			if current == ll.First {
				ll.InsertFirst(data)
			} else {
				prev.Next = node
				node.Prev = prev

				node.Next = current
				current.Prev = node

			}
			ll.Total++
			break
		}

		prev = current
		current = current.Next
	}
}

func (ll *DoubleLinkedList) InsertLast(data string) {
	var node *Node = &Node{Data: data}

	if ll.First == nil {
		ll.First = node
		ll.Last = node
	} else {
		ll.Last.Next = node
		node.Prev = ll.Last
		ll.Last = node
	}

	ll.Total++
}

func (ll *DoubleLinkedList) InsertAfter(data, query string) {
	var node *Node = &Node{Data: data}

	if ll.Total == 0 {
		return
	}

	current := ll.First

	for current != nil {
		if current.Data == query {
			if current == ll.Last {
				ll.InsertLast(data)
			} else {
				node.Prev = current
				current.Next.Prev = node
				node.Next = current.Next
				current.Next = node
			}
			ll.Total++
			break
		}

		current = current.Next
	}

}

func (ll *DoubleLinkedList) DeleteFirst() {
	if ll.Total == 0 {
		return
	}

	if ll.Total == 1 {
		ll.First = &Node{}
		ll.Last = &Node{}
		ll.Total--
		return
	}

	first := ll.First
	second := ll.First.Next
	second.Prev = nil
	ll.First = second
	if first == ll.Last {
		ll.Last = ll.First
	}
	ll.Total--

}

func (ll *DoubleLinkedList) DeleteLast() {
	if ll.Total == 0 {
		return
	}

	if ll.Total == 1 {
		ll.First = &Node{}
		ll.Last = &Node{}
		ll.Total--
		return
	}

	prev := ll.Last.Prev
	prev.Next = nil
	ll.Last = prev
	ll.Total--
}

func (ll *DoubleLinkedList) Delete(query string) {
	if ll.Total == 0 {
		return
	}

	prev := &Node{}
	current := ll.First

	for current != nil {
		if current.Data == query {
			if current == ll.First {
				ll.DeleteFirst()
				return
			} else if current == ll.Last {
				ll.DeleteLast()
				return
			} else {
				next := current.Next
				prev.Next = next
				next.Prev = prev
			}

			ll.Total--
			break
		}
		prev = current
		current = current.Next
	}
}

func (ll *DoubleLinkedList) Reverse() {
	if ll.Total == 0 {
		return
	}

	current := ll.Last

	for current != nil {
		log.Println(current.Data)

		current = current.Prev
	}
}
func (ll *DoubleLinkedList) Traverse() {
	if ll.Total == 0 {
		return
	}
	current := ll.First

	for current != nil {
		log.Println(current.Data)
		current = current.Next
	}
}
