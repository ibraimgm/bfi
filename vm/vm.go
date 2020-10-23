package vm

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ibraimgm/bfi/interpreter/parser"
)

// BFVM is a virtual machine capable of loading and running brainf*ck code
type BFVM struct {
	commands []byte
	tape     []Cell
	jumps    map[int]int
	stdin    io.Reader
	stdout   io.Writer
	position int
}

// LoadFromStream loads the brainf*ck source from the specified reader
// into the virtual machine instance
func (vm *BFVM) LoadFromStream(reader io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(parser.Parse(reader)); err != nil {
		return err
	}

	vm.commands = buf.Bytes()
	vm.position = 0

	s := newStack()

	for i, cmd := range vm.commands {
		cmd, qty := parser.ExtractCommand(cmd)

		if qty > 0 {
			continue
		}

		switch cmd {
		case parser.CmdJump:
			s.push(i)
		case parser.CmdReturn:
			addr, err := s.pop()

			if err != nil {
				return err
			}

			vm.jumps[addr] = i
			vm.jumps[i] = addr
		}
	}

	return nil
}

// LoadFromString parses and loads brainf*ck source into the virtual machine instance
func (vm *BFVM) LoadFromString(source string) error {
	return vm.LoadFromStream(strings.NewReader(source))
}

// SetIO sets the input/output stream used by the virtual machine
func (vm *BFVM) SetIO(in io.Reader, out io.Writer) {
	vm.stdin = in
	vm.stdout = out
}

// GetTapeState returns a copy of the current tape contents
func (vm *BFVM) GetTapeState() []Cell {
	tmp := make([]Cell, len(vm.tape))

	for i, c := range vm.tape {
		tmp[i] = c.Clone()
	}

	return tmp
}

// Run executes the currently loaded brainf*ck code.
// The current position or the values of the cells are not initialized; for that, use Reset().
func (vm *BFVM) Run() error {
	maxCells := len(vm.tape)
	maxCmds := len(vm.commands)
	buffer := make([]byte, 1)

	for i := 0; i < maxCmds; i++ {
		b := vm.commands[i]
		cmd, qty := parser.ExtractCommand(b)

		cell := vm.tape[vm.position]

		switch {
		case cmd == parser.CmdMoveRight && qty > 0:
			vm.position += int(qty)
			if vm.position >= maxCells {
				vm.position -= maxCells
			}

		case cmd == parser.CmdMoveLeft && qty > 0:
			vm.position -= int(qty)
			if vm.position < 0 {
				vm.position += maxCells
			}

		case cmd == parser.CmdInc && qty > 0:
			cell.Add(qty)

		case cmd == parser.CmdDec && qty > 0:
			cell.Subtract(qty)

		case cmd == parser.CmdJump:
			if cell.IsZero() {
				i = vm.jumps[i]
			}

		case cmd == parser.CmdReturn:
			if !cell.IsZero() {
				i = vm.jumps[i] - 1
			}

		case cmd == parser.CmdInput:
			if _, err := vm.stdin.Read(buffer); err != nil {
				return err
			}

			cell.Zero()
			cell.Add(buffer[0])

		case cmd == parser.CmdOutput:
			runes := []rune{rune(cell.ToUint32())}
			fmt.Fprintf(vm.stdout, "%v", string(runes))
		}
	}

	return nil
}

// Reset resets both the position of the tape and the cell values to 0.
func (vm *BFVM) Reset() {
	vm.position = 0

	for _, c := range vm.tape {
		c.Zero()
	}
}

// WithSpecs returns a new VM instance, with the specified cell size and tape size
func WithSpecs(cellSize int, tapeSize int) (*BFVM, error) {
	tape := make([]Cell, tapeSize)
	var err error

	for i := 0; i < tapeSize; i++ {
		tape[i], err = newCell(cellSize)

		if err != nil {
			return nil, err
		}
	}

	return &BFVM{tape: tape, jumps: make(map[int]int), stdin: os.Stdin, stdout: os.Stdout}, nil
}

// WithCellSize returns a new VM instance, with the specified cell size
func WithCellSize(cellSize int) (*BFVM, error) {
	return WithSpecs(cellSize, 3000)
}

// WithSize returns a new VM instance, with the specified tape size
func WithSize(tapeSize int) (*BFVM, error) {
	return WithSpecs(8, tapeSize)
}

// New returns a new VM, with default specs (3000 8-bit cells)
func New() (*BFVM, error) {
	return WithSpecs(8, 3000)
}

// LoadFromStream returns a new machine with default specs
// and with the source loaded from the specified reader
func LoadFromStream(reader io.Reader) (*BFVM, error) {
	machine, err := New()

	if err != nil {
		return nil, err
	}

	err = machine.LoadFromStream(reader)

	if err != nil {
		return nil, err
	}

	return machine, nil
}

// LoadFromString returns a new machine with default specs
// and the specified source preloaded
func LoadFromString(source string) (*BFVM, error) {
	machine, err := New()

	if err != nil {
		return nil, err
	}

	err = machine.LoadFromString(source)

	if err != nil {
		return nil, err
	}

	return machine, nil
}
