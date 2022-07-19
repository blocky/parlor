package internal

type Butler interface {
	SetupTestCalled()
	TearDownTestCalled()
}
