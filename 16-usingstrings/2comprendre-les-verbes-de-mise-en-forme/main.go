package main

import "fmt"

/**
D'autres fonctions du package fmt permettent de contrôler la mise en forme.

`Sprintf(t, ...vals)` : Cette fonction renvoie une chaîne de caractères créée à partir du modèle `t`. Les arguments restants servent de valeurs
aux verbes du modèle.

`Printf(t, ...vals)` : Cette fonction crée une chaîne de caractères à partir du modèle `t`. Les arguments restants servent de valeurs aux
verbes du modèle. La chaîne est affichée sur la sortie standard.

`Fprintf(writer, t, ...vals)` : Cette fonction crée une chaîne de caractères à partir du modèle `t`. Les arguments restants servent de valeurs
aux verbes du modèle. La chaîne est écrite dans un objet `Writer`.

`Errorf(t, ...values)` : Cette fonction génère une erreur à partir du modèle `t`. Les arguments restants servent de valeurs aux verbes du modèle.
Le résultat est une valeur d'erreur dont la méthode `Error` renvoie la chaîne formatée.
**/

func getProductName(index int) (name string, err error) {
	if len(products) > index {
		name = fmt.Sprintf("Name of product: %v", products[index].Name)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}

	return
}

func main() {
	/**
	La fonction `Printf` accepte une chaîne de caractères modèle et une série de valeurs. La chaîne modèle est analysée à la recherche de verbes,
	indiqués par le signe pourcentage (%) suivi d'un spécificateur de format.

	- Le premier verbe est %v et spécifie la représentation par défaut d'un type. Pour une chaîne de caractères, par exemple, %v inclut simplement
	la chaîne dans la sortie.
	- Le verbe %4.2f spécifie le format d'une valeur à virgule flottante, avec 4 chiffres avant la virgule et 2 chiffres après.

	Les valeurs des verbes du modèle sont extraites des arguments restants, utilisés dans l'ordre de leur spécification.
	Dans cet exemple, le verbe %v est utilisé pour formater la valeur `Product.Name` et le verbe %4.2f pour formater la valeur `Product.Price`.
	Ces valeurs sont formatées, insérées dans la chaîne du modèle et affichées dans la console.
	**/
	fmt.Printf("Product: %v, Price: $%4.2f", kayak.Name, kayak.Price)
	fmt.Printf("\n")
	name, _ := getProductName(1)
	fmt.Println("Name : ", name)
	_, err := getProductName(10)
	fmt.Println("Error :", err.Error())
}
