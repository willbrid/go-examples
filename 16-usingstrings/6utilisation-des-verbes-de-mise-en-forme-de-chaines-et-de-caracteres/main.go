package main

import "fmt"

/**
Les verbes de formatage pour les chaînes de caractères et les runes.

%s : Cette commande affiche une chaîne de caractères. Il s’agit du format par défaut, appliqué lorsque la commande %v est utilisée.

%c : Cette commande affiche un caractère. Il convient d’éviter de découper les chaînes en octets individuels.

%U : Cette commande affiche un caractère au format Unicode, de sorte que la sortie commence par `U+` suivi du code hexadécimal du caractère.


Les chaînes de caractères sont faciles à formater, mais il faut être vigilant lors du formatage de chaque caractère.
Certains caractères sont représentés par plusieurs octets ; il est donc impératif de ne pas formater seulement une partie de ces octets.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	name := "kayak"
	Printfln("String : %s", name)
	Printfln("Character : %c", []rune(name)[0])
	Printfln("Unicode : %U", []rune(name)[0])
}
