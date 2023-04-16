package mergeTwoSortedList

type List struct {
	First *ListNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *List) InsertNode(node *ListNode) {
	if l.First == nil {
		l.First = node
		return
	}

	current := l.First

	for current != nil {
		current = current.Next
	}

	current.Next = node
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current1 := list1
	current2 := list2

	list := List{}
	for current1 != nil {
		for current2 != nil {
			if current1.Val < current2.Val {
				list.InsertNode(current1)
			} else {
				list.InsertNode(current2)
			}
			current2 = current2.Next
		}

		current1 = current1.Next
	}

	return list.First
}
