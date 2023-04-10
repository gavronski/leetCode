package validparentheses

import "strings"

// check if number of opened bracekts is same as closed brackets
func IsValid(s string) bool {

	parenthesesStack := Stack{}
	// push all opened brackets to the stack
	for _, char := range s {
		if strings.ContainsRune("{[(", char) {
			parenthesesStack.Push(string(char))
		}
	}

	// pop item from the stack if closed bracket are in a given string
	for _, char := range s {
		if strings.ContainsRune("}])", char) {
			if !parenthesesStack.IsEmpty() {
				parenthesesStack.Pop()
			}
		}
	}

	// if number of openings and closings are equal return true
	return parenthesesStack.IsEmpty()
}

type Stack struct {
	Items []string
}

func (s *Stack) Push(item string) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() {
	s.Items = s.Items[:len(s.Items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
