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
}
