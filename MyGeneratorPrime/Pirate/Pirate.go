package pirate

import (
	"fmt"
	"strings"
)
type Pirate struct {
    Name string
    Prime string
    Img string
}

func New(name, prime, img string) (*Pirate, error){

	P:= Pirate{
		Name: name,
		Prime: prime,
		Img: img,
	}

	if P.Name == "" || P.Prime == "" || P.Img == "" {
		fmt.Print("Erreur toutes les champs ne sont pas remplis ")
	}
	if P.Name != strings.ToUpper(P.Name){
		fmt.Print("Erreur le nom n'est pas écrit en majuscule")
	}
	
	return &P, nil
}
