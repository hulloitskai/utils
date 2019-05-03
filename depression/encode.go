package depression

import (
	"bytes"
	"io"
)

// An Encoder can encode data to depression format.
type Encoder struct{ dst io.Writer }

// NewEncoder creates a new Encoder.
func NewEncoder(dst io.Writer) *Encoder {
	return &Encoder{dst: dst}
}

func (e *Encoder) Write(p []byte) (n int, err error) {
	var buf bytes.Buffer
	for _, b := range p {
		for i, c := range "depresso" {
			if (b>>uint(7-i))&1 == 1 {
				buf.WriteRune(c - ('a' - 'A'))
			} else {
				buf.WriteRune(c)
			}
		}
	}
	if written, err := buf.WriteTo(e.dst); err != nil {
		return int(written) / 8, err
	}
	return len(p), nil
}
