package main

import (
	"app/twosum"
	validparentheses "app/validParentheses"
	"log"
)

func main() {
	// var ll doubleLinkedList.DoubleLinkedList = doubleLinkedList.DoubleLinkedList{}
	// ll.InsertFirst("Tom")

	// ll.InsertFirst("Basia")
	// ll.InsertFirst("Mark")

	// ll.InsertBefore("Ola", "Basia")
	// ll.InsertLast("Adam")
	// ll.InsertAfter("Tamara", "Mark")
	// ll.DeleteFirst()
	// ll.DeleteFirst()
	// ll.DeleteLast()
	// ll.Delete("Basia")
	// ll.Delete("Tom")
	// ll.Delete("Ola")

	// ll.Traverse()
	// log.Println(ll.First, ll.Last)
	// ll.Reverse()

	// var stack = stack.Stack{Limit: 3}
	// var queue = queue.Queue{Limit: 4, Total: 0}
	// items := []int{99, 120, 96, 85, 33333, 1, 40, 41, 123, 123, 33333, 1678}

	// for _, item := range items {
	// 	stack.Push(item)
	// 	queue.Enqueue(item)
	// }
	// queue.Traverse()
	// log.Println(queue.Peek())
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()

	// queue.Traverse()
	// log.Println("Sorting")
	// items = sorting.BubbleSort(items)
	// items = sorting.SelectionSort(items)
	// items = sorting.InsertionSort(items)
	// items = sorting.MergeSort(items)
	// log.Println(items)

	// log.Println(searching.BinarySearch(items, 40))
	// log.Println(searching.BinarySearch(items, 722))
	// log.Println(searching.BSRecursive(items, 40, len(items)-1, 0))
	// log.Println(searching.BSRecursive(items, 722, len(items)-1, 0))
	// log.Println(searching.FindOccurence(items, 33333))
	// log.Println(searching.FindOccurence(items, 722))

	// tree := tree.Tree{}
	// tree.AddChild(nil, "CEO")
	// tree.AddChild(tree.Root, "CEO's Assistant")
	// tree.AddChild(tree.Root, "CTO")
	// tree.AddChild(tree.Root.Children[1], "Senior Architect")
	// tree.AddChild(tree.Root.Children[1], "Senior Software Developer")

	// tree.AddChild(tree.Root.Children[0], "Assistant's supporter")
	// tree.Delete(tree.Root, "CTO")

	// tree.Traverse(tree.Root, 0)
	// stack.Pop()
	// stack.Pop()
	// stack.Pop()
	// stack.Pop()
	// someData := []int{34, 12, 1, 56, 9, 0, 141, 18, 7}
	// sortedData := sorting.MergeSort(someData)
	// fmt.Println(sortedData)
	// log.Println(stack.Items)
	// top := stack.Top()
	// log.Println(top)

	log.Println(twosum.TwoSum([]int{2, 7, 11, 15}, 9))
	log.Println(validparentheses.IsValid("()[]{}"))
}
