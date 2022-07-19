package parlor_test

import (
	"testing"

	"github.com/blocky/parlor"
	"github.com/blocky/parlor/internal/mocks"
)

type ParlorTestSuite struct {
	parlor.Parlor
	butler *mocks.Butler
}

func TestParlorTestSuite(t *testing.T) {
	parlor.Run(t, new(ParlorTestSuite))
}

func (p *ParlorTestSuite) SetupSuite() {
	p.butler = new(mocks.Butler)
	p.butler.On("SetupTest").Return().Once()
}

func (p *ParlorTestSuite) TearDownSuite() {
	p.butler.AssertExpectations(p.T())
}

func (p *ParlorTestSuite) SetupTest() {
	p.butler.SetupTest()
}

func (p *ParlorTestSuite) TearDownTest() {
	p.butler.TearDownTest()
}

func (p *ParlorTestSuite) TestParlor() {
	p.butler.On("SetupTest").Return().Once()

	p.Run("subtest 1", func() {
		p.butler.On("TearDownTest").Return().Once()
	}, p)

	p.butler.On("SetupTest").Return().Once()

	p.Run("subtest 2", func() {
		p.butler.On("TearDownTest").Return().Once()
	}, p)

	p.butler.On("TearDownTest").Return().Once()
}
