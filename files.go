package utils

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

// LoadFile loads a file and returns it as a string, handling reading errors internally
func LoadFile(fname string) (data string) {
    if rawData, err := ioutil.ReadFile(fname); err == nil {
        data = strings.Trim(string(rawData), "\n")
    } else {
        fmt.Println(err)
        os.Exit(2)
    }

    return
}

// LoadTwoStringsFromFile loads the first two strings (separated by \n) from a file
// NOTE: Inefficient when you have (much) more than just two strings in a file.
func LoadTwoStringsFromFile(fname string) (string, string) {
    lines := strings.Split(LoadFile(fname), "\n")
    return lines[0], lines[1]
}
