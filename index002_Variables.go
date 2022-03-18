// Lien : https://www.youtube.com/watch?v=j0dThdwftmQ
// Variables and Constants In Go
// Éditeur : Laurent REYNAUD
// Date : 18-03-22

package main // package principal
import "fmt"

// Importation d'autres packages : fmt -> affichage dans la console

func main() {

	// Types de variables : string, int, float32, bool

	// Déclaration d'une str
	// var firstName string = "Laurent"
	// fmt.Println(firstName)

	// Autre moyen plus rapide pour déclarer des variables (sans var + type)
	// lastName := "Reynaud"
	// age := 45

	// Dans le cas où on n'assigne pas de valeur mais on assigne un type
	// var fullName string // Valeur vide
	// var age int         // Valeur = 0
	// var price float32   // Valeur = 0
	// var tf bool         // Valeur = False
	// fmt.Println(fullName, age, price, tf)

	// Déclaration de plusieurs variables
	var fullName = "Laurent Reynaud"
	var age = 45
	var price = 19.99
	var tf = true
	fmt.Println(fullName, age, price, tf)

	// Déclaration de plusieurs variables de même type
	var name1, name2 string = "Time", "Mary"
	fmt.Println(name1, name2)

	// Constantes
	const PIZZA string = "Pepperoni"
	fmt.Println(PIZZA)

	// Multiples constantes
	const (
		PIZZA1 = "Peperoni"
		PIZZA2 = "Fromage"
		MYNUM  = 77
	)
	fmt.Println(PIZZA2)
}
