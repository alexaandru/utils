package utils

// MaxUint = Plus Infinity (unsigned)
const MaxUint = ^uint(0)

// MinUint = Minus Infinity (unsigned)
const MinUint = 0

// Inf = Plus Infinity
const Inf = int(MaxUint >> 1)

// MinInf = Minus Infinity
const MinInf = -(Inf - 1)

// AbsInt returns the absolute value of an int
func AbsInt(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

// MaxInt returns the maximum int from multiple given ints
func MaxInt(ints ...int) (m int) {
	m = MinInf
	for _, v := range ints {
		if v > m {
			m = v
		}
	}

	return
}

// MaxIntIndex returns the index corresponding to the maximum int
func MaxIntIndex(ints ...int) (i int) {
	m := MinInf
	for k, v := range ints {
		if v >= m {
			m, i = v, k
		}
	}

	return
}

// MinInt returns the minimum int from multiple given ints
func MinInt(ints ...int) (m int) {
	m = Inf
	for _, v := range ints {
		if v < m {
			m = v
		}
	}

	return
}

// MinIntIndex returns the index corresponding to the minimum int
func MinIntIndex(ints ...int) (i int) {
	m := Inf
	for k, v := range ints {
		if v <= m {
			m, i = v, k
		}
	}

	return
}

// NewInts returns a new []int, possibly initialized to some value (if given)
func NewInts(size int, opts ...int) (out []int) {
	out = make([]int, size)

	if len(opts) > 0 {
		defVal := opts[0]
		for k := range out {
			out[k] = defVal
		}
	}

	return
}

// IntsEq compares two []int (int slices) for equality
func IntsEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}

// IntsHas checks if "a" ints slice has a value of "n".
func IntsHas(a []int, n int) bool {
	for _, v := range a {
		if v == n {
			return true
		}
	}

	return false
}

// IntsSum returns a sum of all ints in the slice.
func IntsSum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}

	return
}
