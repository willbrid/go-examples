package main

import (
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

/*
*
Cette méthode ServeHTTP est appelée pour traiter une requête HTTP. La requête est décrite par une valeur Request et la réponse est écrite à l'aide
d'un ResponseWriter, tous deux reçus en tant que paramètres.

Le navigateur effectue deux requêtes HTTP. Le premier est pour /, qui est le composant de chemin de l'URL demandée.
La deuxième demande concerne /favicon.ico, que le navigateur envoie pour obtenir une icône à afficher en haut de la fenêtre ou de l'onglet.
*
*/
func (sh StringHandler) ServeHTTP1(writer http.ResponseWriter, request *http.Request) {
	Printfln("Method : %v", request.Method)
	Printfln("URL : %v", request.URL)
	Printfln("HTTP version : %v", request.Proto)
	Printfln("Host : %v", request.Host)
	for name, val := range request.Header {
		Printfln("Header : %v, Value : %v", name, val)
	}
	Printfln("---")
	io.WriteString(writer, sh.message)
}

/*
*
Champs et méthodes utiles définis par la classe request.URL :

  - Scheme : ce champ renvoie le composant de schéma de l'URL.

  - Host : ce champ renvoie le composant hôte de l'URL, qui peut inclure le port.

  - RawQuery : ce champ renvoie la chaîne de requête à partir de l'URL. Utilisons la méthode Query pour traiter la chaîne de requête dans une map.

  - Path : ce champ renvoie le composant de chemin de l'URL.

  - Fragment : ce champ renvoie le composant fragment de l'URL, sans le caractère #.

  - Hostname() : cette méthode renvoie le composant de nom d'hôte de l'URL sous forme de chaîne.

  - Port() : cette méthode renvoie le composant port de l'URL sous forme de chaîne.

  - Query() : cette méthode renvoie une map[string][]string (une map avec des clés de chaîne et des valeurs de tranche de chaîne),
    contenant les champs de chaîne de requête.

  - User() : cette méthode renvoie les informations de l'utilisateur associées à la requête

  - String() : cette méthode renvoie une représentation sous forme de chaîne de l'URL.

    Les méthodes ResponseWriter (http.ResponseWriter):

  - Header() : cette méthode renvoie un Header, qui est un alias de map[string][]string, qui peut être utilisé pour définir les en-têtes de réponse.

  - WriteHeader(code) : cette méthode définit le code d'état de la réponse, spécifié sous la forme d'un int. Le package net/http définit
    des constantes pour la plupart des codes d'état.

  - Write(data) : cette méthode écrit des données dans le corps de la réponse et implémente l'interface Writer.

*
*/
func (sh StringHandler) ServeHTTP2(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - returning 404")
		/**
		Cette méthode writer.WriteHeader définit le code d'état de la réponse, spécifié sous la forme d'un int. Le package net/http définit
		des constantes pour la plupart des codes d'état.
		**/
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

/*
*
Le package net/http fournit un ensemble de fonctions pratiques qui peuvent être utilisées pour créer des réponses communes aux requêtes HTTP.

  - http.Error(writer, message, code) : cette fonction définit l'en-tête sur le code spécifié, définit l'en-tête Content-Type sur text/plain et
    écrit le message d'erreur dans la réponse. L'en-tête X-Content-Type-Options est également défini pour empêcher les navigateurs
    d'interpréter la réponse comme autre chose que du texte.

  - http.NotFound(writer, request) Cette fonction appelle Error et spécifie un code d'erreur 404.

  - http.Redirect(writer, request, url, code) : cette fonction envoie une réponse de redirection à l'URL spécifiée et avec le code d'état spécifié.

  - http.ServeFile(writer, request, fileName) Cette fonction envoie une réponse contenant le contenu du fichier spécifié.
    L'en-tête Content-Type est défini en fonction du nom de fichier mais peut être remplacé en définissant explicitement
    l'en-tête avant d'appeler la fonction.

*
*/
func (sh StringHandler) ServeHTTP3(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
	io.WriteString(writer, sh.message)
}

/*
*
Le processus d'inspection de l'URL et de sélection d'une réponse peut produire un code complexe difficile à lire et à gérer.
Pour simplifier le processus, le package net/http fournit une implémentation de gestionnaire qui permet de séparer l'URL de la production d'une requête.
*
*/
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}

	/**
	Le package net/http facilite la création d'un serveur HTTP simple, qui peut ensuite être étendu pour ajouter des
	fonctionnalités plus complexes et utiles.

	Cette fonction http.ListenAndServe commence à écouter les requêtes HTTP sur une adresse et transmet les requêtes au gestionnaire (http.handler) spécifié.

	Aucun nom ou adresse n'est spécifié et le numéro de port suit deux-points, ce qui signifie que cette instruction crée un serveur HTTP
	qui écoute les demandes sur le port 5000 sur toutes les interfaces. Lorsqu'une requête arrive, elle est transmise à un gestionnaire (handler),
	qui est chargé de produire une réponse. Les gestionnaires (handlers) doivent implémenter l'interface Handler.

	Les requêtes HTTP sont représentées par la classe Request, définie dans le package net/http :
	- Method : ce champ fournit la méthode HTTP (GET, POST, etc.) sous forme de chaîne. Le package net/http définit des constantes pour les méthodes HTTP,
	  telles que MethodGet et MethodPost.
	- URL : ce champ renvoie l'URL demandée, exprimée sous la forme d'une valeur d'URL.
	- Proto : ce champ renvoie une chaîne qui indique la version de HTTP utilisée pour la requête.
	- Host : ce champ renvoie une chaîne contenant l'hôte demandé.
	- Header : ce champ renvoie une valeur d'en-tête, qui est un alias de map[string][]string et contient les en-têtes de requête.
	  Les clés de map sont les noms des en-têtes et les valeurs sont des tranches de chaîne contenant les valeurs d'en-tête.
	- Trailer : ce champ renvoie une chaîne map[string] qui contient tous les en-têtes supplémentaires inclus dans la requête après le corps.
	- Body : ce champ renvoie un ReadCloser, qui est une interface qui combine la méthode Read de l'interface Reader avec
	  la méthode Close de l'interface Closer.
	**/
	/**
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, world !"})
	if err != nil {
		Printfln("Error : %v", err.Error())
	}
	**/

	/**
	Cette fonction http.Handle crée une règle qui appelle la méthode ServeHTTP spécifiée du Hander spécifié pour les requêtes qui correspondent au modèle.
	La clé de cette fonctionnalité est d'utiliser nil comme argument de la fonction ListenAndServe.

	HandleFunc(pattern, handlerFunc) : cette fonction crée une règle qui appelle la fonction spécifiée pour les requêtes qui correspondent au modèle.
	La fonction est appelée avec les arguments ResponseWriter et Request.

	Les fonctions net/http pour créer des gestionnaires de requêtes :
	- FileServer(root) : cette fonction crée un gestionnaire qui produit des réponses à l'aide de la fonction ServeFile.
	- NotFoundHandler() : cette fonction crée un gestionnaire qui produit des réponses à l'aide de la fonction NotFound.
	- RedirectHandler(url, code) - cette fonction crée un Handler qui produit des réponses à l'aide de la fonction Redirect.
	- StripPrefix(prefix, handler) - cette fonction crée un Handler qui supprime le préfixe spécifié de l'URL de la requête et
	  transmet la requête au Handler spécifié.
	- TimeoutHandler(handler, duration, message) : cette fonction transmet la requête au Handler spécifié mais génère une réponse d'erreur
	  si la réponse n'a pas été produite dans le délai spécifié.
	**/
	http.Handle("/message", StringHandler{message: "Hello, world !"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	/**
	La fonction FileServer crée un gestionnaire qui servira les fichiers et le répertoire est spécifié à l'aide de la fonction Dir.
	(Il est possible de servir des fichiers directement, mais la prudence s'impose car il est facile d'autoriser les requêtes à sélectionner
	des fichiers en dehors du dossier cible. L'option la plus sûre consiste à utiliser la fonction Dir

	Nous servons le contenu dans le dossier statique avec des chemins d'URL qui commencent par des fichiers afin qu'une requête pour /files/store.html,
	par exemple, soit traitée à l'aide du fichier static/store.html. Pour ce faire, nous utilisons la fonction StripPrefix, qui crée un gestionnaire
	qui supprime un préfixe de chemin et transmet la requête à un autre gestionnaire de service.
	**/
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))

	// La fonction ListenAndServeTLS est utilisée pour activer HTTPS, où les arguments supplémentaires spécifient les fichiers de certificat
	// et de clé privée, qui sont nommés certificate.cer et certificate.key
	go func() {
		errHttps := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.key", nil)
		if errHttps != nil {
			Printfln("HTTPS Error : %v", errHttps.Error())
		}
	}()

	// La clé de cette fonctionnalité est d'utiliser nil comme argument de la fonction ListenAndServe
	/**
		Le bloc de fonctions ListenAndServeTLS et ListenAndServe : une goroutine est utilisé pour prendre en charge
		les requêtes HTTP et HTTPS, avec HTTP géré sur le port 5000 et HTTPS sur le port 5500. Les fonctions ListenAndServeTLS et ListenAndServe
		ont été invoquées avec nil comme gestionnaire, ce qui signifie que les requêtes HTTP et HTTPS seront traitées à l'aide du même ensemble de routes.
	**/
	/**
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error : %v", err.Error())
	}
	**/
	/**
	 	Cette fonction http.HandlerFunc crée une règle qui appelle la fonction spécifiée pour les requêtes qui correspondent au modèle.
		La fonction est appelée avec les arguments ResponseWriter et Request.
		Le gestionnaire http.HandlerFunc(HTTPSRedirect) pour HTTP redirige le client vers l'URL HTTPS

	 **/
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error : %v", err.Error())
	}
}
