package main

/**
La fonction `os.Open` ouvre un fichier en lecture et renvoie une valeur `File`, qui représente le fichier ouvert, et une erreur, qui est utilisée
pour indiquer des problèmes d'ouverture du fichier. La structure `File` implémente l'interface `Reader`, qui simplifie la lecture et le traitement
des exemples de données JSON, sans lire l'intégralité du fichier dans une slice d'octets.

Le package `os` définit trois variables `*File`, nommées `Stdin`, `Stdout` et `Stderr`, qui donnent accès à l'entrée standard, à la sortie standard
et à l'erreur standard.

La structure `File` définit des méthodes au-delà de celles requises par l'interface `Reader` qui permettent d'effectuer des lectures à un emplacement
spécifique du fichier.
- ReadAt(slice, offset) : Cette méthode est définie par l'interface `ReaderAt` et effectue une lecture dans la slice spécifique à la position `offset`
spécifiée dans le fichier.

- Seek(offset, how) : Cette méthode est définie par l'interface `Seeker` et déplace le décalage dans le fichier pour la prochaine lecture.
Le décalage est déterminé par la combinaison des deux arguments : le premier argument spécifie le nombre d'octets à décaler et le deuxième argument
détermine la manière dont le décalage est appliqué : une valeur de 0 signifie que le décalage est relatif au début du fichier, une valeur de 1
signifie que le décalage est relatif à la position de lecture actuelle et une valeur de 2 signifie que le décalage est relatif à la fin du fichier.
**/

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
}
