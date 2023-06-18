package main

import (
	"fmt"
	"io"
	"net/http"
)

/**
Les champs et la méthode multipart.FileHeader (l'un des champs de réponse renvoyé par la méthode request.FormFile) :
- Name : ce champ renvoie une chaîne contenant le nom du fichier.
- Size : ce champ renvoie un int64 contenant la taille du fichier.
- Header : ce champ renvoie une chaîne map[string][]string, qui contient les en-têtes de la partie MIME qui contient le fichier.
- Open() : cette méthode retourne un File qui peut être utilisé pour lire le contenu associé à l'en-tête.
**/

/*
*
Le premier résultat de la méthode FormFile est un File, défini dans le package mime/multipart, qui est une interface qui combine
les interfaces Reader, Closer, Seeker et ReaderAt. L'effet est que le contenu du fichier téléchargé peut être traité comme un Reader,
avec un support pour rechercher ou lire à partir d'un emplacement spécifique. Dans cet exemple, l'on copie le contenu du fichier téléchargé
dans le ResponseWriter. Le deuxième résultat de la méthode FormFile est un FileHeader, également défini dans le package mime/multipart.
*
*/
func HandleMultipartForm1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Name : %v, City : %v\n", request.FormValue("name"), request.FormValue("city"))
	fmt.Fprintf(writer, "-------")
	file, header, err := request.FormFile("files")
	if err == nil {
		defer file.Close()
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
		for k, v := range header.Header {
			fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
		}
		fmt.Fprintln(writer, "-------")
		io.Copy(writer, file)
	} else {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

/*
*
Nous devons nous assurer que la méthode ParseMultipartForm est appelée avant d'utiliser le champ MultipartForm.
Le champ MultipartForm renvoie une classe Form, qui est définie dans le package mime/multipart, et qui définit les champs :
- Value : ce champ renvoie une chaîne map[string][]string qui contient les valeurs du formulaire.
- File : ce champ renvoie un map[string][]*FileHeader qui contient les fichiers.
*
*/
func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10000000)
	fmt.Fprintf(writer, "Name: %v, City: %v\n", request.MultipartForm.Value["name"][0], request.MultipartForm.Value["city"][0])
	fmt.Fprintln(writer, "-------")

	for _, header := range request.MultipartForm.File["files"] {
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
		file, err := header.Open()
		if err == nil {
			defer file.Close()
			fmt.Fprintln(writer, "-------")
			io.Copy(writer, file)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func init() {
	http.HandleFunc("/forms/upload", HandleMultipartForm)
}
