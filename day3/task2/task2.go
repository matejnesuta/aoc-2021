package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"regexp"
)

func gammaDecode(list *[]string, index int) {
	zeroSlice := []string{}
	oneSlice := []string{}
	for _, element := range *list {
		if element[index] == '0' {
			zeroSlice = append(zeroSlice, element)
		} else {
			oneSlice = append(oneSlice, element)
		}
	}
	if len(zeroSlice) > len(oneSlice) {
		*list = zeroSlice
	} else {
		*list = oneSlice
	}
}

func epsilonDecode(list *[]string, index int) {
	zeroSlice := []string{}
	oneSlice := []string{}
	for _, element := range *list {
		if element[index] == '0' {
			zeroSlice = append(zeroSlice, element)
		} else {
			oneSlice = append(oneSlice, element)
		}
	}

	if len(zeroSlice) <= len(oneSlice) {
		*list = zeroSlice
	} else {
		*list = oneSlice
	}

}

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
	scanner := bufio.NewScanner(os.Stdin)
	zeroCount := 0
	oneCount := 0
	gamma := []string{}
	epsilon := []string{}
	gammaDec := 0
	epsilonDec := 0
	for scanner.Scan() {
		line := scanner.Text()
		re, _ := regexp.Compile("^([0-1]{12})$")
		if re.MatchString(line) {
			gamma = append(gamma, line)
			if line[0] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		} else {
			perror("Wrong input format. Aborting the program!")
		}
	}
	epsilon = gamma
	for i := 0; i < 12; i++ {
		if len(gamma) > 1 {
			gammaDecode(&gamma, i)
		}
		if len(epsilon) > 1 {
			epsilonDecode(&epsilon, i)
		}
	}
	println(gamma[0])
	println(epsilon[0])
	for i := 0; i < 12; i++ {
		if gamma[0][i]-48 == 1 {
			gammaDec += powInt(2, 11-i)
		}
		if epsilon[0][i]-48 == 1 {
			epsilonDec += powInt(2, 11-i)
		}
	}
	println(gammaDec)
	println(epsilonDec)
	fmt.Printf("The final desired value is: %d\n", gammaDec*epsilonDec)
}
