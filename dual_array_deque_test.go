package datastructure_test

import (
	"testing"

	d "github.com/jun-kagawa/data-structure"
)

func TestDualArrayDeque(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		deque := &d.DualArrayDeque[int]{}
		if deque.Size() != 0 {
			t.Errorf("Expected size 0, got %d", deque.Size())
		}

		// Add at the end
		deque.Add(0, 1) // [1]
		deque.Add(1, 2) // [1, 2]
		deque.Add(2, 3) // [1, 2, 3]

		if deque.Size() != 3 {
			t.Errorf("Expected size 3, got %d", deque.Size())
		}

		if deque.Get(0) != 1 || deque.Get(1) != 2 || deque.Get(2) != 3 {
			t.Errorf("Unexpected values: [%d, %d, %d]", deque.Get(0), deque.Get(1), deque.Get(2))
		}

		// Set value
		old := deque.Set(1, 20)
		if old != 2 || deque.Get(1) != 20 {
			t.Errorf("Set failed: old=%d, current=%d", old, deque.Get(1))
		}

		// Remove value
		val, err := deque.Remove(1) // Remove 20 -> [1, 3]
		if err != nil || val != 20 {
			t.Errorf("Remove failed: val=%d, err=%v", val, err)
		}
		if deque.Size() != 2 {
			t.Errorf("Expected size 2, got %d", deque.Size())
		}
	})

	t.Run("AddFront", func(t *testing.T) {
		deque := &d.DualArrayDeque[int]{}
		deque.Add(0, 3) // [3]
		deque.Add(0, 2) // [2, 3]
		deque.Add(0, 1) // [1, 2, 3]

		if deque.Size() != 3 {
			t.Errorf("Expected size 3, got %d", deque.Size())
		}
		for i := 0; i < 3; i++ {
			if deque.Get(i) != i+1 {
				t.Errorf("Expected %d at %d, got %d", i+1, i, deque.Get(i))
			}
		}
	})

	t.Run("Balance", func(t *testing.T) {
		deque := &d.DualArrayDeque[int]{}
		// Add elements to trigger balance multiple times
		for i := 0; i < 10; i++ {
			deque.Add(deque.Size(), i)
		}
		// Size is 10. Balance should have ensured front and back are somewhat even.
		if deque.Size() != 10 {
			t.Errorf("Expected size 10, got %d", deque.Size())
		}
		for i := 0; i < 10; i++ {
			if deque.Get(i) != i {
				t.Errorf("Expected %d at %d, got %d", i, i, deque.Get(i))
			}
		}

		// Removing from one end should also trigger balance
		for i := 0; i < 8; i++ {
			deque.Remove(0)
		}
		if deque.Size() != 2 {
			t.Errorf("Expected size 2, got %d", deque.Size())
		}
		if deque.Get(0) != 8 || deque.Get(1) != 9 {
			t.Errorf("Expected [8, 9], got [%d, %d]", deque.Get(0), deque.Get(1))
		}
	})
}
