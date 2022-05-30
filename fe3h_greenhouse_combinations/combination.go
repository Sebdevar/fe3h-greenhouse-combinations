package fe3h_greenhouse_combinations

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type Combination struct {
	numberOfSeeds int
	score         int
	minimumYield  int
	maximumYield  int
	seeds         []seedWithAmount
}

type seedWithAmount struct {
	Seed
	amount int
}

func CreateAvailableCombinations(availableSeeds []Seed) (availableCombinations []Combination) {
	availableCombinations = make([]Combination, 0)
	getSeedCombinations(1, availableSeeds, 0, &availableCombinations, []seedWithAmount{})
	getSeedCombinations(2, availableSeeds, 0, &availableCombinations, []seedWithAmount{})
	getSeedCombinations(3, availableSeeds, 0, &availableCombinations, []seedWithAmount{})
	getSeedCombinations(4, availableSeeds, 0, &availableCombinations, []seedWithAmount{})
	getSeedCombinations(5, availableSeeds, 0, &availableCombinations, []seedWithAmount{})
	return
}

func ToCSV(combinations []Combination, writer io.Writer) {
	linesToWrite := [][]string{{"Number of seeds", "Score", "Minimum Yield", "Maximum Yield", "Seeds"}}
	for _, combination := range combinations {
		var seeds []string
		for _, seed := range combination.seeds {
			seeds = append(seeds, fmt.Sprint(seed.name, "X", seed.amount))
		}

		newLine := []string{
			strconv.Itoa(combination.numberOfSeeds),
			strconv.Itoa(combination.score),
			strconv.Itoa(combination.minimumYield),
			strconv.Itoa(combination.maximumYield),
			strings.Join(seeds, ", "),
		}
		linesToWrite = append(linesToWrite, newLine)
	}

	csvWriter := csv.NewWriter(writer)
	err := csvWriter.WriteAll(linesToWrite)
	if err != nil {
		log.Fatal(err)
	}
}

func getSeedCombinations(numberOfSeeds int, availableSeeds []Seed, lookupIndex int, combinations *[]Combination, currentSeeds []seedWithAmount) {
	if getSeedCount(currentSeeds) == numberOfSeeds {
		*combinations = append(*combinations, *createCombination(currentSeeds))
		return
	}
	for index := lookupIndex; index < len(availableSeeds); index++ {
		currentSeeds = addSeed(currentSeeds, availableSeeds[index])
		getSeedCombinations(numberOfSeeds, availableSeeds, index, combinations, currentSeeds)
		currentSeeds = removeLastSeed(currentSeeds)
	}
}

func addSeed(currentSeeds []seedWithAmount, seedToAdd Seed) []seedWithAmount {
	if len(currentSeeds) != 0 && currentSeeds[len(currentSeeds)-1].name == seedToAdd.name {
		currentSeeds[len(currentSeeds)-1].amount++
	} else {
		currentSeeds = append(currentSeeds, seedWithAmount{Seed: seedToAdd, amount: 1})
	}
	return currentSeeds
}

func removeLastSeed(currentSeeds []seedWithAmount) []seedWithAmount {
	if currentSeeds[len(currentSeeds)-1].amount > 1 {
		currentSeeds[len(currentSeeds)-1].amount--
	} else {
		currentSeeds = currentSeeds[:len(currentSeeds)-1]
	}
	return currentSeeds
}

func createCombination(seeds []seedWithAmount) (combination *Combination) {
	combination = new(Combination)
	combination.numberOfSeeds = getSeedCount(seeds)
	combination.seeds = make([]seedWithAmount, len(seeds))
	copy(combination.seeds, seeds)
	combination.score = calculateScore(seeds)
	combination.minimumYield, combination.maximumYield = identifyPossibleYields(combination.score)
	return
}

func getSeedCount(seeds []seedWithAmount) (seedCount int) {
	for _, seed := range seeds {
		seedCount += seed.amount
	}
	return
}

func calculateScore(seeds []seedWithAmount) (score int) {
	var sumOfSeedRanks int
	var sumOfSeedGrades int

	for _, seed := range seeds {
		sumOfSeedRanks += int(seed.rank) * seed.amount
		sumOfSeedGrades += int(seed.grade) * seed.amount
	}

	rankScore := (12 - (sumOfSeedRanks % 12)) * 5
	gradeScore := math.Floor((float64(sumOfSeedGrades) / 5) * 4)

	return rankScore + int(gradeScore)
}

func identifyPossibleYields(score int) (minimumYield int, maximumYield int) {
	switch {
	case score < 21:
		minimumYield = 0
		maximumYield = 1
		break
	case score < 41:
		minimumYield = 1
		maximumYield = 2
		break
	case score < 61:
		minimumYield = 2
		maximumYield = 3
		break
	case score < 71:
		minimumYield = 3
		maximumYield = 4
		break
	case score < 81:
		minimumYield = 3
		maximumYield = 5
		break
	default:
		log.Fatalf("Received a score higher than 80, should not be possible. Score received: %v", score)
	}
	return
}
