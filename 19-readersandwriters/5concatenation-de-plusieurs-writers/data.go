package main

import "io"

func GenerateData(writer io.Writer) {
	var end int
	data := []byte("Kayak, Lifejacket")
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end = i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v - byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error : %v", err.Error())
		}
	}

	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
}
