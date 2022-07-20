package internal

type Butler interface {
	SetupTest()
	TearDownTest()
}
