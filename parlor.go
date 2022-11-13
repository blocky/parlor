package parlor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SetupSubtestSuite interface {
	SetupSubtest()
}

type TearDownSubtestSuite interface {
	TearDownSubtest()
}

type TestingParlor interface {
	suite.TestingSuite
	SetTestingParlor(TestingParlor)
}

type Parlor struct {
	suite.Suite
	testingParlor interface{}
}

func (p *Parlor) SetTestingParlor(tp TestingParlor) {
	p.testingParlor = tp
}

func Run(t *testing.T, parlor TestingParlor) {
	parlor.SetTestingParlor(parlor)
	suite.Run(t, parlor)
}

func (p *Parlor) Run(name string, subtest func()) bool {
	setup := func() {}
	if i, ok := p.testingParlor.(SetupSubtestSuite); ok {
		setup = i.SetupSubtest
	}

	teardown := func() {}
	if i, ok := p.testingParlor.(TearDownSubtestSuite); ok {
		teardown = i.TearDownSubtest
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
