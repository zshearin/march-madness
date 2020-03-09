package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Matchup denotes a matchup
type Matchup struct {
	LowerSeed  int
	HigherSeed int
}

func main() {

	rand.Seed(time.Now().UnixNano())

	regions := []string{"Midwest", "South", "West", "East"}
	var finalFour []int
	for _, value := range regions {
		winner := runRegionSimulation(value)

		finalFour = append(finalFour, winner)

	}

	for key, value := range regions {

		fmt.Printf("%s winner: %d\n", value, finalFour[key])

	}

}

func runRegionSimulation(region string) int {

	fmt.Printf("%s region results: \n", region)
	var confSeeds []int
	for i := 0; i < 16; i++ {
		seed := i + 1
		confSeeds = append(confSeeds, seed)

	}

	round1Matchups := GetMatchups(confSeeds)
	round1Results := GetResultsArrayFromMatchups(round1Matchups)

	PrintMatchupsAndResults(round1Matchups, round1Results, 1)

	round2Matchups := GetMatchups(round1Results)
	round2Results := GetResultsArrayFromMatchups(round2Matchups)

	PrintMatchupsAndResults(round2Matchups, round2Results, 2)

	round3Matchups := GetMatchups(round2Results)
	round3Results := GetResultsArrayFromMatchups(round3Matchups)

	PrintMatchupsAndResults(round3Matchups, round3Results, 3)

	r4m := GetMatchups(round3Results)
	r4r := GetResultsArrayFromMatchups(r4m)

	PrintMatchupsAndResults(r4m, r4r, 4)

	fmt.Printf("\n\n\n")
	return r4r[0]

}

//PrintMatchupsAndResults prints the matchup results
func PrintMatchupsAndResults(matchups []Matchup, results []int, roundNumber int) {

	fmt.Println("======================================")
	fmt.Printf("Round %d matchups and results:\n", roundNumber)
	fmt.Println(matchups)
	fmt.Println(results)
	fmt.Println("======================================")

}

//GetResultsArrayFromMatchups gets the result array from the provided matchup
func GetResultsArrayFromMatchups(matchups []Matchup) []int {

	var results []int

	for _, matchup := range matchups {

		lowerSeed := matchup.LowerSeed
		higherSeed := matchup.HigherSeed

		result := notAnUpset(lowerSeed, higherSeed)

		if result {
			results = append(results, lowerSeed)
		} else {
			results = append(results, higherSeed)
		}
	}

	return results

}

//GetMatchups gets the matchups from the int array
func GetMatchups(round1Results []int) []Matchup {

	//
	var matchups []Matchup

	for len(round1Results) > 1 {

		var first, last int

		first, round1Results = round1Results[0], round1Results[1:]

		last, round1Results = round1Results[len(round1Results)-1], round1Results[:len(round1Results)-1]

		var curMatchup Matchup
		if first < last {
			curMatchup.LowerSeed = first
			curMatchup.HigherSeed = last
		} else {
			curMatchup.LowerSeed = last
			curMatchup.HigherSeed = first
		}

		matchups = append(matchups, curMatchup)
	}

	return matchups

}

//returns true if results expected (not an upset)
func notAnUpset(num1, num2 int) bool {

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
