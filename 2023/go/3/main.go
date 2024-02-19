package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type PartNumber struct {
    value string 
    depth int
    start int
    end int
}

type Symbol struct {
    value rune
    position int
    depth int
}

func isPeriod(r rune) bool {
    return r == '.'
}

func isNewline(r rune) bool {
    return r == '\n'
}

func isAdjacent(p PartNumber, s Symbol) bool {
    d := false 
    pos := false 

    switch s.depth {
    case p.depth - 1, p.depth, p.depth + 1:
        d = true 
    }

    switch {
    case s.position >= p.start - 1 && s.position <= p.end + 1:
        pos = true
    }

    if d == true && pos == true {
        return true
    }

    return false
}

func sumOfParts(p []PartNumber, s []Symbol) int {
    var v int = 0

    for _, cp := range p {
        for _, cs := range s {
            if isAdjacent(cp, cs) {
                r, err := strconv.Atoi(cp.value)

                if err != nil {
                    fmt.Println("Conversion Error: ", err)
                } else {
                    v += r 
                }

                break
            }
        }
    }

    return v
}

func sumOfGears(p []PartNumber, s []Symbol) int {
    var v = 0

    for _, cs := range s {
        if cs.value != '*' {
            continue
        }

        var p1 PartNumber
        var p2 PartNumber

        for _, cp := range p {
            if isAdjacent(cp, cs) {
                if p1.value == ""{
                    p1 = cp
                } else if p2.value == "" {
                    p2 = cp
                    break
                }
            }
        }

        if p1.value != " "&& p2.value != "" {
            p1v, err := strconv.Atoi(p1.value)

            if err != nil {
                fmt.Println("Conversion Error: ", err)
            }

            p2v, err := strconv.Atoi(p2.value)

            if err != nil {
                fmt.Println("Conversion Error: ", err)
            }

            v += p1v * p2v
        }
    }

    return v
}


func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    symbols := []Symbol{}
    parts := []PartNumber{}

    var cp PartNumber // current part number
    cd := 0 // current depth 
    i := 0 // iteration
    pp := false // part number processing

    for {
        r, _, err := reader.ReadRune() 

        if err != nil {
            if pp == true {
                cp.end = i - 1
                parts = append(parts, cp)
            }
            break
        }

        if isPeriod(r) {
            if pp == true {
                cp.end = i - 1
                parts = append(parts, cp)
                pp = false
            }
        } else if unicode.IsNumber(r) {
            if pp != true {
                cp = PartNumber{start: i, depth: cd, value: string(r)} 
                pp = true
            } else {
                cp.value += string(r)
            }
        } else if isNewline(r) {
            if pp == true {
                cp.end = i - 1
                parts = append(parts, cp)
                pp = false
            }
            cd += 1
            i = 0;
            continue
        } else {
            symbols = append(symbols, Symbol{value: r, position: i, depth: cd})

            if pp == true {
                cp.end = i -1
                parts = append(parts, cp)
                pp = false
            } 
        }

        i += 1;
    }

    x := sumOfParts(parts, symbols)
    fmt.Println("Sum or Part Numbers:", x)

    y := sumOfGears(parts, symbols)
    fmt.Println("Gear Ratio:", y)
}
