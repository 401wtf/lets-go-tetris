VERSION = $(shell gobump show -r)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-X github.com/apollohpp/lets-go-tetris.revision=$(CURRENT_REVISION)"
ifdef update
  u=-u
endif

GO ?= GO111MODULE=on go

deps:
	env GO111MODULE=on go mod download

devel-deps: deps
	$(GO) get ${u} \
	  golang.org/x/lint/golint             \
	  github.com/mattn/goveralls           \
	  github.com/motemen/gobump/cmd/gobump \
	  github.com/Songmu/goxz/cmd/goxz      \
	  github.com/Songmu/ghch/cmd/ghch

test: deps
	$(GO) test -a -coverprofile=coverage.out -covermode=count ./...

lint: devel-deps
	$(GO) vet ./...
	golint -set_exit_status ./...

cover: devel-deps
	COVERALLS_TOKEN=${COVERALLS_TOKEN} goveralls -coverprofile=coverage.out --service=travis-ci

bump: devel-deps
	_tools/releng

upload:
	ghr v$(VERSION) dist/v$(VERSION)

release: bump crossbuild upload

.PHONY: deps devel-deps test lint cover build crossbuild build upload release
