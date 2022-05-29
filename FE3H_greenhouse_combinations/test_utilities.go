package FE3H_greenhouse_combinations

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func GetSeeds(amountDesired int) (output [][]string) {
	for i := 0; i < amountDesired; {
		file, err := os.Open("../seeds.csv")
		if err != nil {
			log.Fatal(err)
		}
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
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
