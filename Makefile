ROOT_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
CGO_DEBUG_OPT :=
CGO_OPTS=CGO_CFLAGS="-I$(ROOT_DIR)/cgo/thirdparty/libfaiss-src/c_api" CGO_LDFLAGS="-L$(ROOT_DIR)/cgo/thirdparty/runtimes/osx-arm64/ -lfaiss_c -lm"
BIN_NAME := mo-service
UNAME_S := $(shell uname -s)

.PHONY: cgo
cgo:
	$(info [CGO build])
ifeq ($(UNAME_S),Darwin)
	@cd cgo/thirdparty && sh ./build-faiss-macos.sh
else
	@cd cgo/thirdparty && sh ./build-faiss-linux.sh
endif

# build mo-service binary
.PHONY: build
build:
	$(info [Build binary])
	$(CGO_OPTS) go build -o $(BIN_NAME) ./cmd/mo-service

run:
	./mo-service