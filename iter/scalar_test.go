package iter_test

import (
	"strconv"
	"testing"

	"github.com/shoenig/test/must"

	"github.com/craiggwilson/go-collections/iter"
)

func Test_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:  "all match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i > 0
			},
			expected: true,
		},
		{
			name:  "some match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i < 5
			},
			expected: false,
		},
		{
			name:  "none match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i > 10
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.All(it, tc.predicate)
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Any(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:  "all match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i > 0
			},
			expected: true,
		},
		{
			name:  "some match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i > 5
			},
			expected: true,
		},
		{
			name:  "none match",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool {
				return i > 10
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Any(it, tc.predicate)
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Contains(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		target   int
		expected bool
	}{
		{
			name:     "contains",
			input:    []int{1, 3, 5, 7, 9},
			target:   5,
			expected: true,
		},
		{
			name:     "does not contain",
			input:    []int{1, 3, 5, 7, 9},
			target:   6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Contains(it, tc.target)
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_ElementAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		index    uint
		expected int
		err      error
	}{
		{
			name:     "first",
			input:    []int{1, 3, 5, 7, 9},
			index:    0,
			expected: 1,
		},
		{
			name:     "second",
			input:    []int{1, 3, 5, 7, 9},
			index:    1,
			expected: 3,
		},
		{
			name:     "last",
			input:    []int{1, 3, 5, 7, 9},
			index:    4,
			expected: 9,
		},
		{
			name:  "out of range",
			input: []int{1, 3, 5, 7, 9},
			index: 5,
			err:   iter.ErrOutOfRange,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ElementAt(it, tc.index)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_First(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:  "no elements",
			input: nil,
			err:   iter.ErrEmptyIter,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.First(it)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_FirstOrDefault(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "no elements",
			input:    nil,
			expected: 0,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.FirstOrDefault(it)
			must.NoError(t, err)

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Fold(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		seed     string
		reducer  func(string, int) string
		expected string
	}{
		{
			name:  "no elements",
			input: nil,
			seed:  "Hello!",
			reducer: func(acc string, elem int) string {
				return acc + strconv.Itoa(elem)
			},
			expected: "Hello!",
		},
		{
			name:  "1 element",
			input: []int{1},
			seed:  "Hello!",
			reducer: func(acc string, elem int) string {
				return acc + strconv.Itoa(elem)
			},
			expected: "Hello!1",
		},
		{
			name:  "many elements",
			input: []int{1, 3, 5, 7, 9},
			seed:  "Hello!",
			reducer: func(acc string, elem int) string {
				return acc + strconv.Itoa(elem)
			},
			expected: "Hello!13579",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Fold(it, tc.seed, tc.reducer)
			must.NoError(t, err)

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Last(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:  "no elements",
			input: nil,
			err:   iter.ErrEmptyIter,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Last(it)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_LastOrDefault(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "no elements",
			input:    nil,
			expected: 0,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.LastOrDefault(it)
			must.NoError(t, err)

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Len(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "no elements",
			input:    nil,
			expected: 0,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Len(it)
			must.NoError(t, err)

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Max(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:  "no elements",
			input: nil,
			err:   iter.ErrEmptyIter,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "first",
			input:    []int{9, 1, 3, 5, 7},
			expected: 9,
		},
		{
			name:     "middle",
			input:    []int{1, 3, 5, 9, 7},
			expected: 9,
		},
		{
			name:     "last",
			input:    []int{1, 3, 5, 7, 9},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Max(it)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Min(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:  "no elements",
			input: nil,
			err:   iter.ErrEmptyIter,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "first",
			input:    []int{1, 3, 5, 7, 9},
			expected: 1,
		},
		{
			name:     "middle",
			input:    []int{3, 5, 1, 7, 9},
			expected: 1,
		},
		{
			name:     "last",
			input:    []int{3, 5, 7, 9, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Min(it)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Reduce(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		reducer  func(int, int) int
		expected int
		err      error
	}{
		{
			name:  "no elements",
			input: nil,
			reducer: func(acc int, elem int) int {
				return acc + elem
			},
			err: iter.ErrEmptyIter,
		},
		{
			name:  "1 element",
			input: []int{1},
			reducer: func(acc int, elem int) int {
				return acc + elem
			},
			expected: 1,
		},
		{
			name:  "many elements",
			input: []int{1, 3, 5, 7, 9},
			reducer: func(acc int, elem int) int {
				return acc + elem
			},
			expected: 25,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Reduce(it, tc.reducer)
			if tc.err != nil {
				must.ErrorIs(t, err, tc.err)
			} else {
				must.NoError(t, err)
			}

			must.Eq(t, tc.expected, actual)
		})
	}
}

func Test_Sum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "no elements",
			input:    nil,
			expected: 0,
		},
		{
			name:     "1 element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "many elements",
			input:    []int{1, 3, 5, 7, 9},
			expected: 25,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.Sum(it)
			must.NoError(t, err)

			must.Eq(t, tc.expected, actual)
		})
	}
}
