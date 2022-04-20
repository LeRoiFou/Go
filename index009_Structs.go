// Lien : https://www.youtube.com/watch?v=YAfRxJBELv0
// Cours : Structs In Go - Learn Golang #9

// Dans ce programme on apprend à faire appel aux structures (similaire aux classes)

// Date : 20-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
)

// Structure (similaire à une classe)
type Person struct {
	name     string
	age      int
	favColor string
	weight   int
}

func main() {

	// Assignation de la structure
	var person1 Person

	// Affectation de valeurs aux variables de la structure
	person1.name = "John"
	person1.age = 45
	person1.favColor = "Noir"
	person1.weight = 173
	fmt.Println("Personne n° 1 :", person1)
	fmt.Println("Nom de la personne n° 1 :", person1.name)

	// Changement de la valeur d'une variable
	person1.name = "John Gerald"
	fmt.Println("Nom de la personne n° 1 :", person1.name)

	// Raccourci pour assigner et affecter des valeurs
	person2 := Person{
		name:     "Dean",
		age:      47,
		favColor: "Blanc",
		weight:   185,
	}
	fmt.Println("Personne n° 2 :", person2)
}
