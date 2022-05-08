package calc

import (
	"calculator/stack"
	"fmt"
)

func isOp(o int32) bool {
	if o == '+' || o == '-' || o == '*' || o == '/' {
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
		number += (s[i] - '0') * powerOften(pow-i)
	}
	return number
}

func opOder(op int32) int {
	//fmt.Printf("%c\n", op)
	switch op {
	case '(':
		return 0

	case '+':
		return 1
	case '-':
		return 1

	case '*':
		return 2
	case '/':
		return 2

	}
	panic(fmt.Errorf("%c invalid operator in opOrder", op))
}

func ToPostfix(str string) []int32 { //중위표기식을 후위표기식으로 바꾸는 함수
	//fmt.Println(str)
	opStack := new(stack.Stack)
	numberArr := make([]int32, 100)
	postfix := make([]int32, 0, 100)
	var combo int = 0

	for _, v := range str {
		//fmt.Printf("%c 차례\n\n", v)
		if v >= '0' && v <= '9' {
			numberArr[combo] = v
			combo++
			//fmt.Println(v, "numArr에 임시 저장")
		} else if v == '(' {
			opStack.Push(v)
		} else if v == ')' {
			tempnum := createnumber(numberArr, combo)
			//fmt.Println(tempnum, "숫자 생성")
			postfix = append(postfix, tempnum) //임시로 저장한 숫자배열을 하나의 숫자로 만들어 저장
			combo = 0                          //combo 초기화

			pop, _ := opStack.Pop()
			for pop != '(' {
				postfix = append(postfix, -1) /* 확인! >> 숫자와 연산자를 구별해주기 위해 -1로 간격을 만듦   <<*/
				postfix = append(postfix, pop)
				pop, _ = opStack.Pop()
			}
			//fmt.Println("postfix 현황", postfix)
		} else {
			if combo != 0 {
				tempnum := createnumber(numberArr, combo)
				//fmt.Println(tempnum, "숫자 생성")
				postfix = append(postfix, tempnum) //임시로 저장한 숫자배열을 하나의 숫자로 만들어 저장
				combo = 0                          //combo 초기화
				//fmt.Println("postfix 현황", postfix)
			}
			//연산자들 정렬을 stack을 이용하여 처리
			if opStack.Wsize() == 0 {
				//fmt.Println("opStack is empty")
				opStack.Push(v)
			} else {
				top, _ := opStack.Peek()
				for opOder(v) <= opOder(top) {
					pop, _ := opStack.Pop()
					postfix = append(postfix, -1)
					postfix = append(postfix, pop)
					if opStack.Wsize() == 0 {
						break
					}
					top, _ = opStack.Peek()
				}
				opStack.Push(v)

			}
		}
	}
	if combo != 0 {
		postfix = append(postfix, createnumber(numberArr, combo)) //위 알고리즘으로 수식의 맨 뒤 숫자는 postfix에 포함할 수 없으므로 여기서 추가
	}
	var temp int32
	for opStack.Wsize() != 0 {
		temp, _ = opStack.Pop()
		postfix = append(postfix, -1)
		postfix = append(postfix, temp)
	}

	//fmt.Println("최종 postfix", postfix)
	return postfix
}

func operate(oprd1, oprd2, op int32) int32 {
	switch op {
	case '+':
		return oprd1 + oprd2
	case '-':
		return oprd1 - oprd2
	case '*':
		return oprd1 * oprd2
	case '/':
		return oprd1 / oprd2
	default:
		panic("operate function error!! invalid operator")
	}
}

func Calculate(postfix []int32) int32 {
	numStack := new(stack.Stack)
	var disc int32 = 0 // 연산자와 피연산자 사이를 식별하기 위한 변수
	var oprd1 int32
	var oprd2 int32
	var result int32
	for _, v := range postfix {
		if v == -1 { //다음이 연산자
			disc = 1
			continue
		}
		if disc == 1 { //연산자인 경우
			disc = 0
			oprd2, _ = numStack.Pop()
			oprd1, _ = numStack.Pop()
			result = operate(oprd1, oprd2, v)
			//fmt.Printf("%d %c %d = %d\n", oprd1, v, oprd2, result)
			numStack.Push(result)
		} else {
			numStack.Push(v)
		}

	}
	pop, err := numStack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	return pop
}

func IsValidInfix(infix string) string {
	fixedExp := make([]rune, 0, 100)
	parenthesesStack := new(stack.Stack)
	opCombo := 0
	if isOp(int32(infix[0])) {
		panic("Invalid expression! operator is placed at top of expression")
	}

	for _, v := range infix {

		if v == ' ' {
			continue
		}

		if v >= '0' && v <= '9' {
			fixedExp = append(fixedExp, v)
			opCombo = 0
			continue
		} else if v == '(' {
			fixedExp = append(fixedExp, v)
			parenthesesStack.Push(v)
			opCombo++
			continue
		} else if v == ')' {

			p, _ := parenthesesStack.Pop()
			if p != '(' {
				panic("Invalid expression!  소괄호의 짝이 맞지 않습니다.")
			}
			fixedExp = append(fixedExp, v)
			continue
		}

		if isOp(v) && opCombo == 0 {
			opCombo++
			fixedExp = append(fixedExp, v)

		} else {
			panic("Invalid expression! operators are placed consecutively")
		}
	}
	//fmt.Println(fixedExp)
	return string(fixedExp)
}
