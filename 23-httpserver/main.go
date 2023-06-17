package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, sh.message)
}

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}

	/**
	Le package net/http facilite la création d'un serveur HTTP simple, qui peut ensuite être étendu pour ajouter des
	fonctionnalités plus complexes et utiles.
	**/
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, world !"})
	if err != nil {
		Printfln("Error : %v", err.Error())
	}
}
