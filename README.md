[![GoDoc](https://godoc.org/github.com/blocky/parlor?status.svg)](https://godoc.org/github.com/blocky/parlor)
[![Build Status](https://app.travis-ci.com/blocky/parlor.svg?token=JczzdP6eMqmEqysZ8pDf&branch=main)](https://app.travis-ci.com/blocky/parlor)
[![Go Report Card](https://goreportcard.com/badge/github.com/blocky/parlor)](https://goreportcard.com/report/github.com/blocky/parlor)

# Parlor - opinionated testing tools

Parlor is Blocky's opinionated set of testing tools that augment
[stretchr/testify](https://github.com/stretchr/testify). The package contains a
few additional tools that we find useful for testing including:

- testing that a type conforms to a specific interface.
- setup and teardown for subtests called within a suite.

## A quick tour

### Assert Interfaces

Probably the easiest way to test that a struct conforms to an interface is by
casting. For example, the test below will catch that our struct is missing a
method.

    type Fooer interface {
        Foo()
    }

    type NotAFooer struct{}

    func TestIsFooer(t *testing.T) {
        var i interface{} = &NotAFooer{}
        _, ok := i.(Fooer)
        assert.True(t, ok)
    }

This package gives us a one liner.

    func TestIsFooer_OneLine(t *testing.T) {
        parlor.AssertType[Fooer](t, NotAFooer{})
    }


### Suites

Currently, testify Suite does not provide an easy mechanism for setup and
teardown for subtests. There has be some
[discussion](https://github.com/stretchr/testify/issues/1031), however, the
community decided that the semantics of nested subtests is not general enough
for the library.
At Blocky, however, we decided that we want to have a hook for before and after
subtests (even if they are nested).  And so we use `parlor.Parlor` to
have setup and teardowns for subtests.

For example run parlor as you would with a suite:

    func TestMyTestParlor(t *testing.T) {
        parlor.Run(t, new(MyTestParlor))
    }

You can setup/teardown tests (available in testify) and subtests (only available
in parlor).  You can omit any of these function if they are not needed.

    type MyTestParlor struct {
        parlor.Parlor
    }

    func (p *MyTestParlor) SetupTest() {
        // my setup before a test
    }

    func (p *MyTestParlor) SetupSubtest() {
        // my setup before a subtest
        // specific to parlor
    }

    func (p *MyTestParlor) TearDownTest() {
        // my teardown after a test
    }

    func (p *MyTestParlor) TearDownSubtest() {
        // my teardown after a subtest
        // specific to parlor
    }

Now, your `SetupTest` and `TearDownTest` are called before and
after each test (as was the case with `suite.Suite`) and your
`SetupSubtest` and `TearDownSubtest` are called before and after each `t.Run`.

    func (p *MyTestParlor) ATest() {
        p.Run("subtest 1", func() {
            // subtest 1
        })

        p.Run("subtest 2", func() {
            // subtest 2
        })

        p.Run("subtest 3", func() {
            // subtest 3
        })
    }
