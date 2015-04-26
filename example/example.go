package main

import (
	"bytes"
	"fmt"

	"robhoman/spyio"
)

func main() {
	writer := &bytes.Buffer{}

	bytesWritten := 0
	spyWriter := &spyio.Writer{
		Writer: writer,
		Spy: func(p []byte) {
			bytesWritten += len(p)
		},
	}

	spyWriter.Write([]byte("hello"))
	spyWriter.Write([]byte("world"))

	fmt.Println(bytesWritten)
}
