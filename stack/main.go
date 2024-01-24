package main

import "fmt"

//Stack is LIFO - last in first out

// Define a LIFO stack
type Stack struct {
	items []int
}

// Create a Push method to push item into the stack
func (stack *Stack) push(item int) {
	stack.items = append(stack.items, item)
}

// Create a Pop method to retrieve item from the stack
func (stack *Stack) pop(item int) int {
	items := stack.items
	lastItemPointer := len(items) - 1
	res := items[lastItemPointer]
	fmt.Printf("\nBefore: %v", items)
	items = items[:lastItemPointer]
	fmt.Printf("After: %v", items)
	return res
}

func main() {
	stack := Stack{items: []int{1, 2, 3, 4, 5}}
	stack.push(33)
	stack.pop(33)

}
