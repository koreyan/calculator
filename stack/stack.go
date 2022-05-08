package stack

import "fmt"

type node struct {
	data int32
	next *node
}

type Stack struct {
	top  *node
	size int
}

func Error(s string) error {
	return fmt.Errorf("%s", s)
}

func (stk *Stack) Wsize() int {
	return stk.size
}

func (stk *Stack) Peek() (int32, error) {
	if stk.size == 0 {
		return -1, Error("There is no data")
	}
	return stk.top.data, nil
}

func (stk *Stack) Push(newdata int32) {
	newNode := &node{newdata, stk.top}
	stk.top = newNode
	stk.size++
}

func (stk *Stack) Pop() (int32, error) {
	if stk.size == 0 {
		return -1, Error("There is no data")
	}
	temp := stk.top.data
	stk.top = stk.top.next
	stk.size--
	return temp, nil
}
