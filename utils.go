package utils

import (
    "fmt"
    "os"
    "strconv"
)

// ParseInt converts an int from a string representation to an actual int
func ParseInt(str string, opts ...int) (i int) {
    if raw, err := strconv.ParseInt(str, 10, bitSize(opts)); err == nil {
        i = int(raw)
    } else {
        fmt.Println(err)
        os.Exit(3)
    }

    return
}

// ParseFloat converts a float from a string representation to an actual float
func ParseFloat(str string, opts ...int) (f float64) {
    if raw, err := strconv.ParseFloat(str, bitSize(opts)); err == nil {
        f = float64(raw)
    } else {
        fmt.Println(err)
        os.Exit(3)
    }

    return
}

func bitSize(opts []int) (bs int) {
    if len(opts) > 0 {
        return opts[0]
    }

    return 16
}
