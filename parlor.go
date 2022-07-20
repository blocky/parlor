package parlor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Parlor struct {
	suite.Suite
}

type TestingParlor interface {
	suite.TestingSuite
	suite.SetupTestSuite
	suite.TearDownTestSuite
}

func Run(t *testing.T, parlor TestingParlor) {
	suite.Run(t, parlor)
}

func (p *Parlor) Run(
	name string,
	subtest func(),
	tp TestingParlor,
) bool {
	return p.RunWithSetupAndTeardown(
		name,
		subtest,
		tp.SetupTest,
		tp.TearDownTest,
	)
}

func (p *Parlor) RunWithSetupAndTeardown(
	name string,
	subtest func(),
	setup func(),
	teardown func(),
) bool {
	oldT := p.T()
	defer p.SetT(oldT)

	return oldT.Run(name, func(t *testing.T) {
		p.SetT(t)
		setup()
		defer teardown()
		subtest()
	})
}
