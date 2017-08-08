package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var testData = flag.String("testData", "./data.txt", "path to test data file")
var compass = map[string]string{
	"NL": "W",
	"NR": "E",
	"EL": "N",
	"ER": "S",
	"SL": "E",
	"SR": "W",
	"WL": "S",
	"WR": "N",
}

func setBoundary(maxX int, maxY int) func(int, int) bool {
	return func(x int, y int) bool {
		if x > maxX || x < 0 || y > maxY || y < 0 {
			return true
		}
		return false
	}
}

func strToInt(str string) (int, error) {
	var err error
	var i64 int64
	if i64, err = strconv.ParseInt(str, 10, 32); err == nil {
		return int(i64), nil
	}
	return int(0), err
}

func getStartPoint(data []string) (int, int, string) {
	var x, y int
	var err error
	x, err = strToInt(data[0])
	y, err = strToInt(data[1])
	if len(data) != 3 || err != nil {
		log.Fatalf("Error in file, should be [int int string] but get %v", data)
	}
	return x, y, data[len(data)-1]
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s with error %v", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	boundary := strings.Split(scanner.Text(), " ")
	if len(boundary) != 2 {
		log.Fatalf("Error in the first line %s", boundary)
	}
	maxX, _ := strToInt(boundary[0])
	maxY, _ := strToInt(boundary[1])
	outOfBound := setBoundary(maxX, maxY)

	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), " ")
		x, y, direction := getStartPoint(coordinates)
		scanner.Scan()
		moves := strings.Split(scanner.Text(), " ")

		for _, move := range moves[0] {
			if string(move) != "F" {
				direction = compass[direction+string(move)]
				continue
			}
			var nextX, nextY int
			switch direction {
			case "N":
				nextY = y + 1
				nextX = x
			case "E":
				nextX = x + 1
				nextY = y
			case "S":
				nextY = y - 1
				nextX = x
			case "W":
				nextX = x - 1
				nextY = y
			}
			if outOfBound(nextX, nextY) {
				continue
			}
			x = nextX
			y = nextY
		}
		fmt.Println(x, y, direction)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {

	flag.Parse()

	if *testData == "" {
		log.Fatal("Please provide test file")
	}
	readFile(*testData)
}
