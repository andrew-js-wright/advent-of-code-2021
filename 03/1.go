package main

import (
    "path"
    "fmt"
    "bufio"
    "os"
    "runtime"
)

func createInputByte() uint {
    _, filename, _, _ := runtime.Caller(0)
    inputpath := path.Join(path.Dir(filename), "input.txt")
    readFile, _ := os.Open(inputpath)
    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    binaryCountArray := [12]int{0,0,0,0,0,0,0,0,0,0,0,0}

    for scanner.Scan() {
        value := scanner.Text()
        for index, bit := range value {
            switch bit {
            case '1': binaryCountArray[index]++
            case '0': binaryCountArray[index]--
            }
        }
    }

    var inputBytes uint;
    var oneBit uint = 1

    for _, bitCount := range binaryCountArray {
        inputBytes = inputBytes << 1
        if bitCount > 0 {
            inputBytes = inputBytes | oneBit
        }
    }

    return inputBytes
}

func calcdistance(inputBytes uint) uint {
    var gammaRate uint = inputBytes
    var epsilonRate uint = inputBytes ^ 4095

    return gammaRate * epsilonRate
}

func main() {
    fmt.Println(calcdistance(createInputByte()))
}
