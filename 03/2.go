package main

import (
    "path"
    "fmt"
    "bufio"
    "os"
    "runtime"
    "strconv"
    "strings"
)

type instruction struct {
    direction string
    distance int
}

func makeInstruction(dir, dis string) instruction {
    var distance, _ = strconv.Atoi(dis)
    return instruction{dir, distance}
}

func createinputarray() []instruction {
    _, filename, _, _ := runtime.Caller(0)
    inputpath := path.Join(path.Dir(filename), "input.txt")
    readFile, _ := os.Open(inputpath)
    scanner := bufio.NewScanner(readFile)
    scanner.Split(bufio.ScanLines)

    var inputarray []instruction

    for scanner.Scan() {
        value := scanner.Text()
        fields := strings.Fields(value)
        instruction := makeInstruction(fields[0], fields[1])
        inputarray = append(inputarray, instruction)
    }

    return inputarray
}

func calcdistance(instructions []instruction) int {
    var aim int = 0
    var depth int = 0
    var horizontal int = 0

    for _, instruction := range instructions {
        switch instruction.direction {
            case "forward":
                horizontal += instruction.distance
                depth += (instruction.distance * aim)
            case "down": aim += instruction.distance
            case "up": aim -= instruction.distance
        }
    }

    return horizontal * depth
}

func main() {
    fmt.Println(calcdistance(createinputarray()))
}
