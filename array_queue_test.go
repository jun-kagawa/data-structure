package datastructure_test

import (
	"testing"

	d "github.com/jun-kagawa/data-structure"
)

func TestArrayQueue(t *testing.T) {
	t.Run("Basic Add and Remove", func(t *testing.T) {
		q := d.NewArrayQueue[int]()
		q.Add(1)
		q.Add(2)
		q.Add(3)

		if q.Size() != 3 {
			t.Errorf("Expected size 3, got %d", q.Size())
		}

		v, err := q.Remove()
		if err != nil || v != 1 {
			t.Errorf("Expected (1, nil), got (%d, %v)", v, err)
		}
		v, err = q.Remove()
		if err != nil || v != 2 {
			t.Errorf("Expected (2, nil), got (%d, %v)", v, err)
		}
		v, err = q.Remove()
		if err != nil || v != 3 {
			t.Errorf("Expected (3, nil), got (%d, %v)", v, err)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("Remove from empty queue", func(t *testing.T) {
		q := d.NewArrayQueue[int]()
		_, err := q.Remove()
		if err == nil {
			t.Error("Expected error when removing from empty queue, got nil")
		}
	})

	t.Run("Mixed Add and Remove", func(t *testing.T) {
		q := d.NewArrayQueue[int]()
		q.Add(1)
		q.Add(2)
		v, _ := q.Remove()
		if v != 1 {
			t.Fatal("Expected 1")
		}
		q.Add(3)
		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}

		v, _ = q.Remove()
		if v != 2 {
			t.Errorf("Expected 2, got %d", v)
		}
		v, _ = q.Remove()
		if v != 3 {
			t.Errorf("Expected 3, got %d", v)
		}
	})

	t.Run("Circular Behavior and Resizing", func(t *testing.T) {
		q := d.NewArrayQueue[int]()
		for i := 0; i < 5; i++ {
			q.Add(i)
		}
		q.Remove() // 0
		q.Remove() // 1
		q.Add(5)
		q.Add(6)

		for i := 7; i < 20; i++ {
			q.Add(i)
		}

		expected := 2
		for q.Size() > 0 {
			v, _ := q.Remove()
			if v != expected {
				t.Errorf("Expected %d, got %d", expected, v)
			}
			expected++
		}
	})
}
