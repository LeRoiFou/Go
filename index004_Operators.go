// Lien : https://www.youtube.com/watch?v=VSCZmSqaVDQ
// Cours #4 : Operators In Go

// Date : 01-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Math OPerators")
	fmt.Println()

	// Operations arithmétiques
	fmt.Println("Opérations arithmétiques")
	fmt.Println("10 + 5 = ", 10+5)
	fmt.Println("10 - 5 = ", 10-5)
	fmt.Println("10 * 5 = ", 10*5)
	fmt.Println("10 / 5 = ", 10/5)
	fmt.Println("10 % 5 = ", 10%5)
	fmt.Println()

	// Recours au module math
	fmt.Println("Module math")
	var base, exponent float64
	base = 2
	exponent = 3
	fmt.Println("2 puissance 3 = ", math.Pow(base, exponent))
	fmt.Println()

	// Incrémentation
	fmt.Println("Incrémentations")
	num := 10
	num++
	fmt.Println("10 + 1 = ", num)
	num2 := 10
	num2--
	fmt.Println("10 - 1 = ", num2)
	fmt.Println()

	// Assignement des signes arithmétiques : =, +=, *=, /=
	fmt.Println("Signes arithmétiques")
	myNum := 10
	myNum += 5
	fmt.Println("10 + 5  = ", myNum)
	fmt.Println()

	// Assignement des signes de comparaison : ==, !=, <, >, <=, >+
	fmt.Println("Signes de comparaison")
	fmt.Println("5 == 5 : ", 5 == 5)
	fmt.Println("5 != 6 : ", 5 != 6)
	fmt.Println("5 < 6 : ", 5 < 6)
	fmt.Println("6 > 5 : ", 6 > 5)
	fmt.Println("5 <= 6 : ", 5 <= 6)
	fmt.Println("6 >= 5 : ", 6 >= 5)
	fmt.Println()
}
