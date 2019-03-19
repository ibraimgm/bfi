package dummy_test

import (
	"testing"

	"github.com/ibraimgm/bfi/dummy"
)

func TestDummy(t *testing.T) {
	if dummy.Dummy() != 0 {
		t.Errorf("we have a problem!")
	}
}
