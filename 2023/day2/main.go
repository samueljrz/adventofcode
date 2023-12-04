package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id   int
	bags []Bag
}

func loadGames(input string, id int) Game {
	gameAns := Game{id: id, bags: []Bag{}}

	games := strings.Split(input, ": ")
	setOfGames := strings.Split(games[1], ";")

	reRed := regexp.MustCompile(`(\d+) red`)
	reGreen := regexp.MustCompile(`(\d+) green`)
	reBlue := regexp.MustCompile(`(\d+) blue`)

	for _, game := range setOfGames {
		matchesRed := reRed.FindStringSubmatch(game)
		matchesGreen := reGreen.FindStringSubmatch(game)
		matchesBlue := reBlue.FindStringSubmatch(game)
		var tmpRed, tmpGreen, tmpBlue int
		if len(matchesRed) > 1 {
			tmpRed, _ = strconv.Atoi(matchesRed[1])
		}
		if len(matchesGreen) > 1 {
			tmpGreen, _ = strconv.Atoi(matchesGreen[1])
		}
		if len(matchesBlue) > 1 {
			tmpBlue, _ = strconv.Atoi(matchesBlue[1])
		}
		gameAns.bags = append(gameAns.bags, Bag{red: tmpRed, green: tmpGreen, blue: tmpBlue})
	}

	return gameAns
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	mul := 0
	id := 1
	var games []Game

	defaultBag := Bag{red: 12, green: 13, blue: 14}

	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, loadGames(line, id))
		id++
	}
	for _, game := range games {
		isImpossible := false
		currentGameRedMax := 0
		currentGameGreenMax := 0
		currentGameBlueMax := 0
		for _, bag := range game.bags {
			if bag.blue > defaultBag.blue || bag.green > defaultBag.green || bag.red > defaultBag.red {
				isImpossible = true
			}
			if bag.blue > currentGameBlueMax {
				currentGameBlueMax = bag.blue
			}
			if bag.green > currentGameGreenMax {
				currentGameGreenMax = bag.green
			}
			if bag.red > currentGameRedMax {
				currentGameRedMax = bag.red
			}
		}
		if !isImpossible {
			sum += game.id
		}
		mul += currentGameRedMax * currentGameGreenMax * currentGameBlueMax
	}
	fmt.Println(sum)
	fmt.Println(mul)
}
