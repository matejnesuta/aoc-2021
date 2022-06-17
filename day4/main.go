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

var numbers = [...]int {7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,
	18,20,8,19,3,26,1}

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

type coordinate_t struct{
	y int
	x int
}

type lookup_t struct{
	coordinates[]coordinate_t
}

type board_t struct{
	sum int
	lastCall int
	matrix[5][5]int
	lookup map[int]coordinate_t
} 

func parse(mainSlice*[]board_t){
	var b board_t
	words := strings.Fields(matrixes)
	var iteration int
	b.lookup = map[int]coordinate_t{}
	for _, n := range words{
		number, _ := strconv.Atoi(n)
		b.matrix[iteration/5][iteration%5] = number
		var coordinate coordinate_t
		coordinate.x = iteration%5
		coordinate.y = iteration/5
		b.lookup[number]=coordinate
		if iteration == 24 {
			*mainSlice = append(*mainSlice, b)
			iteration = 0
			fmt.Println(b.lookup)
			// for _, i := range b.lookup{
			// 	fmt.Println(i)
			// }
			for k := range b.lookup {
				delete(b.lookup, k)
			}
			
		}else {
			iteration ++
		}
	}
}

func main(){
	var mainSlice [] board_t
	parse(&mainSlice)

	fmt.Println(mainSlice[0])
	fmt.Println(mainSlice[1])
	fmt.Println(mainSlice[2])

}
