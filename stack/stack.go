package stack

import "log"

type Stack struct {
	Items []int
	Limit int
}

func (stack *Stack) Push(item int) {
	length := len(stack.Items)
	if length >= stack.Limit {
		log.Println("Stack is overflow")
		return
	}

	stack.Items = append(stack.Items, item)
}

func (stack *Stack) Pop() {
	if stack.isEmpty() {
		log.Println("Stack is underflow")
		return
	}

	stack.Items = stack.Items[:len(stack.Items)-1]
}

func (stack *Stack) isEmpty() bool {
	return len(stack.Items) == 0
}

func (stack *Stack) Top() int {
	if stack.isEmpty() {
		return 0
	}

	length := len(stack.Items)
	return stack.Items[length-1]
}
