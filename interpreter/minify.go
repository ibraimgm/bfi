package interpreter

import (
	"io"

	"github.com/ibraimgm/bfi/interpreter/token"
)

type minifyReader struct {
	r io.Reader
}

func (m *minifyReader) Read(buffer []byte) (int, error) {
	n, err := m.r.Read(buffer)
	valid := 0

	for i := 0; i < n; i++ {
		if token.IsValid(rune(buffer[i])) {
			buffer[valid] = buffer[i]
			valid++
		}
	}

	return valid, err
}

// Minify returns a io.Reader that removes all special characters from the source stream and leave only
// valid brainfuck commands
func Minify(source io.Reader) io.Reader {
	return &minifyReader{source}
}
