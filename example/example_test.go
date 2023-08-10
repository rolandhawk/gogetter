package example_test

import (
	"testing"

	"github.com/rolandhawk/gogetter/example"
)

func TestBasicType(t *testing.T) {
	var b *example.BasicType
	b.GetBool()
}
