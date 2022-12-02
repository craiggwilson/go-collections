package iter

import (
	"errors"
)

var ErrEmptyIter = errors.New("contains no elements")

var ErrOutOfRange = errors.New("out of range")

func Err[S any](err error) Iterer[S] {
	return ItererFunc[S](func() Iter[S] {
		return &errIter[S]{err}
	})
}

type errIter[S any] struct {
	err error
}

func (it *errIter[S]) Next() (S, bool) {
	var def S
	return def, false
}

func (it *errIter[S]) Close() error {
	return it.err
}
