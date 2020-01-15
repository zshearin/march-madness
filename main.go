package main

import "os"
import "fmt"
import "math/rand"
import "strconv"
import "time"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need at least two arguments")
		return
	}

	//Read in variables
	var number1, number2 = os.Args[1], os.Args[2]
	num1, err1 := strconv.Atoi(number1)
	num2, err2 := strconv.Atoi(number2)

	if err1 != nil {
		fmt.Println("First argument given is not a number")
		return
	}
	if err2 != nil {
		fmt.Println("Second argument given is not a number")
		return
	}
	//make num1 the smaller number and num2 the bigger number (seeding processing later)
	if num1 > num2 {
		temp := num1
		num1 = num2
		num2 = temp
	}

	//Generate threshold and random number to assess against (probability based on seed value)
	threshold := calculateThreshold(num1, num2)
	random := getRandomNumber()

	printSeedingCalculations(random, threshold)
	printWinner(num1, num2, random, threshold)
}


func getRandomNumber() float64 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()
	return random
}

//Assess random value relative to threshold
// For a 1 and 16 seed matchup, the 1 seed has a 16/17 chance
// of winning and the 16 seed has a 1/17 chance
func calculateThreshold(num1, num2 int) float64 {
	sum := num1 + num2
	if num1 > num2 {
		num2 = num1
	}
	threshold := float64(num2) / float64(sum)
	return threshold
}

func printWinner(num1, num2 int, random, threshold float64) {
        if random < threshold {
                fmt.Printf("Seed %d wins", num1)
        } else {
		phrase := getPhrase(num2 - num1)
		fmt.Printf("Seed %d wins - %s", num2, phrase)
        }
}

func getPhrase(difference int) string {
	if difference > 9 {
		return "and your bracket is done"
	} else if difference > 5 {
		return "crazy prediction here"
	} else if difference > 3 {
		return "quality upset prediction"
	} else {
		return "barely an upset"
	}
}

func printSeedingCalculations(random, threshold float64) {
	//Print out calculated values
        fmt.Printf("==============================================================\n")
	fmt.Printf("Threshold: %f\n", threshold)
	fmt.Printf("Random # : %f (if lower than threshold, lower seed wins)\n", random)
        fmt.Printf("==============================================================\n")

}
