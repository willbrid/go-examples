package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/**
Le package net/http définit la fonction SetCookie, qui ajoute un en-tête Set-Cookie à la réponse envoyée au client.
SetCookie(writer, cookie) : cette fonction ajoute un en-tête Set-Cookie au ResponseWriter spécifié. Le cookie est décrit à l'aide d'un pointeur
vers une classe Cookie. Les cookies sont décrits à l'aide de la classe Cookie, qui est définie dans le package net/http et définit les champs :
- Name : ce champ représente le nom du cookie, exprimé sous forme de chaîne.
- Value : ce champ représente la valeur du cookie, exprimée sous forme de chaîne.
- Path : ce champ facultatif spécifie le chemin du cookie.
- Domain : ce champ facultatif spécifie l'hôte/domaine sur lequel le cookie sera défini.
- Expires : ce champ spécifie l'expiration du cookie, exprimée sous la forme d'une valeur time.Time.
- MaxAge : ce champ spécifie le nombre de secondes jusqu'à ce que le cookie expire, exprimé sous la forme d'un entier.
- Secure : lorsque ce champ booléen est vrai, le client n'enverra le cookie que sur des connexions HTTPS.
- HttpOnly : lorsque ce champ booléen est vrai, le client empêchera le code JavaScript d'accéder au cookie.
- SameSite : ce champ spécifie la stratégie d'origine croisée pour le cookie à l'aide des constantes SameSite,
             qui définissent SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode et SameSiteNoneMode.

La classe Cookie est également utilisée pour obtenir l'ensemble de cookies qu'un client envoie, ce qui est fait à l'aide des méthodes Request :
- Cookie(name) : cette méthode renvoie un pointeur vers la valeur Cookie avec le nom spécifié et une erreur qui indique
                 qu'il n'y a pas de cookie correspondant.
- Cookies() : cette méthode renvoie une tranche de pointeurs Cookie.
**/

/*
*
Cet exemple configure une route /cookies, pour laquelle la fonction GetAndSetCookie définit un cookie nommé counter avec une valeur initiale de zéro.
Lorsqu'une requête contient le cookie, la valeur du cookie est lue, analysée en entier et incrémentée afin qu'elle puisse être utilisée
pour définir une nouvelle valeur de cookie. La fonction énumère également les cookies dans la requête et écrit les champs Name et Value dans la réponse.
*
*/
func GetAndSetCookie(writer http.ResponseWriter, request *http.Request) {
	counterVal := 1
	counterCookie, err := request.Cookie("counter")
	if err == nil {
		counterVal, _ = strconv.Atoi(counterCookie.Value)
		counterVal++
	}

	http.SetCookie(writer, &http.Cookie{Name: "counter", Value: strconv.Itoa(counterVal)})

	if len(request.Cookies()) > 0 {
		for _, c := range request.Cookies() {
			fmt.Fprintf(writer, "Cookie Name: %v, Value: %v", c.Name, c.Value)
		}
	} else {
		fmt.Fprintln(writer, "Request contains no cookies")
	}
}

func init() {
	http.HandleFunc("/cookies", GetAndSetCookie)
}
