package parlor

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TypePredicate[T any](obj any) bool {
	_, ok := obj.(T)
	return ok
}

func AssertType[T any](t *testing.T, obj any) {
	ok := TypePredicate[T](obj)

	assert.Truef(
		t,
		ok,
		"obj of type %T is not %v",
		obj,
		reflect.TypeOf((*T)(nil)).String()[1:],
	)
}
