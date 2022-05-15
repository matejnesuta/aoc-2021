package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func main() {
    b := 0
    c := 0
    scanner := bufio.NewScanner(os.Stdin)
    increases := 0
    counter := 0
    prevMeasurement := -1
    measurement := 0
    for scanner.Scan() {
        counter++ 
        a, err := strconv.Atoi(scanner.Text())
        if err != nil {
            os.Stderr.WriteString("Wrong input. Aborting the program!")
            os.Exit(1)
        } else if counter >= 3 {
            measurement = a + b + c
            if prevMeasurement == -1{ 
                fmt.Printf("%d (N/A - no previous measurement)\n", measurement)
            } else if measurement > prevMeasurement {
                fmt.Printf("%d (increased)\n", measurement)
                increases++
            } else if measurement < prevMeasurement {
                fmt.Printf("%d (decreased)\n", measurement)
            } else if measurement == prevMeasurement {
                fmt.Printf("%d (decreased)\n", measurement)
            }
            prevMeasurement = measurement
        } 
        c = b
        b = a 
    }  
    fmt.Printf("\nThe total number of increases is: %d\n",increases)
}
