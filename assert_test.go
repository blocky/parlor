package parlor_test

import (
	"testing"

	"github.com/blocky/parlor"
	"github.com/stretchr/testify/assert"
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

type NotAFooer struct{}

func TestIsFooer(t *testing.T) {
	parlor.AssertType[Fooer](t, NotAFooer{})
}
