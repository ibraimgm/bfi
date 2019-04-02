package parser

import (
	"bufio"
	"io"

	"github.com/ibraimgm/bfi/interpreter"
	"github.com/ibraimgm/bfi/interpreter/token"
)

type parseReader struct {
	r *bufio.Reader
}

const emptyToken = '\x00'

func (p *parseReader) Read(buffer []byte) (int, error) {
	n := 0
	maxSize := len(buffer)
	lastCmd := emptyToken
	lastQty := byte(0)

	for {
		b, err := p.r.ReadByte()

		if err == nil {
			currToken := rune(b)

			if currToken == lastCmd {
				if lastQty == 63 {
					buffer[n] = EncodeCommand(lastCmd, lastQty)
					n++
					lastQty = 0
				}

				if n == maxSize {
					err = p.r.UnreadByte()
					return n, err
				}

				lastQty++
			} else {
				if lastCmd != emptyToken {
					buffer[n] = EncodeCommand(lastCmd, lastQty)
					n++
					lastQty = 0

					if n == maxSize {
						err = p.r.UnreadByte()
						return n, err
					}
				}

				switch currToken {
				case token.MoveRight:
					fallthrough
				case token.MoveLeft:
					fallthrough
				case token.Inc:
					fallthrough
				case token.Dec:
					lastCmd = currToken
					lastQty = 1
				default:
					buffer[n] = EncodeCommand(currToken, 0)
					lastCmd = emptyToken
					n++
				}
			}
		}

		if n >= maxSize {
			return n, err
		}

		if err != nil {
			if lastQty > 0 {
				buffer[n] = EncodeCommand(lastCmd, lastQty)
				n++
			}

			return n, err
		}
	}
}

// Parse returns a new io.Reader that transform the source code into a smaller
// representation of itself
func Parse(source io.Reader) io.Reader {
	minified := interpreter.Minify(source)
	buffered := bufio.NewReader(minified)
	return &parseReader{buffered}
}
