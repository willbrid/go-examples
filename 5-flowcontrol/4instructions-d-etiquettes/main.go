package main

import "fmt"

func main() {
	// Target : les instructions d'étiquette permettent à l'exécution de passer à un point différent.
	// Les étiquettes sont définies par un nom, suivi de deux points, puis d'une instruction de code standard.
	// Lorsque l'exécution atteint le mot-clé goto, elle passe à l'instruction portant l'étiquette spécifiée, puis
	// continue vers les instructions suivantes.
	var counter7 int = 0
target:
	fmt.Println("Counter : ", counter7)
	counter7++
	if counter7 < 5 {
		goto target
	}
}
