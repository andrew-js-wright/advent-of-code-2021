package main

import (
    "path"
    "fmt"
    "bufio"
    "os"
    "runtime"
    "strconv"
)

func createinputarray() []int64 {
    _, filename, _, _ := runtime.Caller(0)
    inputpath := path.Join(path.Dir(filename), "input.txt")
    readFile, _ := os.Open(inputpath)
    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    var inputarray []int64

    for scanner.Scan() {
        value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
        inputarray = append(inputarray, value)
    }

    return inputarray
}

func countincreases(values []int64) int {
    var count = 0

    for index, value := range values {
        if index == 0 {
            continue
        }

        if values[index-1] < value {
            count = count + 1
        }
    }

    return count
}

func main() {
    fmt.Println(countincreases(createinputarray()))
}
