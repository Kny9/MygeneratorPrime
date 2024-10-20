package Csvparser

import (
	pdf "MyGeneratorPrime/MyGeneratorPrime/Pdf"
	pirates "MyGeneratorPrime/MyGeneratorPrime/Pirate"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ParseCSV(filepath string) {
	data, err := os.Open(filepath)
	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print("Erreur dans la lecture du csv", err, "\n")
	}
	for _, line := range records {

		pirate,err := pirates.New(line[1], line[0], line[2])
		pdf.Createpdf(pirate.Name, pirate.Prime, pirate.Img)
		if err != nil {
			log.Fatal("Error")
		}
	}
}