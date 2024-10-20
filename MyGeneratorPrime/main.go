package main

import (
	pdf "MyGeneratorPrime/MyGeneratorPrime/Pdf"
	pirate "MyGeneratorPrime/MyGeneratorPrime/Pirate"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	switch args [0] {
	case "-cli":
		if len(args) != 7 {
			log.Fatal("Not enough arguments")
			return
		}
		if args[1] != "-name" || args[3] != "-prime" || args[5] != "-img"{
			fmt.Print("entrer -cli, le nom, la prime et l'image ")
			return
		} else {
			name := args[2]
			prime := args[4]
			img := args[6]
			
			pirate, err := pirate.New(name, prime, img)

			if err != nil {
				return
			}

			pdf.Createpdf(pirate.Name, pirate.Prime, pirate.Img)
		}
	}
	
}