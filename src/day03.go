package main
import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    dat, err := os.ReadFile("./src/day03.txt")
    if err != nil {
        panic(err)
    }
    puzzle := strings.Split(string(dat), "\r\n")
    result := 0
    p_len := len(puzzle) - 1
    for n, line := range puzzle {
        starts := -1
        ends := -1
        for i, s := range line {
            if s >= '0' && s <= '9' {
                if starts == -1 {
                    starts = i
                }

                if i == len(line) - 1 {
                    ends = i + 1
                }
            } else if starts != -1{
                ends = i
            }

            if starts != -1 && ends != -1 {
                isGear := false
                num_s := line[starts:ends]
                if (starts - 1 > -1 && line[starts - 1] != '.') || (ends != len(line) && line[ends] != '.') {
                    isGear = true
                } else {
                    // check for gears
                    for j := starts - 1; j <= ends; j++ {
                        if j < 0 || j > len(line) {
                            continue
                        }

                        // line above
                        if n != 0 && j < len(puzzle[n-1]) && puzzle[n-1][j] != '.'{
                            isGear = true
                            break
                        }

                        // line below
                        if n < p_len && j < len(puzzle[n+1]) && puzzle[n+1][j] != '.' {
                            isGear = true
                            break
                        }
                    }
                }

                if isGear {
                    num_d, _ := strconv.Atoi(string(num_s))
                    result += num_d
                }

                starts = -1
                ends = -1
            }
        }
    }
    fmt.Println("Total:", result)
}
