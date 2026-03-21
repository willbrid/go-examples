package main

import (
	"os"
	"path/filepath"
)

/**
Si nous souhaitons lire et écrire des fichiers situés ailleurs, nous devons spécifier leurs chemins d'accès. Le problème est que tous les systèmes
d'exploitation prenant en charge Go ne gèrent pas les chemins de fichiers de la même manière.
Le package `os` fournit un ensemble de fonctions qui renvoient les chemins d'accès aux emplacements courants :
- Getwd() : Cette fonction renvoie le répertoire de travail courant, sous forme de chaîne de caractères, ainsi qu'une erreur indiquant un problème
lors de la récupération de cette valeur.
- UserHomeDir() : Cette fonction renvoie le répertoire personnel de l'utilisateur et une erreur indiquant un problème lors de la récupération du chemin.
- UserCacheDir() : Cette fonction renvoie le répertoire par défaut des données mises en cache spécifiques à l'utilisateur et une erreur
indiquant un problème lors de la récupération du chemin.
- UserConfigDir() : Cette fonction renvoie le répertoire par défaut des données de configuration spécifiques à l'utilisateur et une erreur
indiquant un problème lors de la récupération du chemin.
- TempDir() : Cette fonction renvoie le répertoire par défaut des fichiers temporaires et une erreur indiquant un problème lors de
la récupération du chemin.

Une fois que nous avons obtenu un chemin d'accès, nous pouvons le traiter comme une chaîne de caractères et simplement y ajouter des segments
supplémentaires ou, pour éviter les erreurs, utiliser les fonctions fournies par le package `path/filepath` pour manipuler les chemins d'accès.

Fonctions du package `path/filepath` pour les chemins d'accès :

- Abs(chemin) : Cette fonction renvoie un chemin absolu, utile si nous avons un chemin relatif, comme un nom de fichier.

- IsAbs(chemin) : Cette fonction renvoie vrai si le chemin spécifié est absolu.

- Base(chemin) : Cette fonction renvoie le dernier élément du chemin.

- Clean(chemin) : Cette fonction nettoie les chaînes de chemin en supprimant les séparateurs dupliqués et les références relatives.

- Dir(chemin) : Cette fonction renvoie tous les éléments du chemin sauf le dernier.

- EvalSymlinks(chemin) : Cette fonction évalue un lien symbolique et renvoie le chemin résultant.

- Ext(chemin) : Cette fonction renvoie l'extension du fichier à partir du chemin spécifié, supposée être le suffixe suivant le dernier point
dans la chaîne de chemin.

- FromSlash(chemin) : Cette fonction remplace chaque barre oblique par le séparateur de fichiers de la plateforme.

- ToSlash(chemin) : Cette fonction remplace le séparateur de fichiers de la plateforme par des barres obliques.

- Join(...éléments) : Cette fonction combine plusieurs éléments en utilisant le séparateur de fichiers de la plateforme.

- Match(modèle, chemin) : Cette fonction renvoie vrai si le chemin correspond au modèle spécifié.

- Split(chemin) : Cette fonction renvoie les composants de part et d'autre du dernier séparateur de chemin dans le chemin spécifié.

- SplitList(chemin) : Cette fonction divise un chemin en ses composants, qui sont renvoyés sous forme de tranche de chaîne.

- VolumeName(chemin) : Cette fonction renvoie le composant volume du chemin spécifié ou une chaîne vide si le chemin ne contient pas de volume.
**/

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "Myapp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))
}
