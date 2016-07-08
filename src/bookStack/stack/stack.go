package stack

import "errors"

// Stack implements data structure
type Stack []interface{}

// Pop removes element from stack
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}

// Push adds element to stack
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top Returns the value on the top of the stack
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

// Cap returns the capacity of the stack
func (stack Stack) Cap() int {
	return cap(stack)
}

// Len returns the length of the stack
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty returns true if the stack is empty
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
