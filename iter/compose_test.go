package iter_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"github.com/craiggwilson/go-collections/iter"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		left     []int
		right    []int
		expected []int
	}{
		{
			name:     "both with elements",
			left:     []int{1, 3, 5, 7, 9},
			right:    []int{2, 4, 6},
			expected: []int{1, 3, 5, 7, 9, 2, 4, 6},
		},
		{
			name:     "empty right",
			left:     []int{1, 3, 5, 7, 9},
			right:    nil,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "empty left",
			left:     nil,
			right:    []int{1, 3, 5, 7, 9},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "empty both",
			left:     nil,
			right:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			left := iter.FromSlice(tc.left)
			right := iter.FromSlice(tc.right)

			actual, err := iter.ToSlice(iter.Concat(left, right))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestDistinct(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 3, 5, 7, 9},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "all duplicates",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "some duplicates",
			input:    []int{1, 5, 5, 9, 5},
			expected: []int{1, 5, 9},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.Distinct(it))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "keep all",
			input:     []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool { return true },
			expected:  []int{1, 3, 5, 7, 9},
		},
		{
			name:      "keep none",
			input:     []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool { return false },
			expected:  nil,
		},
		{
			name:      "keep some",
			input:     []int{1, 3, 5, 7, 9},
			predicate: func(i int) bool { return i > 1 && i < 9 },
			expected:  []int{3, 5, 7},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.Filter(it, tc.predicate))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestSelect(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		selector func(int) int
		expected []int
	}{
		{
			name:     "identity",
			input:    []int{1, 3, 5, 7, 9},
			selector: func(i int) int { return i },
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "add2",
			input:    []int{1, 3, 5, 7, 9},
			selector: func(i int) int { return i + 2 },
			expected: []int{3, 5, 7, 9, 11},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.Select(it, tc.selector))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestSelectMany(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		selector func(int) iter.Iterer[int]
		expected []int
	}{
		{
			name:     "identity",
			input:    []int{1, 3, 5, 7, 9},
			selector: func(i int) iter.Iterer[int] { return iter.FromSlice([]int{i}) },
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "doubling",
			input:    []int{1, 3, 5, 7, 9},
			selector: func(i int) iter.Iterer[int] { return iter.Repeat(i, 2) },
			expected: []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.SelectMany(it, tc.selector))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestSkip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		skip     int
		expected []int
	}{
		{
			name:     "all",
			input:    []int{1, 3, 5, 7, 9},
			skip:     5,
			expected: nil,
		},
		{
			name:     "none",
			input:    []int{1, 3, 5, 7, 9},
			skip:     0,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "some",
			input:    []int{1, 3, 5, 7, 9},
			skip:     3,
			expected: []int{7, 9},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.Skip(it, tc.skip))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestTake(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		take     int
		expected []int
	}{
		{
			name:     "all",
			input:    []int{1, 3, 5, 7, 9},
			take:     5,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "none",
			input:    []int{1, 3, 5, 7, 9},
			take:     0,
			expected: nil,
		},
		{
			name:     "some",
			input:    []int{1, 3, 5, 7, 9},
			take:     3,
			expected: []int{1, 3, 5},
		},
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := iter.FromSlice(tc.input)

			actual, err := iter.ToSlice(iter.Take(it, tc.take))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}

func TestZip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		left     []int
		right    []int
		zipper   func(int, int) int
		expected []int
	}{
		{
			name:     "same length",
			left:     []int{1, 3, 5},
			right:    []int{2, 4, 6},
			zipper:   func(a, b int) int { return a + b },
			expected: []int{3, 7, 11},
		},
		{
			name:     "right shorter",
			left:     []int{1, 3, 5, 7, 9},
			right:    []int{2, 4, 6},
			zipper:   func(a, b int) int { return a + b },
			expected: []int{3, 7, 11},
		},
		{
			name:     "left shorter",
			left:     []int{1, 3, 5},
			right:    []int{2, 4, 6, 8, 10},
			zipper:   func(a, b int) int { return a + b },
			expected: []int{3, 7, 11},
		},
		{
			name:     "empty right",
			left:     []int{1, 3, 5, 7, 9},
			right:    nil,
			zipper:   func(a, b int) int { return a + b },
			expected: nil,
		},
		{
			name:     "empty left",
			left:     nil,
			right:    []int{1, 3, 5, 7, 9},
			expected: nil,
		},
		{
			name:     "empty both",
			left:     nil,
			right:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			left := iter.FromSlice(tc.left)
			right := iter.FromSlice(tc.right)

			actual, err := iter.ToSlice(iter.Zip(left, right, tc.zipper))
			must.NoError(t, err)
			must.Eq(t, tc.expected, actual)
		})
	}
}
