package set

import "github.com/craiggwilson/go-collections/iter"

type ReadOnly[T comparable] interface {
	iter.Iterer[T]

	Contains(T) bool
	Len() int
}

type Set[T comparable] interface {
	ReadOnly[T]

	Add(T)
	Clear()
	Remove(T)
}
