package calc

import (
	"calculator/stack"
	"fmt"
)

func isOp(o int32) bool {
	if o == '+' || o == '-' || o == '*' || o == '/' || o == '(' || o == ')' || o == ' ' {
		return true
	}
	return false
}

func powerOften(p int) int32 {
	var pow int32 = 1
	for i := 0; i < p; i++ {
		pow *= 10
	}
	return pow
}
func createnumber(s []int32, combo int) int32 { //배열의 숫자를 하나의 숫자로 만드는 함수
	var number int32 = 0
	pow := combo - 1
	for i := 0; i < combo; i++ {
		number += s[i] * powerOften(pow-i)
	}
	return number
}

func opOder(op int32) int {
	switch op {
	case '(':
		return 0
	case '+':
	case '-':
		return 1
	case '*':
	case '/':
		return 2
	}
	panic("invalid operator in opOrder")
}

func ToPostfix(str string) []int32 { //중위표기식을 후위표기식으로 바꾸는 함수
	opStack := &stack.Stack{nil, 0}
	numberArr := make([]int32, 100)
	postfix := make([]int32, 100)
	var postidx int = 0
	var numarridx int = 0
	var combo int = 0

	for _, v := range str {
		if v >= '0' && v <= '9' {
			combo++
			numberArr[numarridx] = v
		} else {
			disting := isOp(v) //연산자인지 검사
			if disting == false {
				panic(fmt.Errorf("Invalid operator"))
			} else {
				postfix[postidx] = createnumber(numberArr, combo) //임시로 저장한 숫자배열을 하나의 숫자로 만들어 저장
				postidx++
				postfix[postidx] = -1 /* 확인! >> 숫자와 연산자를 구별해주기 위해 -1로 간격을 만듦   <<*/
				postidx++
				combo = 0 //combo 초기화
				//연산자들 정렬을 stack을 이용하여 처리
				if opStack.Size == 0 || v == '(' {
					opStack.Push(v)
				} else if v == ')' {
					pop, _ := opStack.Pop()
					for pop != '(' {
						postfix[postidx] = pop
						postidx++
						pop, _ = opStack.Pop()
					}
				} else {
					top, _ := opStack.Peek()
					for opOder(v) < opOder(top) {
						pop, _ := opStack.Pop()
						postfix[postidx] = pop
						postidx++
						top, _ = opStack.Peek()
					}
					opStack.Push(v)

				}
			}
		}
	}
	return postfix
}
