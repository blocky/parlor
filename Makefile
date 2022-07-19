GOFLAGS=GOFLAGS=-mod=mod
GO=$(GOFLAGS) go
GOMOD=$(GO) mod
GOTEST=$(GO) test -count=1
GOTIDY=$(GOMOD) tidy

default: test

test: mock
	@$(GOTEST) -v ./... && $(GOTIDY)

mock:
	cd internal && mockery --all