package main

import "testing"

func Test_checkArgs_ReturnsErrorOnSingleArgument(t *testing.T) {
    mockArgs := []string{"first"}
    if processArgs(mockArgs) == nil {
        t.Fail()
    }
}

func Test_checkArgs_ReturnsNilOnTwoArguments(t *testing.T) {
    mockArgs := []string{"first", "second"}
    if processArgs(mockArgs) != nil {
        t.Fail()
    }
}
