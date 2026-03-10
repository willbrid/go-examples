package main

import "io"

/**
Le code définit un type structure nommé CustomReader qui sert d'enveloppe à un reader. L'implémentation de la méthode `Read` génère un résultat
indiquant la quantité de données lues et le nombre total d'opérations de lecture.

Le constructeur NewCustomWriter encapsule un Writer dans une structure CustomWriter, qui rend compte de ses opérations d'écriture.
**/

type CustomReader struct {
	reader    io.Reader
	readCount int
}

type CustomWriter struct {
	writer     io.Writer
	writeCount int
}

func NewCustomReader(reader io.Reader) *CustomReader {
	return &CustomReader{reader: reader, readCount: 0}
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
	count, err = cr.reader.Read(slice)
	cr.readCount++
	Printfln("Custom Reader : %v bytes", count)
	if err == io.EOF {
		Printfln("Total Reads : %v", cr.readCount)
	}

	return
}

func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.writeCount++
	Printfln("Custom Writer: %v bytes", count)
	return
}

func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		err = closer.Close()
	}
	Printfln("Total Writes: %v", cw.writeCount)
	return
}
