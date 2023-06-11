package main

import (
	"encoding/json"
	"os"
	"strings"
)

type ConfigData struct {
	Username           string
	AdditionalProducts []Product
}

var config ConfigData

func LoadConfig() (err error) {
	/**
	Cette fonction os.ReadFile ouvre le fichier spécifié et lit son contenu. Les résultats sont une tranche d'octet contenant le contenu du fichier
	et une erreur indiquant des problèmes d'ouverture ou de lecture du fichier.
	La fonction LoadConfig utilise la fonction ReadFile pour lire le contenu du fichier config.json. Le fichier sera lu à partir du répertoire de travail
	actuel lors de l'exécution de l'application, ce qui signifie que nous pouvons ouvrir le fichier uniquement avec son nom.
	Le contenu du fichier est renvoyé sous la forme d'une tranche d'octets, qui est convertie en chaîne.
	**/
	data, err := os.ReadFile("config.json")
	if err == nil {
		stringData := string(data)
		decoder := json.NewDecoder(strings.NewReader(stringData))
		err = decoder.Decode(&config)
	}

	return
}

func LoadConfigWithOpen() (err error) {
	/**
	La fonction Open ouvre un fichier en lecture et renvoie une valeur File, qui représente le fichier ouvert, et une erreur,
	qui est utilisée pour indiquer les problèmes d'ouverture du fichier. La structure File implémente l'interface Reader, ce qui simplifie
	la lecture et le traitement des exemples de données JSON, sans lire l'intégralité du fichier dans une tranche d'octet.
	**/
	file, err := os.Open("config.json")
	if err == nil {
		/**
		Le mot clé defer peut être utilisé pour appeler la méthode Close lorsque la fonction englobante se termine.
		Nous pouvons simplement appeler la méthode Close à la fin de la fonction si nous préférons, mais l'utilisation du mot clé defer garantit
		que le fichier est fermé même lorsqu'une fonction revient plus tôt.
		**/
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
	}

	return
}

/*
*
Si nous recevons une erreur dans cet exemple, la cause probable est que les emplacements spécifiés ne correspondent pas à la structure de votre fichier JSON.
Dans un premier temps, en particulier sous Linux, assurons-nous d'avoir enregistré le fichier avec les caractères CR et LR, ce que nous pouvons
faire dans Visual Studio Code en cliquant sur l'indicateur LR en bas à droite de la fenêtre.
*
*/
func LoadConfigWithSpecificLocation() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()

		nameSlice := make([]byte, 5)
		/**
		Cette méthode file.ReadAt est définie par l'interface ReaderAt et effectue une lecture dans la tranche spécifique à
		l'offset de position spécifié dans le fichier.
		**/
		file.ReadAt(nameSlice, 20)
		config.Username = string(nameSlice)

		/**
		Cette méthode file.Seek est définie par l'interface Seeker et déplace le décalage dans le fichier pour la lecture suivante.
		Le décalage est déterminé par la combinaison des deux arguments : le premier argument spécifie le nombre d'octets à décaler, et
		le deuxième argument détermine comment le décalage est appliqué. Une valeur de 0 signifie que le décalage est relatif au début du fichier,
		une valeur de 1 signifie que le décalage est relatif à la position de lecture actuelle et une valeur de 2 signifie que le décalage est relatif
		à la fin du fichier.
		**/
		file.Seek(55, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config.AdditionalProducts)
	}

	return
}

/*
*
La fonction LoadConfig est invoquée par une fonction d'initialisation init, qui assure la lecture du fichier de configuration.
*
*/
func init() {
	// err := LoadConfig()
	err := LoadConfigWithSpecificLocation()
	if err != nil {
		Printfln("Error Loading Config : %v", err.Error())
	} else {
		Printfln("Username : %v", config.Username)
		Products = append(Products, config.AdditionalProducts...)
	}
}
