package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/**
Le paquet net/http/cookiejar contient une implémentation de l'interface CookieJar qui stocke les cookies en mémoire.
Les Cookiejar sont créées avec une fonction constructeur :
- New(options) : cette fonction crée un nouveau CookieJar, configuré avec une classe Options.
                 La fonction renvoie également une erreur qui signale des problèmes lors de la création du CookieJar.
La fonction New accepte une clase net/http/cookiejar/Options, qui est utilisée pour configurer le cookiejar.
Il n'y a qu'un seul champ Options, PublicSuffixList, qui est utilisé pour spécifier une implémentation de l'interface avec le même nom,
qui fournit un support pour empêcher les cookies d'être définis trop largement, ce qui peut entraîner des violations de la vie privée.
Si plusieurs valeurs Client sont requises mais que les cookies doivent être partagés, un seul CookieJar peut être utilisé.
**/

func init() {
	http.HandleFunc("/cookie", func(writer http.ResponseWriter, request *http.Request) {
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
	})
}
