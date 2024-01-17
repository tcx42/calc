package eval

import (
	"errors"
	"math"
	sy "scicalc/pkg/shuntingYard"
	"scicalc/pkg/stack"
	"strconv"
)

func Eval(input string) (float64, error) {
	if !sy.ValidateBrackets(input) {
		return 0, errors.New("Invalid Brackets")
	}
	a := sy.ParseExprString(input)
	b := sy.ShuntingYard(a)
	c := calc(b)
	return c, nil
}

func calc(rpnInput []string) float64 {
	s := stack.NewStack[float64]()
	for i := 0; i < len(rpnInput); i++ {
		if v, e := strconv.ParseFloat(rpnInput[i], 64); e == nil {
			s.Push(v)
			continue
		}
		b := s.Pop()
		a := s.Pop()
		c := operate(a, b, rpnInput[i])
		s.Push(c)
	}
	return s.Pop()
}

func operate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "^":
		return math.Pow(a, b)
	default:
		return 0
	}
}
