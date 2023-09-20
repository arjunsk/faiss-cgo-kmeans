RACE_OPT :=
CGO_DEBUG_OPT :=
CGO_OPTS=CGO_CFLAGS="-I$(ROOT_DIR)/cgo/test" CGO_LDFLAGS="-L$(ROOT_DIR)/cgo/test -lmo -lm"
BIN_NAME := mo-service

.PHONY: cgo
cgo:
	@(cd cgo/test; ${MAKE} ${CGO_DEBUG_OPT})

# build mo-service binary
.PHONY: build
build: cgo
	$(info [Build binary])
	$(CGO_OPTS) go build  $(RACE_OPT) -o $(BIN_NAME) ./cmd/mo-service