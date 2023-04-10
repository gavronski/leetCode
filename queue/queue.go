package queue

import "log"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Head  *Node
	Limit int
	Total int
}

func (q *Queue) Enqueue(data int) {
	var node *Node = &Node{Data: data}

	if q.Total >= q.Limit {
		log.Println("Queue is full")
		return
	}

	if q.Head == nil {
		q.Head = node
	} else {
		current := q.Head

		for current.Next != nil {
			current = current.Next
		}

		current.Next = node
	}

	q.Total++
}

func (q *Queue) Dequeue() {
	if q.isEmpty() {
		log.Println("Queue is empty")
		return
	}

	head := q.Head
	if head.Next != nil {
		q.Head = head.Next
	} else {
		q.Head = nil
	}
	q.Total--

}

func (q *Queue) Peek() *Node {
	if !q.isEmpty() {
		return q.Head
	}

	return nil
}

func (q *Queue) isEmpty() bool {
	return q.Total == 0
}

func (q *Queue) Traverse() {
	if q.isEmpty() {
		return
	}

	current := q.Head

	for current != nil {
		log.Println(current.Data)
		current = current.Next
	}
}
