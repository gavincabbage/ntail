package main

import (
    "fmt"
    "os"
    "errors"
)

func processArgs(args []string) error {
    if len(args) != 2 {
        return errors.New("incorrect number of arguments")
    }
    return nil
}

func main() {

    if e := processArgs(os.Args); e != nil {
        panic(e)
    }

    fmt.Println(os.Args)

}
