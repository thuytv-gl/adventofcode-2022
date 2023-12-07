package main
import (
    "fmt"
    "os"
    "strings"
    "unicode"
    "strconv"
)

// ============================================================
//              Part One
// ============================================================
type PartOne struct {}

func (self PartOne) head(s string) string {
    for _, c := range s {
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }
    }
    return ""
}

func (self PartOne) tail(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        c :=s[i] 
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }
    }
    return ""
}

func (self PartOne) start() {
    dat, err := os.ReadFile("./src/day01.txt")
    if err != nil {
        panic(err)
    }
    puzzle := strings.Split(string(dat), "\r\n")
    result := 0
    for _, p := range puzzle {
        n1 := self.head(p)
        n2 := self.tail(p)
        if n, err := strconv.Atoi(n1+n2); err == nil {
            result = result + n
        }
    }
    fmt.Println("Part one:", result)
}

// ============================================================
//              Part Two
// ============================================================
type PartTwo struct {}
var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func (self PartTwo) head(s string) string {
    s_len := len(s)
    for i, c := range s {
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }
        for n, nt := range numbers {
            if i + len(nt) < s_len &&  s[i:i+len(nt)] == nt {
                return strconv.Itoa(n+1)
            }
        }
    }
    return ""
}

func (self PartTwo) tail(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        c :=s[i] 
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }

        for n, nt := range numbers {
            if (i - len(nt) > -2) && s[i-len(nt)+1:i+1] == nt {
                return strconv.Itoa(n+1)
            }
        }
    }
    return ""
}

func (self PartTwo) start() {
    dat, err := os.ReadFile("./src/day01_2.txt")
    if err != nil {
        panic(err)
    }
    puzzle := strings.Split(string(dat), "\r\n")
    result := 0
    for _, p := range puzzle {
        n1 := self.head(p)
        n2 := self.tail(p)
        if n, err := strconv.Atoi(n1+n2); err == nil {
            result = result + n
        }
    }
    fmt.Println("Part two:", result)
}

func main() {
    PartOne{}.start()
    PartTwo{}.start()
}
