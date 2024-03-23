package pkg_test

import (
	"testing"

	"github.com/conorkenn/data_structures_go/pkg/stack.go"
)

func TestPush(t *testing.T) {
	myStack := stack.NewStack()

	myStack.Push("123")

	want := false
	got := myStack.IsEmpty()

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

}

func TestSize(t *testing.T) {
	myStack := stack.NewStack()

	myStack.Push("1")
	myStack.Push("2")
	want := 2
	got := myStack.Size()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTop(t *testing.T) {
	myStack := stack.NewStack()

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	want := 3
	got, _ := myStack.Top()

	if got.(int) != want {
		t.Errorf("got %d want %d", got, want)
	}
	if myStack.Size() != 3 {
		t.Errorf("stack should be len 3")
	}
}

func TestPop(t *testing.T) {
	myStack := stack.NewStack()

	myStack.Push(1)
	myStack.Push(1)
	myStack.Push(3)
	want := 3
	got, _ := myStack.Pop()

	if got.(int) != want {
		t.Errorf("got %d want %d", got, want)
	}
	if myStack.Size() != 2 {
		t.Errorf("stack should be len 2")
	}
}
