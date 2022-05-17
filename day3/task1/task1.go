package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"regexp"
)

func perror(msg string) {
	os.Stderr.WriteString(msg)
	os.Exit(1)
}

func powInt(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}

func main() {
	word := []int {0,0,0,0,0,0,0,0,0,0,0,0}
	gamma := 0 
	epsilon := 0 
	word[0]=0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		re, _ := regexp.Compile("^([0-1]{12})$")
		if  re.MatchString(line) {
			for i, rune := range line {
				if rune == '0'{
					word[i]--
				} else {
					word[i]++
				}
			}

		} else {
			perror("Wrong input format. Aborting the program!")
		}
	}
	for index, element := range word {
			if element > 0 {
				epsilon+=powInt(2, 11-index)
			} else {
				gamma+=powInt(2, 11-index)
			}
		}
		fmt.Printf("The final result is: %d\n", gamma*epsilon)
}
