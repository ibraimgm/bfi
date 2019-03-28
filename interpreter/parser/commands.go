package parser

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
