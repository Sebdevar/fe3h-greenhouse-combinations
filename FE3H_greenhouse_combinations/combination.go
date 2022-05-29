package FE3H_greenhouse_combinations

import (
	"log"
	"math"
)

type Combination struct {
	numberOfSeeds int
	score         int
	minimumYield  int
	maximumYield  int
	seeds         []Seed
}

func CreateAvailableCombinations(availableSeeds []Seed) (availableCombinations []Combination) {
	availableCombinations = make([]Combination, 0)
	getSeedCombinations(1, availableSeeds, 0, &availableCombinations, []Seed{})
	getSeedCombinations(2, availableSeeds, 0, &availableCombinations, []Seed{})
	getSeedCombinations(3, availableSeeds, 0, &availableCombinations, []Seed{})
	getSeedCombinations(4, availableSeeds, 0, &availableCombinations, []Seed{})
	getSeedCombinations(5, availableSeeds, 0, &availableCombinations, []Seed{})
	return
}

func getSeedCombinations(numberOfSeeds int, availableSeeds []Seed, lookupIndex int, combinations *[]Combination, currentSeeds []Seed) {
	if len(currentSeeds) == numberOfSeeds {
		*combinations = append(*combinations, *createCombination(currentSeeds))
		return
	}
	for index := lookupIndex; index < len(availableSeeds); index++ {
		currentSeeds = append(currentSeeds, availableSeeds[index])
		getSeedCombinations(numberOfSeeds, availableSeeds, index, combinations, currentSeeds)
		currentSeeds = currentSeeds[:len(currentSeeds)-1]
	}
}

func createCombination(seeds []Seed) (combination *Combination) {
	combination = new(Combination)
	combination.numberOfSeeds = len(seeds)
	combination.seeds = make([]Seed, len(seeds))
	copy(combination.seeds, seeds)
	combination.score = calculateScore(seeds)
	combination.minimumYield, combination.maximumYield = identifyPossibleYields(combination.score)
	return
}

func calculateScore(seeds []Seed) (score int) {
	var sumOfSeedRanks int
	var sumOfSeedGrades int

	for _, seed := range seeds {
		sumOfSeedRanks += int(seed.rank)
		sumOfSeedGrades += int(seed.grade)
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
