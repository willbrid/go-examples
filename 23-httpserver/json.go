package main

import (
	"encoding/json"
	"net/http"
)

/*
*
Cette fonction HandleJsonRequest utilise les fonctionnalités JSON pour encoder la tranche de valeurs Product.
*
*/
func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Products)
}

/*
*
La fonction d'initialisation crée une route, ce qui signifie que les requêtes pour /json seront traitées par la fonction HandleJsonRequest.
*
*/
func init() {
	http.HandleFunc("/json", HandleJsonRequest)
}
