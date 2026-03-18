package main

import (
	"encoding/json"
	"os"
)

/**
La structure `File` implémente également l'interface `Closer` qui définit une méthode `Close`. Le mot-clé `defer` peut être utilisé pour appeler
la méthode `Close` lorsque la fonction englobante se termine.

La lecture à partir d'emplacements spécifiques nécessite une connaissance de la structure du fichier.
**/

type ConfigData struct {
	Username           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
	}

	return
}

func LoadConfigFromSpecificLocation() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()

		nameSlice := make([]byte, 5)
		file.ReadAt(nameSlice, 20)
		Config.Username = string(nameSlice)
		file.Seek(55, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config.AdditionalProducts)
	}

	return
}

func init() {
	// err := LoadConfig()
	err := LoadConfigFromSpecificLocation()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.Username)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
