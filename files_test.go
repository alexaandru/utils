package utils

import (
    "io/ioutil"
    "strings"
    "testing"
)

func TestLoadFile(t *testing.T) {
    str1 := LoadFile("files_test.go")
    if str2, err := ioutil.ReadFile("files_test.go"); err == nil {
        if str1 != strings.Trim(string(str2), "\n") {
            t.Error("LoadFile failed")
        }
    } else {
        t.Error("Some error ocurred while trying to load test data")
    }
}

func TestLoadTwoStringsFromFile(t *testing.T) {
    str1, str2 := LoadTwoStringsFromFile("files_test.go")
    if str1 != "package utils" {
        t.Error("LoadTwoStringsFromFile failed")
    }
    if str2 != "" {
        t.Error("LoadTwoStringsFromFile failed")
    }
}
