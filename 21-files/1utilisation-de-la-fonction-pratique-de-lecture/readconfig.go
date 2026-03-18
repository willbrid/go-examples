package main

import (
	"encoding/json"
	"os"
	"strings"
)

/**
La fonction LoadConfig utilise la fonction ReadFile pour lire le contenu du fichier config.json. Le fichier sera lu à partir du répertoire de travail
actuel lors de l'exécution de l'application, ce qui signifie que nous pouvons ouvrir le fichier uniquement avec son nom.
Le contenu du fichier est renvoyé sous la forme d'une slice d'octets, qui est convertie en chaîne.
**/

type ConfigData struct {
	Username           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		Printfln(string(data))
	}

	return
}

func LoadAndDecodeConfig() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&Config)
	}

	return
}

func init() {
	// err := LoadConfig()
	err := LoadAndDecodeConfig()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.Username)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
