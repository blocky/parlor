package parlor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestingParlor interface {
	suite.TestingSuite
	suite.SetupTestSuite
	suite.TearDownTestSuite
}

func Run(t *testing.T, parlor TestingParlor) {
	suite.Run(t, parlor)
}

func RunSubtest(
	parlor TestingParlor,
	name string,
	subtest func(),
) bool {
	oldT := parlor.T()
	defer parlor.SetT(oldT)

	return oldT.Run(name, func(t *testing.T) {
		parlor.SetT(t)
		parlor.SetupTest()
		defer parlor.TearDownTest()
		subtest()
	})
}
