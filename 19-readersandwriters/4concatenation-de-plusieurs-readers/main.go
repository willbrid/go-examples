package main

import (
	"io"
	"strings"
)

/**
Utilisation des readers et writers spécialisés
Outre les interfaces de lecture et d'écriture de base, le package `io` fournit des implémentations spécialisées.

Pipe() : Cette fonction renvoie un PipeReader et un PipeWriter, permettant de connecter des fonctions nécessitant un Reader et un Writer.

MultiReader(...readers) : Cette fonction définit un paramètre variadique autorisant la spécification d'un nombre quelconque de valeurs de Reader.
Le résultat est un Reader qui transmet le contenu de chacun de ses paramètres dans l'ordre de leur définition.

MultiWriter(...writers) : Cette fonction définit un paramètre variadique autorisant la spécification d'un nombre quelconque de valeurs de Writer.
Le résultat est un Writer qui envoie les mêmes données à tous les Writers spécifiés.

LimitReader(r, limit) : Cette fonction crée un Reader qui effectuera une fin de fichier (EOF) après le nombre d'octets spécifié.


La fonction `MultiReader` concentre les données provenant de plusieurs readers afin qu'elles puissent être traitées séquentiellement.
Le reader renvoyé par la fonction `MultiReader` répond à la méthode `Read` avec le contenu de l'un des readers sous-jacents. Lorsque le
premier reader renvoie `EOF`, le contenu est lu depuis le deuxième reader.
Ce processus se poursuit jusqu'à ce que le dernier reader sous-jacent renvoie `EOF`.
**/

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
}
