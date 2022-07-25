package parlor_test

import (
	"testing"

	"github.com/blocky/parlor"
)

type ParlorYesSetupNoTearDown struct {
	parlor.Parlor
	setupCallCount int
}

func TestParlorYesSetupNoTearDown(t *testing.T) {
	parlor.Run(t, new(ParlorYesSetupNoTearDown))
}

func (p *ParlorYesSetupNoTearDown) SetupSuite() {
	p.setupCallCount = 0
}

func (p *ParlorYesSetupNoTearDown) SetupTest() {
	p.setupCallCount += 1
}

func (p *ParlorYesSetupNoTearDown) TestParlor() {
	p.Equal(1, p.setupCallCount, "should be called once before the test")

	p.Run("subtest 1", func() {
		p.Equal(2, p.setupCallCount, "should be called again before subtest 1")
	})

	p.Run("subtest 2", func() {
		p.Equal(3, p.setupCallCount, "should be called again before subtest 2")
	})
}

func (p *ParlorYesSetupNoTearDown) TearDownSuite() {
	p.Equal(3, p.setupCallCount, "should not be called again")
}

type ParlorNoSetupYesTearDown struct {
	parlor.Parlor
	tearDownCallCount int
}

func TestParlorNoSetupYesTearDown(t *testing.T) {
	parlor.Run(t, new(ParlorNoSetupYesTearDown))
}

func (p *ParlorNoSetupYesTearDown) SetupSuite() {
	p.tearDownCallCount = 0
}

func (p *ParlorNoSetupYesTearDown) TearDownTest() {
	p.tearDownCallCount += 1
}

func (p *ParlorNoSetupYesTearDown) TestParlor() {
	p.Equal(0, p.tearDownCallCount, "tear down not yet called")

	p.Run("subtest 1", func() {})
	p.Equal(1, p.tearDownCallCount, "should be called after subtest 1")

	p.Run("subtest 2", func() {})
	p.Equal(2, p.tearDownCallCount, "should be called after subtest 2")
}

func (p *ParlorNoSetupYesTearDown) TearDownSuite() {
	p.Equal(3, p.tearDownCallCount, "one more time after test parlor completes")
}
