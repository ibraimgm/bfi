package vm

import (
	"fmt"
)

// InvalidCellTypeError indicates that a wrong cell size was specified.
// To avoid this problem, make sure you use one of the CellType constans defined in this package.
type InvalidCellTypeError int

func (err InvalidCellTypeError) Error() string {
	return fmt.Sprintf("invalid cell type error (received %v). Try using one of the Cell<size> constants, like 'Cell8'.", int(err))
}
