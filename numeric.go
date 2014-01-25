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
