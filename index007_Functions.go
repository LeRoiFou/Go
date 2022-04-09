// Lien : https://www.youtube.com/watch?v=FQqNKpFomro
// Cours : Functions - Learn Golang #7

// Date : 09-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
)

func myFunction(firstName string, age int) {
	fmt.Println("C'est ma fonction !")
	fmt.Println("Salut", firstName, "! Tu as", age, "ans")
}

func main() {
	myFunction("Laurent", 45)
}
