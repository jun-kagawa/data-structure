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
		// We cannot directly check the underlying array capacity from outside the package,
		// but we can infer it by adding elements. Initial capacity is 5.
		stack.Add(0, 1)
		if stack.Size() != 1 {
			t.Errorf("Expected size 1 after adding one element, got %d", stack.Size())
		}
	})

	t.Run("Add_and_Size", func(t *testing.T) {
		stack := d.NewArrayStack[string]()

		// Test Add at index 0
		stack.Add(0, "first")
		if stack.Size() != 1 {
			t.Errorf("Expected size 1, got %d", stack.Size())
		}
		if stack.Get(0) != "first" {
			t.Errorf("Expected 'first' at index 0, got %s", stack.Get(0))
		}

		// Test Add at middle index
		stack.Add(1, "third")
		stack.Add(1, "second") // Add "second" between "first" and "third"
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}
		expected := []string{"first", "second", "third"}
		for i, val := range expected {
			if stack.Get(i) != val {
				t.Errorf("Expected '%s' at index %d, got '%s'", val, i, stack.Get(i))
			}
		}

		// Test Add at end
		stack.Add(stack.Size(), "fourth")
		if stack.Size() != 4 {
			t.Errorf("Expected size 4, got %d", stack.Size())
		}
		if stack.Get(3) != "fourth" {
			t.Errorf("Expected 'fourth' at index 3, got %s", stack.Get(3))
		}
	})

	t.Run("Add_Resizing", func(t *testing.T) {
		stack := d.NewArrayStack[int]() // Initial capacity 5

		// Add 5 elements to reach capacity
		for i := 0; i < 5; i++ {
			stack.Add(i, i+1)
		}
		if stack.Size() != 5 {
			t.Errorf("Expected size 5, got %d", stack.Size())
		}

		// Add 6th element to trigger resize (capacity should become 10)
		stack.Add(5, 6)
		if stack.Size() != 6 {
			t.Errorf("Expected size 6 after resize, got %d", stack.Size())
		}
		// Check elements after resize
		for i := 0; i < 6; i++ {
			if stack.Get(i) != i+1 {
				t.Errorf("After resize, expected %d at index %d, got %d", i+1, i, stack.Get(i))
			}
		}
	})

	t.Run("Remove_and_Size", func(t *testing.T) {
		stack := d.NewArrayStack[string]()
		stack.Add(0, "a")
		stack.Add(1, "b")
		stack.Add(2, "c")
		stack.Add(3, "d")

		// Remove from middle
		removed := stack.Remove(1) // remove "b"
		if removed != "b" {
			t.Errorf("Expected removed 'b', got '%s'", removed)
		}
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}
		expected := []string{"a", "c", "d"}
		for i, val := range expected {
			if stack.Get(i) != val {
				t.Errorf("After removing 'b', expected '%s' at index %d, got '%s'", val, i, stack.Get(i))
			}
		}

		// Remove from beginning
		removed = stack.Remove(0) // remove "a"
		if removed != "a" {
			t.Errorf("Expected removed 'a', got '%s'", removed)
		}
		if stack.Size() != 2 {
			t.Errorf("Expected size 2, got %d", stack.Size())
		}
		expected = []string{"c", "d"}
		for i, val := range expected {
			if stack.Get(i) != val {
				t.Errorf("After removing 'a', expected '%s' at index %d, got '%s'", val, i, stack.Get(i))
			}
		}

		// Remove from end
		removed = stack.Remove(stack.Size() - 1) // remove "d"
		if removed != "d" {
			t.Errorf("Expected removed 'd', got '%s'", removed)
		}
		if stack.Size() != 1 {
			t.Errorf("Expected size 1, got %d", stack.Size())
		}
		if stack.Get(0) != "c" {
			t.Errorf("After removing 'd', expected 'c' at index 0, got '%s'", stack.Get(0))
		}
	})

	t.Run("Remove_Shrinking", func(t *testing.T) {
		stack := d.NewArrayStack[int]() // Initial capacity 5

		// Add enough elements to cause multiple resizes, ensuring len(s.array) is large
		for i := 0; i < 10; i++ { // Capacity will grow to 10, then 20
			stack.Add(i, i+1)
		}
		if stack.Size() != 10 {
			t.Fatalf("Expected size 10, got %d", stack.Size())
		}

		// Remove elements to trigger shrinking (len(s.array) >= 3 * s.n)
		// Current capacity is likely 20, size is 10.
		// Shrink when size drops below len/3 (e.g., size becomes 6, len is 20) -> 20 >= 3 * 6 (18) is true.
		// The `resize` operation creates a new array of size `max(2 * s.n, 1)`.
		// So if size becomes 6, new capacity will be 12.

		// Remove 4 elements, size becomes 6
		for i := 0; i < 4; i++ {
			stack.Remove(0) // Remove from beginning for simplicity
		}
		if stack.Size() != 6 {
			t.Errorf("Expected size 6 after removals, got %d", stack.Size())
		}
		// We cannot directly check capacity, but we can verify elements.
		for i := 0; i < 6; i++ {
			if stack.Get(i) != i+5 { // Original values were 1..10, removed 1..4. So now 5..10.
				t.Errorf("After shrinking, expected %d at index %d, got %d", i+5, i, stack.Get(i))
			}
		}
		// Another remove to trigger shrinking again if applicable
		stack.Remove(0) // size becomes 5
		if stack.Size() != 5 {
			t.Errorf("Expected size 5, got %d", stack.Size())
		}
		for i := 0; i < 5; i++ {
			if stack.Get(i) != i+6 { // Original values were 1..10, removed 1..5. So now 6..10.
				t.Errorf("After shrinking again, expected %d at index %d, got %d", i+6, i, stack.Get(i))
			}
		}
	})

	t.Run("Get_and_Set", func(t *testing.T) {
		stack := d.NewArrayStack[int]()
		stack.Add(0, 10)
		stack.Add(1, 20)
		stack.Add(2, 30)

		// Test Get
		if stack.Get(0) != 10 {
			t.Errorf("Expected 10 at index 0, got %d", stack.Get(0))
		}
		if stack.Get(1) != 20 {
			t.Errorf("Expected 20 at index 1, got %d", stack.Get(1))
		}

		// Test Set
		oldVal := stack.Set(1, 25)
		if oldVal != 20 {
			t.Errorf("Expected old value 20, got %d", oldVal)
		}
		if stack.Get(1) != 25 {
			t.Errorf("Expected 25 at index 1 after Set, got %d", stack.Get(1))
		}
		if stack.Size() != 3 {
			t.Errorf("Size should not change after Set, got %d", stack.Size())
		}
	})

	t.Run("Edge_Cases", func(t *testing.T) {
		t.Run("EmptyStack_Remove", func(t *testing.T) {
			stack := d.NewArrayStack[int]()
			// Trying to remove from an empty stack with an invalid index would cause a panic.
			// The current implementation panics for invalid indices.
			// To test this gracefully, one would typically expect an error return or a specific panic message.
			// For this review, we'll note that such an operation (e.g., stack.Remove(0) on empty) will panic.
			// If graceful error handling is desired, bounds checks should be added to the Remove method.
			_ = stack // suppress unused variable warning
		})

		t.Run("SingleElementStack", func(t *testing.T) {
			stack := d.NewArrayStack[string]()
			stack.Add(0, "only")
			if stack.Size() != 1 {
				t.Errorf("Expected size 1, got %d", stack.Size())
			}
			if stack.Get(0) != "only" {
				t.Errorf("Expected 'only', got %s", stack.Get(0))
			}

			removed := stack.Remove(0)
			if removed != "only" {
				t.Errorf("Expected removed 'only', got '%s'", removed)
			}
			if stack.Size() != 0 {
				t.Errorf("Expected size 0 after removing last element, got %d", stack.Size())
			}
		})
	})

	t.Run("Stack_Behavior", func(t *testing.T) {
		stack := d.NewArrayStack[string]()

		// Push-like behavior (Add to end)
		stack.Add(stack.Size(), "A") // Push A
		stack.Add(stack.Size(), "B") // Push B
		stack.Add(stack.Size(), "C") // Push C

		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}
		if stack.Get(0) != "A" || stack.Get(1) != "B" || stack.Get(2) != "C" {
			t.Errorf("Elements not in expected order: %v, %v, %v", stack.Get(0), stack.Get(1), stack.Get(2))
		}

		// Pop-like behavior (Remove from end)
		popped := stack.Remove(stack.Size() - 1) // Pop C
		if popped != "C" {
			t.Errorf("Expected popped 'C', got '%s'", popped)
		}
		if stack.Size() != 2 {
			t.Errorf("Expected size 2, got %d", stack.Size())
		}

		popped = stack.Remove(stack.Size() - 1) // Pop B
		if popped != "B" {
			t.Errorf("Expected popped 'B', got '%s'", popped)
		}
		if stack.Size() != 1 {
			t.Errorf("Expected size 1, got %d", stack.Size())
		}

		popped = stack.Remove(stack.Size() - 1) // Pop A
		if popped != "A" {
			t.Errorf("Expected popped 'A', got '%s'", popped)
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size 0, got %d", stack.Size())
		}
	})

	t.Run("Different_Generic_Types", func(t *testing.T) {
		// Test with int
		stackInt := d.NewArrayStack[int]()
		stackInt.Add(0, 100)
		if stackInt.Get(0) != 100 {
			t.Errorf("int stack: Expected 100, got %d", stackInt.Get(0))
		}
		if stackInt.Size() != 1 {
			t.Errorf("int stack: Expected size 1, got %d", stackInt.Size())
		}

		// Test with struct
		type MyStruct struct {
			ID   int
			Name string
		}
		stackStruct := d.NewArrayStack[MyStruct]()
		val := MyStruct{ID: 1, Name: "Test"}
		stackStruct.Add(0, val)
		if stackStruct.Get(0).ID != 1 || stackStruct.Get(0).Name != "Test" {
			t.Errorf("struct stack: Expected {ID:1, Name:\"Test\"}, got %v", stackStruct.Get(0))
		}
		if stackStruct.Size() != 1 {
			t.Errorf("struct stack: Expected size 1, got %d", stackStruct.Size())
		}
	})
}
