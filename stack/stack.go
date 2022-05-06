package stack

import "fmt"

type node struct {
	data int32
	next *node
}

type Stack struct {
	Top  *node
	Size int
}

func Error(s string) error {
	return fmt.Errorf("%s", s)
}

func (stk *Stack) Wsize() int {
	return stk.Size
}

func (stk *Stack) Peek() (int32, error) {
	if stk.Size == 0 {
		return -1, Error("There is no data")
	}
	return stk.Top.data, nil
}

func (stk *Stack) Push(newdata int32) {
	newNode := &node{newdata, stk.Top}
	stk.Top = newNode
	stk.Size++
}

func (stk *Stack) Pop() (int32, error) {
	if stk.Size == 0 {
		return -1, Error("There is no data")
	}
	temp := stk.Top.data
	stk.Top = stk.Top.next
	stk.Size--
	return temp, nil
}
