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


La fonction `MultiWriter` combine plusieurs writers afin que les données soient envoyées à chacun d'eux.
Dans cet exemple, les objets Writer sont des valeurs de type `string.Builder` implémentant l'interface `Writer`.
La fonction `MultiWriter` permet de créer un objet `Writer`, de sorte que l'appel à la méthode `Write` entraîne l'écriture des mêmes données
dans les trois objets `Writer`.
**/

func main() {
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}
