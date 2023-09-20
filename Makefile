ROOT_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
RACE_OPT :=
CGO_DEBUG_OPT :=
CGO_OPTS=CGO_CFLAGS="-I$(ROOT_DIR)/cgo/thirdparty/libfaiss-src/c_api" CGO_LDFLAGS="-L$(ROOT_DIR)/cgo/thirdparty/runtimes/osx-arm64/ -lfaiss_c -lm"
BIN_NAME := mo-service

# build mo-service binary
.PHONY: build
build:
	$(info [Build binary])
	$(CGO_OPTS) go build  $(RACE_OPT) -o $(BIN_NAME) ./cmd/mo-service