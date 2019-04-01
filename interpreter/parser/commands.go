package parser

import (
	"github.com/ibraimgm/bfi/interpreter/token"
)

const maskCmdType byte = 3 << 6
const maskCmdQty byte = 0xFF >> 2

// List of constants representing the parsed commands
const (
	CmdMoveRight byte = 0
	CmdMoveLeft  byte = 1
	CmdInc       byte = 2
	CmdDec       byte = 3

	CmdOutput byte = 0
	CmdInput  byte = 1
	CmdJump   byte = 2
	CmdReturn byte = 3
)

// ExtractCommand return the relevant byte parts of the identified command
// and the number of times that the command repeats
func ExtractCommand(value byte) (byte, byte) {
	return (value & maskCmdType) >> 6, value & maskCmdQty
}

// EncodeCommand return the byte equivalent of the command represented by
// cmdToken, repeated by qty times. Non-repeating commands ignore the value of
// qty.
func EncodeCommand(cmdToken rune, qty byte) byte {
	var cmd byte
	q := qty

	switch cmdToken {
	case token.MoveRight:
		cmd = CmdMoveRight
	case token.MoveLeft:
		cmd = CmdMoveLeft
	case token.Inc:
		cmd = CmdInc
	case token.Dec:
		cmd = CmdDec
	case token.Output:
		cmd = CmdOutput
	case token.Input:
		cmd = CmdInput
	case token.Jump:
		cmd = CmdJump
	case token.Return:
		cmd = CmdReturn
	}

	switch cmdToken {
	case token.Output:
		fallthrough
	case token.Input:
		fallthrough
	case token.Jump:
		fallthrough
	case token.Return:
		q = 0
	}

	return (cmd << 6) | q
}
