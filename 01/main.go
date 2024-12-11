package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// ReadInput reads the input file and returns two sorted slices of integers.
func ReadInput(file string) ([]int, []int, error) {
    data, err := os.ReadFile(file)
    if err != nil {
        return nil, nil, err
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    var arr1, arr2 []int

    for _, line := range lines {
        splitted := strings.Fields(line)
        if len(splitted) < 2 {
            continue
        }
        na, err := strconv.Atoi(splitted[0])
        if err != nil {
            return nil, nil, err
        }
        nb, err := strconv.Atoi(splitted[1])
        if err != nil {
            return nil, nil, err
        }
        arr1 = append(arr1, na)
        arr2 = append(arr2, nb)
    }

    sort.Ints(arr1)
    sort.Ints(arr2)

    return arr1, arr2, nil
}

// Solution1 calculates the distance based on the input file.
func Solution1(file string) (int, error) {
    arr1, arr2, err := ReadInput(file)
    if err != nil {
        return 0, err
    }

    distance := 0
    for i := range arr1 {
        distance += abs(arr2[i] - arr1[i])
    }

    return distance, nil
}

// Solution2 calculates the weighted sum of repeated elements based on the input file.
func Solution2(file string) (int, error) {
    arr1, arr2, err := ReadInput(file)
    if err != nil {
        return 0, err
    }

    rightRepeatCount := make(map[int]int)
    for _, v := range arr2 {
        rightRepeatCount[v]++
    }

    distance := 0
    for _, v := range arr1 {
        distance += rightRepeatCount[v] * v
    }

    return distance, nil
}

// abs returns the absolute value of an integer.
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    exampleFile := "01/example.txt"
    inputFile := "01/input.txt"

    if result, err := Solution1(exampleFile); err != nil {
        fmt.Printf("Error in Solution1 with example file: %v\n", err)
    } else {
        fmt.Printf("Part 01 example: %d\n", result)
    }

    if result, err := Solution1(inputFile); err != nil {
        fmt.Printf("Error in Solution1 with input file: %v\n", err)
    } else {
        fmt.Printf("Part 01: %d\n", result)
    }

    if result, err := Solution2(exampleFile); err != nil {
        fmt.Printf("Error in Solution2 with example file: %v\n", err)
    } else {
        fmt.Printf("Part 02 example: %d\n", result)
    }

    if result, err := Solution2(inputFile); err != nil {
        fmt.Printf("Error in Solution2 with input file: %v\n", err)
    } else {
        fmt.Printf("Part 02: %d\n", result)
    }
}