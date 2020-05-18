.PHONY: test help default

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

default: test

help:
	@echo 'Management commands for goexim:'
	@echo
	@echo 'Usage:'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

test:
	go test -coverprofile cp.out ./...

test-coverage:
	go tool cover -html=cp.out

