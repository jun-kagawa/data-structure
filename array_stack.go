package datastructure


type ArrayStack[T any] struct {
	n int
	array []T
}

func NewArrayStack[T any] () *ArrayStack[T] {
	return &ArrayStack[T]{
		n: 0,
		array: make([]T, 5),
	}
}

func (s *ArrayStack[T]) Size() int {
	return s.n
}

func (s *ArrayStack[T]) Get(i int) T {
	return s.array[i]
}

func (s *ArrayStack[T]) Set(i int, x T) T {
	y := s.array[i]
	s.array[i] = x
	return y
}

func (s *ArrayStack[T]) Add(i int, x T) {
	if s.n + 1 > len(s.array) {
		s.resize()
	}
	for j := s.n; j > i; j-- {
		s.array[j] = s.array[j-1]
	}
	s.array[i] = x
	s.n++
}

func (s *ArrayStack[T]) Remove(i int) T {
	x := s.array[i]
	s.array = s.array[:i+copy(s.array[i:], s.array[i+1:])]
	s.n--
	if len(s.array) >= 3 * s.n {
		s.resize()
	}
	return x
}

func (s *ArrayStack[T]) resize() {
	array := make([]T, max(2 * s.n, 1))
	for i := 0; i < s.n; i++ {
		array[i] = s.array[i]
	}
	s.array = array
}

