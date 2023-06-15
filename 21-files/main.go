package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path filepath.WalkDir callback : %v, Size: %v", path, info.Size())
	return
}

func main() {
	for _, p := range Products {
		Printfln("Product : %v, Category : %v, Price : $%.2f", p.Name, p.Category, p.Price)
	}

	/**
	La fonction WriteFile fournit un moyen pratique d'écrire un fichier entier en une seule étape et créera le fichier s'il n'existe pas.
	Cette fonction os.WriteFile crée un fichier avec le nom, le mode et les autorisations spécifiés et écrit le contenu de la tranche d'octets spécifiée.
	Si le fichier existe déjà, son contenu sera remplacé par la tranche d'octets. Le résultat est une erreur qui signale tout problème
	lors de la création du fichier ou de l'écriture des données.
	Les autorisations de fichier sont plus largement utilisées et suivent le style UNIX d'autorisation de fichier, qui se compose de trois chiffres
	qui définissent l'accès pour le propriétaire, le groupe et les autres utilisateurs du fichier. Chaque chiffre est la somme des autorisations
	qui doivent être accordées, où read a une valeur de 4, write a une valeur de 2 et execute a une valeur de 1.
	Si le fichier spécifié existe déjà, la méthode WriteFile remplace son contenu.
	**/
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	dataStr := fmt.Sprintf("Time : %v, Total : $%.2f\n", time.Now().Format("Mon 15:04:05"), total)
	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if err == nil {
		fmt.Println("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}

	/**
		La fonction OpenFile ouvre un fichier et renvoie un pointeur sur la classe File. Contrairement à la fonction Open, la fonction OpenFile accepte
		un ou plusieurs indicateurs qui spécifient comment le fichier doit être ouvert. Les indicateurs sont définis comme des constantes dans le package os.
		- O_RDONLY : cet indicateur ouvre le fichier en lecture seule afin qu'il puisse être lu mais pas écrit.
		- O_WRONLY : cet indicateur ouvre le fichier en écriture seule afin qu'il puisse être écrit mais pas lu.
		- O_RDWR : cet indicateur ouvre le fichier en lecture-écriture afin qu'il puisse être écrit et lu à partir de celui-ci.
	    - O_APPEND : ce indicateur ajoutera les écritures à la fin du fichier.
	    - O_CREATE : ce indicateur créera le fichier s'il n'existe pas.
	    - O_EXCL : cet indicateur est utilisé conjointement avec O_CREATE pour s'assurer qu'un nouveau fichier est créé. Si le dossier
	      existe déjà, ce indicateur déclenchera une erreur.
	    - O_SYNC : cet indicateur active les écritures synchrones, de sorte que les données sont écrites sur le périphérique de stockage avant
	      la fonction/méthode d'écriture revient.
	    - O_TRUNC : cet indicateur tronque le contenu existant dans le fichier.
		**/
	file, err1 := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err1 == nil {
		defer file.Close()
		/**
		Cette méthode file.WriteString écrit une chaîne dans le fichier. Il s'agit d'une méthode pratique qui convertit la chaîne en une tranche d'octets,
		appelle la méthode Write et renvoie les résultats qu'elle reçoit.
		- Seek(offset, how) : cette méthode définit l'emplacement pour les opérations suivantes.
		- Write(slice) : cette méthode écrit le contenu de la tranche d'octets spécifiée dans le fichier.
		  Les résultats sont le nombre d'octets écrits et une erreur qui indique des problèmes d'écriture des données.
		- WriteAt(slice, offset) - cette méthode écrit les données dans la tranche à l'emplacement spécifié et est l'équivalent de la méthode ReadAt
		  pour la lecture.
		**/
		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err1.Error())
	}

	/**
	La classe File implémente l'interface Writer, qui permet d'utiliser un fichier avec les fonctions de formatage et de traitement des chaînes.
	Cela signifie également que les fonctionnalités JSON peuvent être utilisées pour écrire des données JSON dans un fichier.
	Cet exemple sélectionne les valeurs Product avec une valeur Price inférieure à 100, les place dans une tranche et utilise un encodeur JSON
	pour écrire cette tranche dans un fichier nommé cheap.json.
	**/
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file1, err2 := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err2 == nil {
		defer file1.Close()
		encoder1 := json.NewEncoder(file1)
		encoder1.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err2.Error())
	}

	/**
	Cette fonction os.Create équivaut à appeler OpenFile avec les indicateurs O_RDWR, O_CREATE et O_TRUNC. Les résultats sont le fichier,
	qui peut être utilisé pour la lecture et l'écriture, et une erreur qui est utilisée pour indiquer des problèmes lors de la création du fichier.
	Notons que cette combinaison d'indicateurs signifie que si un fichier existe avec le nom spécifié, il sera ouvert et son contenu sera supprimé.
	**/
	file2, err3 := os.Create("file.json")
	if err3 == nil {
		defer file2.Close()
		encoder2 := json.NewEncoder(file2)
		encoder2.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err3.Error())
	}

	/**
	Cette fonction os.CreateTemp crée un nouveau fichier dans le répertoire avec le nom spécifié. Si le nom est la chaîne vide, le répertoire temporaire
	du système est utilisé, obtenu à l'aide de la fonction TempDir. Le fichier est créé avec un nom qui contient une séquence aléatoire de caractères,
	comme illustré dans le texte après le tableau. Le fichier est ouvert avec les indicateurs O_RDWR, O_CREATE et O_EXCL. Le fichier n'est pas supprimé
	lorsqu'il est fermé.
	**/
	file3, err4 := os.CreateTemp("", "tempfile-*.json")
	if err4 == nil {
		defer file3.Close()
		encoder3 := json.NewEncoder(file3)
		encoder3.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err4.Error())
	}

	/**
	Le package os fournit un ensemble de fonctions qui renvoient les chemins des emplacements communs.
	- Getwd() : cette fonction renvoie le répertoire de travail actuel, exprimé sous forme de chaîne, et une erreur indiquant des problèmes
	  d'obtention de la valeur.
	- UserHomeDir() : cette fonction renvoie le répertoire personnel de l'utilisateur et une erreur qui indique des problèmes pour obtenir le chemin.
	- UserCacheDir() : cette fonction renvoie le répertoire par défaut pour les données mises en cache spécifiques à l'utilisateur et une erreur
	  qui indique des problèmes d'obtention du chemin.
	- UserConfigDir() : cette fonction renvoie le répertoire par défaut pour les données de configuration spécifiques à l'utilisateur et une erreur
	  qui indique des problèmes d'obtention du chemin.
	- TempDir() : cette fonction renvoie le répertoire par défaut pour les fichiers temporaires et une erreur qui indique des problèmes pour
	  obtenir le chemin.

	filepath.Abs(path) : cette fonction renvoie un chemin absolu, ce qui est utile si nous avons un chemin relatif, tel qu'un nom de fichier.
	filepath.IsAbs(path) : cette fonction renvoie true si le chemin spécifié est absolu.
	filepath.Match(pattern, path) : cette fonction renvoie vrai si le chemin correspond au modèle spécifié.
	**/
	path, err5 := os.UserHomeDir()
	if err5 == nil {
		// Cette fonction filepath.Join combine plusieurs éléments à l'aide du séparateur de fichiers de la plateforme (système d'opération).
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	// Cette fonction filepath.VolumeName renvoie le composant de volume du chemin spécifié ou la chaîne vide si le chemin ne contient pas de volume.
	Printfln("Volume name: %v", filepath.VolumeName(path))
	// Cette fonction filepath.Dir renvoie tous les éléments sauf le dernier du chemin.
	Printfln("Dir component: %v", filepath.Dir(path))
	// Cette fonction filepath.Base renvoie le dernier élément du chemin.
	Printfln("File component: %v", filepath.Base(path))
	// Cette fonction filepath.Ext renvoie l'extension de fichier à partir du chemin spécifié, qui est supposé être le suffixe suivant le point final
	// dans la chaîne de chemin.
	Printfln("File extension: %v", filepath.Ext(path))

	/**
		- os.MkdirTemp(parentDir,name) : cette fonction est similaire à CreateTemp mais crée un répertoire plutôt qu'un fichier. Une chaîne aléatoire
		est ajoutée à la fin du nom spécifié ou à la place d'un astérisque, et le nouveau répertoire est créé dans le parent spécifié. Les résultats
		sont le nom du répertoire et une erreur indiquant des problèmes.
		- os.Chdir(dir) : cette fonction remplace le répertoire de travail courant par le répertoire spécifié. Le résultat est une erreur qui indique des
		  problèmes lors de la modification.
	    - os.Mkdir(name, modePerms) : cette fonction crée un répertoire avec le nom et le mode/permissions spécifiés. Le résultat est une erreur
		  qui est nulle si le répertoire est créé ou qui décrit un problème le cas échéant.
		- os.Remove(name) : cette fonction supprime le fichier ou le répertoire spécifié. Le résultat est une erreur qui décrit
		  tous les problèmes qui surviennent.
		- os.RemoveAll(name) : cette fonction supprime le fichier ou le répertoire spécifié. Si le nom spécifie un répertoire,
		  tous les enfants qu'il contient sont également supprimés. Le résultat est une erreur qui décrit tous les problèmes qui surviennent.
		- os.Rename(ancien, nouveau) : cette fonction renomme le fichier ou le dossier spécifié. Le résultat est une erreur qui décrit tous les
		  problèmes qui surviennent.
		- os.Symlink(old, new) : cette fonction crée un lien symbolique vers le fichier spécifié. Le résultat est une erreur qui décrit tous
		  les problèmes qui surviennent.
		**/
	path1, err6 := os.UserHomeDir()
	if err6 == nil {
		// Cette fonction filepath.Join combine plusieurs éléments à l'aide du séparateur de fichiers de la plateforme (système d'opération).
		path1 = filepath.Join(path1, "MyApp", "MyTempFile.json")
	}
	/**
	Cette fonction os.MkdirAll effectue la même tâche que Mkdir mais crée tous les répertoires parents dans le chemin spécifié.
	**/
	err7 := os.MkdirAll(filepath.Dir(path1), 0766)
	if err7 == nil {
		file4, err8 := os.OpenFile(path1, os.O_CREATE|os.O_WRONLY, 0666)
		if err8 == nil {
			defer file4.Close()
			encoder4 := json.NewEncoder(file4)
			encoder4.Encode(cheapProducts)
		} else {
			Printfln("Error: %v", err8.Error())
		}
	} else {
		Printfln("Error: %v", err7.Error())
	}

	/**
	Getwd renvoie un nom de chemin racine correspondant au répertoire courant. Si le répertoire courant peut être atteint via plusieurs chemins
	(en raison de liens symboliques), Getwd peut renvoyer n'importe lequel d'entre eux.
	**/
	path2, err9 := os.Getwd()
	if err9 == nil {
		// Cette fonction os.ReadDir lit le répertoire spécifié et renvoie une tranche DirEntry, chacune décrivant un élément du répertoire.
		dirEntries, err10 := os.ReadDir(path2)
		if err10 == nil {
			for _, dentry := range dirEntries {
				/**
				Cette méthode dentry.Name renvoie le nom du fichier ou du répertoire décrit par la valeur DirEntry.
				Cette méthode dentry.IsDir renvoie true si la valeur DirEntry représente un répertoire.
				**/
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
				/**
				Cette méthode dentry.Type renvoie une valeur FileMode, qui est un alias de uint32, qui décrit davantage le fichier et
				les autorisations du fichier ou du répertoire représenté par la valeur DirEntry.
				**/
				fileMode := dentry.Type()
				Printfln("FileMode -> IsDir : %v, IsRegular : %v", fileMode.IsDir(), fileMode.IsRegular())
				/**
				Cette méthode dentry.Info renvoie une valeur FileInfo qui fournit des détails supplémentaires sur le fichier ou le répertoire représenté
				par la valeur DirEntry.
				**/
				fileinfo, err := dentry.Info()
				if err == nil {
					/**
					Cette méthode fileinfo.Name renvoie une chaîne contenant le nom du fichier ou du répertoire.
					Cette méthode fileinfo.Size renvoie la taille du fichier, exprimée sous la forme d'une valeur int64.
					Cette méthode fileinfo.Mode renvoie le mode de fichier et les paramètres d'autorisation pour le fichier ou le répertoire.
					Cette méthode fileinfo.ModTime renvoie l'heure de la dernière modification du fichier ou du répertoire.
					**/
					Printfln("FileInfo -> Name : %v, Size : %v, Mode : %v, ModTime : %v", fileinfo.Name(), fileinfo.Size(), fileinfo.Mode(), fileinfo.ModTime())
				}
			}
		}

		/**
		Cette fonction os.Stat accepte une chaîne de chemin. Il renvoie une valeur FileInfo qui décrit le fichier et une erreur, qui indique
		des problèmes lors de l'inspection du fichier.
		**/
		pathFileinfo, err11 := os.Stat(path2)
		if err11 == nil {
			Printfln("Path FileInfo -> Name : %v, Size : %v, Mode : %v, ModTime : %v", pathFileinfo.Name(), pathFileinfo.Size(), pathFileinfo.Mode(), pathFileinfo.ModTime())
		}
	} else {
		Printfln("Error: %v", err9.Error())
	}

	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		/**
		Le package os définit une fonction nommée IsNotExist, accepte une erreur et renvoie true s'il indique que
		l'erreur indique qu'un fichier n'existe pas
		**/
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}
	}

	path3, err12 := os.Getwd()
	if err12 == nil {
		configFile := filepath.Join(path3, "config.json")
		/**
		Cette fonction filepath.Match fait correspondre un seul chemin à un modèle. Les résultats sont un booléen, qui indique s'il y a une correspondance,
		et une erreur, qui indique des problèmes avec le modèle ou avec l'exécution de la correspondance.
		**/
		isMatchConfigFile, err13 := filepath.Match("^/*/*/*/*/*/*.json", configFile)
		if err13 == nil {
			Printfln("Config file : %v - Is it Match config file : %v", configFile, isMatchConfigFile)
		} else {
			Printfln("Error: %v", err13.Error())
		}

		/**
		Cette fonction filepath.Glob trouve tous les fichiers qui correspondent au modèle spécifié. Les résultats sont une tranche de chaîne
		contenant les chemins correspondants et une erreur indiquant des problèmes lors de l'exécution de la recherche.
		- * : Ce terme correspond à n'importe quelle séquence de caractères, à l'exception du séparateur de chemin.
		- * : Ce terme correspond à n'importe quel caractère unique, à l'exception du séparateur de chemin.
		- [a-Z] : Ce terme correspond à n'importe quel caractère dans la plage spécifiée.
		**/
		matches, err14 := filepath.Glob(filepath.Join(path3, "*.json"))
		if err14 == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	} else {
		Printfln("Error: %v", err12.Error())
	}

	path4, err15 := os.Getwd()
	if err15 == nil {
		// Cette fonction filepath.WalkDir appelle la fonction spécifiée pour chaque fichier et répertoire dans le répertoire spécifié.
		// callback(path string, dir os.DirEntry, dirErr error) (err error)
		err16 := filepath.WalkDir(path4, callback)
		if err16 != nil {
			Printfln("Error err 16 : %v", err16.Error())
		}
	} else {
		Printfln("Error: %v", err15.Error())
	}
}
