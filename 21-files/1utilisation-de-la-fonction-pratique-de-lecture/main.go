package main

/**
Le package clé lors du traitement des fichiers est le package `os`. Ce package donne accès aux fonctionnalités du système d'exploitation,
y compris le système de fichiers, de manière à masquer la plupart des détails d'implémentation, ce qui signifie que les mêmes fonctions peuvent
être utilisées pour obtenir les mêmes résultats quel que soit le système d'exploitation utilisé.
Les fonctionnalités fournies par le package `os` sont solides et fiables et permettent d’écrire du code Go utilisable sur différentes
plateformes sans modification.
Les fonctions fournies par le package `os` pour la lecture de fichiers :
- ReadFile(name) : Cette fonction ouvre le fichier spécifié et lit son contenu. Les résultats sont une slice d'octets contenant le contenu du fichier
et une erreur indiquant des problèmes d'ouverture ou de lecture du fichier. Elle offre un moyen pratique de lire le contenu complet d'un fichier
dans une slice d'octets en une seule étape.
- Open(name) : Cette fonction ouvre le fichier spécifié en lecture. Le résultat est une structure `File` et une erreur indiquant
des problèmes d’ouverture du fichier.

L'une des raisons les plus courantes de lire un fichier est de charger des données de configuration.
**/

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
}
