package FE3H_greenhouse_combinations

import (
	"reflect"
	"testing"
)

func Test_CreateAvailableCombinations(testFramework *testing.T) {
	type args struct {
		availableSeeds []Seed
	}

	mockSeeds, err := CreateAvailableSeeds(GetSeeds(10))
	if err != nil {
		testFramework.Fatal(err)
	}

	tests := []struct {
		name string
		args args
		want []Combination
	}{
		{
			name: "Given an empty input",
			args: args{
				availableSeeds: []Seed{},
			},
			want: []Combination{},
		},
		{
			name: "Given a single seed",
			args: args{
				availableSeeds: mockSeeds[0:1],
			},
			want: []Combination{
				{numberOfSeeds: 1, score: 15, minimumYield: 0, maximumYield: 1, seeds: []Seed{{"Western Fodlan Seeds", 1, 9}}},
				{numberOfSeeds: 2, score: 31, minimumYield: 1, maximumYield: 2, seeds: []Seed{{"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}}},
				{numberOfSeeds: 3, score: 47, minimumYield: 2, maximumYield: 3, seeds: []Seed{{"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}}},
				{numberOfSeeds: 4, score: 63, minimumYield: 3, maximumYield: 4, seeds: []Seed{{"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}}},
				{numberOfSeeds: 5, score: 19, minimumYield: 0, maximumYield: 1, seeds: []Seed{{"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}, {"Western Fodlan Seeds", 1, 9}}},
			},
		},
	}
	for _, testCase := range tests {
		testFramework.Run(testCase.name, func(test *testing.T) {
			got := CreateAvailableCombinations(testCase.args.availableSeeds)
			if err != nil {
				test.Errorf("An error occured: %v ", err)
			} else if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("\n CreateAvailableCombinations() = %v\n want %v", got, testCase.want)
			}
		})
	}
}
