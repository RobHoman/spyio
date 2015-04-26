package spyio

import (
	"io"
)

type Writer struct {
	Writer io.Writer
	Spy    func(p []byte)
}

func (sw *Writer) Write(p []byte) (int, error) {
	bytesRead, err := sw.Writer.Write(p)
	sw.Spy(p)

	return bytesRead, err
}
