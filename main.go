package main

import (
	"FE3H-greehouse-combinations/fe3h_greenhouse_combinations"
	"log"
	"os"
)

func main() {
	availableSeeds, err := fe3h_greenhouse_combinations.GetDefaultAvailableSeeds()
	if err != nil {
		log.Fatal(err)
	}

	availableCombinations := fe3h_greenhouse_combinations.CreateAvailableCombinations(availableSeeds)

	file, err := os.Create("combinations.csv")
	if err != nil {
		log.Fatal(err)
	}

	fe3h_greenhouse_combinations.ToCSV(availableCombinations, file)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
