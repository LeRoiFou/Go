// Lien : https://www.youtube.com/watch?v=uACH9abGVlA
// Cours : Maps With Key Value Pairs - Learn Golang #8

// Date : 12-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
)

func main() {

	// map : map[type de clé]type de valeur{...}
	var pizza = map[string]string{
		"John": "Pepperoni",
		"Mary": "Cheese",
		"Tim":  "Mushroom"}

	fmt.Println(pizza)
	fmt.Println(pizza["John"])

	// map : raccourci
	// toppings := map[string]string{
	// 	"John": "Pepperoni"}

	// Changer une valeur du map
	pizza["John"] = "Peppers"
	fmt.Println(pizza["John"])

	// Supprimer une valeur du map
	delete(pizza, "Tim")
	fmt.Println(pizza)

	// Boucle avec les clés et les valeurs
	for key, value := range pizza {
		fmt.Println(key, value)
	}
}
