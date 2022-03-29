// Lien : https://www.youtube.com/watch?v=KHrQYq6OLMk
// cours n° 3 : Arrays and Slices

// Création de listes (arrays)

// Date : 29-03-22
// Éditeur : Laurent REYNAUD

package main

import "fmt"

func main() {

	// Arrays (liste)
	var firstNames = [3]string{"John", "April", "Wes"}
	fmt.Println(firstNames[1]) // Composant n° 1

	// Arrays : nombre de composants non déterminés
	var numies = []int{1, 2, 3, 4, 5}
	fmt.Println(numies[0]) // Composant n° 0

	// Arrays : instruction courte ('var' non mentionnée)
	lastNames := [2]string{"Elder", "Smith"}
	fmt.Println(lastNames[0])

	// Arrays : instruction courte et nbre de composants non déterminés
	ourNames := []string{"Elder", "Browne"}
	fmt.Println(ourNames)

	// Arrays : changement d'une valeur d'un composant
	ourNames[1] = "Smith"
	fmt.Println(ourNames)

	// Arrays : valeur par défaut si le nombre de composant < nombre prévu
	var ourNummies = [5]int{1, 2}
	fmt.Println(ourNummies)

	// Arrays : position de certaines valeurs dans l'array
	var moreNummies = [10]int{0: 41, 4: 99}
	fmt.Println(moreNummies)

	// Slices : découpage d'un array
	var toppings = [5]string{"Pepperoni", "Oignons", "Fromage", "Suprême", "Salade"}
	fmt.Println(toppings)
	toppingSlice := toppings[0:2]
	fmt.Println(toppingSlice)

	// Slices : modification d'une valeur d'un composant
	toppingSlice[1] = "Poivre"
	fmt.Println(toppingSlice)

	// Slices : ajout d'un nouveau composant
	toppingSlice = append(toppingSlice, "Pomme")
	fmt.Println(toppingSlice)

	// Slices : remplacement de valeurs
	otherSlice := toppings[3:4]
	fmt.Println(otherSlice)

	// Slices : ajout de deux arrays
	otherSlice = append(otherSlice, toppingSlice...)
	fmt.Println(otherSlice)
}
