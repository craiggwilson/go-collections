package dict

import "github.com/craiggwilson/go-collections/iter"

type ReadOnly[K comparable, V any] interface {
	iter.Iterer[iter.KeyValuePair[K, V]]

	Contains(K) bool
	Keys() iter.Iterer[K]
	Len() int
	Value(K) (V, bool)
	Values() iter.Iterer[V]
}

type Dict[K comparable, V any] interface {
	ReadOnly[K, V]

	Add(K, V)
	Remove(K)
}
