package vm_test

import (
	"strings"
	"testing"

	"github.com/ibraimgm/bfi/vm"
)

type expectedCell struct {
	cell  int
	value byte
}

func TestLoadFromString(t *testing.T) {
	testCases := []struct {
		source string
		cells  []expectedCell
	}{
		{
			source: `+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++`,
			cells:  []expectedCell{{cell: 0, value: 65}},
		},
		{
			source: `++++++++[->+>+>+<<<]+++[->>>-<<<]+>+`,
			cells: []expectedCell{
				{cell: 0, value: 1},
				{cell: 1, value: 9},
				{cell: 2, value: 8},
				{cell: 3, value: 5},
			},
		},
		{
			source: `>+++++++++>>>>>>>>>>++++++++[-<<<<<<<<<<<>[>]<-[->+>+<<]>>[-<<+>>]<<+[<]>>>>>>>>>>>]<<[<]>[[-<+>]>]`,
			cells: []expectedCell{
				{cell: 0, value: 9},
				{cell: 1, value: 8},
				{cell: 2, value: 7},
				{cell: 3, value: 6},
				{cell: 4, value: 5},
				{cell: 5, value: 4},
				{cell: 6, value: 3},
				{cell: 7, value: 2},
				{cell: 8, value: 1},
				{cell: 9, value: 0},
			},
		},
	}

	for i, test := range testCases {
		machine, err := vm.LoadFromString(test.source)

		if err != nil {
			t.Errorf(err.Error())
		}

		machine.Run()
		results := machine.GetTapeState()

		for j, cell := range test.cells {
			expected := cell.value
			received := results[cell.cell].ToUint8()

			if received != expected {
				t.Errorf("Case %v, cell %v, value mismatch. Expected \"%v\", received \"%v\"", i, j, expected, received)
			}
		}

		machine.Reset()
		results = machine.GetTapeState()
		for j, cell := range test.cells {
			received := results[cell.cell].ToUint8()

			if received != 0 {
				t.Errorf("Case %v, cell %v, value mismatch. Expected \"0\" after reset, received \"%v\"", i, j, received)
			}
		}
	}
}

func TestGetTapeStateReturnsACopy(t *testing.T) {
	testCases := []struct {
		source   string
		expected byte
	}{
		{
			source:   `+++`,
			expected: 3,
		},
		{
			source:   `+-`,
			expected: 0,
		},
	}

	for i, test := range testCases {
		machine, err := vm.LoadFromString(test.source)

		if err != nil {
			t.Errorf(err.Error())
		}

		machine.Run()
		stateA := machine.GetTapeState()

		if stateA[0].ToUint8() != test.expected {
			t.Errorf("Case %v, initial value mismatch. Expected \"%v\", received \"%v\"", i, test.expected, stateA[0].ToUint8())
		}

		stateA[0].Inc()
		stateB := machine.GetTapeState()

		if stateA[0].ToUint8() == stateB[0].ToUint8() {
			t.Errorf("Case %v, invalid mutation. A is \"%v\", B is \"%v\"", i, stateA[0].ToUint8(), stateB[0].ToUint8())
		}
	}
}

func TestIO(t *testing.T) {
	testCases := []struct {
		source  string
		inputs  string
		outputs string
		cells   []expectedCell
	}{
		{
			source:  ",,,.",
			inputs:  "ABC",
			outputs: "C",
			cells: []expectedCell{
				{cell: 0, value: 67},
			},
		},
		{
			source:  "++>,<[->+<]>>,>,--[<]>.>.>.",
			inputs:  "ABC",
			outputs: "CBA",
			cells: []expectedCell{
				{cell: 0, value: 0},
				{cell: 1, value: 67},
				{cell: 2, value: 66},
				{cell: 3, value: 65},
			},
		},
		{
			source:  "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			inputs:  "",
			outputs: "Hello World!\n",
			cells: []expectedCell{
				{cell: 0, value: 0},
				{cell: 1, value: 0},
				{cell: 2, value: 72},
				{cell: 3, value: 100},
				{cell: 4, value: 87},
				{cell: 5, value: 33},
				{cell: 6, value: 10},
			},
		},
		{
			source:  "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+.",
			inputs:  "",
			outputs: "Hello, World!",
			cells: []expectedCell{
				{cell: 0, value: 172},
				{cell: 1, value: 108},
				{cell: 2, value: 44},
				{cell: 3, value: 33},
				{cell: 4, value: 87},
			},
		},
		{
			source:  ">>>>++++++++++[->++++++++++[-<<<+<+<+>>>>>]<]<<+++++<++<--[.>]",
			inputs:  "",
			outputs: "bfi",
			cells: []expectedCell{
				{cell: 0, value: 98},
				{cell: 1, value: 102},
				{cell: 2, value: 105},
			},
		},
	}

	for i, test := range testCases {
		machine, err := vm.LoadFromString(test.source)

		if err != nil {
			t.Errorf(err.Error())
		}

		reader := strings.NewReader(test.inputs)
		writer := strings.Builder{}
		machine.SetIO(reader, &writer)
		machine.Run()

		if writer.String() != test.outputs {
			t.Errorf("Case %v, expected output to be \"%v\", but it was \"%v\".", i, test.outputs, writer.String())
		}

		results := machine.GetTapeState()
		for j, cell := range test.cells {
			expected := cell.value
			received := results[cell.cell].ToUint8()

			if received != expected {
				t.Errorf("Case %v, cell %v, value mismatch. Expected \"%v\", received \"%v\"", i, j, expected, received)
			}
		}
	}
}
