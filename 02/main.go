package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
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

func hasMistake(a, b int) bool {
    r := Abs(a - b)
    return r > 3 || r == 0
}

func Remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func AllIncreasing(line []int) bool {
    return sort.SliceIsSorted(line, func(a, b int) bool { return a < b })
}

func AllDecreasing(line []int) bool {
    return sort.SliceIsSorted(line, func(a, b int) bool { return a > b })
}

func IsSafe(line []int, selfCorrect int) bool {
    lim := len(line) - 1
    fixedCount := 0
    errors := 0

    for i := 0; i < lim; i++ {
        a := line[i]
        b := line[i + 1]

        if !(AllIncreasing(line) || AllDecreasing(line)) || hasMistake(a, b) {
            errors++
            if fixedCount < selfCorrect {
                if IsSafe(Remove(line, i), 0) || IsSafe(Remove(line, i + 1), 0) {
                    errors--
                    fixedCount++
                }
            }
        }
    }

    return errors == 0
}

func Solution(file string, selfCorrect int) int {
    safeCount := 0
    reports := ParseInput(file)

    for _, line := range reports {
        if IsSafe(line, selfCorrect) {
            safeCount++
        }
    }

    return safeCount
}

func main() {
    fmt.Println("Part 01 example:", Solution("02/example.txt", 0))
    fmt.Println("Part 01:", Solution("02/input.txt", 0))
    fmt.Println("Part 02 example:", Solution("02/example.txt", 1))
    fmt.Println("Part 02:", Solution("02/input.txt", 1))
}