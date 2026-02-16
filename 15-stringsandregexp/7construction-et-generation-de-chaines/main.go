package main

import (
	"fmt"
	"strings"
)

/**
Le package `strings` fournit deux fonctions pour générer des chaînes de caractères et un type struct dont les méthodes peuvent être utilisées
pour construire efficacement des chaînes de caractères progressivement.

Join(slice, sep) : Cette fonction combine les éléments de la tranche de chaînes spécifiée, en insérant la chaîne de séparation spécifiée
entre les éléments.

Repeat(s, count) : Cette fonction génère une chaîne en répétant la chaîne s un nombre de fois spécifié.


Le package strings fournit le type `Builder`, qui n'exporte pas de champs mais propose un ensemble de méthodes permettant de construire
efficacement des chaînes de caractères de manière progressive.

WriteString(s) : Cette méthode ajoute la chaîne « s » à la chaîne en cours de construction.

WriteRune(r) : Cette méthode ajoute le caractère « r » à la chaîne en cours de construction.

WriteByte(b) : Cette méthode ajoute l’octet « b » à la chaîne en cours de construction.

String() : Cette méthode renvoie la chaîne créée par le générateur.

Reset() : Cette méthode réinitialise la chaîne créée par le générateur.

Len() : Cette méthode renvoie le nombre d’octets utilisés pour stocker la chaîne créée par le générateur.

Cap() : Cette méthode renvoie le nombre d’octets alloués par le générateur.

Grow(size) : Cette méthode augmente le nombre d’octets alloués par le générateur pour stocker la chaîne en cours de construction.
**/

func main() {
	text := "It was a boat. A small boat."
	// Cette instruction retourn un slice basé sur le séparateur d'espace
	elements := strings.Fields(text)
	// Cette fonction combine les éléments dans le slice de chaîne spécifiée, avec la chaîne de séparation spécifiée placée entre les éléments.
	joinResult := strings.Join(elements, "--")
	fmt.Println("Join :", joinResult)
	// Cette fonction repète la chaine "good" 3 fois et retourne une chaine concatenée de ces 3 occurrences de la chaine "good"
	repeatResult := strings.Repeat("good", 3)
	fmt.Println("Repeat :", repeatResult)

	text1 := "It was a boat. A small boat."
	var builder strings.Builder
	for _, sub := range strings.Fields(text1) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		// Cette méthode builder.WriteString ajoute la chaîne sub à la chaîne en cours de construction.
		builder.WriteString(sub)
		// Cette méthode builder.WriteString ajoute le caractère ' ' à la chaîne en cours de construction.
		builder.WriteRune(' ')
	}
	fmt.Println("Builder accumulated string result :", builder.String())
	// Cette méthode renvoie le nombre d'octets utilisés pour stocker la chaîne créée par le générateur.
	fmt.Println("Builder String Len :", builder.Len())
	// Cette méthode renvoie le nombre d'octets qui ont été alloués par le générateur.
	fmt.Println("Builder String Cap :", builder.Cap())
	// Cette méthode réinitialise la chaîne créée par le générateur.
	builder.Reset()
	fmt.Println("Builder accumulated string result :", builder.String())

	/**
	Créer la chaîne à l'aide du `Builder` est plus efficace que d'utiliser l'opérateur de concaténation sur des valeurs de chaîne classiques,
	surtout si la méthode `Grow` est utilisée pour allouer de l'espace de stockage à l'avance.

	Il convient d'utiliser des pointeurs lors du passage de valeurs Build`er entre les fonctions et les méthodes ; sinon,
	les gains d'efficacité seront perdus lors de la copie du Builder.
	**/
}
