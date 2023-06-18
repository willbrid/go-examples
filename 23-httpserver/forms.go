package main

import (
	"net/http"
	"strconv"
)

/**
Champs de données et méthodes du formulaire de requête :

- Form : ce champ renvoie une chaîne map[string][]string contenant les données de formulaire analysées et les paramètres de la chaîne de requête.
         La méthode ParseForm doit être appelée avant que ce champ ne soit lu.

- PostForm : ce champ est similaire à Form mais exclut les paramètres de chaîne de requête afin que
             seules les données du corps de la requête soient contenues dans la map.
			 La méthode ParseForm doit être appelée avant que ce champ ne soit lu.

- MultipartForm  : ce champ renvoie un formulaire multipart représenté à l'aide de la classe Form définie dans le package mime/multipart.
                   La méthode ParseMultipartForm doit être appelée avant que ce champ ne soit lu.

- FormValue(key) : cette méthode renvoie la première valeur de la clé de formulaire spécifiée et renvoie la chaîne vide s'il n'y a pas de valeur.
                   La source de données pour cette méthode est le champ Form, et l'appel de la méthode FormValue appelle automatiquement
				   ParseForm ou ParseMultipartForm pour analyser le formulaire.

- PostFormValue(key) : cette méthode renvoie la première valeur de la clé de formulaire spécifiée et renvoie la chaîne vide s'il n'y a pas de valeur.
                       La source de données pour cette méthode est le champ PostForm et l'appel de la méthode PostFormValue appelle automatiquement
                       ParseForm ou ParseMultipartForm pour analyser le formulaire.

- FormFile(key) : cette méthode donne accès au premier fichier avec la clé spécifiée dans le formulaire. Les résultats sont un fichier et une en-tête
                  de fichier, tous deux définis dans le package mime/multipart, et une erreur. L'appel de cette fonction entraîne
				  l'invocation des fonctions ParseForm ou ParseMultipartForm pour analyser le formulaire.

- ParseForm() : cette méthode analyse un formulaire et remplit les champs Form et PostForm. Le résultat est une erreur qui décrit tout problème d'analyse.

- ParseMultipart : cette méthode analyse un formulaire multipart MIME et remplit le champ MultipartForm.

- Form(max) : L'argument spécifie le nombre maximal d'octets à allouer aux données du formulaire, et le résultat est une erreur
              qui décrit tout problème de traitement du formulaire.
**/

func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		index, _ := strconv.Atoi(request.PostFormValue("index"))
		p := Product{}

		p.Name = request.PostFormValue("name")
		p.Category = request.PostFormValue("category")
		p.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
		Products[index] = p
	}

	http.Redirect(writer, request, "/templates/", http.StatusTemporaryRedirect)
}

func init() {
	http.HandleFunc("/forms/edit", ProcessFormData)
}
