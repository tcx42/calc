package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[string]()
	if s.Len != 0 {
		t.Fail()
	}
	s.Push("a")
	if s.Len != 1 {
		t.Fail()
	}
	if s.Peek() != "a" {
		t.Fail()
	}
	if s.Pop() != "a" {
		t.Fail()
	}
	s.Push("b")
	s.Push("c")
	s.Push("d")
	if s.Len != 3 {
		t.Fail()
	}
	if s.Pop() != "d" {
		t.Fail()
	}
	if s.Len != 2 {
		t.Fail()
	}
	if s.Pop() != "c" {
		t.Fail()
	}
	if s.Pop() != "b" {
		t.Fail()
	}
	if s.Len != 0 {
		t.Fail()
	}
	s.Peek()
	s.Pop()
}
