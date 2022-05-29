package FE3H_greenhouse_combinations

import (
	"log"
	"strconv"
)

type Seed struct {
	name  string
	grade int64
	rank  int64
}

func CreateAvailableSeeds(data [][]string) (availableSeeds []Seed, err error) {
	availableSeeds = make([]Seed, 0)
	for _, entry := range data {
		if len(entry) != 3 {
			log.Fatalf("Expected 3 strings in data entry for Seed, Got: %v", entry)
		}

		var seed Seed

		seed.name = entry[0]
		seed.grade, err = strconv.ParseInt(entry[1], 10, 64)
		if err != nil {
			return nil, err
		}
		seed.rank, err = strconv.ParseInt(entry[2], 10, 64)
		if err != nil {
			return nil, err
		}

		availableSeeds = append(availableSeeds, seed)
	}
	return
}
