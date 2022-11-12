package slicelist

import "github.com/craiggwilson/go-collections/iter"
import "github.com/craiggwilson/go-collections/list"

var _ list.List[int] = (*SliceList[int])(nil)

func FromSlice[T comparable](slice []T, opts ...Opt[T]) *SliceList[T] {
	l := New(opts...)
	for _, e := range slice {
		l.Add(e)
	}
	return l
}

func New[T any](opts ...Opt[T]) *SliceList[T] {
	var o options[T]
	for _, opt := range opts {
		opt(&o)
	}

	var s SliceList[T]
	if o.initialCapacity > 0 {
		s.values = make([]T, 0, o.initialCapacity)
	}

	return &s
}

type SliceList[T any] struct {
	values []T
}

func (l *SliceList[T]) Add(v T) {
	l.values = append(l.values, v)
}

func (l *SliceList[T]) InsertAt(idx int, v T) {
	l.values = append(l.values[:idx], append([]T{v}, l.values[idx+1:]...)...)
}

func (l *SliceList[T]) Iter() iter.Iter[T] {
	return iter.FromSlice[T](l.values).Iter()
}

func (l *SliceList[T]) Len() int {
	return len(l.values)
}

func (l *SliceList[T]) RemoveAt(idx int) {
	l.values = append(l.values[:idx], l.values[idx+1:]...)
}

func (l *SliceList[T]) Reverse() {
	length := len(l.values) - 1
	for i := 0; i < len(l.values)/2; i++ {
		l.values[i], l.values[length-i] = l.values[length-i], l.values[i]
	}
}

func (l *SliceList[T]) Value(idx int) T {
	return l.values[idx]
}
