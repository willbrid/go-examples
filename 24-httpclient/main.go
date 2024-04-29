package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	/**
	Le serveur est démarré dans une goroutine pour l'empêcher de se bloquer et permettre l'envoi de la requête HTTP au sein de la même application.
	L'on utilise la fonction time.Sleep pour s'assurer que la goroutine a le temps de démarrer le serveur.
	**/
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	Printfln("Starting HTTP Server")

	/**
	L'argument de la fonction Get est une chaîne qui contient l'URL à demander. Les résultats sont une valeur de réponse et
	une erreur qui signale tout problème d'envoi de la demande.
	Les champs et méthodes définis par la structure de réponse :
	- StatusCode : ce champ renvoie le code d'état de la réponse, exprimé sous la forme d'un entier.
	- Status : ce champ renvoie une chaîne contenant la description du statut.
	- Proto : ce champ renvoie une chaîne contenant le protocole HTTP de réponse.
	- Header : ce champ renvoie une chaîne map[string][]string qui contient les en-têtes de réponse.
	- Body : ce champ renvoie un ReadCloser, qui est un Reader qui définit une méthode Close et qui donne accès au corps de la réponse.
	- Trailer : ce champ renvoie une chaîne map[string][]string qui contient les bandes-annonces de réponse.
	- ContentLength : ce champ renvoie la valeur de l'en-tête Content-Length, analysée en une valeur int64.
	- TransferEncoding : ce champ renvoie l'ensemble des valeurs d'en-tête Transfer-Encoding.
	- Close : ce champ booléen renvoie true si la réponse contient un en-tête de connexion défini sur close,
			  ce qui indique que la connexion HTTP doit être fermée.
	- Uncompressed : ce champ renvoie vrai si le serveur a envoyé une réponse compressée qui a été décompressée par le package net/http.
	- Request : ce champ renvoie la requête qui a été utilisée pour obtenir la réponse.
	- TLS : ce champ fournit des détails sur la connexion HTTPS.
	- Cookies() : cette méthode renvoie un []*Cookie, qui contient les en-têtes Set-Cookie dans la réponse.
	- Location() : cette méthode renvoie l'URL de l'en-tête Location de la réponse et
	               une erreur qui indique quand la réponse ne contient pas cet en-tête.
	- Write(writer) : cette méthode écrit un résumé de la réponse au Writer spécifié.

	L'on utilise la fonction ReadAll définie dans le package io pour lire la réponse Body dans une tranche d'octet et l'on écrit sur la sortie standard.
	**/
	response1, err1 := http.Get("http://localhost:5000/html")
	if err1 == nil && response1.StatusCode == http.StatusOK {
		response1.Write(os.Stdout)
		Printfln("\n")
	} else {
		Printfln("Error : %v", err1.Error())
	}

	/**
	Cette fonction http.Get envoie une requête GET à l'URL HTTP ou HTTPS spécifiée. Les résultats sont une réponse et une erreur
	qui signale des problèmes avec la demande.
	**/
	response2, err2 := http.Get("http://localhost:5000/html")
	if err2 == nil && response2.StatusCode == http.StatusOK {
		data, err := io.ReadAll(response2.Body)
		if err == nil {
			defer response2.Body.Close()
			os.Stdout.Write(data)
		}
		response2.Write(os.Stdout)
		Printfln("\n")
	} else {
		Printfln("Error : %v", err2.Error())
	}

	response3, err3 := http.Get("http://localhost:5000/json")
	if err3 == nil && response3.StatusCode == http.StatusOK {
		defer response3.Body.Close()
		data := []Product{}
		jsonErr := json.NewDecoder(response3.Body).Decode(&data)
		if jsonErr == nil {
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", jsonErr.Error())
		}
	} else {
		Printfln("Error : %v", err3.Error())
	}

	formData := map[string][]string{
		"name":     {"Kayak "},
		"category": {"Watersports"},
		"price":    {"279"},
	}
	/**
	Cette fonction http.PostForm envoie une requête POST à l'URL HTTP ou HTTPS spécifiée, avec l'en-tête Content-Type défini
	sur application/x-www-form-urlencoded. Le contenu du formulaire est fourni par une map[string][]string. Les résultats sont une
	réponse et une erreur qui signale des problèmes avec la demande.
	**/
	response4, err4 := http.PostForm("http://localhost:5000/echo", formData)
	if err4 == nil && response4.StatusCode == http.StatusOK {
		io.Copy(os.Stdout, response4.Body)
		defer response4.Body.Close()
	} else {
		Printfln("Error: %v, Status Code: %v", err4.Error(), response4.StatusCode)
	}

	var builder strings.Builder
	jsonErr := json.NewEncoder(&builder).Encode(Products[0])
	if jsonErr == nil {
		/**
		Cette fonction http.Post envoie une requête POST à l'URL HTTP ou HTTPS spécifiée, avec la valeur d'en-tête Content-Type spécifiée.
		Le contenu du formulaire est fourni par le Reader spécifié. Les résultats sont une réponse et une erreur qui signale des problèmes avec la requête.
		**/
		response5, err5 := http.Post("http://localhost:5000/echo", "application/json", strings.NewReader(builder.String()))
		if err5 == nil && response5.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, response5.Body)
			defer response5.Body.Close()
		} else {
			Printfln("Error: %v", err5.Error())
		}
	} else {
		Printfln("Error: %v", jsonErr.Error())
	}

	/**
	La classe Client est utilisée lorsqu'un contrôle est requis sur une requête HTTP et définit les champs et les méthodes :
	- Transport : ce champ permet de sélectionner le transport qui sera utilisé pour envoyer la requête HTTP.
	              Le package net/http fournit un transport par défaut.
	- CheckRedirect : ce champ est utilisé pour spécifier une stratégie personnalisée pour traiter les redirections répétées.
	- Jar : ce champ renvoie un CookieJar, qui sert à gérer les cookies.
	- Timeout : ce champ est utilisé pour définir un délai d'attente pour la requête, spécifié sous forme de time.Duration.
	- Do(request) : cette méthode envoie la requête spécifiée, renvoyant une réponse et une erreur indiquant des problèmes lors de l'envoi de la requête.
	- CloseIdleConnections() : cette méthode ferme toutes les requêtes HTTP inactives qui sont actuellement ouvertes et inutilisées.
	- Get(url) : cette méthode est appelée par la fonction Get.
	- Head(url) : cette méthode est appelée par la fonction Head.
	- Post(url, contentType, reader) : cette méthode est appelée par la fonction Post.
	- PostForm(url, data) : cette méthode est appelée par la fonction PostForm.

	La classe Request qui décrit la requête HTTP. Les champs et méthodes de Request les plus utiles pour les requêtes des clients :
	- Method : ce champ de chaîne spécifie la méthode HTTP qui sera utilisée pour la requête.
	           Le package net/http définit des constantes pour les méthodes HTTP, telles que MethodGet et MethodPost.
	- URL : ce champ URL spécifie l'URL à laquelle la requête sera envoyée. .
	- Header : ce champ permet de préciser les en-têtes de la requête. Les en-têtes sont spécifiés dans une chaîne map[string][],
	           et le champ sera nil lorsqu'une valeur Request est créée à l'aide de la syntaxe de structure littérale.
	- ContentLength : ce champ est utilisé pour définir l'en-tête Content-Length à l'aide d'une valeur int64.
	- TransferEncoding : ce champ est utilisé pour définir l'en-tête Transfer-Encoding à l'aide d'une tranche de chaînes.
	- Body : ce champ ReadCloser spécifie la source du corps de la requête. Si nous avons un lecteur qui ne définit pas de méthode Close,
	         la fonction io.NopCloser peut être utilisée pour créer un ReadCloser dont la méthode Close ne fait rien.

	Le moyen le plus simple de créer une valeur d'URL consiste à utiliser la fonction Parse fournie par le package net/url, qui analyse une chaîne :
	- Parse(string) : cette méthode analyse une chaîne dans une URL. Les résultats sont la valeur de l'URL et une erreur indiquant des
	  				  problèmes d'analyse de la chaîne.
	**/

	/**
	Cette liste crée une nouvelle requête à l'aide de la syntaxe littérale, puis définit les champs Méthode, URL et Corps.
	La méthode est définie de sorte qu'une requête POST soit envoyée, l'URL est créée à l'aide de la fonction Parse et le champ Body
	est défini à l'aide de la fonction io.NopCloser, qui accepte un Reader et renvoie un ReadCloser, qui est le type requis par le Structure de requête.
	Le champ Header se voit attribuer une carte qui définit l'en-tête Content-Type. Un pointeur vers la requête est passé à la méthode Do du client
	affecté à la variable DefaultClient, qui envoie la requête.
	**/
	var builder1 strings.Builder
	jsonErr1 := json.NewEncoder(&builder1).Encode(Products[0])
	if jsonErr1 == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-type": {"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(builder1.String())),
			}
			response, err := http.DefaultClient.Do(&req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Parse Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", jsonErr1.Error())
	}

	/**
	Le Client garde une trace des cookies qu'il reçoit du serveur et les inclut automatiquement dans les requêtes ultérieures.

	Le package mime/multipart peut être utilisé pour créer un corps de requête encodé en tant que multipart/form-data,
	ce qui permet à un formulaire de contenir en toute sécurité des données binaires, telles que le contenu d'un fichier.
	**/
	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)
	fieldWriter, err := formWriter.CreateFormField("name")
	if err == nil {
		io.WriteString(fieldWriter, "Alice")
	}
	fieldWriter, err = formWriter.CreateFormField("city")
	if err == nil {
		io.WriteString(fieldWriter, "New York")
	}
	fileWriter, err := formWriter.CreateFormFile("codeFile", "printer.go")
	if err == nil {
		fileData, err := os.ReadFile("./printer.go")
		if err == nil {
			fileWriter.Write(fileData)
		}
	}
	formWriter.Close()
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/form", &buffer)
	if err == nil {
		req.Header["Content-Type"] = []string{formWriter.FormDataContentType()}
	}
}
