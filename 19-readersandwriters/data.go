package main

import "io"

func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v - byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error : %v", err.Error())
		}
	}
}

func GenerateDataImprove(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v - byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error : %v", err.Error())
		}
	}
	/**
		C'est mieux de vérifier si un Writer implémente l'interface Closer dans le code qui produit les données.
		Cette approche fournit des gestionnaires cohérents de Writers qui définissent une méthode Close, qui inclut certains des types les plus utiles.
		Cela permet également de changer la goroutine afin qu'elle exécute la fonction GenerateData sans avoir besoin d'une fonction anonyme

	**/
	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
}

func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	Printfln("Read data: %v", string(data))
}
