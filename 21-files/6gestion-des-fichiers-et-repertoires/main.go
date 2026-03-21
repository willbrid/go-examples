package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

/**
Fonctions du package `os` pour la gestion des fichiers et des répertoires :

- Chdir(dir) : Cette fonction modifie le répertoire de travail courant pour le répertoire spécifié. En cas de problème lors de la modification,
une erreur est renvoyée.

- Mkdir(name, modePerms) : Cette fonction crée un répertoire avec le nom et les permissions spécifiés. En cas de problème,
une erreur est renvoyée : `nil` si le répertoire est créé, ou un message d'erreur est affiché.

- MkdirAll(name, modePerms) : Cette fonction effectue la même opération que `Mkdir`, mais crée tous les répertoires parents dans le chemin spécifié.

- MkdirTemp(parentDir, name) : Cette fonction est similaire à `CreateTemp`, mais crée un répertoire au lieu d'un fichier.
Une chaîne de caractères aléatoire est ajoutée à la fin du nom spécifié ou à la place d'un astérisque, et le nouveau répertoire est créé
dans le répertoire parent spécifié. Les résultats sont le nom du répertoire et un message d'erreur.

- Remove(name) : Cette fonction supprime le fichier ou le répertoire spécifié. En cas de problème, une erreur est renvoyée.

- RemoveAll(name) : Cette fonction supprime le fichier ou le répertoire spécifié. Si le nom désigne un répertoire, tous ses sous-répertoires
sont également supprimés. Un message d'erreur décrivant les éventuels problèmes rencontrés s'affiche.

- Rename(old, new) : Cette fonction renomme le fichier ou le dossier spécifié. Un message d'erreur décrivant les éventuels problèmes
rencontrés s'affiche.

- Symlink(old, new) : Cette fonction crée un lien symbolique vers le fichier spécifié. Un message d'erreur décrivant les éventuels problèmes
rencontrés s'affiche.
**/

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "Myapp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	err = os.MkdirAll(filepath.Dir(path), 0766)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}
