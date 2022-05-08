package main

import (
	"bufio"
	"calculator/calc"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("계산기 수식을 적어주세요\n")
	reader := bufio.NewReader(os.Stdin)
	exp, err := reader.ReadString('\n')
	exp = strings.TrimSpace(exp) // 문자열의 쓸모없는 양쪽 공백을 없앰
	exp = calc.IsValidInfix(exp)
	//fmt.Println(exp)
	if err != nil {
		panic("수식이 올바르지 않습니다.")
	}

	postfix := calc.ToPostfix(exp)
	fmt.Printf("결과 : %d", calc.Calculate(postfix))

}
