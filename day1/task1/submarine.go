package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func main() {
    prevMeasurement := -1
    scanner := bufio.NewScanner(os.Stdin)
    increases := 0
    for scanner.Scan() {
        measurement, err := strconv.Atoi(scanner.Text())
        if err != nil {
            os.Stderr.WriteString("Wrong input. Aborting the program!")
            os.Exit(1)
        } else if prevMeasurement == -1 {
            fmt.Printf("%d (N/A - no previous measurement)\n", measurement)
        } else if measurement > prevMeasurement {
            fmt.Printf("%d (increased)\n", measurement)
            increases++
        } else if measurement < prevMeasurement {
            fmt.Printf("%d (decreased)\n", measurement)
        } 
        prevMeasurement = measurement
    }  
    fmt.Printf("\nThe total number of increases is: %d\n",increases)
}
