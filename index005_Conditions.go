// Lien : https://www.youtube.com/watch?v=_WCE_VA-IEM
// Cours #5 : Logic If Else Statements - Learn Golang #5

// Date : 07-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
)

func main() {

	var firstName string = "Mary"

	if firstName == "John" {
		fmt.Println("Hello", firstName)
	} else if firstName == "Tim" {
		fmt.Println("What's up", firstName, "!")
	} else {
		fmt.Println("Sorry", firstName, "I don't know you!")
	}

	myNum := 12
	if myNum > 30 {
		fmt.Println(myNum, "is greater than 30")
	} else {
		fmt.Println(myNum, "is less than 30")
	}

}
