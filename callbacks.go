package utils

// IdentIntFunc returns the identity function
func IdentIntFunc(i int) int {
    return i
}

// ConstIntFunc returns a constant function
func ConstIntFunc(x int) func(int) int {
    return func(i int) int { return x }
}

// MultIntFunc returns a constant multiplier function
func MultIntFunc(x int) func(int) int {
    return func(i int) int { return i * x }
}
