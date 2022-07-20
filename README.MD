[![GoDoc](https://godoc.org/github.com/blocky/parlor?status.svg)](https://godoc.org/github.com/blocky/parlor)
[![Build Status](https://app.travis-ci.com/blocky/parlor.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://app.travis-ci.com/blocky/parlor)
[![Go Report Card](https://goreportcard.com/badge/github.com/blocky/parlor)](https://goreportcard.com/report/github.com/blocky/parlor)

# Parlor - a Testify Suite wrapper
Parlor is a simple wrapper for the Testify Suite by Stretchr
https://github.com/stretchr/testify

## The Issue
Currently, Suite does not allow the functions SetupTest() and TearDownTest()
to be run before and after subtests
```
type MyTestSuite structure {
	suite.Suite
}

func (s *MyTestSuite) SetupTest() {
	// do setup before a test
}

func (s *MyTestSuite) TearDownTest() {
	// teardown after a test
}

func (s *MyTestSuite) Test() {
	// SetupTest() called

	s.Run("subtest 1", func() {
		// do subtest 1
	})

	s.Run("subtest 2", func() {
		// do subtest 2
	})

	// TearDownTest() called
}
```

## An Error-Prone Solution
One solution is to call the suite's SetupTest and TearDownTest before and after
all subtests, but this is prone to error if we forget the function calls
```
...

func (s *MyTestSuite) Test() {
	// SetupTest() called

	s.SetupTest()
	s.Run("subtest 1", func() {
		// do subtest 1
	})
	s.TearDownTest()

	s.SetupTest()
	s.Run("subtest 2", func() {
		// do subtest 2
	})
	s.TearDownTest()

	// TearDownTest() called
}
```

## A Better Solution
Parlor attempts to use Suite interfaces and similar paradigms to make an effective
and minimal solution

### Embedding parlor.Parlor
Suite's `suite.Suite` implements the `Suite.TestingSuite` interface. We follow
the paradigms of `suite.Suite`, and embed a `suite.Suite` into your testing struct. The Parlor
struct embeds a `suite.Suite` in itself, so you can embed a `parlor.Parlor` instead
```
// old way
type MyTestSuite struct {
	suite.Suite
}

// new way
type MyTestParlor struct {
	parlor.Parlor
}
```

### Add your Suite to the Testing Runtime
To add your test struct to the test runtime, simply follow the same `Run`
pattern as the Suite package dictates
```
...

//old way
func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}

// new way
func TestMyTestParlor(t *testing.T) {
	parlor.Run(t, new(MyTestParlor))
}
```

### Implementing the SetupTestSuite & TearDownTestSuite Interfaces
Now, define the `SetupTest` and `TearDownTest` functions for your test struct
```
type MyTestParlor struct {
	parlor.Parlor
}

func TestMyTestParlor(t *testing.T) {
	parlor.Run(t, new(MyTestParlor))
}

func (p *Parlor) SetupTest() {
	// my setup before a test or subtest
}

func (p *Parlor) TearDownTest() {
	// my teardown after a test or subtest
}
```

### Using the Custom Run Function
Now, to ensure that your `SetupTest` and `TearDownTest` are called before and
after a test or subtest, use the special `Parlor.Run` function
```
...

func (p *MyTestParlor) Test() {
	p.Run("subtest 1", func() {
		// subtest 1
	}, p)

	p.Run("subtest 2", func() {
		// subtest 2
	}, p)

	p.Run("subtest 3", func() {
		// subtest 3
	}, p)
}
```
Now your subtests, using the custom `Run` function of `parlor.Parlor`, will
call the `SetupTest` and `TearDownTest` functions of your test struct for every
subtest

### Notes
+ your test struct must implement the `SetupTest` and `TearDownTest`
functions or else `parlor.Parlor`'s `Run` function will not compile from
unimplemented interfaces

+ At the very beginning and very end of your test struct's test functions, `SetupTest`
will be called twice, as well as at the end `TearDownTest` will also be called
twice. Parlor does not rewrite Suite to reorganize the existing testing
interface calls (and solve the tricky issue of nested subtests), but simply adds
in calls to `SetupTest` and `TearDownTest`.