package parlor_test

import (
	"fmt"
	"testing"

	"github.com/blocky/parlor"
)

type ParlorTestSuite struct {
	parlor.Parlor

	str string
}

func TestParlorTestSuite(t *testing.T) {
	parlor.Run(t, new(ParlorTestSuite))
}

func (p *ParlorTestSuite) SetupTest() {
	fmt.Println("setuptest")
	p.str = "setup"
}

func (p *ParlorTestSuite) TearDownTest() {
	fmt.Println("teardowntest")
	p.str = "teardown"
}

func (p *ParlorTestSuite) TestParlor() {
	p.Equal("setup", p.str)

	p.Run("subtest 1", func() {
		p.Equal("setup", p.str)
	}, p)
	p.Equal("teardown", p.str)

	p.Run("subtest 2", func() {
		p.Equal("setup", p.str)
	}, p)
	p.Equal("teardown", p.str)
}
