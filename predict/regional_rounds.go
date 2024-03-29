package predict

import (
	"fmt"
	"math/rand"
	"strings"
)

//Matchup denotes a matchup
type Matchup struct {
	LowerSeed  int
	HigherSeed int
}

func RunRegionSimulationsAndGetFinalFour(regions []string) []int {
	var finalFour []int
	for _, regionName := range regions {
		winner := runRegionSimulation(regionName)
		finalFour = append(finalFour, winner)
	}
	return finalFour
}

func runRegionSimulation(region string) int {
	fmt.Println("==========================================================")
	underline := strings.Repeat("-", 16+len(region))
	fmt.Printf("%s region results: \n", region)
	fmt.Println(underline)

	var confSeeds []int
	for i := 0; i < 16; i++ {
		seed := i + 1
		confSeeds = append(confSeeds, seed)
	}

	matchupsRound1 := GetMatchups(confSeeds)

	resultsRound1 := GetMatchupResults(matchupsRound1)
	matchupsRound2 := GetMatchups(resultsRound1)
	PrintRoundMatchups(matchupsRound2, 2)

	resultsRound2 := GetMatchupResults(matchupsRound2)
	matchupsRound3 := GetMatchups(resultsRound2)
	PrintRoundMatchups(matchupsRound3, 3)

	resultsRound3 := GetMatchupResults(matchupsRound3)
	matchupsRound4 := GetMatchups(resultsRound3)
	PrintRoundMatchups(matchupsRound4, 4)

	resultsRound4 := GetMatchupResults(matchupsRound4)

	winner := resultsRound4[0]
	fmt.Printf("Winner of region: %d\n\n", winner)
	return winner
}

//PrintRoundMatchups prints the matchups for a round
func PrintRoundMatchups(matchups []Matchup, roundNumber int) {

	fmt.Printf("Round %d matchups: ", roundNumber)
	for _, value := range matchups {
		fmt.Printf("%d vs. %d   ", value.LowerSeed, value.HigherSeed)
	}
	fmt.Printf("\n")
}

//GetMatchupResults gets the result array from the provided matchup
func GetMatchupResults(matchups []Matchup) []int {

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
