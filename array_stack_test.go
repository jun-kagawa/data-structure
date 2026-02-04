package datastructure_test

import (
	"fmt"
	"testing"

	d "github.com/jun-kagawa/data-structure"
)

func TestArrayStack(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		stack := d.NewArrayStack[string]()
		stack.Add(0, "a")
		stack.Add(1, "b")
		stack.Add(2, "c")
		stack.Add(3, "d")
		stack.Add(4, "e")
		stack.Add(5, "f")
		stack.Add(6, "g")
		fmt.Println(stack)
		stack.Remove(5)
		stack.Remove(0)
		fmt.Println(stack)
		stack.Add(0, "z")
		fmt.Println(stack)
	})
}

