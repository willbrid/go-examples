package main

import (
	"bufio"
	"fmt"
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

func processData2(reader io.Reader, writer io.Writer) {
	/**
	Cette fonction io.Copy copie les données d'un Reader vers un Writer jusqu'à ce qu'EOF soit renvoyé ou qu'une autre erreur soit rencontrée.
	Les résultats sont le nombre de copies d'octets et une erreur utilisée pour décrire tout problème.
	**/
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func processData3(reader io.Reader, writer io.Writer) {
	var buffer []byte
	/**
	Cette fonction io.CopyBuffer effectue la même tâche que Copy mais lit les données dans le tampon spécifié avant
	qu'elles ne soient transmises à writer.
	**/
	count, err := io.CopyBuffer(writer, reader, buffer)
	if err == nil {
		Printfln("Read %v bytes", count)
		Printfln("Len of buffer : %v", len(buffer))
	} else {
		Printfln("Error : %v", err.Error())
	}
}

func processData4(reader io.Reader, writer io.Writer) {
	var num int64 = 4
	/**
	Cette fonction io.CopyN copie le nombre d'octets du Reader vers le Writer. Les résultats sont le nombre de copies d'octets et
	une erreur utilisée pour décrire tout problème. Le 3ème paramètre permet de spécifier le nombre de bytes à copier.
	**/
	count, err := io.CopyN(writer, reader, num)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error : %v", err.Error())
	}
}

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	/**
	La méthode fmt.Fscanf analyse le texte lu à partir du reader, stockant les valeurs successives séparées par des espaces dans des arguments
	successifs déterminés par le format. Il renvoie le nombre d'éléments analysés avec succès. Les retours à la ligne dans l'entrée doivent
	correspondre aux retours à la ligne dans le format.
	**/
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	/**
	La méthode fmt.Fscan analyse le texte lu à partir du reader, stockant les valeurs successives séparées par des espaces dans des arguments successifs.
	Les retours à la ligne comptent comme un espace. Il renvoie le nombre d'éléments analysés avec succès.
	Si c'est moins que le nombre d'arguments, err indiquera pourquoi.
	**/
	return fmt.Fscan(reader, val)
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	// La méthode fmt.Fprintf formate selon un spécificateur de format (template) et écrit dans writer.
	// Il renvoie le nombre d'octets écrits et toute erreur d'écriture rencontrée.
	fmt.Fprintf(writer, template, vals...)
}

/*
*
La classe strings.Replacer peut être utilisée pour effectuer des remplacements sur une chaîne et envoyer le résultat modifié à un Writer.
La méthode WriteString effectue ses substitutions et écrit la chaîne modifiée.
*
*/
func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	// La méthode WriteString écrit la chaine str dans le writer avec tous les remplacements effectués.
	replacer.WriteString(writer, str)
}

func main() {
	Printfln("Product : %v, Price : %v", kayak.Name, kayak.Price)

	// Le package strings fournit une fonction constructeur NewReader, qui accepte une chaîne comme argument.
	r := strings.NewReader("Kayak")
	processData(r)

	// Le résultat de la fonction strings.NewReader est un pointeur : *strings.Reader.
	r1 := strings.NewReader("Kayak")
	var builder1 strings.Builder
	processData1(r1, &builder1)
	Printfln("String builder contents 1 : %s", builder1.String())

	r2 := strings.NewReader("Kayak")
	var builder2 strings.Builder
	processData2(r2, &builder2)
	Printfln("String builder contents 2 : %s", builder2.String())

	r3 := strings.NewReader("Kayak")
	var builder3 strings.Builder
	processData3(r3, &builder3)
	Printfln("String builder contents 3 : %s", builder3.String())

	r4 := strings.NewReader("Kayak")
	var builder4 strings.Builder
	processData4(r4, &builder4)
	Printfln("String builder contents 4 : %s", builder4.String())

	r5 := strings.NewReader("Kayak")
	dataRead5, err5 := io.ReadAll(r5)
	if err5 == nil {
		Printfln("Len of dataRead5 : %v", len(dataRead5))
	} else {
		Printfln("Error : %v", err5.Error())
	}

	/**
	Cette fonction io.Pipe renvoie un PipeReader et un PipeWriter, qui peuvent être utilisés pour connecter des fonctions nécessitant un Reader et un Writer
	**/
	pipeReader1, pipeWriter1 := io.Pipe()
	/**
	Les canaux sont synchrones, de sorte que la méthode PipeWriter.Write se bloquera jusqu'à ce que les données soient lues à partir du canal.
	Cela signifie que le PipeWriter doit être utilisé dans une goroutine différente du reader pour éviter que l'application ne se bloque.
	**/
	go func() {
		GenerateData(pipeWriter1)
		// Cette méthode pipeWriter.Close ferme le Reader ou le Writer.
		pipeWriter1.Close()
	}()
	ConsumeData(pipeReader1)

	pipeReader2, pipeWriter2 := io.Pipe()
	go GenerateDataImprove(pipeWriter2)
	ConsumeData(pipeReader2)

	r6 := strings.NewReader("Kayak")
	r7 := strings.NewReader("Lifejacket")
	r8 := strings.NewReader("Canoe")
	/**
	La fonction MultiReader concentre les entrées de plusieurs Reader afin qu'ils puissent être traités en séquence.
	Le Reader renvoyé par la fonction MultiReader répond à la méthode Read avec le contenu de l'une des valeurs Reader sous-jacentes.
	Lorsque le premier Reader renvoie EOF, le contenu est lu à partir du deuxième Reader. Ce processus se poursuit jusqu'à ce que le
	Reader sous-jacent final renvoie EOF
	**/
	concatReader := io.MultiReader(r6, r7, r8)
	ConsumeData(concatReader)

	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	/**
	La fonction MultiWriter combine plusieurs Writer afin que les données soient envoyées à chacun d'eux.
	Les MultiWriter dans cet exemple sont des valeurs string.Builder, qui étaient et qui implémentent l'interface Writer.
	La fonction MultiWriter est utilisée pour créer un Writer, de sorte que l'appel de la méthode Write entraînera l'écriture
	des mêmes données dans les trois Writer individuels.
	**/
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateDataImprove(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

	r9 := strings.NewReader("Kayak")
	r10 := strings.NewReader("Lifejacket")
	r11 := strings.NewReader("Canoe")
	concatReader1 := io.MultiReader(r9, r10, r11)
	var writer1 strings.Builder
	/**
	La fonction TeeReader renvoie un Reader qui renvoie les données qu'il reçoit à un Writer.
	La fonction TeeReader est utilisée pour créer un Reader qui renverra les données à un strings.Builder et qui implémente l'interface Writer.
	**/
	teeReader := io.TeeReader(concatReader1, &writer1)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer1.String())

	r12 := strings.NewReader("Kayak")
	r13 := strings.NewReader("Lifejacket")
	r14 := strings.NewReader("Canoe")
	concatReader2 := io.MultiReader(r12, r13, r14)
	/**
	La fonction LimitReader est utilisée pour limiter la quantité de données pouvant être obtenues à partir d'un Reader.
	Le premier argument de la fonction LimitReader est le Reader qui fournira les données. Le deuxième argument est
	le nombre maximum d'octets pouvant être lus. Le Reader renvoyé par la fonction LimitReader enverra EOF lorsque
	la limite est atteinte, à moins que le Reader sous-jacent n'envoie EOF en premier.
	**/
	limited := io.LimitReader(concatReader2, 5)
	ConsumeData(limited)

	/**
		C'est la taille de la tranche d'octet transmise à la fonction Read qui détermine la façon dont les données sont consommées.
		Dans ce cas, la taille de la tranche est de 5, ce qui signifie qu'un maximum de 5 octets est lu pour chaque appel à la fonction Read.
		Deux lectures n'ont pas obtenu 5 octets de données. L'avant-dernière lecture a produit 3 octets car les données source ne sont pas parfaitement
		divisibles par cinq et il restait trois octets de données. La lecture finale a renvoyé 0 octet mais a reçu l'erreur EOF, indiquant que
		la fin des données avait été atteinte. Au total, la lecture de 28 octets a nécessité 7 lectures.

		La lecture de petites quantités de données peut être problématique lorsqu'il y a une grande quantité de surcharge associée à chaque opération.
		Ce n'est pas un problème lors de la lecture d'une chaîne stockée en mémoire, mais la lecture de données à partir d'autres sources de données,
		telles que des fichiers, peut être plus coûteuse et il peut être préférable d'effectuer un plus petit nombre de lectures plus importantes.
		Cela se fait en introduisant un tampon dans lequel une grande quantité de données est lue pour répondre à plusieurs demandes de données plus petites.
	**/
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer2 strings.Builder
	slice := make([]byte, 5)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer2.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer2.String())

	/**
	La fonction NewReader utilisée, qui crée un reader avec la taille de tampon par défaut. Le reader mis en mémoire tampon remplit
	sa mémoire tampon et utilise les données qu'il contient pour répondre aux appels à la méthode Read.

	La taille de la mémoire tampon par défaut est de 4 096 octets, ce qui signifie que le reader mis en mémoire tampon a pu lire toutes les données
	en une seule opération de lecture, plus une lecture supplémentaire pour produire le résultat EOF.
	L'introduction de la mémoire tampon réduit la surcharge associée aux opérations de lecture, mais au détriment de la mémoire utilisée
	pour mettre les données en mémoire tampon.

	Les fonctions NewReader et NewReaderSize renvoient des valeurs bufio.Reader, qui implémentent l'interface io.Reader et
	qui peuvent être utilisées comme wrappers pour d'autres types de méthodes Reader.
	**/
	text1 := "It was a boat. A small boat."
	var reader1 io.Reader = NewCustomReader(strings.NewReader(text1))
	var writer3 strings.Builder
	slice1 := make([]byte, 5)
	// Cette fonction bufio.NewReader renvoie un reader tamponné avec la taille de tampon par défaut (qui est de 4 096 octets au moment de l'écriture).
	reader1 = bufio.NewReader(reader1)
	for {
		count1, err1 := reader1.Read(slice1)
		if count1 > 0 {
			writer3.Write(slice1[0:count1])
		}
		if err1 != nil {
			break
		}
	}
	Printfln("Read data: %v", writer3.String())

	text2 := "It was a boat. A small boat."
	var reader2 io.Reader = NewCustomReader(strings.NewReader(text2))
	var writer4 strings.Builder
	slice2 := make([]byte, 5)
	buffered := bufio.NewReader(reader2)
	for {
		count2, err2 := buffered.Read(slice2)
		if count2 > 0 {
			/**
			Cette méthode buffered.Buffered renvoie un entier qui indique le nombre d'octets pouvant être lus à partir du tampon.
			Cette méthode buffered.Size renvoie la taille du tampon, exprimée en entier.
			Cette méthode buffered.Discard(count) ignore le nombre d'octets spécifié.
			Cette méthode Peek(count) renvoie le nombre d'octets spécifié sans les supprimer de la mémoire tampon, ce qui signifie qu'ils seront renvoyés par
			les appels ultérieurs à la méthode Read.
			Cette méthode Reset(reader) supprime les données dans la mémoire tampon et effectue les lectures suivantes à partir du Reader spécifié.
			**/
			Printfln("Buffer size : %v, buffered : %v", buffered.Size(), buffered.Buffered())
			writer4.Write(slice2[0:count2])
		}
		if err2 != nil {
			break
		}
	}
	Printfln("Read data: %v", writer4.String())

	text3 := "It was a boat. A small boat."
	var builder5 strings.Builder
	var writer5 = NewCustomWriter(&builder5)
	for i := 0; true; {
		end := i + 5
		if end >= len(text3) {
			writer5.Write([]byte(text3[i:]))
			break
		}
		writer5.Write([]byte(text3[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder5.String())

	/**
	Le Writer mis en mémoire tampon conserve les données dans un tampon et les transmet au Writer sous-jacent uniquement
	lorsque le tampon est plein ou lorsque la méthode Flush est appelée.
	La transition vers un Writer tamponné n'est pas entièrement transparente car il est important d'appeler la méthode Flush
	pour s'assurer que toutes les données sont écrites.
	**/
	text4 := "It was a boat. A small boat."
	var builder6 strings.Builder
	// Cette fonction bufio.NewWriterSize renvoie un Writer mis en mémoire tampon avec la taille de mémoire tampon spécifiée.
	var writer6 *bufio.Writer = bufio.NewWriterSize(NewCustomWriter(&builder6), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text4) {
			writer6.Write([]byte(text4[i:]))
			// Cette méthode writer.Flush() écrit le contenu du tampon dans le Writer sous-jacent.
			writer6.Flush()
			break
		}
		writer6.Write([]byte(text4[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder6.String())

	/**
	Le processus d'analyse lit les octets du reader et utilise le template de scan pour analyser les données reçues.
	Le template de scan contient deux chaînes et une valeur float64.
	**/
	var (
		reader3        io.Reader = strings.NewReader("Kayak Watersports $279.00")
		name, category string
		price          float64
		scanTemplate   string = "%s %s $%f"
	)
	_, err := scanFromReader(reader3, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}

	/**
	La boucle for appelle la fonction scanSingle, qui utilise la fonction Fscan pour lire une chaîne à partir du Reader.
	Les valeurs sont lues jusqu'à ce que EOF soit retourné, moment auquel la boucle est terminée.
	**/
	var reader4 io.Reader = strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader4, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error : %v", err.Error())
			}
			break
		}
		Printfln("Value : %v", str)
	}

	// La fonction writeFormatted utilise la fonction fmt.Fprintf pour écrire une chaîne formatée avec un template dans un Writer.
	var writer7 strings.Builder
	template := "Name : %s, Category : %s, Price : $%.2f"
	writeFormatted(&writer7, template, "Kayak", "Watersports", float64(279))
	Printfln("Value from writer : %v", writer7.String())

	text5 := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer8 strings.Builder
	writeReplaced(&writer8, text5, subs...)
	Printfln("Value from replacer : %v", writer8.String())
}
