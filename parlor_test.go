package parlor_test

import (
	"testing"

	"github.com/blocky/parlor"
	"github.com/stretchr/testify/suite"
)

type ParlorTestSuite struct {
	suite.Suite

	str string
}

func TestParlorTestSuite(t *testing.T) {
	parlor.Run(t, new(ParlorTestSuite))
}

func (p *ParlorTestSuite) SetupTest() {
	p.str = "setup"
}

func (p *ParlorTestSuite) TearDownTest() {
	p.str = "teardown"
}

func (p *ParlorTestSuite) TestParlor() {
	p.Equal("setup", p.str)

	parlor.RunSubtest(p, "subtest 1", func() {
		p.Equal("setup", p.str)
	})
	p.Equal("teardown", p.str)

	parlor.RunSubtest(p, "subtest 2", func() {
		p.Equal("setup", p.str)
	})
	p.Equal("teardown", p.str)
}
