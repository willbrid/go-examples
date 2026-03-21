package main

import (
	"os"
	"path/filepath"
)

/**
Si notre projet repose sur le traitement de fichiers créés par un autre processus, nous devons alors explorer le système de fichiers.
Fonction du package `os` pour lister les répertoires
- ReadDir(name) : Cette fonction lit le répertoire spécifié et renvoie une tranche `DirEntry`, dont chaque élément décrit un élément du répertoire.
Le résultat de la fonction `ReadDir` est une tranche de valeurs qui implémentent l'interface `DirEntry`.

Les méthodes définies par l'interface `DirEntry` :
- Name() : Cette méthode renvoie le nom du fichier ou du répertoire décrit par la valeur `DirEntry`.
- IsDir() : Cette méthode renvoie vrai si la valeur `DirEntry` représente un répertoire.
- Type() : Cette méthode renvoie une valeur `FileMode` (alias de uint32) qui décrit plus en détail le fichier et ses permissions.
- Info() : Cette méthode renvoie une valeur `FileInfo` fournissant des informations supplémentaires sur le fichier ou le répertoire.

L'interface `FileInfo`, qui résulte de la méthode `Info`, permet d'obtenir des informations détaillées sur un fichier ou un répertoire.
Les méthodes les plus utiles définies par l'interface `FileInfo` sont les suivantes :
- Name() : Cette méthode renvoie une chaîne de caractères contenant le nom du fichier ou du répertoire.
- Size() : Cette méthode renvoie la taille du fichier, exprimée sous forme d'entier 64 bits.
- Mode() : Cette méthode renvoie le mode d'accès et les permissions du fichier ou du répertoire.
- ModTime() : Cette méthode renvoie la date et l'heure de dernière modification du fichier ou du répertoire.

Nous pouvons également obtenir une valeur `FileInfo` concernant un fichier unique à l'aide de la fonction :
- Stat(path) : Cette fonction accepte une chaîne de caractères représentant un chemin d'accès. Elle renvoie une valeur `FileInfo` décrivant le
fichier et une erreur indiquant des problèmes lors de l'inspection du fichier.

Déterminer si un fichier existe
Le package `os` définit une fonction nommée `IsNotExist`, qui accepte une erreur et renvoie vrai si l'erreur indique que le fichier n'existe pas.


Recherche de fichiers à l'aide d'un modèle
Le package `path/filepath` définit la fonction `Glob`, qui renvoie tous les noms d'un répertoire correspondant à un modèle spécifié.
Les Fonctions du package `path/filepath` pour localiser des fichiers selon un modèle :

- Match(pattern, name) : Cette fonction compare un chemin d'accès à un modèle. Elle renvoie un booléen indiquant s'il y a correspondance ou
une erreur signalant un problème avec le modèle ou lors de la recherche.
- Glob(pathPatten) : Cette fonction recherche tous les fichiers correspondant au modèle spécifié. Elle renvoie une liste de chaînes de
caractères contenant les chemins d'accès correspondants et une erreur signalant un problème lors de la recherche.

Syntaxe des modèles de recherche pour les fonctions `path/filepath`
* : Ce terme correspond à toute séquence de caractères, à l'exception du séparateur de chemin.
? : Ce terme correspond à n'importe quel caractère, à l'exception du séparateur de chemin.
[a-Z] : Ce terme correspond à n'importe quel caractère de la plage spécifiée.


Traitement de tous les fichiers d'un répertoire
Une alternative à l'utilisation de modèles consiste à énumérer tous les fichiers d'un emplacement spécifique, ce qui peut être réalisé à
l'aide de la fonction définie dans le package `path/filepath` :
- WalkDir(directory, func) : Cette fonction appelle la fonction spécifiée pour chaque fichier et répertoire du répertoire spécifié.

La fonction `callback` appelée par `WalkDir` reçoit une chaîne de caractères contenant le chemin d'accès, une valeur `DirEntry` fournissant des détails
sur le fichier ou le répertoire, et une erreur indiquant des problèmes d'accès à ce fichier ou répertoire. Le résultat de la fonction `callback`
est une erreur empêchant la fonction `WalkDir` d'accéder au répertoire courant en renvoyant la valeur spéciale `SkipDir`.
**/

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	path, err := os.Getwd()
	if err == nil {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}

	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}
	}

	path1, err1 := os.Getwd()
	if err1 == nil {
		matches, err1 := filepath.Glob(filepath.Join(path1, "*.json"))
		if err1 == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err1 != nil {
		Printfln("Error %v", err.Error())
	}

	// Cet exemple utilise la fonction `WalkDir` pour énumérer le contenu du répertoire actuel et affiche le chemin et la taille
	// de chaque fichier trouvé.
	path2, err2 := os.Getwd()
	if err2 == nil {
		err2 = filepath.WalkDir(path2, callback)
	} else {
		Printfln("Error %v", err.Error())
	}
}
