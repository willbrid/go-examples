package main

import (
	"bufio"
	"io"
	"strings"
)

/**
Le paquet bufio permet d'ajouter des tampons aux readers et aux writers.

L'exemple ci-dessous utilise le nouveau type `CustomReader` comme enveloppe autour d'un reader de type chaîne de caractères.
C'est la taille du slice d'octets passée à la fonction Read qui détermine la manière dont les données sont lues. Ici, la taille du slice est de
cinq octets, ce qui signifie qu'un maximum de cinq octets est lu à chaque appel de la fonction Read. Deux lectures n'ont pas permis d'obtenir
cinq octets de données. L'avant-dernière lecture a produit trois octets car les données sources ne sont pas divisibles par cinq et il restait
trois octets. La dernière lecture a renvoyé zéro octet et a reçu l'erreur EOF, indiquant que la fin des données avait été atteinte.

La lecture de petites quantités de données peut s'avérer problématique lorsque chaque opération engendre une surcharge importante. Ce problème
ne se pose pas lors de la lecture d'une chaîne de caractères stockée en mémoire, mais la lecture de données provenant d'autres sources, comme
des fichiers, peut être plus coûteuse. Dans ce cas, il est souvent préférable d'effectuer un plus grand nombre de lectures, mais plus volumineuses.
Pour ce faire, on utilise un tampon dans lequel une grande quantité de données est lue afin de répondre à plusieurs requêtes de données plus petites.

Les fonctions du package `bufio` pour la création de readers tamponnés.
- NewReader(r) : Cette fonction renvoie un reader tamponné avec la taille de tampon par défaut (4 096 octets au moment de l’écriture).
- NewReaderSize(r, size) : Cette fonction renvoie un reader tamponné avec la taille de tampon spécifiée.

Les résultats produits par `NewReader` et `NewReaderSize` implémentent l'interface `Reader` mais introduisent un tampon, ce qui peut réduire
le nombre d'opérations de lecture effectuées sur la source de données sous-jacente.

Dans l'exemple ci-dessous, l'on utilise la fonction `bufio.NewReader`, qui crée un reader avec la taille de tampon par défaut. Ce reader remplit
son tampon et utilise les données qu'il contient pour répondre aux appels à la méthode `Read`.
La taille de tampon par défaut est de 4 096 octets, ce qui signifie que le reader a pu lire toutes les données en une seule opération de lecture,
plus une lecture supplémentaire pour obtenir le résultat EOF. L'utilisation du tampon réduit la surcharge associée aux opérations de lecture,
au prix toutefois de la mémoire utilisée pour mettre les données en mémoire tampon.


Utilisation des méthodes supplémentaires de lecture tamponnée
Les fonctions NewReader et NewReaderSize renvoient des valeurs bufio.Reader, qui implémentent l'interface io.Reader.
Ces méthodes peuvent servir de wrappers pour d'autres types de méthodes Reader, en introduisant facilement un tampon de lecture.
La structure bufio.Reader définit des méthodes supplémentaires qui utilisent directement le tampon.
- Buffered() : Cette méthode renvoie un entier indiquant le nombre d'octets pouvant être lus depuis le tampon.
- Discard(count) : Cette méthode supprime le nombre d'octets spécifié.
- Peek(count) : Cette méthode renvoie le nombre d'octets spécifié sans les supprimer du tampon ; ils seront donc renvoyés par les
appels suivants à la méthode Read.
- Reset(reader) : Cette méthode supprime les données du tampon et effectue les lectures suivantes à partir du reader spécifié.
- Size() : Cette méthode renvoie la taille du tampon, exprimée en entier.


Écritures mises en mémoire tampon
Le package bufio prend également en charge la création de writers utilisant une mémoire tampon, grâce aux fonctions suivantes :
- NewWriter(w) : Cette fonction renvoie une valeur Writer avec tampon et une taille de tampon par défaut (4 096 octets au moment de l’écriture).
- NewWriterSize(w, size) : Cette fonction renvoie une valeur Writer avec tampon et une taille de tampon spécifiée.
Les résultats produits par ces fonctions implémentent l'interface Writer, ce qui permet d'introduire facilement un tampon d'écriture.

Le type de données spécifique renvoyé par ces fonctions est `bufio.Writer`, qui définit les méthodes ci-dessous de gestion du tampon et de son contenu :
- Available() : Cette méthode renvoie le nombre d'octets disponibles dans le tampon.
- Buffered() : Cette méthode renvoie le nombre d'octets écrits dans le tampon.
- Flush() : Cette méthode écrit le contenu du tampon dans l'objet Writer sous-jacent.
- Reset(writer) : Cette méthode efface les données du tampon et effectue les écritures suivantes dans l'objet Writer spécifié.
- Size() : Cette méthode renvoie la capacité du tampon en octets.
**/

func main() {
	Printfln("Opération de lecture sans tampon.......")
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("#0 Read data: %v", writer.String())

	Printfln("Opération de lecture avec tampon.......")
	text1 := "It was a boat. A small boat."
	var reader1 io.Reader = NewCustomReader(strings.NewReader(text1))
	var writer1 strings.Builder
	reader1 = bufio.NewReader(reader1)
	slice1 := make([]byte, 5)
	for {
		count1, err1 := reader1.Read(slice1)
		if count1 > 0 {
			writer1.Write(slice1[0:count1])
		}
		if err1 != nil {
			break
		}
	}
	Printfln("#1 Read data: %v", writer1.String())

	Printfln("Opération de lecture avec tampon et utilisation des méthodes Size et Buffered de bufio.Reader.......")
	text2 := "It was a boat. A small boat."
	var reader2 io.Reader = NewCustomReader(strings.NewReader(text2))
	var writer2 strings.Builder
	bufferedReader := bufio.NewReader(reader2)
	slice2 := make([]byte, 5)
	for {
		count2, err2 := bufferedReader.Read(slice2)
		if count2 > 0 {
			Printfln("#2 Buffer size: %v, buffered: %v", bufferedReader.Size(), bufferedReader.Buffered())
			writer2.Write(slice2[0:count2])
		}
		if err2 != nil {
			break
		}
	}
	Printfln("#2 Read data: %v", writer2.String())

	Printfln("Opération d'écriture sans tampon.......")
	// Cet exemple écrit cinq octets à la fois dans le Writer, qui est pris en charge par un Builder du package strings.
	text3 := "It was a boat. A small boat."
	var builder3 strings.Builder
	var writer3 = NewCustomWriter(&builder3)
	for i := 0; true; {
		end := i + 5
		if end >= len(text3) {
			writer3.Write([]byte(text3[i:]))
			break
		}
		writer3.Write([]byte(text3[i:end]))
		i = end
	}
	Printfln("#3 Written data: %v", builder3.String())

	Printfln("Opération d'écriture avec tampon.......")
	/**
	Le Writer mis en mémoire tampon conserve les données dans une mémoire tampon et les transmet au Writer sous-jacent uniquement lorsque
	la mémoire tampon est pleine ou lorsque la méthode Flush est appelée.
	La transition vers un Writer avec tampon n'est pas totalement transparente, car il est important d'appeler la
	méthode Flush de `bufio.Writer` pour garantir l'écriture de toutes les données.
	**/
	text4 := "It was a boat. A small boat."
	var builder4 strings.Builder
	var writer4 = bufio.NewWriterSize(NewCustomWriter(&builder4), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text4) {
			writer4.Write([]byte(text4[i:]))
			writer4.Flush()
			break
		}
		writer4.Write([]byte(text4[i:end]))
		i = end
	}
	Printfln("#4 Written data: %v", builder4.String())
}
