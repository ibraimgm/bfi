package vm

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestCellValues(t *testing.T) {
	testCases := make([]Cell, 4)
	testCases[0], _ = newCell(Cell8)
	testCases[1], _ = newCell(Cell16)
	testCases[2], _ = newCell(Cell32)
	testCases[3], _ = newCell(Cell64)

	checkValue := func(n int, expected uint8, c Cell) {
		if c.ToUint8() != expected {
			t.Errorf("Case %v, expected cell value to be %v, received \"%v\"", n, expected, c.ToUint8())
		}

		if expected == 0 && !c.IsZero() {
			t.Errorf("Case %v, IsZero should be \"true\".", n)
		}

		se := fmt.Sprintf("%d", expected)

		if c.String() != se {
			t.Errorf("Case %v, expected print value to be \"%v\", received \"%v\"", n, se, c.String())
		}
	}

	for i, c := range testCases {
		checkValue(i, 0, c)

		c.Inc()
		checkValue(i, 1, c)

		c.Add(17)
		checkValue(i, 18, c)

		c.Dec()
		checkValue(i, 17, c)

		c.Subtract(12)
		checkValue(i, 5, c)

		c.Zero()
		checkValue(i, 0, c)
	}
}

func TestCellCast(t *testing.T) {
	// Note: math constants are untyped.
	// See https://github.com/golang/go/issues/24523 for details
	var max8 uint8 = math.MaxUint8
	var max16 uint16 = math.MaxUint16
	var max32 uint32 = math.MaxUint32
	var max64 uint64 = math.MaxUint64

	// small hack to force a maximum value and test for overflow
	hack := &cellImpl{^uint64(0)}
	c := Cell(hack)

	if c.ToUint64() != max64 {
		t.Errorf("Wrong int64 value, expected \"%d\", received \"%d\"", max64, c.ToUint64())
	}

	if c.ToUint32() != max32 {
		t.Errorf("Wrong int32 value, expected \"%d\", received \"%d\"", max32, c.ToUint32())
	}

	if c.ToUint16() != max16 {
		t.Errorf("Wrong int16 value, expected \"%d\", received \"%d\"", max16, c.ToUint16())
	}

	if c.ToUint8() != max8 {
		t.Errorf("Wrong int8 value, expected \"%d\", received \"%d\"", max8, c.ToUint8())
	}
}

func TestNewCellError(t *testing.T) {
	c, err := newCell(CellType(-1))

	if c != nil {
		t.Errorf("Should have returned a nil cell")
	}

	if err == nil {
		t.Errorf("The error should not be nil")
	}

	if e2, ok := err.(InvalidCellTypeError); !ok {
		t.Errorf("Wrong error type. Expected \"InvalidCellTypeError\", received \"%T\"", e2)
	} else if !strings.Contains(e2.Error(), "-1") {
		t.Errorf("Wrong error message. Received \"%v\"", e2.Error())
	}
}
