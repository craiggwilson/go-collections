package frozenlist

import (
	"github.com/craiggwilson/go-collections/iter"
	"github.com/craiggwilson/go-collections/list"
)

var _ list.ReadOnly[int] = (*Frozen[int])(nil)

func NewFrozen[T any](l list.ReadOnly[T]) *Frozen[T] {
	return &Frozen[T]{l}
}

type Frozen[T any] struct {
	l list.ReadOnly[T]
}

func (l *Frozen[T]) ElementAt(idx int) T {
	return l.l.ElementAt(idx)
}

func (l *Frozen[T]) Iter() iter.Iter[T] {
	return l.l.Iter()
}

func (l *Frozen[T]) Len() int {
	return l.l.Len()
}
