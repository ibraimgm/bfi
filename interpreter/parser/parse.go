package parser

import (
	"bufio"
	"io"

	"github.com/ibraimgm/bfi/interpreter"
	"github.com/ibraimgm/bfi/interpreter/token"
)

const emptyToken = '\x00'

type parseState struct {
	reader  *bufio.Reader
	buffer  []byte
	maxSize int
	current int
	lastCmd rune
	lastQty byte
}

func initState(buffer []byte, reader *bufio.Reader) *parseState {
	return &parseState{reader, buffer, len(buffer), 0, emptyToken, 0}
}

func (s *parseState) read() (rune, error) {
	b, err := s.reader.ReadByte()
	return rune(b), err
}

func (s *parseState) encodeCombined() {
	s.buffer[s.current] = EncodeCommand(s.lastCmd, s.lastQty)
	s.current++
	s.lastQty = 0
}

func (s *parseState) undoWhenOverflow() (int, bool, error) {
	if s.current == s.maxSize {
		err := s.reader.UnreadByte()
		return s.current, true, err
	}

	return s.current, false, nil
}

func (s *parseState) combineCommand() (int, bool, error) {
	if s.lastQty == 63 {
		s.encodeCombined()
	}

	read, undo, err := s.undoWhenOverflow()

	if !undo {
		s.lastQty++
	}

	return read, undo, err
}

func (s *parseState) commitPending() (int, bool, error) {
	if s.hasPendingCommand() {
		s.encodeCombined()
		return s.undoWhenOverflow()
	}

	return s.current, false, nil
}

func (s *parseState) shouldCombine(currToken rune) bool {
	return currToken == s.lastCmd
}

func (s *parseState) hasPendingCommand() bool {
	return s.lastCmd != emptyToken && s.lastQty > 0
}

func (s *parseState) initCombined(currToken rune) {
	s.lastCmd = currToken
	s.lastQty = 1
}

func (s *parseState) encode(currToken rune) {
	s.buffer[s.current] = EncodeCommand(currToken, 0)
	s.lastCmd = emptyToken
	s.current++
}

type parseReader struct {
	r *bufio.Reader
}

func (p *parseReader) Read(buffer []byte) (int, error) {
	st := initState(buffer, p.r)

	for {
		//check if we can read a new token
		if currToken, err := st.read(); err == nil {

			// if possible combine the read token, until the 6-bit limit
			// might undo the read if the buffer overflows
			if st.shouldCombine(currToken) {
				if read, undo, err := st.combineCommand(); undo {
					return read, err
				}
			} else {
				// if it is not combinable, commit the last pending combinable command.
				// this might undo the read if the buffer overflows
				if read, undo, err := st.commitPending(); undo {
					return read, err
				}

				// actual processing of a new (or non-combinable) token
				// this might either start a new combination or directly encode the
				// token, if it is not combinable
				switch currToken {
				case token.MoveRight:
					fallthrough
				case token.MoveLeft:
					fallthrough
				case token.Inc:
					fallthrough
				case token.Dec:
					st.initCombined(currToken)
				default:
					st.encode(currToken)
				}
			}
		} else {
			// in case of error or EOF, finish any remaining pending command
			// before exiting
			if st.hasPendingCommand() {
				st.encodeCombined()
			}

			return st.current, err
		}

		// check for a filled buffer
		if st.current >= st.maxSize {
			return st.current, nil
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
