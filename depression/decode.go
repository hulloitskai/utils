package depression

import (
	"bytes"
	"io"
)

// An Decoder can encode data to depression format.
type Decoder struct{ dst io.Writer }

// NewDecoder creates a new Decoder.
func NewDecoder(dst io.Writer) *Decoder {
	return &Decoder{dst: dst}
}

func (d *Decoder) Write(p []byte) (n int, err error) {
	buf := bytes.NewBuffer(make([]byte, len(p)))
	for i := 0; i < len(p); i += 8 {
		var b byte
		for j := i; j < (i + 8); j++ {
			if j != 0 {
				b = b << 1
			}
			if p[j] <= 'Z' {
				b++
			}
		}
		buf.WriteByte(b)
	}
	if n, err := buf.WriteTo(d.dst); err != nil {
		return int(n) * 8, err
	}
	return len(p), nil
}
