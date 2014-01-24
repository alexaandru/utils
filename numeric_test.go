package utils

import "testing"

func TestAbsInt(t *testing.T) {
    assertEqual(42, AbsInt(42), "Expected AbsInt to return %d, got %d", t)
    assertEqual(42, AbsInt(-42), "Expected AbsInt to return %d, got %d", t)
}

func TestMaxInt(t *testing.T) {
    assertEqual(4, MaxInt(1, 4, 2, 3), "Expected MaxInt to return %d, got %d", t)
    assertEqual(-4, MaxInt(-6, -5, -4), "Expected MaxInt to return %d, got %d", t)
}

func TestMaxIntIndex(t *testing.T) {
    assertEqual(1, MaxIntIndex(1, 4, 2, 3), "Expected MaxIntIndex to return %d, got %d", t)
    assertEqual(2, MaxIntIndex(-6, -5, -4), "Expected MaxIntIndex to return %d, got %d", t)
}

func TestMinInt(t *testing.T) {
    assertEqual(1, MinInt(1, 4, 2, 3), "Expected MinInt to return %d, got %d", t)
    assertEqual(-6, MinInt(-6, -5, -4), "Expected MinInt to return %d, got %d", t)
}

func TestMinIntIndex(t *testing.T) {
    assertEqual(0, MinIntIndex(1, 4, 2, 3), "Expected MinIntIndex to return %d, got %d", t)
    assertEqual(0, MinIntIndex(-6, -5, -4), "Expected MinIntIndex to return %d, got %d", t)
}
