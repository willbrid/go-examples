package main

import "io"

/**
Utilisation des readers et writers spécialisés
Outre les interfaces de lecture et d'écriture de base, le package `io` fournit des implémentations spécialisées.

Pipe() : Cette fonction renvoie un PipeReader et un PipeWriter, permettant de connecter des fonctions nécessitant un Reader et un Writer.

MultiReader(...readers) : Cette fonction définit un paramètre variadique autorisant la spécification d'un nombre quelconque de valeurs de Reader.
Le résultat est un Reader qui transmet le contenu de chacun de ses paramètres dans l'ordre de leur définition.

MultiWriter(...writers) : Cette fonction définit un paramètre variadique autorisant la spécification d'un nombre quelconque de valeurs de Writer.
Le résultat est un Writer qui envoie les mêmes données à tous les Writers spécifiés.

LimitReader(r, limit) : Cette fonction crée un Reader qui effectuera une fin de fichier (EOF) après le nombre d'octets spécifié.


Les pipes servent à connecter le code qui consomme des données via un reader et le code qui produit du code via un writer.
La fonction `io.Pipe` renvoie un `io.PipeReader` et un `io.PipeWriter`. Les structures `io.PipeReader` et `io.PipeWriter` implémentent l'interface `Closer`.
Close() : Cette méthode ferme le reader ou le writer. Le comportement exact dépend de l'implémentation, mais en général, toute lecture ultérieure
d'un reader fermé renverra zéro octet et l'erreur EOF, tandis que toute écriture ultérieure d'un writer fermé renverra une erreur.

Puisque `PipeWriter` implémente l'interface `Writer`, l'on peux l'utiliser comme argument de la fonction `GenerateData`, puis appeler la
méthode `Close` une fois la fonction terminée afin que le lecteur reçoive `EOF`.

Les pipes sont synchrones, de sorte que la méthode `PipeWriter.Write` se bloque jusqu'à ce que les données soient lues depuis le pipe.
Cela signifie que `PipeWriter` doit être utilisé dans une goroutine différente de celle du reader afin d'éviter un blocage de l'application.

Notez les parenthèses à la fin de cette instruction. Elles sont nécessaires lors de la création d'une goroutine pour une fonction anonyme,
mais il est facile de les oublier.

La structure `PipeReader` implémente l'interface `Reader`, ce qui signifie que l'on peut l'utiliser comme argument de la fonction `ConsumeData`.
La fonction `ConsumeData` est exécutée dans la goroutine principale, ce qui signifie que l'application ne se terminera pas tant que la fonction
n'aura pas terminé son exécution.

Ainsi, les données sont écrites dans le pipe à l'aide de `PipeWriter` et lues depuis ce pipe à l'aide de `PipeReader`. Lorsque la
fonction `GenerateData` a terminé son exécution, la méthode `Close` est appelée sur `PipeWriter`, ce qui provoque la fin du fichier (EOF)
lors de la prochaine lecture par `PipeReader`.

Le résultat met en évidence la synchronisation des flux de données. La fonction GenerateData appelle la méthode Write du processus d'écriture,
puis se bloque jusqu'à la lecture des données. C'est pourquoi le premier message affiché provient du processus de lecture : ce dernier traite
les données par octets, ce qui implique deux opérations de lecture avant que l'appel initial à la méthode Write, utilisée pour envoyer quatre
octets, ne soit terminé et que le message de la fonction GenerateData ne s'affiche.
**/

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)

	// Amélioration
	pipeReader1, pipeWriter1 := io.Pipe()
	go GenerateDataImprove(pipeWriter1)
	ConsumeData(pipeReader1)
}
