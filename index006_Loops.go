// Lien : https://www.youtube.com/watch?v=_WCE_VA-IEM
// Cours #6 : For Loops - Learn Golang #6

// Date : 07-04-22
// Éditeur : Laurent REYNAUD

package main

import (
	"fmt"
)

func main() {

	// Boucle : de 0 à 4, incrémenter de 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Boucle d'un array avec un index de 0 à 2 par nom
	var names = [3]string{"John", "April", "Wes"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	// Boucle d'un array avec uniquement affichage des noms
	var names2 = [3]string{"John", "April", "Wes"}
	for _, name := range names2 {
		fmt.Println(name)
	}

}
