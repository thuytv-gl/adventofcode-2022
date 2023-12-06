package main
import (
    "fmt"
    "os"
    "strings"
    "unicode"
    "strconv"
)

func head(s string) string {
    for _, c := range s {
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }
    }
    return ""
}

func tail(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        c :=s[i] 
        if unicode.IsNumber(rune(c)) {
            return string(c)
        }
    }
    return ""
}

func main() {
    dat, err := os.ReadFile("./src/day01.txt")
    if err != nil {
        panic(err)
    }
    puzzle := strings.Split(string(dat), "\r\n")
    result := 0
    for _, p := range puzzle {
        n1 := head(p)
        n2 := tail(p)
        if n, err := strconv.Atoi(n1+n2); err == nil {
            result = result + n
        }
    }
    fmt.Println(result)
}
