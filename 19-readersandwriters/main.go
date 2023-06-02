package main

import (
	"io"
	"strings"
)

func processData(reader io.Reader) {
	/**
	Nous précisons le nombre maximum d'octets que nous souhaitons recevoir en définissant la taille de
	la tranche (slice) d'octets transmise à la fonction Read.
	**/
	b := make([]byte, 2)
	for {
		/**
		Cette méthode reader.Read lit les données dans []byte. La méthode renvoie le nombre d'octets lus, exprimé sous la forme d'un int, et une erreur.
		**/
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read : %v - bytes : %v", count, string(b[0:count]))
		}
		/**
		Le paquet io définit une erreur spéciale nommée EOF, qui est utilisée pour signaler quand le Reader atteint la fin des données.
		**/
		if err == io.EOF {
			break
		}
	}
}

/*
*
En règle générale, les méthodes Reader et Writer sont implémentées pour les pointeurs afin que le passage d'un Reader
ou d'un Writer à une fonction ne crée pas de copie. Nous n'avons pas eu à utiliser l'opérateur d'adresse pour le Reader car
le résultat de la fonction strings.NewReader est un pointeur : *strings.Reader.
*
*/
func processData1(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			/**
			Cette méthode writer.Write écrit les données à partir de la tranche d'octets spécifiée. La méthode renvoie le nombre d'octets
			qui ont été écrits et une erreur. L'erreur sera non nulle si le nombre d'octets écrits est inférieur à la longueur de la tranche.
			**/
			writer.Write(b[0:count])
			Printfln("Read : %v - bytes : %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	Printfln("Product : %v, Price : %v", kayak.Name, kayak.Price)

	// Le package strings fournit une fonction constructeur NewReader, qui accepte une chaîne comme argument.
	r := strings.NewReader("Kayak")
	processData(r)

	// Le résultat de la fonction strings.NewReader est un pointeur : *strings.Reader.
	r1 := strings.NewReader("Kayak")
	var builder strings.Builder
	processData1(r1, &builder)
	Printfln("String builder contents : %s", builder.String())
}
