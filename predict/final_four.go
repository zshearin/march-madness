package main

import (
	"fmt"
	"strconv"
	"strings"
)

//CalculateAndPrintFinalFourResults calculates and prints final four results
func CalculateAndPrintFinalFourResults(finalFour []int, regions []string) {

	for index, regionName := range regions {

		space := strings.Repeat(" ", 8-len(regionName))

		fmt.Printf("%s winner: %s%d\n", regionName, space, finalFour[index])

	}

	fmt.Printf("\nFinal Four Matchups and Results: \n\n")
	var champ1, champ2, champion int
	var champ1Str, champ2Str, championStr string

	bracket1Winner := finalFour[0]
	bracket2Winner := finalFour[1]
	bracket3Winner := finalFour[2]
	bracket4Winner := finalFour[3]

	result1 := finalFourResult(bracket1Winner, bracket4Winner)
	if result1 {
		champ1Str = regions[0]
		champ1 = finalFour[0]
	} else {
		champ1Str = regions[3]
		champ1 = finalFour[3]
	}

	result2 := finalFourResult(bracket2Winner, bracket3Winner)

	if result2 {
		champ2Str = regions[1]
		champ2 = finalFour[1]
	} else {
		champ2Str = regions[2]
		champ2 = finalFour[2]
	}

	result3 := finalFourResult(champ1, champ2)

	if result3 {
		champion = champ1
		championStr = champ1Str
	} else {
		champion = champ2
		championStr = champ2Str
	}

	printChamp1 := strconv.Itoa(champ1) + " (" + champ1Str + ")"
	printChamp2 := strconv.Itoa(champ2) + " (" + champ2Str + ")"

	finalChampion := strconv.Itoa(champion) + " (" + championStr + ")"

	printResults(printChamp1, printChamp2, finalChampion, finalFour, regions)

}

func printResults(printChamp1, printChamp2, finalChampion string, finalFour []int, regions []string) {

	var initialRound []string

	for i, value := range regions {
		curString := strconv.Itoa(finalFour[i]) + " (" + value + ")"
		curBlank := strings.Repeat("_", 15-len(curString))
		newString := curString + curBlank
		initialRound = append(initialRound, newString)
	}

	afterChamp1 := strings.Repeat("_", 12-len(printChamp1))
	afterChamp2 := strings.Repeat("_", 12-len(printChamp2))

	fmt.Printf("___ %s\n", initialRound[0])
	fmt.Printf("                   |___ %s %s\n", printChamp1, afterChamp1)
	fmt.Printf("___ %s|                 |\n", initialRound[3])
	fmt.Printf("                                     |___ %s____ \n", finalChampion)
	fmt.Printf("___ %s                  |\n", initialRound[1])
	fmt.Printf("                   |___ %s %s|\n", printChamp2, afterChamp2)
	fmt.Printf("___ %s|\n", initialRound[2])
}

//returns true if results expected (not an upset)
func finalFourResult(num1, num2 int) bool {
	threshold := .5

	random := getRandomNumber()

	return random < threshold
}
