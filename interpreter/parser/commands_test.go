package parser_test

import (
	"testing"

	"github.com/ibraimgm/bfi/interpreter/parser"
	"github.com/ibraimgm/bfi/interpreter/token"
)

func TestExtractCommand(t *testing.T) {
	testCases := []struct {
		value, command, quantity byte
	}{
		{value: 34, command: parser.CmdMoveRight, quantity: 34},
		{value: 109, command: parser.CmdMoveLeft, quantity: 45},
		{value: 138, command: parser.CmdInc, quantity: 10},
		{value: 193, command: parser.CmdDec, quantity: 1},
		{value: 0, command: parser.CmdOutput},
		{value: 64, command: parser.CmdInput},
		{value: 128, command: parser.CmdJump},
		{value: 192, command: parser.CmdReturn},
	}

	for i, test := range testCases {
		cmd, qty := parser.ExtractCommand(test.value)

		if cmd != test.command {
			t.Errorf("Case %v, received command \"%v\", expected \"%v\"", i, cmd, test.command)
		}

		if qty != test.quantity {
			t.Errorf("Case %v, received quantity \"%v\", expected \"%v\"", i, qty, test.quantity)
		}
	}

}

func TestEncodeCommand(t *testing.T) {
	testCases := []struct {
		cmd      rune
		qty      byte
		expected byte
	}{
		{cmd: token.MoveRight, qty: 63, expected: 63},
		{cmd: token.MoveLeft, qty: 63, expected: 127},
		{cmd: token.Inc, qty: 4, expected: 132},
		{cmd: token.Dec, qty: 3, expected: 195},
		{cmd: token.Output, qty: 7, expected: 0},
		{cmd: token.Input, qty: 2, expected: 64},
		{cmd: token.Jump, qty: 6, expected: 128},
		{cmd: token.Return, qty: 1, expected: 192},
	}

	for i, test := range testCases {
		encoded := parser.EncodeCommand(test.cmd, test.qty)

		if encoded != test.expected {
			t.Errorf("Case %v, token %v, received \"%v\", expected \"%v\"", i, test.cmd, encoded, test.expected)
		}
	}
}
