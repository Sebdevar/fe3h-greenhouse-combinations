package FE3H_greehouse_combinations

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func getSeeds(amountDesired int) (output [][]string) {
	file, err := os.Open("seeds.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	for i := 0; i < amountDesired; {
		reader := csv.NewReader(file)

		for ; i < amountDesired || err == io.EOF; i++ {

			entry, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			output = append(output, entry)
		}
	}
	return
}

func Test_createAvailableSeeds(testFramework *testing.T) {
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
				data: getSeeds(1),
			},
			want: []Seed{
				{name: "Western Fodlan Seeds", grade: 1, rank: 9},
			},
		},
		{
			name: "Given 20 entries",
			args: args{
				data: getSeeds(50),
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
				{"Boa-Fruit Seeds", 5, 3},
			},
		},
	}
	for _, testCase := range tests {
		testFramework.Run(testCase.name, func(test *testing.T) {
			got, err := createAvailableSeeds(testCase.args.data)
			if err != nil {
				test.Errorf("An error occured: %v ", err)
			} else if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("createAvailableSeeds() = %v, want %v", got, testCase.want)
			}
		})
	}
}
