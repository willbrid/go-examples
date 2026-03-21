package main

import (
	"encoding/json"
	"os"
)

/**
Le système d'exploitation offre également quelques fonctions permettant de créer de nouveaux fichiers :
- `Create(name)` : Cette fonction est équivalente à l’appel de `OpenFile` avec les options `O_RDWR`, `O_CREATE` et `O_TRUNC`. Elle renvoie
le fichier, utilisable en lecture et en écriture, ainsi qu’une erreur signalant les problèmes rencontrés lors de sa création.
Notons que cette combinaison d’options implique que si un fichier existe déjà sous le nom spécifié, il sera ouvert et son contenu supprimé.

- `CreateTemp(dirName, fileName)` : Cette fonction crée un nouveau fichier dans le répertoire spécifié, sous le nom indiqué. Si le nom est une
chaîne vide, le répertoire temporaire du système, obtenu grâce à la fonction `TempDir`, est utilisé. Le fichier est créé avec un nom contenant
une séquence aléatoire de caractères. Le fichier est ouvert avec les options `O_RDWR`, `O_CREATE` et `O_EXCL`. Il n’est pas supprimé à sa fermeture.
Cette fonction peut s'avérer utile, mais il est important de comprendre que son but est de générer un nom de fichier aléatoire et que, pour le reste,
le fichier créé est un fichier ordinaire. Ce fichier n'est pas supprimé automatiquement et restera sur le périphérique de stockage après l'exécution
de l'application.
Dans l'exemple ci-dessous, l'emplacement du fichier temporaire est spécifié par "/tmp". Si une chaîne vide est utilisée, le fichier sera créé dans le répertoire temporaire par défaut,
obtenu via la variable d'environnement `TempDir`. Le nom du fichier peut contenir un astérisque (*). Dans ce cas, la partie aléatoire du nom de fichier
le remplacera. Si le nom de fichier ne contient pas d'astérisque, la partie aléatoire sera ajoutée à la fin.
**/

func main() {
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file, err := os.CreateTemp("/tmp", "tempfile-*.json")
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
