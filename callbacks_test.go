package utils

import (
    "fmt"
    "testing"
)

func TestIdentIntFunc(t *testing.T) {
    x := 42
    f := IdentIntFunc(x)

    assertEqual(x, f, "Expected IdentIntFunc to return %d, got %d", t)
}

func TestConstIntFunc(t *testing.T) {
    x := 42
    f := ConstIntFunc(x)

    assertEqual(x, f(1), "Expected ConstIntFunc to return %d, got %d", t)
    assertEqual(x, f(2), "Expected ConstIntFunc to return %d, got %d", t)
}

func TestMultIntFunc(t *testing.T) {
    x := 42
    f := MultIntFunc(42)

    assertEqual(x, f(1), "Expected MultIntFunc to return %d, got %d", t)
    assertEqual(10*x, f(10), "Expected MultIntFunc to return %d, got %d", t)
}

func assertEqual(a, b int, msg string, t *testing.T) {
    if a != b {
        t.Error(fmt.Sprintf(msg, a, b))
    }
}

func assertEqualF32(a, b float32, msg string, t *testing.T) {
    if a != b {
        t.Error(fmt.Sprintf(msg, a, b))
    }
}

func assertEqualF64(a, b float64, msg string, t *testing.T) {
    if a != b {
        t.Error(fmt.Sprintf(msg, a, b))
    }
}
