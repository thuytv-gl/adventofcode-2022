package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func ParseInput(file string) ([][]int) {
    data, err := os.ReadFile(file)
    if err != nil {
        return nil
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    var reports [][]int

    for _, line := range lines {
        splitted := strings.Fields(line)
        var report []int
        for _, n := range splitted {
            num, _ := strconv.Atoi(n);
            report = append(report, num)
        }
        reports = append(reports, report)
    }
    return reports
}

func Abs(num int) int {
    if num > 0 {
        return num
    }
    return -num
}

func hasMistake(a, b int, direction bool) bool {
    r := Abs(a - b)
    return (a > b) != direction || r > 3 || r == 0
}

func GetErrorIndexes(line []int) []int {
    lim := len(line) - 1
    var errors []int
    direction := line[0] > line[1]

    for i := 0; i < lim; i++ {
        a := line[i]
        b := line[i + 1]

        if hasMistake(a, b, direction) {
            errors = append(errors, i + 0)
        }
    }

    return errors
}

func Solution(file string, maxerrors int) int {
    safeCount := 0
    reports := ParseInput(file)

    for _, line := range reports {
        errors := GetErrorIndexes(line)
        if len(errors) <= maxerrors {
            safeCount++
        }
        fmt.Println(errors, line)
    }

    return safeCount
}

func main() {
    fmt.Println("Part 01 example:", Solution("02/example.txt", 0))
    // fmt.Println("Part 01:", Solution("02/input.txt", 0))
    fmt.Println("Part 02 example:", Solution("02/example.txt", 1))
    // fmt.Println("Part 02:", Solution("02/input.txt", 1))
}