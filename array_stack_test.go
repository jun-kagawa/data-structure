package datastructure_test

import (
	"testing"

	d "github.com/jun-kagawa/data-structure"
)

func TestArrayStack(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		stack := d.NewArrayStack[int]()
		if stack == nil {
			t.Fatal("NewArrayStack should not return nil")
		}
		if stack.Size() != 0 {
			t.Errorf("Expected initial size 0, got %d", stack.Size())
		}
		stack.Add(0, 1)
		if stack.Size() != 1 {
			t.Errorf("Expected size 1 after adding one element, got %d", stack.Size())
		}
	})

	t.Run("Add_and_Size", func(t *testing.T) {
		stack := d.NewArrayStack[string]()
		stack.Add(0, "first")
		if stack.Size() != 1 {
			t.Errorf("Expected size 1, got %d", stack.Size())
		}
		if stack.Get(0) != "first" {
			t.Errorf("Expected 'first' at index 0, got %s", stack.Get(0))
		}

		stack.Add(1, "third")
		stack.Add(1, "second")
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}
		expected := []string{"first", "second", "third"}
		for i, val := range expected {
			if stack.Get(i) != val {
				t.Errorf("Expected '%s' at index %d, got '%s'", val, i, stack.Get(i))
			}
		}
	})

	t.Run("Remove_and_Size", func(t *testing.T) {
		stack := d.NewArrayStack[string]()
		stack.Add(0, "a")
		stack.Add(1, "b")
		stack.Add(2, "c")
		stack.Add(3, "d")

		removed, err := stack.Remove(1)
		if err != nil || removed != "b" {
			t.Errorf("Expected ('b', nil), got ('%s', %v)", removed, err)
		}
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}

		_, err = stack.Remove(10) // Out of bounds
		if err == nil {
			t.Error("Expected error for out of bounds removal, got nil")
		}
	})

	t.Run("Edge_Cases", func(t *testing.T) {
		t.Run("EmptyStack_Remove", func(t *testing.T) {
			stack := d.NewArrayStack[int]()
			_, err := stack.Remove(0)
			if err == nil {
				t.Error("Expected error when removing from empty stack")
			}
		})
	})
}
