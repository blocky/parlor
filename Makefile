GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go
GOMOD=$(GO) mod
GOTEST=$(GO) test -count=1
GOTIDY=$(GOMOD) tidy

default: test

test:
	($(GOTEST) -v ./... && $(GOTIDY))