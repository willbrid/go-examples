package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(products) > index {
		/**
		Cette fonction Sprintf renvoie une chaîne, qui est créée en traitant le modèle en premier argument.
		Les arguments restants sont utilisés comme valeurs pour les verbes modèles.
		**/
		name = fmt.Sprintf("Name of Product: %v", products[index].Name)
	} else {
		/**
		Cette fonction crée une erreur en traitant le modèle en premier argument. Les arguments restants sont utilisés comme valeurs pour les verbes modèles.
		Le résultat est une valeur d'erreur dont la méthode Error renvoie la chaîne formatée.
		**/
		err = fmt.Errorf("error for index %v", index)
	}

	return
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func (p Product) String() string {
	return fmt.Sprintf("Product : %v, Price : $%4.2f", p.Name, p.Price)
}

func main() {
	fmt.Println("Product : ", kayak.Name, " - Price : ", kayak.Price)
	fmt.Print("Product : ", kayak.Name, " - Price : ", kayak.Price, "\n")
	/**
	La fonction Printf accepte une chaîne de modèle et une série de valeurs. Le modèle est analysé à la recherche de verbes,
	qui sont indiqués par le signe de pourcentage (le caractère %) suivi d'un spécificateur de format.
	**/
	fmt.Printf("Product : %v, Price : $%4.2f \n", kayak.Name, kayak.Price)

	name, _ := getProductName(1)
	fmt.Println(name)
	_, err := getProductName(10)
	fmt.Println(err.Error())

	/**
	Ce verbe %v affiche le format par défaut de la valeur. La modification du verbe avec un signe plus (%+v) inclut les noms de champ
	lors de l'écriture des valeurs de structure.
	**/
	Printfln("Value: %v", kayak)
	Printfln("Value with fields: %+v", kayak)
	/**
	Ce verbe %#v affiche une valeur dans un format qui pourrait être utilisé pour recréer la valeur dans un fichier de code Go.
	**/
	Printfln("Go syntax: %#v", kayak)
	// Ce verbe affiche le type Go d'une valeur.
	Printfln("Type: %T", kayak)

	fmt.Println("Display Kayak product", kayak.String())

	numberInt := 250
	// Ce verbe %b affiche une valeur entière sous forme de chaîne binaire.
	Printfln("Binary: %b", numberInt)
	// Il (%d) s'agit du format par défaut pour les valeurs entières, appliqué lorsque le verbe %v est utilisé.
	Printfln("Decimal: %d", numberInt)
	// Ces (%o, %O) verbes affichent une valeur entière sous forme de chaîne octale. Le verbe %O ajoute le préfixe 0o.
	Printfln("Octal: %o, %O", numberInt, numberInt)
	/**
	Ces verbes %x et %X affichent une valeur entière sous forme de chaîne hexadécimale.
	Les lettres A–F sont affichées en minuscules par le verbe %x et en majuscules par le verbe %X.
	**/
	Printfln("Hexadecimal: %x, %X", numberInt, numberInt)

	numberFloat := 279.00
	// Ce verbe %b affiche une valeur à virgule flottante avec un exposant et sans décimale.
	Printfln("Decimalless with exponent: %b", numberFloat)
	/**
	Ces verbes %e et %E affichent une valeur à virgule flottante avec un exposant et une décimale. Le %e utilise un indicateur d'exposant
	en minuscule, tandis que %E utilise un indicateur en majuscule.
	**/
	Printfln("Decimal with exponent: %e", numberFloat)
	/**
	Ces verbes %f et %F affichent une valeur à virgule flottante avec une décimale mais pas d'exposant. Les verbes %f et %F produisent la même sortie.
	**/
	Printfln("Decimal without exponent: %f", numberFloat)
	// Ces verbes %x et %X affichent une valeur à virgule flottante en notation hexadécimale, avec des lettres minuscules (%x) ou majuscules (%X).
	Printfln("Hexadecimal: %x, %X", numberFloat, numberFloat)
	/**
	Le format des valeurs à virgule flottante peut être contrôlé en modifiant le verbe pour spécifier la largeur
	(le nombre de caractères utilisés pour exprimer la valeur) et la précision (le nombre de chiffres après la décimale)
	les espaces sont utilisés pour le remplissage lorsque la largeur spécifiée est supérieure au nombre de caractères requis pour afficher la valeur.
	**/
	Printfln("Decimal without exponent : >>%8.2f<<", numberFloat)
	Printfln("Decimal without exponent : >>%.2f<<", numberFloat)
	// Ce modificateur (le signe plus) affiche toujours un signe, positif ou négatif, pour les valeurs numériques.
	Printfln("Sign: >>%+.2f<<", numberFloat)
	/**
	Ce modificateur utilise des zéros plutôt que des espaces comme remplissage lorsque la largeur est supérieure au
	nombre de caractères requis pour afficher la valeur.
	**/
	Printfln("Zeros for Padding: >>%010.2f<<", numberFloat)
	// Ce modificateur (le symbole de soustraction) ajoute un remplissage à droite du nombre, plutôt qu'à gauche.
	Printfln("Right Padding: >>%-8.2f<<", numberFloat)

	name1 := "Kayak"
	// Ce verbe %s affiche une chaîne. C'est le format par défaut, appliqué lorsque le verbe %v est utilisé.
	Printfln("String: %s", name1)
	// Ce verbe %c affiche un caractère. Des précautions doivent être prises pour éviter de découper les chaînes en octets individuels
	Printfln("Character: %c", []rune(name1)[0])
	// Ce verbe %U affiche un caractère au format Unicode afin que la sortie commence par U+ suivi d'un code de caractère hexadécimal.
	Printfln("Unicode : %U", []rune(name1)[0])
	// Ce verbe %t formate les valeurs booléennes et affiche true ou false.
	Printfln("Bool1 : %t", len(name1) > 1)
	Printfln("Bool2 : %t", len(name1) > 100)
	// Ce verbe %p affiche une représentation hexadécimale de l'emplacement de stockage du pointeur.
	Printfln("Pointer : %p", &name1)
}
