package utils

import "testing"

func TestParseInt(t *testing.T) {
    assertEqual(127, ParseInt("127", 8), "Expected ParseInt to return %d, got %d", t)
    assertEqual(256, ParseInt("256"), "Expected ParseInt to return %d, got %d", t)
    assertEqual(256, ParseInt("256", 16), "Expected ParseInt to return %d, got %d", t)
    assertEqual(65536, ParseInt("65536", 32), "Expected ParseInt to return %d, got %d", t)
    assertEqual(4294967296, ParseInt("4294967296", 64), "Expected ParseInt to return %d, got %d", t)

    // FIXME: also need negative tests (that it fails to parse)
}

func TestParseFloat(t *testing.T) {
    assertEqualF64(65536.55, ParseFloat("65536.550", 64), "Expected ParseFloat to return %f, got %f", t)

    // FIXME: also add teats for float32 parsing
    // FIXME: also need negative tests (that it fails to parse)
}
