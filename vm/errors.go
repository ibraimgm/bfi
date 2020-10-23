package vm

import (
	"fmt"
)

// InvalidCellSizeError indicates that a wrong cell size was specified.
type InvalidCellSizeError int

func (err InvalidCellSizeError) Error() string {
	return fmt.Sprintf("invalid cell size: %v", int(err))
}
