package parlor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/parlor"
)

type CallCounter struct {
	Test    int
	Subtest int
}

func (c CallCounter) Assert(t *testing.T, test, subtest int) {
	assert.Equal(t, test, c.Test)
	assert.Equal(t, subtest, c.Subtest)

}

// this test makes sure everything compiles and total call count is correct
// if we just have a setup
type WithSetup struct {
	parlor.Parlor
	setup CallCounter
}

func (p *WithSetup) SetupSubtest() {
	p.setup.Subtest += 1
}

func (p *WithSetup) TestParlor() {
	p.Run("subtest 1", func() {})
	p.Run("subtest 2", func() {})
}

func TestParlorWithSetup(t *testing.T) {
	p := new(WithSetup)
	parlor.Run(t, p)
	p.setup.Assert(t, 0, 2)
}

// this test makes sure everything compiles and total call count it correct
// if we just have a tear down
type WithTearDown struct {
	parlor.Parlor
	tearDown CallCounter
}

func (p *WithTearDown) TearDownSubtest() {
	p.tearDown.Subtest += 1
}

func (p *WithTearDown) TestParlor() {
	p.Run("subtest 1", func() {})
	p.Run("subtest 2", func() {})
}

func TestParlorWithTearDown(t *testing.T) {
	p := new(WithTearDown)
	parlor.Run(t, p)
	p.tearDown.Assert(t, 0, 2)
}

// this test just makes sure everything compiles if we don't have a setup or a
// tear down for a subtest but are using parlor
type NoSetupOrTearDown struct {
	parlor.Parlor
}

func (p *NoSetupOrTearDown) TestParlor() {
	p.Run("subtest 1", func() {})
	p.Run("subtest 2", func() {})
}

func TestParlorNoSetupOrTeardown(t *testing.T) {
	parlor.Run(t, new(NoSetupOrTearDown))
}

// this is the real test, we will make sure to count every setup and tear down
// call and as we move through the life cycle, we make sure that all calls are
// happening as expected.
type ParlorTestCallCounts struct {
	parlor.Parlor
	setup    CallCounter
	tearDown CallCounter
}

func (p *ParlorTestCallCounts) AssertSetup(test, subtest int) {
	p.setup.Assert(p.T(), test, subtest)
}

func (p *ParlorTestCallCounts) AssertTeardown(test, subtest int) {
	p.tearDown.Assert(p.T(), test, subtest)
}

func (p *ParlorTestCallCounts) SetupTest() {
	p.setup.Test += 1
}

func (p *ParlorTestCallCounts) SetupSubtest() {
	p.setup.Subtest += 1
}

func (p *ParlorTestCallCounts) TearDownTest() {
	p.tearDown.Test += 1
}

func (p *ParlorTestCallCounts) TearDownSubtest() {
	p.tearDown.Subtest += 1
}

func (p *ParlorTestCallCounts) SetupSuite() {
	p.AssertSetup(0, 0)
	p.AssertTeardown(0, 0)
}

func (p *ParlorTestCallCounts) TearDownSuite() {
	p.AssertSetup(1, 2)
	p.AssertTeardown(1, 2)
}

func (p *ParlorTestCallCounts) TestParlor() {
	p.AssertSetup(1, 0)
	p.AssertTeardown(0, 0)

	p.Run("subtest 1", func() {
		p.AssertSetup(1, 1)
		p.AssertTeardown(0, 0)
	})

	p.AssertSetup(1, 1)
	p.AssertTeardown(0, 1)

	p.Run("subtest 2", func() {
		p.AssertSetup(1, 2)
		p.AssertTeardown(0, 1)
	})

	p.AssertSetup(1, 2)
	p.AssertTeardown(0, 2)
}

func TestParlorCallCounts(t *testing.T) {
	p := new(ParlorTestCallCounts)
	parlor.Run(t, p)
	p.AssertSetup(1, 2)
	p.AssertTeardown(1, 2)
}
