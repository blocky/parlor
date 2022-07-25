package parlor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Parlor struct {
	suite.Suite
	testingParlor interface{}
}

func (p *Parlor) SetTestingParlor(tp TestingParlor) {
	p.testingParlor = tp
}

type TestingParlor interface {
	suite.TestingSuite
	SetTestingParlor(TestingParlor)
}

func Run(t *testing.T, parlor TestingParlor) {
	parlor.SetTestingParlor(parlor)
	suite.Run(t, parlor)
}

func (p *Parlor) Run(name string, subtest func()) bool {
	setup := func() {}
	if i, ok := p.testingParlor.(suite.SetupTestSuite); ok {
		setup = i.SetupTest
	}

	teardown := func() {}
	if i, ok := p.testingParlor.(suite.TearDownTestSuite); ok {
		teardown = i.TearDownTest
	}

	oldT := p.T()
	defer p.SetT(oldT)

	return oldT.Run(name, func(t *testing.T) {
		p.SetT(t)
		setup()
		defer teardown()
		subtest()
	})
}
