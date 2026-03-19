package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/**
Le package os inclut également des fonctions d'écriture de fichiers.
- WriteFile(name, slice, modePerms) : Cette fonction crée un fichier avec le nom, le mode et les permissions spécifiés, et y écrit le contenu
de la slice d'octets indiquée. Si le fichier existe déjà, son contenu sera remplacé par celui de la slice d'octets. En cas de problème lors
de la création du fichier ou de l'écriture des données, une erreur sera levée.

- OpenFile(name, flag, modePerms) : Cette fonction ouvre le fichier dont le nom est spécifié, en utilisant les options (flags) pour contrôler
son ouverture. Si un nouveau fichier est créé, le mode et les permissions spécifiés seront appliqués. Le résultat est une valeur de type `File`
permettant d'accéder au contenu du fichier, ainsi qu'une erreur signalant un problème lors de son ouverture.

Utilisation de la fonction d'écriture simplifiée
La fonction WriteFile offre un moyen pratique d'écrire un fichier entier en une seule étape et créera le fichier s'il n'existe pas.
Les deux premiers arguments de la fonction WriteFile sont le nom du fichier et une tranche d'octets contenant les données à écrire.
Le troisième argument combine deux paramètres du fichier : le mode et les permissions. Le mode permet de spécifier des caractéristiques particulières
pour le fichier, mais la valeur zéro est utilisée pour les fichiers ordinaires, comme dans l'exemple. Nous pouvons trouver la liste des valeurs
des modes de fichier et leurs paramètres à l'adresse https://golang.org/pkg/io/fs/#FileMode.


Utilisation de la structure File pour écrire dans un fichier
La fonction OpenFile ouvre un fichier et renvoie une valeur de type `File`. Contrairement à la fonction `Open`, `OpenFile` accepte un ou plusieurs
indicateurs (flags) spécifiant le mode d'ouverture du fichier. Ces indicateurs sont définis comme des constantes dans le package `os`.
Il convient d'être prudent avec ces indicateurs, car tous ne sont pas pris en charge par tous les systèmes d'exploitation.
Les indicateurs d'ouverture de fichier :
- O_RDONLY : Ce paramètre ouvre le fichier en lecture seule, permettant la lecture mais pas l'écriture.
- O_WRONLY : Ce paramètre ouvre le fichier en écriture seule, permettant l'écriture mais pas la lecture.
- O_RDWR : Ce paramètre ouvre le fichier en lecture-écriture, permettant la lecture et l'écriture.
- O_APPEND : Ce paramètre ajoute les données à la fin du fichier.
- O_CREATE : Ce paramètre crée le fichier s'il n'existe pas.
- O_EXCL : Ce paramètre est utilisé conjointement avec O_CREATE pour garantir la création d'un nouveau fichier. Si le fichier existe déjà,
ce paramètre génère une erreur.
- O_SYNC : Ce paramètre active les écritures synchrones, de sorte que les données sont écrites sur le périphérique de stockage avant que
la fonction/méthode d'écriture ne se termine.
- O_TRUNC : Ce paramètre tronque le contenu existant du fichier.

La structure `File` définit les méthodes permettant d'écrire des données dans un fichier une fois celui-ci ouvert.
- Seek(offset, how) : Cette méthode définit l'emplacement pour les opérations suivantes.
- Write(slice) : Cette méthode écrit le contenu de la slice d'octets spécifiée dans le fichier. Elle renvoie le nombre d'octets écrits et
une erreur indiquant les problèmes rencontrés lors de l'écriture des données.
- WriteAt(slice, offset) : Cette méthode écrit les données de la slice à l'emplacement spécifié et est l'équivalent de la méthode ReadAt.
- WriteString(str) : Cette méthode écrit une chaîne de caractères dans le fichier. Il s'agit d'une méthode pratique qui convertit la chaîne en
une slice d'octets, appelle la méthode `Write` et renvoie le résultat.


Écriture de données JSON dans un fichier
La structure File implémente l'interface `Writer`, ce qui permet d'utiliser un fichier avec les fonctions de formatage et de traitement des chaînes de
caractères. Cela signifie également que les fonctionnalités JSON peuvent être utilisées pour écrire des données JSON dans un fichier.
**/

func main() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}

	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)
	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if err == nil {
		Printfln("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}

	dataStr1 := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)
	file1, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer file1.Close()
		file1.WriteString(dataStr1)
	} else {
		Printfln("Error: %v", err.Error())
	}

	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file2, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		defer file2.Close()
		encoder := json.NewEncoder(file2)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
