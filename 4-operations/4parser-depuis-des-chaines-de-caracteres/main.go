package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	Formatage depuis une chaine de caractères

	ParseBool(str) : cette fonction convertit une chaîne de caractères en une valeur booléenne.
	    Les valeurs booléennes reconnues sont : « true », « false », « TRUE », « FALSE », « True », « False », « T », « F », « 0 » et « 1 ».
	ParseFloat(str, size) : cette fonction analyse une chaîne de caractères et la convertit en une valeur à virgule flottante de la taille spécifiée.
	ParseInt(str, base, size) : cette fonction convertit une chaîne de caractères en un entier 64 bits (int64) avec la base et
	    la taille spécifiées. Les valeurs de base acceptables sont : 2 pour le binaire, 8 pour l’octal, 16 pour
		l’hexadécimal et 10 pour le hexadécimal.
	ParseUint(str, base, size) : cette fonction analyse une chaîne de caractères et la convertit en une valeur entière non signée avec
	    la base et la taille spécifiées.
	Atoi(str) : cette fonction convertit une chaîne de caractères en un entier décimal (base 10) et est
	    équivalente à l'appel de ParseInt(str, 10, 0).

	0b : ce préfixe désigne une valeur binaire, par exemple 0b1100100.

	0o : ce préfixe désigne une valeur octale, par exemple 0o144.

	0x : ce préfixe désigne une valeur hexadécimale, par exemple 0x64.
	**/
	var val1 string = "true"
	var val2 string = "false"
	var val3 string = "not true"
	var val4 string = "T"
	var bool1, b1Err = strconv.ParseBool(val1)
	var bool2, b2Err = strconv.ParseBool(val2)
	var bool3, b3Err = strconv.ParseBool(val3)
	fmt.Println("Bool 1 : ", bool1, b1Err)
	fmt.Println("Bool 2 : ", bool2, b2Err)
	fmt.Println("Bool 3 : ", bool3, b3Err)
	if bool4, b4Err := strconv.ParseBool(val4); b4Err == nil {
		fmt.Println("Parsed value : ", bool4)
	} else {
		fmt.Println("Cannot parse", val4)
	}

	var val5 string = "100"
	int1, int1err := strconv.ParseInt(val5, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value: ", int1)
	} else {
		fmt.Println("Cannot parse : ", val5)
	}

	var val6 string = "100"
	var int2 int64
	var int2err error
	int2, int2err = strconv.ParseInt(val6, 10, 0)
	if int2err == nil {
		var intResult int = int(int2)
		fmt.Println("Parsed value : ", intResult)
	} else {
		fmt.Println("Cannot parse : ", val6, int2err)
	}

	var val7 string = "0b1100100"
	int3, int3err := strconv.ParseInt(val7, 0, 8)
	if int3err == nil {
		smallInt := int8(int3)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int3err)
	}

	var val8 string = "100"
	int4, int4err := strconv.Atoi(val8)
	if int4err == nil {
		var intResult int = int(int4)
		fmt.Println("Parsed value : ", intResult)
	} else {
		fmt.Println("Cannot parse : ", val7, int4err)
	}

	var val9 string = "48.95"
	float1, float1err := strconv.ParseFloat(val8, 64)
	if float1err == nil {
		fmt.Println("Parsed value : ", float1)
	} else {
		fmt.Println("Cannot parse : ", val9, float1err)
	}
}
