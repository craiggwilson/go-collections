package iter

func Empty[S any]() Iterer[S] {
	return Generate(func() (result S, ok bool) {
		return
	})
}
