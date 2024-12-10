package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"sort"
)

func ReadInput(file string) ([]int, []int) {
	dat, err := os.ReadFile(file)
    check(err)
	lines := strings.Split(string(dat), "\n")
	var arr1 []int
	var arr2 []int

	for _, line := range(lines) {
		splted := strings.Split(line, "   ")
		na, _ := strconv.Atoi(splted[0])
		nb, _ := strconv.Atoi(splted[1])
		arr1 = append(arr1, na)
		arr2 = append(arr2, nb)
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	return arr1, arr2
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Solution1(file string) int {
	arr1, arr2 := ReadInput(file)
	distance := 0
	for i := range(len(arr1)) {
		if arr2[i] > arr1[i] {
			distance += arr2[i] - arr1[i]
		} else {
			distance -= arr2[i] - arr1[i]
		}
	}
	return distance;
}

func Solution2(file string) int {
	arr1, arr2 := ReadInput(file)
	rightRepeatCount := make(map[int]int)
	distance := 0

	for _, v := range(arr2) {
		rightRepeatCount[v] = rightRepeatCount[v] + 1
	}

	for _, v := range(arr1) {
		distance +=  rightRepeatCount[v] * v;
	}

	return distance
}

func main() {
	fmt.Println("Part 01 example: ", Solution1("01/example.txt"))
	fmt.Println("Part 01: ", Solution1("01/input.txt"))
	fmt.Println("Part 02 example: ", Solution2("01/example.txt"))
	fmt.Println("Part 02: ", Solution2("01/input.txt"))
}