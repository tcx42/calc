package eval

import (
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	input := "(4/2-1*5)^2"
	out := 9.
	res, e := Eval(input)
	if res != out || e != nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}

	input = "((4/2-1)*5)^2"
	out = 25
	res, e = Eval(input)
	if res != out || e != nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}

	input = "3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3"
	out = 3.0001220703125
	res, e = Eval(input)
	if res != out || e != nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}

	input = "3.1 + 1.9"
	out = 5
	res, e = Eval(input)
	if res != out || e != nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}

	input = "(3.1*[ + 1.9)"
	out = 0
	res, e = Eval(input)
	if e == nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}

	input = "(3.1*[1-2] + 1.9)"
	out = -1.2
	res, e = Eval(input)
	if math.Abs(res-out) > 1e-15 || e != nil {
		t.Error("input:", input, "result:", res, "expected:", out, "error:", e)
	}
}
