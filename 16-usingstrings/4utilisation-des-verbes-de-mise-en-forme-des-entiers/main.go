package main

import (
	"fmt"
)

/**
Les verbes de formatage pour les valeurs entières, quelle que soit leur taille.

%b : Ce verbe affiche une valeur entière sous forme de chaîne binaire.

%d : Ce verbe affiche une valeur entière sous forme de chaîne décimale. Il s'agit du format par défaut pour les valeurs entières,
appliqué lorsque le verbe %v est utilisé.

%o, %O : Ces verbes affichent une valeur entière sous forme de chaîne octale. Le verbe %O ajoute le préfixe 0o.

%x, %X : Ces verbes affichent une valeur entière sous forme de chaîne hexadécimale. Les lettres A à F sont affichées en minuscules par
le verbe %x et en majuscules par le verbe %X.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	number := 50
	Printfln("Binary : %b", number)
	Printfln("Decimal : %d", number)
	Printfln("Octal : %o, %O", number, number)
	Printfln("Hexadecimal : %x, %X", number, number)
	Printfln("Decimal : %d", int(products[0].Price))
}
