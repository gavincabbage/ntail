package main

// TODO
// - remove wildcard from end of regex and fix that
// - fix repetitive for loop in parseSpecFile

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "io"
    "errors"
    "strings"
)

const (
    specRegex string = "/[[:alnum:]._-]+@[[:alnum:]._/-]+:[[:alnum:]._/-]+|[[:alnum:]._/-]+:[[:alnum:]._/-]+|[[:alnum:]._/-]+|.*/g"
)

func parseSpecFile(reader *bufio.Reader) ([]string, error) {
    lines := []string{}
    for line, eof := reader.ReadString('\n'); eof != io.EOF; line, eof = reader.ReadString('\n') {
        if eof != nil {
            return lines, errors.New("problem reading spec file")
        }
        line = strings.TrimSpace(line)
        fmt.Println(line)
        if match, err := regexp.MatchString(specRegex, line); err != nil {
            return lines, err
        } else if match {
            lines = append(lines, line)
        } else {
            return lines, errors.New("spec files syntax incorrect")
        }
    }
    return lines, nil
}

func main() {

    numArgs := len(os.Args)
    if numArgs < 0 || numArgs > 2 {
        panic("incorrect number of arguments")
    }
    fmt.Println(os.Args)

    specFile := os.Stdin
    if numArgs == 2 {
        var err error
        specFile, err = os.Open(os.Args[1])
        if err != nil {
            panic("could not open spec file")
        }
    }
    reader := bufio.NewReader(specFile)

    lines, err := parseSpecFile(reader)
    if err != nil {
        panic(err)
    }
    fmt.Println(lines)

    // TODO
    // - now we have the spec file, time to tail -f them all...

}
