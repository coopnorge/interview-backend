SHELL=bash
.SHELLFLAGS=-ec -o pipefail

GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

ifeq ($(GOHOSTOS), windows)
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api/logistics -name *.proto")
else
	API_PROTO_FILES=$(shell find 'api\/logistics' -name *.proto)
endif

.PHONY: init
init:
	rm -rf internal/app/logistics/api/v1
	buf generate
