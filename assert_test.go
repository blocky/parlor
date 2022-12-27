package parlor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/parlor"
)

type Fooer interface {
	Foo()
}

type AFooer int

func (AFooer) Foo() {}

func TestTypePredicate(t *testing.T) {
	assert.True(t, parlor.TypePredicate[Fooer](AFooer(0)))
	assert.False(t, parlor.TypePredicate[Fooer](int(0)))

}
