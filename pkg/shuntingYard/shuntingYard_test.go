package shuntingyard

import (
	"testing"
)

func TestShuntingYard(t *testing.T) {
	input := []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"}
	expected := []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}
	result := ShuntingYard(input)
	if len(expected) != len(result) {
		t.Log("Expected: ", expected)
		t.Log("Result: ", result)
		t.Fatal()
	}
	for i, e := range expected {
		if e != result[i] {
			t.Fail()
		}
	}
	t.Log("Expected: ", expected)
	t.Log("Result: ", result)
}

func TestParseExprString(t *testing.T) {
	inStr := " 3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3"
	exp := []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"}
	res := ParseExprString(inStr)
	for i := 0; i < len(exp) && i < len(res); i++ {
		if exp[i] != res[i] {
			t.Error("i ", i, "exp ", exp[i], "res ", res[i])
		}
	}
}
