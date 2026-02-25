package main

import "fmt"

/**
Le verbe décrit s'applique aux pointeurs.

%p : Ce verbe affiche une représentation hexadécimale de l’emplacement de stockage du pointeur.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	name := "Kayak"
	Printfln("Pointer : %p", &name)
}
