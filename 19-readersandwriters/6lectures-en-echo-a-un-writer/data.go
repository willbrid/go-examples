package main

import "io"

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
