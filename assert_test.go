package parlor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/parlor"
)

type Fooer interface {
	Foo()
}

type AFooer struct{}

func (AFooer) Foo() {}

type NotAFooer struct{}

func TestTypePredicate(t *testing.T) {
	assert.True(t, parlor.TypePredicate[Fooer](AFooer{}))
	assert.False(t, parlor.TypePredicate[Fooer](NotAFooer{}))

}
