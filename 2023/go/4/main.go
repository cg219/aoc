package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
    id int
    numbers [10]int
    score int
}

func splitWords(input <- chan string, out chan <- string, re *regexp.Regexp) {
    for line := range input {
        parts := re.Split(line, -1)

        for _, val := range parts {
            out <- strings.TrimSpace(val)
        }
    }
    close(out)
}

func sortData(input <- chan string, out chan <- int, re *regexp.Regexp) {
    for chunk := range input {

        switch {
        case strings.Contains(chunk, "Card"):
            matches := re.FindAllString(chunk, 1)

            for _, id := range matches {
                v, err := strconv.Atoi(id) 
                if err != nil {
                    out <- 00
                }

                out <- v
            }

        default:
            for _, number := range strings.Fields(chunk) {
                v, err := strconv.Atoi(number) 
                if err != nil {
                    out <- 00
                }

                out <- v
            }
        }
    }
    close(out)
}

func loadData(out chan <- string, reader *bufio.Reader) {
    for {
        w, err := reader.ReadString('\n')

        if err != nil {
            break
        }

        out<- w
    }

    close(out)
}

func makeGames(input <- chan int, out chan <- Game) {
    dataPosition := 0
    buffer := [36]int{}

    for number := range input {
        buffer[dataPosition] = number
        dataPosition += 1

        if dataPosition == 36  {
            game := Game{ id: buffer[0], numbers: [10]int{} }

            for i := 0; i < 10; i++ {
                game.numbers[i] = buffer[i + 1]
            }

            for i := 11; i < 36; i++ {
                game.UpdateScore(buffer[i])
            }

            out <- game
            
            buffer = [36]int{} 
            dataPosition = 0
        }
    }

    close(out)
}

func (game * Game) UpdateScore(entry int) {
    for _, number := range game.numbers {
        if number == entry {
            if game.score == 0 {
                game.score = 1
            } else {
                game.score *= 2
            }
            break
        }
    }
}

func main() {
    file, err := os.Open("data.txt")    

    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    lineRegex := regexp.MustCompile("[:|]+")
    digitRegex := regexp.MustCompile(`\d+`)

    dataChannel := make(chan string)
    wordsChannel := make(chan string)
    outputChannel := make(chan Game)
    gameChannel := make(chan int)

    go loadData(wordsChannel, reader)
    go splitWords(wordsChannel, dataChannel, lineRegex)
    go sortData(dataChannel, gameChannel, digitRegex)
    go makeGames(gameChannel, outputChannel)

    totalScore := 0
    for game := range outputChannel {
        totalScore += game.score
        fmt.Println(game)
    }

    fmt.Println(totalScore)
}
