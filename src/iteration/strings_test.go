package iteration

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	result := strings.Compare("sub", "sub")
	expected := 0
	if result != expected{
		t.Errorf("error type")

	}
}
