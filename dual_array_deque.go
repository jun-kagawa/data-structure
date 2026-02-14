package datastructure

type DualArrayDeque[T any] struct {
	front ArrayStack[T]
	back  ArrayStack[T]
}

func (d *DualArrayDeque[T]) Size() int {
	return d.front.Size() + d.back.Size()
}

func (d *DualArrayDeque[T]) Get(i int) T {
	if i < d.front.Size() {
		return d.front.Get(d.front.Size() - i - 1)
	} else {
		return d.back.Get(i - d.front.Size())
	}
}

func (d *DualArrayDeque[T]) Set(i int, x T) T {
	if i < d.front.Size() {
		return d.front.Set(d.front.Size()-i-1, x)
	} else {
		return d.back.Set(i-d.front.Size(), x)
	}
}

func (d *DualArrayDeque[T]) Add(i int, x T) {
	if i < d.front.Size() {
		d.front.Add(d.front.Size()-i, x)
	} else {
		d.back.Add(i-d.front.Size(), x)
	}
	d.balance()
}

func (d *DualArrayDeque[T]) Remove(i int) (T, error) {
	var x T
	var err error
	if i < d.front.Size() {
		x, err = d.front.Remove(d.front.Size() - i - 1)
		if err != nil {
			return x, err
		}
	} else {
		x, err = d.back.Remove(i - d.front.Size())
		if err != nil {
			return x, err
		}
	}
	d.balance()
	return x, nil
}

func (d *DualArrayDeque[T]) balance() {
	f := d.front.Size()
	b := d.back.Size()
	if 3*f < b || 3*b < f {
		n := f + b
		nf := n / 2

		af := make([]T, nf)
		for i := 0; i < nf; i++ {
			af[nf-i-1] = d.Get(i)
		}

		nb := n - nf
		ab := make([]T, nb)
		for i := 0; i < nb; i++ {
			ab[i] = d.Get(nf + i)
		}

		d.front.ReplaceAll(af)
		d.back.ReplaceAll(ab)
	}
}
