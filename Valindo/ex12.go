package main

import "fmt"

const stackSize = 5

const bottomOfStack = -1

type stack struct {
	position int
	element  [stackSize]int
}

func (s *stack) push(item int) {
	if s.position == stackSize {
		fmt.Println("Stack is full")
		return
	}
	s.position += 1
	s.element[s.position] = item
	return
}

func (s *stack) pop() {
	if s.position == bottomOfStack {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Value popped is ", s.element[s.position])
	s.position -= 1
	return
}

func main() {
	var s stack
	s.position = bottomOfStack
	s.push(1)
	s.push(2)
	s.pop()
	s.pop()
	s.push(6)
	s.pop()
}
