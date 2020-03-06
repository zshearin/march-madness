package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	for conf := 0; conf < 4; conf++ {
		var round1Winners []int

		for i := 0; i < 8; i++ {
			lowerSeed := i + 1
			higherSeed := 17 - lowerSeed

			result := upsetOrNot(lowerSeed, higherSeed)

			if result {
				round1Winners = append(round1Winners, lowerSeed)
			} else {
				round1Winners = append(round1Winners, higherSeed)
			}

		}
		fmt.Printf("Conference %d:\n", conf)
		fmt.Println(round1Winners)
	}
}

func upsetOrNot(num1, num2 int) bool {

	//make num1 the smaller number and num2 the bigger number (seeding processing later)
	if num1 > num2 {
		temp := num1
		num1 = num2
		num2 = temp
	}

	//Generate threshold and random number to assess against (probability based on seed value)
	threshold := calculateThreshold(num1, num2)
	random := getRandomNumber()

	//	printSeedingCalculations(random, threshold)
	//	printWinner(num1, num2, random, threshold)

	return random < threshold

}

func getRandomNumber() float64 {
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
