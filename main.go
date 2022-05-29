package main

import (
	"FE3H-greehouse-combinations/FE3H_greenhouse_combinations"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("seeds.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	entries, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	availableSeeds, err := FE3H_greenhouse_combinations.CreateAvailableSeeds(entries)
	if err != nil {
		log.Fatal(err)
	}

	availableCombinations := FE3H_greenhouse_combinations.CreateAvailableCombinations(availableSeeds)

	fmt.Println(len(availableCombinations))
}
