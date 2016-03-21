package main

import "testing"

func Test_processArgsReturnsErrorOnSingleArgument(t *testing.T) {
    mockArgs := []string{"first"}
    if processArgs(mockArgs) == nil {
        t.Fail()
    }
}

func Test_processArgsReturnsNilOnTwoArguments(t *testing.T) {
    mockArgs := []string{"first", "second"}
    if processArgs(mockArgs) != nil {
        t.Fail()
    }
}
