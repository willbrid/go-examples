package main

import (
	"fmt"
	"io"
	"strings"
)

/**
le package `fmt` permet d'appliquer les fonctionnalités de scan aux readers et aux writers.

Lecture des valeurs d'un reader
Le package `fmt` fournit des fonctions permettant de lire les valeurs d'un reader et de les convertir en différents types.
Dans l'exemple #0 ci-dessous le processus d'analyse lit les octets du reader et utilise le modèle d'analyse pour traiter les données reçues.
Ce modèle contient deux chaînes de caractères et une valeur float64.

Une technique utile lors de l'utilisation d'un reader consiste à analyser les données progressivement à l'aide d'une boucle.
Cette approche est particulièrement efficace lorsque les octets arrivent au fil du temps, comme lors de la lecture d'une connexion HTTP.


Écriture de chaînes formatées dans une valeur `Writer`
Le package `fmt` fournit également des fonctions pour écrire des chaînes formatées dans une valeur `Writer`.


Utilisation d'un Replacer avec un Writer
La structure `strings.Replacer` peut être utilisée pour effectuer des remplacements sur une chaîne de caractères et afficher le résultat
modifié dans un Writer.
**/

func scanFromReader(reader io.Reader, template string, vals ...any) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val any) (int, error) {
	return fmt.Fscan(reader, val)
}

// La fonction writeFormatted utilise la fonction `fmt.Fprintf` pour écrire une chaîne formatée avec un modèle dans une valeur Writer.
func writeFormatted(writer io.Writer, template string, vals ...any) {
	fmt.Fprintf(writer, template, vals...)
}

// La méthode `WriteString` de la structure `strings.Replacer` effectue ses substitutions et écrit la chaîne modifiée.
func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func main() {
	Printfln("Exemple #0.......")
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("#1 Error: %v", err.Error())
	} else {
		Printfln("#1 Name: %v", name)
		Printfln("#1 Category: %v", category)
		Printfln("#1 Price: %.2f", price)
	}

	Printfln("Exemple #1.......")
	/**
	La boucle for appelle la fonction `scanSingle`, qui utilise la fonction `Fscan` pour lire une chaîne de caractères depuis le reader.
	La lecture se poursuit jusqu'à ce que l'erreur `EOF` soit renvoyée, après quoi la boucle s'arrête.
	**/
	reader1 := strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader1, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("#1 Value: %v", str)
	}

	Printfln("Exemple #2.......")
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())

	Printfln("Exemple #3.......")
	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer1 strings.Builder
	writeReplaced(&writer1, text, subs...)
	fmt.Println(writer1.String())
}
