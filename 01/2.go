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

func countwindowincreases(values []int64) int {
    var count = 0
    windowSize := 3
    relativeWindowSize := windowSize - 1
    var previousWindowValue int64 = 0

    for index, _ := range values {

        if index < relativeWindowSize {
            // not yet reached the first window
            continue
        }

        // calculate the current window value
        var currentWindowValue int64 = 0
        for i:= (index - relativeWindowSize); i<=index; i++ {
            currentWindowValue += values[i]
        }

        // check if thit window is greater than the last
        if previousWindowValue != 0 && previousWindowValue < currentWindowValue {
            count = count + 1
        }

        // store this window's value for the next iteration
        previousWindowValue = currentWindowValue
    }

    return count
}

func main() {
    fmt.Println(countwindowincreases(createinputarray()))
}
