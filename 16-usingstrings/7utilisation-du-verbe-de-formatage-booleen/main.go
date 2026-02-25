package main

import "fmt"

/**
Le verbe utilisé pour formater les valeurs booléennes. Il s'agit du format booléen par défaut, c'est-à-dire celui qui sera utilisé par le verbe %v.

%t : Ce verbe formate les valeurs booléennes et affiche `true` ou `false`.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	name := "Kayak"
	Printfln("Bool : %t", len(name) > 1)
	Printfln("Bool : %t", len(name) > 100)
	Printfln("Bool : %t", products[0].Price > 50.00)
}
