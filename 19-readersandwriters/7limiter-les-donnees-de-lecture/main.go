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


La fonction `LimitReader` est utilisée pour limiter la quantité de données pouvant être obtenues à partir d'un reader.
Le premier argument de la fonction `LimitReader` est le reader qui fournira les données. Le second argument est le nombre maximal d'octets pouvant être lus.
Le reader renvoyé par la fonction `LimitReader` enverra un signal `EOF` lorsque la limite sera atteinte, sauf si le reader sous-jacent envoie
déjà un signal `EOF`.
**/

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
}
