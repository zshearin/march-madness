package main

import "os"
import "fmt"
import "math/rand"
import "strconv"
import "time"

func main() {
	//Read in variables
	var number1, number2 = os.Args[1], os.Args[2]
	var sum int
	//Pass to int from string
	num1, err1 := strconv.Atoi(number1)
	num2, err2 := strconv.Atoi(number2)
	if err1 != nil {
	}
	if err2 != nil {
	}
	//make num1 the smaller number and num2 the bigger number (seeding processing later)
	if num1 > num2 {
		temp := num1
		num1 = num2
		num2 = temp
	}

	sum = num1 + num2
	var threshold, random float64

	//Generate threshold and random number to assess against (probability based on seed value)

	threshold = float64(num2) / float64(sum)
	rand.Seed(time.Now().UnixNano())
        random = rand.Float64()

	//Print out calculated values
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("=====================================================\n")
	fmt.Printf("Threshold: %f = %d / %d (higher seed over total)\n", threshold, num2, sum)
	fmt.Printf("Random #:  %f           (if lower than threshold, lower seed wins)\n", random)
	fmt.Printf("=====================================================\n")

	//Assess random value relative to threshold
	// For a 1 and 16 seed matchup, the 1 seed has a 16/17 chance
	// of winning and the 16 seed has a 1/17 chance

	if random < threshold {
		fmt.Printf("seed %d wins", num1)
	} else {
		fmt.Printf("seed %d wins (yikes)", num2)
	}

}
