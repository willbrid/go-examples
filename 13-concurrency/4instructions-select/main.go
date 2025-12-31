package main

import (
	"fmt"
	"time"
)

/**
Le mot-clé `select` permet de regrouper les opérations d'envoi ou de réception de données via des canaux, ce qui autorise la
création d'agencements complexes de goroutines et de canaux.
Nous pouvons combiner des instructions `case` avec des opérations d'envoi et de réception dans une même instruction `select`.
Lors de l'exécution de cette instruction, l'environnement d'exécution Go construit une liste combinée d'instructions `case` exécutables
sans blocage et en sélectionne une au hasard, qui peut être une instruction d'envoi ou de réception.
**/

func main() {
	/**
	L'utilisation la plus simple des instructions `select` consiste à recevoir des données d'un canal sans blocage, garantissant ainsi
	qu'une goroutine n'aura pas à attendre lorsque le canal est vide.

	Une instruction SELECT a une structure similaire à une instruction SWITCH, à ceci près que les instructions CASE sont des opérations
	de canal. Lors de l'exécution de l'instruction SELECT, chaque opération de canal est évaluée jusqu'à ce qu'une opération non bloquante
	soit trouvée. L'opération de canal est alors effectuée, et les instructions incluses dans l'instruction CASE sont exécutées.
	Si aucune opération de canal n'est bloquante, les instructions de la clause DEFAULT sont exécutées.

	L'instruction `SELECT` évalue ses instructions `CASE` une seule fois, d'où l'utilisation d'une boucle `for` dans l'exemple ci-dessous.
	Cette boucle exécute l'instruction `SELECT` en continu, qui reçoit les valeurs du canal dès qu'elles sont disponibles.
	Si aucune valeur n'est disponible, la clause `default` est exécutée, ce qui introduit une pause.
	L'opération sur le canal dans l'exemple ci-dessous vérifie si le canal est fermé et, le cas échéant, utilise le mot-clé `goto`
	pour accéder à une instruction étiquetée, située en dehors de la boucle `for`.

	Les délais introduits par la méthode `time.Sleep` créent un léger décalage entre le débit d'envoi des valeurs via le canal et leur
	débit de réception. Par conséquent, l'instruction `select` est parfois exécutée alors que le canal est vide. Au lieu de se bloquer,
	comme ce serait le cas lors d'une opération classique sur un canal, l'instruction `select` exécute les instructions de la clause `default`.
	Une fois le canal fermé, la boucle s'arrête.
	**/
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("#1 Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("#1 Channel has been closed")
				goto alldone1
			}
		default:
			fmt.Println("-- #1 No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone1:
	fmt.Println("#1 All values received")

	/**
	Une instruction `SELECT` peut être utilisée pour recevoir des données sans blocage. Toutefois, cette fonctionnalité s'avère plus utile
	lorsqu'il existe plusieurs canaux, par lesquels les valeurs sont transmises à des débits différents. Une instruction `SELECT` permet
	alors au récepteur d'obtenir les valeurs de n'importe quel canal qui les diffuse, sans bloquer aucun canal en particulier.

	Dans cet exemple, l'instruction `SELECT` est utilisée pour recevoir des valeurs provenant de deux canaux : l'un contenant les valeurs
	DispatchNotification et l'autre les valeurs Product. À chaque exécution de l'instruction `SELECT`, celle-ci parcourt les instructions
	`CASE`, constituant ainsi une liste de celles permettant de lire une valeur sans blocage. Une instruction `CASE` est ensuite sélectionnée
	aléatoirement dans cette liste et exécutée. Si aucune instruction `CASE` ne peut être exécutée, la clause `DEFAULT` est appliquée.
	Il convient d'être vigilant lors de la gestion des canaux fermés, car ils renvoient une valeur NIL pour toute opération de réception
	effectuée après leur fermeture, se basant sur l'indicateur de fermeture pour signaler leur inactivité. Malheureusement, cela signifie
	que les instructions `CASE` des canaux fermés seront systématiquement sélectionnées par les instructions `SELECT`, car elles sont toujours
	prêtes à fournir une valeur sans blocage, même si cette valeur est inutile.

	Si la clause par défaut est omise, l'instruction SELECT sera bloquée jusqu'à ce qu'une valeur soit reçue sur l'un des canaux.
	Cela peut s'avérer utile, mais ne gère pas les canaux susceptibles d'être fermés.

	La gestion des canaux fermés nécessite deux mesures :
	- La première consiste à empêcher l'instruction `SELECT` de sélectionner un canal une fois celui-ci fermé. Pour ce faire, on peut affecter
	la valeur nil à la variable `channel`. Un canal de valeur nil n'est jamais prêt et ne sera pas sélectionné, ce qui permet à l'instruction
	`SELECT` de passer à d'autres instructions `CASE` dont les canaux peuvent encore être ouverts.
	- La seconde mesure consiste à sortir de la boucle `for` lorsque tous les canaux sont fermés. Sans cela, l'instruction `SELECT`
	exécuterait indéfiniment la clause `default`. L'exemple ci-dessous utilise une variable `int`, décrémentée à chaque fermeture de canal.
	Lorsque le nombre de canaux ouverts atteint zéro, une instruction `goto` sort de la boucle.
	**/
	notificationChannel := make(chan DispatchNotification, 100)
	go DispatchOrders((chan<- DispatchNotification)(notificationChannel))
	productChannel := make(chan *Product)
	go EnumerateProducts((chan<- *Product)(productChannel))
	var openChannels int = 2
	for {
		select {
		case details, ok := <-notificationChannel:
			if ok {
				fmt.Println("#2 Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("#2 Dispatch channel has been closed")
				notificationChannel = nil
				openChannels--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("#2 Product :", product.Name)
			} else {
				fmt.Println("#2 Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone2
			}
			fmt.Println("-- #2 No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone2:
	fmt.Println("#2 All values received")

	/**
	Le canal présenté dans l'exemple ci-dessous est créé avec une petite mémoire tampon, et les valeurs ne sont reçues qu'après un court
	délai. Ainsi, la fonction `EnumerateWithSelectProducts` peut envoyer des valeurs via ce canal sans se bloquer tant que la mémoire tampon
	n'est pas pleine.
	**/
	productChannel1 := make(chan *Product, 5)
	go EnumerateWithSelectProducts((chan<- *Product)(productChannel1))
	time.Sleep(time.Second)
	for p := range productChannel1 {
		fmt.Println("#3 Received product :", p.Name)
	}

	/**
	Cet exemple ci-dessous comporte deux canaux avec de petits tampons. Comme pour la réception, l'instruction `select` construit
	une liste des canaux par lesquels une valeur peut être envoyée sans blocage, puis en choisit un au hasard dans cette liste.
	Si aucun canal n'est utilisable, la clause `default` est exécutée. Dans cet exemple (fonction EnumerateWithSelectCasesProducts du fichier `productdispatch.go`),
	aucune clause `default` n'est présente ; l'instruction `select` sera donc bloquée jusqu'à ce qu'un des canaux puisse recevoir une valeur.

	Les valeurs provenant du canal ne sont reçues qu'une seconde après la création de la goroutine exécutant la fonction
	`EnumerateWithSelectCasesProducts`, ce qui signifie que seuls les tampons déterminent si l'envoi vers un canal sera bloquant.
	**/
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go EnumerateWithSelectCasesProducts(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("#4 Channel 1 received product :", p.Name)
	}
	for p := range c2 {
		fmt.Println("#4 Channel 2 received product :", p.Name)
	}
}
