package parser_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ibraimgm/bfi/interpreter/parser"
)

type expectedCommand struct {
	cmdType byte
	cmdQty  byte
}

func TestCompile(t *testing.T) {
	testCases := []struct {
		source   string
		commands []expectedCommand
	}{
		{
			source: `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`,
			commands: []expectedCommand{
				{cmdType: parser.CmdInc, cmdQty: 8},
				{cmdType: parser.CmdJump},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 4},
				{cmdType: parser.CmdJump},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 2},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 3},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 3},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 1},
				{cmdType: parser.CmdMoveLeft, cmdQty: 4},
				{cmdType: parser.CmdDec, cmdQty: 1},
				{cmdType: parser.CmdReturn},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 1},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 1},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdDec, cmdQty: 1},
				{cmdType: parser.CmdMoveRight, cmdQty: 2},
				{cmdType: parser.CmdInc, cmdQty: 1},
				{cmdType: parser.CmdJump},
				{cmdType: parser.CmdMoveLeft, cmdQty: 1},
				{cmdType: parser.CmdReturn},
				{cmdType: parser.CmdMoveLeft, cmdQty: 1},
				{cmdType: parser.CmdDec, cmdQty: 1},
				{cmdType: parser.CmdReturn},
				{cmdType: parser.CmdMoveRight, cmdQty: 2},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdDec, cmdQty: 3},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdInc, cmdQty: 7},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdInc, cmdQty: 3},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveRight, cmdQty: 2},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveLeft, cmdQty: 1},
				{cmdType: parser.CmdDec, cmdQty: 1},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveLeft, cmdQty: 1},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdInc, cmdQty: 3},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdDec, cmdQty: 6},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdDec, cmdQty: 8},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveRight, cmdQty: 2},
				{cmdType: parser.CmdInc, cmdQty: 1},
				{cmdType: parser.CmdOutput},
				{cmdType: parser.CmdMoveRight, cmdQty: 1},
				{cmdType: parser.CmdInc, cmdQty: 2},
				{cmdType: parser.CmdOutput},
			},
		},
	}

	for i, test := range testCases {
		buf := new(bytes.Buffer)
		buf.ReadFrom(parser.Parse(strings.NewReader(test.source)))
		compiled := buf.Bytes()

		for j, value := range compiled {
			cmd, qty := parser.ExtractCommand(value)

			if cmd != test.commands[j].cmdType {
				t.Errorf("Case %v, byte %v, command mismatch. Received \"%v\", expected \"%v\"", i, j, cmd, test.commands[j].cmdType)
			}

			if qty != test.commands[j].cmdQty {
				t.Errorf("Case %v, byte %v, qty mismatch. Received \"%v\", expected \"%v\"", i, j, qty, test.commands[j].cmdQty)
			}
		}
	}
}
