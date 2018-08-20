package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// No validation s is valid rever polish expression
func Eval(s []string) int {

	if len(s) <= 0 {
		return 0
	}

	stack := []string{}
	for _, v := range s {
		if IsDigit(v) {
			stack = append(stack, v)
		} else {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, Operate(a, b, v))
		}
	}
	exprVal, _ := strconv.ParseInt(stack[0], 10, 32)
	return int(exprVal)
}

func IsDigit(s string) bool {
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

func Operate(a, b, op string) string {
	op1, _ := strconv.ParseInt(a, 10, 32) // omitting validation
	op2, _ := strconv.ParseInt(b, 10, 32) // omitting validation

	var v int64

	switch op {
	case "+":
		v = op1 + op2
	case "-":
		v = op1 - op2
	case "*":
		v = op1 * op2
	case "/":
		v = op1 / op2
	default:
		panic(op + " operator not allowed")
	}

	return strconv.Itoa(int(v))
}

func main() {
	v := Eval([]string{"22", "1", "+", "3", "*"})
	fmt.Println(v)

	v = Eval([]string{"4", "13", "5", "/", "+"})
	fmt.Println(v)

}
