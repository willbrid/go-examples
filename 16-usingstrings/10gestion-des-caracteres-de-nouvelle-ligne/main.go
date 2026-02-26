package main

import "fmt"

/**
Le package `fmt` fournit des fonctions pour l'analyse de chaînes de caractères, c'est-à-dire le processus d'analyse syntaxique de chaînes
contenant des valeurs séparées par des espaces.

Scan(...vals) : Cette fonction lit du texte depuis l'entrée standard et stocke les valeurs séparées par des espaces dans les arguments spécifiés.
Les sauts de ligne sont interprétés comme des espaces, et la fonction lit jusqu'à ce qu'elle ait reçu des valeurs pour tous ses arguments.
Le résultat est le nombre de valeurs lues et un message d'erreur décrivant les éventuels problèmes.

Scanln(...vals) : Cette fonction fonctionne de la même manière que `Scan`, mais arrête la lecture lorsqu'elle rencontre un caractère de saut de ligne.

Scanf(template, ...vals) : Cette fonction fonctionne de la même manière que `Scan`, mais utilise une chaîne de caractères modèle pour sélectionner
les valeurs de l'entrée reçue.

Fscan(reader, ...vals) : Cette fonction lit des valeurs séparées par des espaces depuis le reader spécifié. Les sauts de ligne sont interprétés
comme des espaces, et la fonction renvoie le nombre de valeurs lues et un message d'erreur décrivant les éventuels problèmes.

Fscanln(reader, ...vals) : Cette fonction fonctionne de la même manière que `Fscan`, mais interrompt la lecture dès qu'elle rencontre
un caractère de nouvelle ligne.

Fscanf(reader, template, ...vals) : Cette fonction fonctionne de la même manière que `Fscan`, mais utilise un modèle pour sélectionner
les valeurs de la chaîne de caractères reçue.

Sscan(str, ...vals) : Cette fonction analyse la chaîne de caractères spécifiée à la recherche de valeurs séparées par des espaces,
qui sont affectées aux arguments restants. Le résultat est le nombre de valeurs analysées et un message d'erreur décrivant les éventuels problèmes.

Sscanf(str, template, ...vals) : Cette fonction fonctionne de la même manière que `Sscan`, mais utilise un modèle pour sélectionner les valeurs
de la chaîne.

Sscanln(str, template, ...vals) : Cette fonction fonctionne de la même manière que `Sscanf`, mais interrompt l'analyse de la chaîne dès
qu'elle rencontre un caractère de nouvelle ligne.


Le choix de la fonction d'analyse à utiliser dépend de la source de la chaîne à analyser, de la manière dont les sauts de ligne sont gérés et
de l'utilisation éventuelle d'un modèle.


La fonction Scan continue de rechercher des valeurs jusqu'à ce qu'elle ait reçu le nombre attendu, et la première pression sur la touche `Entrée`
est interprétée comme un séparateur et non comme la fin de la saisie. Les fonctions dont le nom se termine par ln, comme Scanln, modifient
ce comportement.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	var (
		name     string
		category string
		price    float64
	)

	/**
	Lorsque nous appuyons pour la première fois sur la touche `Entrée`, le saut de ligne interrompt la saisie, laissant la fonction `Scanln` avec
	moins de valeurs que nécessaire.
	**/
	fmt.Print("Enter text to scan : ")
	n, err := fmt.Scanln(&name, &category, &price)

	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error :%v", err.Error())
	}
}
