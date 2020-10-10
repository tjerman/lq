GO         = go
GOGET      = $(GO) get -u
GOTEST    ?= go test
GOFLAGS   ?= -mod=vendor
GOPATH    ?= $(HOME)/go

export GOFLAGS

run:
	$(GO) run main/main.go
