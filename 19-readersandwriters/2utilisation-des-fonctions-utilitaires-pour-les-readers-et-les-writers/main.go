package main

import (
	"io"
	"strings"
)

/**
Le package `io` contient un ensemble de fonctions qui offrent des moyens supplémentaires de lire et d'écrire des données.

Copy(w, r) : Cette fonction copie les données d'un reader vers un writer jusqu'à ce que la fin du fichier (EOF) soit atteinte ou qu'une autre
erreur survienne. Elle renvoie le nombre d'octets copiés et un message d'erreur décrivant le problème rencontré.

CopyBuffer(w, r, buffer) : Cette fonction effectue la même opération que `Copy`, mais lit les données dans le tampon spécifié avant de les transmettre
au writer.

CopyN(w, r, count) : Cette fonction copie `count` octets du reader vers le writer. Elle renvoie le nombre d'octets copiés et un message d'erreur
décrivant le problème rencontré.

ReadAll(r) : Cette fonction lit les données du reader spécifié jusqu'à ce que la fin du fichier (EOF) soit atteinte. Elle renvoie une tranche
d'octets contenant les données lues et un message d'erreur décrivant le problème rencontré.

ReadAtLeast(r, byteSlice, min) : Cette fonction lit au moins le nombre d'octets spécifié depuis le reader et les place dans la tranche d'octets.
Une erreur est signalée si le nombre d'octets lus est inférieur à celui spécifié.

ReadFull(r, byteSlice) : Cette fonction remplit la tranche d'octets spécifiée avec des données. Le résultat est le nombre d'octets lus et une erreur.
Une erreur est signalée si la fin du fichier (EOF) est atteinte avant que suffisamment d'octets pour remplir la tranche n'aient été lus.

WriteString(w, str) : Cette fonction écrit la chaîne de caractères spécifiée dans un writer.
**/

func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}
