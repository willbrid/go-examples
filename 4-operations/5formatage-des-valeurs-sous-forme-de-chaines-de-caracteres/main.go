package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
			Formatage d'une valeur en chaine de caractères

		FormatBool(val) : cette fonction renvoie la chaîne de caractères « true » ou « false » selon la valeur du booléen spécifié.
		FormatInt(val, base) : cette fonction renvoie une représentation sous forme de chaîne de caractères de la valeur int64 spécifiée, exprimée dans la base spécifiée.

		FormatUint(val, base) : Cette fonction renvoie une représentation sous forme de chaîne de caractères de la valeur uint64 spécifiée, exprimée dans la base spécifiée.

		FormatFloat(val, format, précision, taille) : Cette fonction renvoie une représentation sous forme de chaîne de caractères de la valeur float64 spécifiée,
		    exprimée selon le format, la précision et la taille spécifiés.
			Les formats possibles sont :
			- f : La valeur à virgule flottante sera exprimée sous la forme ±ddd.ddd sans exposant, par exemple 49,95.
			- e, E : La valeur à virgule flottante sera exprimée sous la forme ±ddd.ddde±dd, par exemple 4,995e+01 ou 4,995E+01. La casse de la lettre indiquant l'exposant est déterminée par la casse de la rune
	utilisée comme argument de formatage.
			- g, G La valeur à virgule flottante sera exprimée au format e/E pour les grands exposants ou au format f pour les petits exposants.

		Itoa(val) : Cette fonction renvoie une représentation sous forme de chaîne de caractères de la valeur int spécifiée, exprimée en base 10.
			**/
	val9 := true
	val10 := false
	str1 := strconv.FormatBool(val9)
	str2 := strconv.FormatBool(val10)
	fmt.Println("Formatted value 1 : " + str1)
	fmt.Println("Formatted value 2 : " + str2)

	val11 := 275
	base10String1 := strconv.FormatInt(int64(val11), 10)
	base10String2 := strconv.Itoa(val11)
	base2String := strconv.FormatInt(int64(val11), 2)
	fmt.Println("Base 10 : " + base10String1)
	fmt.Println("Base 10 : " + base10String2)
	fmt.Println("Base 2 : " + base2String)

	val12 := 49.95
	Fstring := strconv.FormatFloat(val12, 'f', 2, 64)
	Estring := strconv.FormatFloat(val12, 'e', -1, 64)
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
}
