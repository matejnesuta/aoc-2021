package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "os"
	// "regexp"
	// "strconv"
	"strings"
)

var numbers = [...]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22,
	18, 20, 8, 19, 3, 26, 1}

var matrixes string = `22 13 17 11 0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

type coordinate_t struct {
	y int
	x int
}

type lookup_t struct {
	coordinates []coordinate_t
}

type checksum_t struct {
	rows    [5]int
	columns [5]int
}

type board_t struct {
	sum      int
	lastCall int
	matrix   [5][5]int
	lookup   map[int]coordinate_t
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func search(mainSlice *[]board_t) int {
	var checksum []checksum_t = make([]checksum_t, len(*mainSlice))
	for _, i := range numbers {
		for position, board := range *mainSlice {
			if val, ok := board.lookup[i]; ok {
				fmt.Println("value before decrement:	", board.sum)
				board.sum -= i
				fmt.Println("value after decrement:	", board.sum)
				checksum[position].rows[(val).y] += intPow(2, (val).x)
				checksum[position].columns[(val).x] += intPow(2, (val).y)
				board.lastCall = i
				// fmt.Println(board.lookup, board.lookup)
				// fmt.Println("pozice:", position, val)
				for x := 0; x < 5; x++ {
					if checksum[position].columns[x] == 31 ||
						checksum[position].rows[x] == 31 {
						// fmt.Println(board.sum, board.lastCall, board.matrix, position, iteration)
						return board.sum * board.lastCall
					}
				}
			}
			// fmt.Println("--------------------------------------")
		}
	}
	return -1
}

func parse(mainSlice *[]board_t) {
	var b board_t
	words := strings.Fields(matrixes)
	var iteration int
	b.lookup = map[int]coordinate_t{}
	var linker = map[int]coordinate_t{}
	for _, n := range words {
		if iteration == 0 {
			for k := range linker {
				delete(linker, k)
			}
			linker = make(map[int]coordinate_t)
		}
		number, _ := strconv.Atoi(n)
		b.matrix[iteration/5][iteration%5] = number
		b.sum += number
		var coordinate coordinate_t
		coordinate.x = iteration % 5
		coordinate.y = iteration / 5

		linker[number] = coordinate
		if iteration == 24 {
			copy := func(b map[int]coordinate_t) map[int]coordinate_t { return b }
			b.lookup = copy(linker)
			*mainSlice = append(*mainSlice, b)
			iteration = 0

			// for m, k := range (*mainSlice)[0].lookup {
			// 	fmt.Println(m, k)
			// }
			// fmt.Println(b.matrix)
		} else {
			iteration++
		}
	}
}

func main() {
	var mainSlice []board_t
	parse(&mainSlice)
	for _, i := range mainSlice{
		fmt.Println("lol")
		for m, k := range i.lookup{
			fmt.Println(m,k)
		}
		fmt.Println(i.matrix)
	}
	var result int = search(&mainSlice)
	fmt.Println(len(mainSlice))
	fmt.Printf("The result is:	%d\n", result)

}
