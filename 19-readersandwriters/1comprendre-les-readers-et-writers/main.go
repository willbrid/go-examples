package main

import (
	"io"
	"strings"
)

/**
L'interface Reader définit une seule méthode :
Read(byteSlice) : Cette méthode lit des données dans le tableau d’octets spécifié []byte. Elle renvoie le nombre d’octets lus
(exprimé sous forme d’entier) et une erreur.

L'interface `Reader` ne précise pas la provenance des données ni leur mode d'obtention ; elle définit uniquement la méthode Read.
Ces détails sont laissés à la discrétion des types qui implémentent l'interface, et la bibliothèque standard propose des implémentations
de `Reader` pour différentes sources de données.

L'un des Readers les plus simples utilise une chaîne de caractères comme source de données.

Chaque type de `Reader` est créé différemment. Par exemple pour créer un Reader basé sur une chaîne de caractères, le package `strings` fournit
une fonction constructeur `NewReader`, qui accepte une chaîne de caractères comme argument.


L'interface `Writer` définit une seule méthode :
Write(byteSlice) : cette méthode écrit les données de la tranche d'octets spécifiée. Elle renvoie le nombre d'octets écrits et une erreur.
Cette erreur sera différente de `nil` si le nombre d'octets écrits est inférieur à la longueur de la tranche.

L'interface `Writer` ne fournit aucun détail sur la manière dont les données écrites sont stockées, transmises ou traitées ;
ces opérations sont laissées à la charge des types qui implémentent l'interface.


En règle générale, les méthodes `Reader` et `Writer` sont implémentées pour les pointeurs afin que le passage d'un Reader ou d'un Writer à une fonction
n'entraîne pas la création d'une copie. Pour l'exemple avec le Reader l'on n'a pas eu besoin d'utiliser l'opérateur d'adresse pour le Reader car
le résultat de la fonction `strings.NewReader` est un pointeur.
**/

/**
Pour illustrer l'utilisation de l'interface, l'on utilise le résultat de la fonction `NewReader` comme argument d'une fonction acceptant un
`io.Reader`. Dans cette fonction, l'on utilise la méthode `Read` pour lire des octets de données. L'on spécifie le nombre maximal d'octets à
recevoir en définissant la taille du segment d'octets passé à la fonction `Read`. Les résultats de la fonction `Read` indiquent
le nombre d'octets lus et la présence éventuelle d'erreurs.
Le package `io` définit une erreur spéciale nommée `EOF`, qui signale la fin des données lues par le `Reader`. Si l'erreur renvoyée par
la fonction `Read` est une erreur `EOF`, la boucle `for` qui lit les données depuis le `Reader` est interrompue.
Ainsi, la boucle `for` appelle la fonction `Read` pour lire au maximum deux octets à la fois et les écrit. Lorsque la fin de la chaîne est
atteinte, la fonction `Read` renvoie l'erreur `EOF`, ce qui entraîne l'arrêt de la boucle `for`.
**/

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

/**
La structure `strings.Builder` implémente l'interface `io.Writer`, ce qui signifie que l'on peut écrire des octets dans un `Builder`
puis appeler sa méthode `String` pour créer une chaîne à partir de ces octets. Les `Writer` renvoient une erreur s'ils ne parviennent pas
à écrire toutes les données de la tranche. L'on vérifie le résultat de l'erreur et interromps la boucle `for` si une erreur est renvoyée.
Cependant, comme le `Writer` de cet exemple construit une chaîne en mémoire, le risque d'erreur est minime.
Notons que l'on utilise l'opérateur d'adresse pour passer un pointeur vers le `Builder` à la fonction `processDataWithWriter` car la
La structure `strings.Builder` implémente l'interface `io.Writer` avec pour receveur `*strings.Builder` (un pointeur sur `strings.Builder`).
**/

func processDataWithWriter(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak")
	processData(r)

	r1 := strings.NewReader("Kayak")
	var builder strings.Builder
	processDataWithWriter(r1, &builder)
	Printfln("String builder contents: %s", builder.String())
}
