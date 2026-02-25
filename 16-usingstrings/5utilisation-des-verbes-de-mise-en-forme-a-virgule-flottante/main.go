package main

import "fmt"

/**
Les verbes de formatage pour les valeurs à virgule flottante, qui peuvent être appliqués aux valeurs float32 et float64.

%b : Ce verbe affiche une valeur à virgule flottante avec un exposant et sans décimale.

%e, %E : Ces verbes affichent une valeur à virgule flottante avec un exposant et une décimale. %e utilise une minuscule pour l'exposant,
tandis que %E utilise une majuscule.

%f, %F : Ces verbes affichent une valeur à virgule flottante avec une décimale mais sans exposant. Les verbes %f et %F produisent le même résultat.

%g : Ce verbe s'adapte à la valeur affichée. Le format %e est utilisé pour les valeurs avec de grands exposants, et le format %f dans les autres cas.
Il s'agit du format par défaut, appliqué lorsque le verbe %v est utilisé.

%G : Ce verbe s'adapte à la valeur affichée. Le format %E est utilisé pour les valeurs avec de grands exposants, et le format %f dans les autres cas.

%x, %X : Ces verbes affichent une valeur à virgule flottante en notation hexadécimale, avec des lettres minuscules (%x) ou majuscules (%X).


Le résultat des verbes ci-dessus peut être modifié à l'aide des modificateurs.

+ : Ce modificateur (le signe plus) affiche toujours un signe : positif ou négatif, pour les valeurs numériques.

0 : Ce modificateur utilise des zéros, plutôt que des espaces, comme espacement lorsque la largeur est supérieure au nombre de
caractères nécessaires à l'affichage de la valeur.

- : Ce modificateur (le symbole moins) ajoute un espacement à droite du nombre, et non à gauche.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	number := 279.00
	Printfln("Decimalless with exponent : %b", number)
	Printfln("Decimal with exponent : %e, %E", number, number)
	Printfln("Decimal without exponent : %f", number)
	Printfln("Decimal : %g, %G", products[0].Price, products[0].Price)
	Printfln("Hexadecimal : %x, %X", number, number)

	/**
	Le format des valeurs à virgule flottante peut être contrôlé en modifiant le verbe pour spécifier la
	largeur (le nombre de caractères utilisés pour exprimer la valeur) et la précision (le nombre de chiffres après la virgule).
	Des espaces sont utilisés pour le remplissage lorsque la largeur spécifiée est supérieure au nombre de caractères nécessaires
	pour afficher la valeur.
	**/
	Printfln("Decimal without exponent : >>%8.2f<<", number)

	/**
	La largeur peut être omise si seule la précision vous intéresse.
	La valeur de la largeur est omise, mais le caractère point reste obligatoire pour mentionner la précision.
	**/
	Printfln("Decimal without exponent : >>%.2f<<", number)

	/**
	Utilisation des modificateurs pour modifier la mise en forme d'une valeur à virgule flottante.
	**/
	Printfln("Sign : >>%+.2f<<", number)
	Printfln("Zeros for Padding : >>%010.2f<<", number)
	Printfln("Right Padding : >>%-8.2f<<", number)
}
