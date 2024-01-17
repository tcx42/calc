package shuntingyard

import (
	"regexp"
	"scicalc/pkg/stack"
	"strconv"
)

type operator struct {
	prec      int
	isLeftAss bool
}

var operators = map[string]operator{
	"+": {2, true},
	"-": {2, true},
	"*": {3, true},
	"/": {3, true},
	"^": {4, false},
}

func isNum(s string) bool {
	_, e := strconv.ParseFloat(s, 64)
	return e == nil
}

func isOperator(s string) bool {
	return operators[s].prec != 0
}

func ParseExprString(s string) []string {
	regex := regexp.MustCompile(`(\d+(\.\d+)*E[+-]{0,1}\d+)|(\d+(\,|\.)\d+(\,|\.)\d+)|(\d+(\,|\.)\d+)|(\d+\.)|(\.\d+)|\d+|[\+\-\*\/\^]|[()\[\]{}]`)
	out := regex.FindAllString(s, -1)
	return out
}

func ValidateBrackets(s string) bool {
	cs := stack.NewStack[rune]()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			cs.Push(c)
			continue
		}
		if c == ')' || c == ']' || c == '}' {
			if cs.Len > 0 && matchBrackets(string(c), string(cs.Peek())) {
				cs.Pop()
				continue
			}
			return false
		}
	}
	if cs.Len == 0 {
		return true
	}
	return false
}

func ShuntingYard(input []string) []string {
	out := stack.NewStack[string]()
	ops := stack.NewStack[string]()
	for _, s := range input {
		if isNum(s) {
			out.Push(s)
			continue
		}
		if isOperator(s) {
			for ops.Len > 0 && (operators[ops.Peek()].prec > operators[s].prec ||
				(operators[ops.Peek()].prec == operators[s].prec && operators[s].isLeftAss)) {
				out.Push(ops.Pop())
			}
			ops.Push(s)
			continue
		}
		if s == "(" || s == "[" || s == "{" {
			ops.Push(s)
			continue
		}
		if s == ")" || s == "]" || s == "}" {
			for !matchBrackets(s, ops.Peek()) {
				if ops.Len <= 0 {
					break
				}
				out.Push(ops.Pop())
			}
			if ts := ops.Peek(); ts == "(" || ts == "[" || ts == "{" {
				ops.Pop()
			}
		}
	}
	for ops.Len > 0 {
		out.Push(ops.Pop())
	}
	var rpnExp = make([]string, out.Len)
	for i := out.Len - 1; i >= 0; i-- {
		rpnExp[i] = out.Pop()
	}
	return rpnExp
}

func matchBrackets(b1, b2 string) bool {
	switch b1 {
	case "(":
		return b2 == ")"
	case ")":
		return b2 == "("
	case "[":
		return b2 == "]"
	case "]":
		return b2 == "["
	case "{":
		return b2 == "}"
	case "}":
		return b2 == "{"
	default:
		return false
	}
}
