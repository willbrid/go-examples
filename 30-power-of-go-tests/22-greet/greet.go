package greet

import (
	"bufio"
	"fmt"
	"io"
)

func Greet(in io.Reader, out io.Writer) {
	fmt.Fprint(out, "Your name? ")
	scanner := bufio.NewScanner(in)
	if !scanner.Scan() {
		return
	}
	fmt.Fprintf(out, "Hello, %s!\n", scanner.Text())
}

/**
bufio.NewScanner prend un flux d'entrée (un io.Reader) et renvoie un objet scanner qui lira le flux. Nous pouvons ensuite appeler sa méthode
Scan pour lui demander d'effectuer une analyse. Si une ligne a été analysée avec succès, la méthode renvoie true, et nous pouvons alors
récupérer la ligne en appelant Text().
Le problème ici est de savoir quoi faire si Scan renvoie false, ce qui signifie qu'aucune entrée n'a été lue. Or, nous ne pouvons pas
renvoyer une erreur de Greet sans modifier sa signature, et nous ne voulons pas que tous ceux qui appellent cette fonction aient à vérifier
une erreur, ce qui se produira rarement dans le programme réel : ce serait irritant.
Il est également absurde de choisir un nom par défaut si nous ne pouvons pas en lire un de la part de l'utilisateur. Dans ce cas,
la meilleure solution est probablement de ne rien faire et de renvoyer la ligne.
**/
