package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Context struct {
	Request *http.Request
	Data    []Product
}

var htmlTemplates *template.Template

/*
*
Notons que l'en-tête Content-Type n'a pas été défini lors de l'utilisation d'un modèle pour générer une réponse.
Lors de la diffusion de fichiers, l'en-tête Content-Type est défini en fonction de l'extension de fichier,
mais ce n'est pas possible dans cette situation lorsque le contenu est écrit directement dans le ResponseWriter.
Lorsqu'une réponse n'a pas d'en-tête Content-Type, les 512 premiers octets de contenu écrits dans ResponseWriter sont transmis à la
fonction DetectContentType, qui implémente l'algorithme MIME Sniffing défini par https://mimesniff.spec.whatwg.org.
Le processus de détection ne peut pas détecter tous les types de contenu, mais il fonctionne bien avec les types Web standard,
tels que HTML, CSS et JavaScript. La fonction DetectContentType renvoie un type MIME, qui est utilisé comme valeur pour l'en-tête Content-Type.
*
*/
func HandleTemplateRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if path == "" {
		path = "products.html"
	}
	Printfln("Path : %v", path)
	t := htmlTemplates.Lookup(path)
	if t == nil {
		http.NotFound(writer, request)
	} else {
		err := t.Execute(writer, Context{Request: request, Data: Products})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

/*
*
La fonction d'initialisation charge tous les modèles avec l'extension html dans le dossier des modèles et configure une route afin que
les requêtes commençant par /templates/ soient traitées par la fonction HandleTemplateRequest. Cette fonction recherche le modèle,
revient au fichier products.html si aucun chemin de fichier n'est spécifié, exécute le modèle et écrit la réponse.
*
*/
func init() {
	var err error
	htmlTemplates = template.New("all")
	htmlTemplates.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})
	htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
	if err == nil {
		http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(HandleTemplateRequest)))
	} else {
		panic(err)
	}
}
