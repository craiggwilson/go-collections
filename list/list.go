package list

import "github.com/craiggwilson/go-collections/iter"

type ReadOnly[T any] interface {
	iter.Iterer[T]

	ElementAt(int) T
	Len() int
}

type List[T any] interface {
	ReadOnly[T]

	Add(T)
	InsertAt(int, T)
	RemoveAt(int)
}
