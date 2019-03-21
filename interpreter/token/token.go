package token

// MoveRight is the rune used to mark when the tape should move to the right cell.
const MoveRight = '>'

// MoveLeft is the rune used to mark when the tape should move to the left cell.
const MoveLeft = '<'

// Inc is the rune used to mark when the the current cell value should be incremented.
const Inc = '+'

// Dec is the rune used to mark when the the current cell value should be decremented.
const Dec = '-'

// Output is the rune used to mark when the the current cell value should be prited to screen.
const Output = '.'

// Input is the rune used to mark when the the user input should be stored on the current cell.
const Input = ','

// Jump mark the need to go past the matching ] if the cell under the pointer is zero.
const Jump = '['

// Return jumps back to the matching [ if the cell under the pointer is nonzero.
const Return = ']'

// IsValid returns true if the rune c is a valid brainfuck token.
func IsValid(c rune) bool {
	switch c {
	case MoveRight, MoveLeft, Inc, Dec, Output, Input, Jump, Return:
		return true
	}

	return false
}
