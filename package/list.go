package linkedList

import "fmt"

type List[T comparable] struct {
	Value T
	Next  *List[T]
}

func (node *List[T]) Append(Value T) {
	current := node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &List[T]{Value: Value}
}

func (node *List[T]) Len() int {
	len := 0
	current := node
	for current.Next != nil {
		current = current.Next
		len++
	}
	return len
}

func (node *List[T]) DelVal(Value T) {
	current := node
	var prev *List[T]
	for current != nil {
		if current.Value == Value {
			if prev == nil {
				if current.Next != nil {
					*node = *current.Next
				} else {
					node.Value = *new(T)
					node.Next = nil
				}
			} else {
				prev.Next = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}
}

func (node *List[T]) Reverse() {
	current := node
	var prev, Next *List[T]
	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}
	if prev != nil {
		*node = *prev
	}
}

func (node *List[T]) String() string {
	current := node
	result := ""
	for current != nil {
		result += fmt.Sprintf("%v -> ", current.Value)
		current = current.Next
	}
	result += "nil"
	return result
}
func main() {

	fmt.Println("vim-go")
}
