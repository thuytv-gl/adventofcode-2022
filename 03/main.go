package main

import (
    "fmt"
    "regexp"
    "os"
    "strconv"
)

func Solution1(file string) int {
    data, _ := os.ReadFile(file)
    r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
    matches := r.FindAllString(string(data), -1)
    r1, _ := regexp.Compile(`\d{1,3}`)
    result := 0;
    for _, mul := range matches {
        nums := r1.FindAllString(mul, 2)
        num1, _ := strconv.Atoi(nums[0])
        num2, _ := strconv.Atoi(nums[1])
        result += num1 * num2;
    }
    return result;
}

func IsInDont(donts [][]int, item []int) bool {
    for _, dont := range donts {
        if item[0] > dont[0] && item[1] < dont[1] {
            return true
        }
    }
    return false
}

func Solution2(file string) int {
    data, _ := os.ReadFile(file)
    r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
    dontx, _ := regexp.Compile(`(?s)don't\(\).*?do\(\)|don't\(\).*?$`)
    muls := r.FindAllIndex(data, -1)
    r1, _ := regexp.Compile(`\d{1,3}`)

    donts := dontx.FindAllIndex(data, -1)

    result := 0
    for _, mul := range muls {
        if !IsInDont(donts, mul) {
            nums := r1.FindAllString(string(data[mul[0]:mul[1]]), 2)
            num1, _ := strconv.Atoi(nums[0])
            num2, _ := strconv.Atoi(nums[1])
            result += num1 * num2;
        }
    }

    return result;
}

func main() {
    fmt.Println("Part 01 example:", Solution1("03/example.txt"))
    fmt.Println("Part 01:", Solution1("03/input.txt"))
    fmt.Println("Part 02 example:", Solution2("03/example.txt"))
    fmt.Println("Part 02:", Solution2("03/input.txt"))
}