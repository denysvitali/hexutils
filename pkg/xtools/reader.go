package xtools

import (
	"bufio"
	"io"
)

type Reader struct {
	r        *bufio.Reader
	lineSize uint
	addr     uint

	styleOptions StyleOptions
}

func New(input io.Reader) Reader {
	return Reader{
		r:            bufio.NewReader(input),
		lineSize:     8 * 2,
		styleOptions: DefaultStyle,
	}
}

func (r *Reader) ReadLine() ([]byte, int, error) {
	buffer := make([]byte, r.lineSize)
	read, err := r.r.Read(buffer)
	if err != nil {
		return nil, 0, err
	}
	return buffer, read, nil
}
