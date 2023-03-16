package main

import (
	"math/rand"
	"time"

	"github.com/zshearin/march-madness/predict"
)

func main() {

	//Set seed based on time the application is run
	rand.Seed(time.Now().UnixNano())

	//Order of regions here determines who plays in the final four
	// 1st vs 2nd, 3rd vs 4th - then winners play each other
	regions := []string{"South", "East", "Midwest", "West"}

	finalFour := predict.RunRegionSimulationsAndGetFinalFour(regions)

	predict.CalculateAndPrintFinalFourResults(finalFour, regions)
}
