package vm

import (
	"errors"
	"fmt"
)

// CellType define the type of the cell used on the tape
type CellType int

// List the available cell types. They differ only in size, and the
// original brainf*ck assumes a 8 bit cell (Cell8)
const (
	Cell8 CellType = iota
	Cell16
	Cell32
	Cell64
)

func ParseCellType(value int) (CellType, error) {
	switch value {
	case 8:
		return Cell8, nil
	case 16:
		return Cell16, nil
	case 32:
		return Cell32, nil
	case 64:
		return Cell64, nil
	default:
		return 0, errors.New(fmt.Sprintf("invalid cell type: %v", value))
	}
}

// Cell represents a cell in the tape. A cell might have 8, 16, 32 or 64
// bits (this is defined in the tape creation).
//
// Be wary that by using the ToUint* methods, you must take care to ensure
// you are using the correct integer type.
type Cell interface {
	fmt.Stringer
	Inc()
	Add(value byte)
	Dec()
	Subtract(value byte)
	Zero()
	IsZero() bool
	ToUint8() uint8
	ToUint16() uint16
	ToUint32() uint32
	ToUint64() uint64
	Clone() Cell
}

type cellImpl struct {
	inner interface{}
}

func (c *cellImpl) Inc() {
	c.Add(1)
}

func (c *cellImpl) Add(value byte) {
	switch c.inner.(type) {
	case uint8:
		c.inner = c.inner.(uint8) + value
	case uint16:
		c.inner = c.inner.(uint16) + uint16(value)
	case uint32:
		c.inner = c.inner.(uint32) + uint32(value)
	case uint64:
		c.inner = c.inner.(uint64) + uint64(value)
	}
}

func (c *cellImpl) Dec() {
	c.Subtract(1)
}

func (c *cellImpl) Subtract(value byte) {
	switch c.inner.(type) {
	case uint8:
		c.inner = c.inner.(uint8) - value
	case uint16:
		c.inner = c.inner.(uint16) - uint16(value)
	case uint32:
		c.inner = c.inner.(uint32) - uint32(value)
	case uint64:
		c.inner = c.inner.(uint64) - uint64(value)
	}
}

func (c *cellImpl) Zero() {
	switch c.inner.(type) {
	case uint8:
		c.inner = uint8(0)
	case uint16:
		c.inner = uint16(0)
	case uint32:
		c.inner = uint32(0)
	case uint64:
		c.inner = uint64(0)
	}
}

func (c *cellImpl) IsZero() bool {
	switch c.inner.(type) {
	case uint8:
		return c.inner.(uint8) == 0
	case uint16:
		return c.inner.(uint16) == 0
	case uint32:
		return c.inner.(uint32) == 0
	case uint64:
		return c.inner.(uint64) == 0
	}

	return false
}

func (c *cellImpl) String() string {
	return fmt.Sprintf("%v", c.inner)
}

func (c *cellImpl) ToUint8() uint8 {
	return uint8(c.ToUint64())
}

func (c *cellImpl) ToUint16() uint16 {
	return uint16(c.ToUint64())
}

func (c *cellImpl) ToUint32() uint32 {
	return uint32(c.ToUint64())
}

func (c *cellImpl) ToUint64() uint64 {
	switch c.inner.(type) {
	case uint8:
		return uint64(c.inner.(uint8))
	case uint16:
		return uint64(c.inner.(uint16))
	case uint32:
		return uint64(c.inner.(uint32))
	default:
		return c.inner.(uint64)
	}
}

func (c *cellImpl) Clone() Cell {
	other := *c
	return &other
}

func newCell(t CellType) (Cell, error) {
	switch t {
	case Cell8:
		return &cellImpl{uint8(0)}, nil
	case Cell16:
		return &cellImpl{uint16(0)}, nil
	case Cell32:
		return &cellImpl{uint32(0)}, nil
	case Cell64:
		return &cellImpl{uint64(0)}, nil
	default:
		return nil, InvalidCellTypeError(t)
	}
}
