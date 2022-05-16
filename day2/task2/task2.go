package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func perror(msg string) {
	os.Stderr.WriteString(msg)
	os.Exit(1)
}

func main() {
	var position [2]int
	position[0] = 0
	position[1] = 0
	aim := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.IndexByte(line, ' '); idx >= 0 {
			direction := line[:idx]
			change, err := strconv.Atoi(line[idx+1:])
			if err != nil {
				perror("Given position is not an integer. Aborting!")
			}
			if direction == "forward" {
				position[1] += change
				position[0] += change * aim
			} else if direction == "up" {
				aim -= change
			} else if direction == "down" {
				aim += change
			} else {
				perror("The given word is not a valid direction. Aborting!")
			}

		} else {
			perror("Wrong input format. Aborting the program!")
		}
	}
	fmt.Printf("The final result is: %d\n", position[0]*position[1])
}
