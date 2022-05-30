package fe3h_greenhouse_combinations

import (
	"reflect"
	"testing"
)

func Test_CreateAvailableSeeds(testFramework *testing.T) {
	type args struct {
		data [][]string
	}

	tests := []struct {
		name string
		args args
		want []Seed
	}{
		{
			name: "Given an empty input",
			args: args{
				data: [][]string{},
			},
			want: []Seed{},
		},
		{
			name: "Given a single entry",
			args: args{
				data: GetSeeds(1),
			},
			want: []Seed{
				{name: "Western Fodlan Seeds", grade: 1, rank: 9},
			},
		},
		{
			name: "Given 20 entries",
			args: args{
				data: GetSeeds(20),
			},
			want: []Seed{
				{"Western Fodlan Seeds", 1, 9},
				{"Mixed Herb Seeds", 1, 27},
				{"Vegetable Seeds", 1, 33},
				{"Mixed Fruit Seeds", 1, 44},
				{"Root Vegetable Seeds", 1, 49},
				{"Albinean Seeds", 2, 20},
				{"Morfis Seeds", 2, 23},
				{"Southern Fodlan Seeds", 2, 37},
				{"Eastern Fodlan Seeds", 2, 42},
				{"Northern Fodlan Seeds", 2, 53},
				{"Pale-Blue Flower Seeds", 3, 1},
				{"White Flower Seeds", 3, 5},
				{"Green Flower Seeds", 3, 10},
				{"Purple Flower Seeds", 3, 16},
				{"Red Flower Seeds", 3, 24},
				{"Blue Flower Seeds", 3, 38},
				{"Yellow Flower Seeds", 3, 55},
				{"Nordsalat Seeds", 4, 3},
				{"Morfis-Plum Seeds", 4, 18},
				{"Boa-Fruit Seeds", 5, 31},
			},
		},
	}
	for _, testCase := range tests {
		testFramework.Run(testCase.name, func(test *testing.T) {
			got, err := CreateAvailableSeeds(testCase.args.data)
			if err != nil {
				test.Errorf("An error occured: %v ", err)
			} else if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("\n CreateAvailableSeeds() = %v\n want %v", got, testCase.want)
			}
		})
	}
}
