package frozendict

import (
	"github.com/craiggwilson/go-collections/dict"
	"github.com/craiggwilson/go-collections/iter"
)

var _ dict.ReadOnly[int, int] = (*FrozenDict[int, int])(nil)

func New[K comparable, V any](l dict.ReadOnly[K, V]) *FrozenDict[K, V] {
	return &FrozenDict[K, V]{l}
}

type FrozenDict[K comparable, V any] struct {
	d dict.ReadOnly[K, V]
}

func (d *FrozenDict[K, V]) Contains(k K) bool {
	return d.d.Contains(k)
}

func (d *FrozenDict[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return d.d.Iter()
}

func (d *FrozenDict[K, V]) Keys() iter.Iterer[K] {
	return d.d.Keys()
}

func (d *FrozenDict[K, V]) Len() int {
	return d.d.Len()
}

func (d *FrozenDict[K, V]) Value(k K) (V, bool) {
	return d.d.Value(k)
}

func (d *FrozenDict[K, V]) Values() iter.Iterer[V] {
	return d.d.Values()
}
